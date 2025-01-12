package repository

import (
	"aitu-moment/db"
	"aitu-moment/models"
	"encoding/json"

	"fmt"
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

type Thread struct {
	ID             int64     `db:"thread_id" json:"id"`
	CreatorID      int64     `db:"creator_id" json:"creator_id"`
	Content        string    `db:"content" json:"content"`
	CreateDate     time.Time `db:"create_date" json:"create_date"`
	UpVotes        int       `db:"up_votes" json:"up_votes"`
	ParentThreadID *int64    `db:"parent_thread_id" json:"parent_thread_id,omitempty"`
}

type ThreadFilter struct {
	CreatorID      *int64
	ParentThreadID *int64
	StartDate      *time.Time
	EndDate        *time.Time
	MinUpVotes     *int
	Search         string
	Page           int
	PageSize       int
	OrderBy        *string
	Order          *string
}

func (r *UserRepository) FetchThreads(filter ThreadFilter) ([]Thread, int, error) {
	log.Debug("Building query")

	qb := db.QueryBuilder{
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
		qb.AddWhereCondition("creator_id =", *filter.CreatorID)
	}

	if filter.ParentThreadID != nil {
		qb.AddWhereCondition("parent_thread_id =", *filter.ParentThreadID)
	}

	if filter.StartDate != nil {
		qb.AddWhereCondition("create_date >=", *filter.StartDate)
	}

	if filter.EndDate != nil {
		qb.AddWhereCondition("create_date <=", *filter.EndDate)
	}

	if filter.MinUpVotes != nil {
		qb.AddWhereCondition("up_votes >=", *filter.MinUpVotes)
	}

	if filter.Search != "" {
		qb.AddWhereCondition("content ILIKE", "%"+filter.Search+"%")
	}

	if *filter.OrderBy != "" && *filter.Order != "" {
		qb.AddOrderBy(*filter.OrderBy, *filter.Order)
	}

	qb.Query.WriteString(`
        )
        SELECT 
            t.*,
            COUNT(*) OVER() as total_count
        FROM filtered_threads t`)

	qb.AddLimit(filter.PageSize)
	qb.AddOffset((filter.Page - 1) * filter.PageSize)

	jsonArgs, err := json.Marshal(qb.Args)

	if err != nil {
		log.Error("error mashalizing args: ", err)
	}

	//query, args, err := sqlx.In(qb.Query.String(), qb.Args...)

	log.WithFields(logrus.Fields{
		"query":      qb.Query.String(),
		"args_count": len(qb.Args),
		"args":       string(jsonArgs),
	}).Debug("Executing a query")

	type ThreadsRes struct {
		Total          int       `db:"total_count"`
		ID             int64     `db:"thread_id" json:"id"`
		CreatorID      int64     `db:"creator_id" json:"creator_id"`
		Content        string    `db:"content" json:"content"`
		CreateDate     time.Time `db:"create_date" json:"create_date"`
		UpVotes        int       `db:"up_votes" json:"up_votes"`
		ParentThreadID *int64    `db:"parent_thread_id" json:"parent_thread_id,omitempty"`
	}

	rows := make([]ThreadsRes, 0)

	err = r.DB.Select(&rows, qb.Query.String(), qb.Args...)

	if err != nil {
		log.WithError(err).Error("Error executing query")
		return nil, 0, fmt.Errorf("error querying threads: %w", err)
	}

	var totalCount int
	threads := make([]Thread, 0)
	for _, row := range rows {
		var thread Thread
		totalCount = row.Total
		thread.ID = row.ID
		thread.CreatorID = row.CreatorID
		thread.Content = row.Content
		thread.CreateDate = row.CreateDate
		thread.UpVotes = row.UpVotes
		thread.ParentThreadID = row.ParentThreadID
		threads = append(threads, thread)
	}

	log.WithFields(logrus.Fields{
		"threads":       threads,
		"threads_found": len(threads),
		"total_count":   totalCount,
	}).Info("Successfully fetched threads")

	log.Info("HELLO")
	log.Info(threads)

	return threads, totalCount, nil
}
