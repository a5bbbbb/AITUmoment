package repository

import (
	"aitu-moment/models"
	"database/sql"
	//"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)


var log = logrus.New()

type UserRepository struct {
	DB *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) GetAllUsers() ([]models.User, error) {
	query := " SELECT u.*, ep.name as program_name FROM users u JOIN educational_programs ep ON u.educational_program = ep.id"
	var users []models.User
	err := r.DB.Select(&users, query)
	return users, err
}

func (r *UserRepository) GetUser(userId int64) (models.User, error) {
	var user models.User
	query := ` 
        SELECT u.*, ep.name as program_name 
        FROM users u 
        JOIN educational_programs ep ON u.educational_program = ep.id 
        WHERE u.id = $1`

	err := r.DB.Get(&user, query, userId)
	return user, err
}

func (r *UserRepository) CreateUser(user models.User) (int64, error) {
	var insertedID int
	err := r.DB.QueryRow(
		"INSERT INTO users (username, educational_program) VALUES ($1, $2) RETURNING id",
		user.Name, user.EducationalProgram,
	).Scan(&insertedID)
	return int64(insertedID), err

}




type QueryBuilder struct {
	Query  strings.Builder
	Args   []interface{}
	Count  int
}


type Thread struct {
	ID            int64      `db:"thread_id" json:"id"`
	CreatorID     int64      `db:"creator_id" json:"creator_id"`
	Content       string     `db:"content" json:"content"`
	CreateDate    time.Time  `db:"create_date" json:"create_date"`
	UpVotes       int        `db:"up_votes" json:"up_votes"`
	ParentThreadID *int64    `db:"parent_thread_id" json:"parent_thread_id,omitempty"`
}


type ThreadFilter struct {
	CreatorID     *int64
	ParentThreadID *int64
	StartDate     *time.Time
	EndDate       *time.Time
	MinUpVotes    *int
	Search        string
	Page          int
	PageSize      int
}

func (qb *QueryBuilder) Add(condition string, value interface{}) {
	qb.Count++
	qb.Query.WriteString(fmt.Sprintf(" AND %s $%d", condition, qb.Count))
	qb.Args = append(qb.Args, value)
}

func (r *UserRepository) FetchThreads(filter ThreadFilter) ([]Thread, int, error) {
    log.Debug("Building query")

    qb := QueryBuilder{
        Query: strings.Builder{},
        Args:  make([]interface{}, 0),
    }

    qb.Query.WriteString(`
        WITH filtered_threads AS (
            SELECT 
                thread_id, creator_id, content, create_date, up_votes, parent_thread_id
            FROM threads
            WHERE 1=1
    `)

    // Add filter conditions
    if filter.CreatorID != nil {
        qb.Add("creator_id =", *filter.CreatorID)
        log.WithField("creator_id", *filter.CreatorID).Debug("Added creator_id filter")
    }

    if filter.ParentThreadID != nil {
        qb.Add("parent_thread_id =", *filter.ParentThreadID)
        log.WithField("parent_thread_id", *filter.ParentThreadID).Debug("Added parent_thread_id filter")
    }

    if filter.StartDate != nil {
        qb.Add("create_date >=", *filter.StartDate)
        log.WithField("start_date", *filter.StartDate).Debug("Added start_date filter")
    }

    if filter.EndDate != nil {
        qb.Add("create_date <=", *filter.EndDate)
        log.WithField("end_date", *filter.EndDate).Debug("Added end_date filter")
    }

    if filter.MinUpVotes != nil {
        qb.Add("up_votes >=", *filter.MinUpVotes)
        log.WithField("min_upvotes", *filter.MinUpVotes).Debug("Added min_upvotes filter")
    }

    if filter.Search != "" {
        qb.Add("content ILIKE", "%"+filter.Search+"%")
        log.WithField("search", filter.Search).Debug("Added search filter")
    }

    qb.Query.WriteString(`
        )
        SELECT 
            t.*,
            COUNT(*) OVER() as total_count
        FROM filtered_threads t
        ORDER BY create_date DESC
        LIMIT $` + strconv.Itoa(qb.Count+1) + ` OFFSET $` + strconv.Itoa(qb.Count+2))

    qb.Args = append(qb.Args, filter.PageSize, (filter.Page-1)*filter.PageSize)

    log.WithFields(logrus.Fields{
        "query": qb.Query.String(),
        "args_count": len(qb.Args),
    }).Debug("Executing query")

    rows, err := r.DB.Queryx(qb.Query.String(), qb.Args...)
    if err != nil {
        log.WithError(err).Error("Error executing query")
        return nil, 0, fmt.Errorf("error querying threads: %w", err)
    }
    defer rows.Close()

    var threads []Thread
    var totalCount int

    // Modified this part to directly scan into Thread struct
    for rows.Next() {
        var t Thread
        var total sql.NullInt64
        
        // Create a map of column pointers
        columns, _ := rows.ColumnTypes()
        values := make([]interface{}, len(columns))
        
        for i := range columns {
            switch columns[i].Name() {
            case "total_count":
                values[i] = &total
            case "thread_id":
                values[i] = &t.ID
            case "creator_id":
                values[i] = &t.CreatorID
            case "content":
                values[i] = &t.Content
            case "create_date":
                values[i] = &t.CreateDate
            case "up_votes":
                values[i] = &t.UpVotes
            case "parent_thread_id":
                values[i] = &t.ParentThreadID
            default:
                var v interface{}
                values[i] = &v
            }
        }

        if err := rows.Scan(values...); err != nil {
            log.WithError(err).Error("Error scanning row")
            return nil, 0, fmt.Errorf("error scanning thread: %w", err)
        }

        if total.Valid {
            totalCount = int(total.Int64)
        }

        threads = append(threads, t)
    }

    log.WithFields(logrus.Fields{
        "threads": threads,
        "threads_found": len(threads),
        "total_count": totalCount,
    }).Info("Successfully fetched threads")

    return threads, totalCount, nil
}


