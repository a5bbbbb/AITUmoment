package repository

import (
	"aitu-moment/db"
	"aitu-moment/models"

	"github.com/jmoiron/sqlx"
)

type ThreadRepo struct {
	db *sqlx.DB
}

func NewThreadRepo() *ThreadRepo {
	return &ThreadRepo{db: db.GetDB()}
}

func (r *ThreadRepo) GetParentThreads(userID *int) (*[]models.Thread, error) {

	query := `
        SELECT 
            t.*,
            u.public_name as creator_name,
            CASE 
                WHEN uv.id IS NOT NULL THEN true
                ELSE false
            END as has_upvote
        FROM threads t
        JOIN users u ON t.creator_id = u.id
        LEFT JOIN upvotes uv ON t.thread_id = uv.threadID AND uv.userID = $1
        WHERE t.parent_thread_id IS NULL
        ORDER BY t.create_date DESC
    `

	var threads []models.Thread
	err := r.db.Select(&threads, query, userID)

	return &threads, err
}

func (r *ThreadRepo) GetSubThreads(userID int, parentThread int) (*[]models.Thread, error) {
	query := `
        SELECT 
            t.*,
            u.public_name as creator_name,
            CASE 
                WHEN uv.id IS NOT NULL THEN true
                ELSE false
            END as "has_upvote"
        FROM threads t
        JOIN users u ON t.creator_id = u.id
        LEFT JOIN upvotes uv ON t.thread_id = uv.threadID AND uv.userID = $1
        WHERE t.parent_thread_id = $2
    `

	var threads []models.Thread
	err := r.db.Select(&threads, query, userID, parentThread)

	return &threads, err
}

func (r *ThreadRepo) GetThread(threadID int, userID int) (*models.Thread, error) {
	query := `
        SELECT 
            t.*,
            u.public_name as creator_name,
            CASE 
                WHEN uv.id IS NOT NULL THEN true
                ELSE false
            END as has_upvote
        FROM threads t
        JOIN users u ON t.creator_id = u.id
        LEFT JOIN upvotes uv ON t.thread_id = uv.threadID AND uv.userID = $1
        WHERE t.thread_id = $2
    `

	var thread models.Thread
	err := r.db.Get(&thread, query, userID, threadID)

	return &thread, err
}

func (r *ThreadRepo) SaveThread(thread *models.Thread) (*int, error) {
	insertThreadQuery := `
                INSERT INTO threads (
                    creator_id,
                    content,
                    create_date,
                    up_votes,
                    parent_thread_id
                ) VALUES (
                    :creator_id,
                    :content,
                    CURRENT_TIMESTAMP,
                    0,
                    :parent_thread_id
                ) RETURNING thread_id
    `

	var threadId int
	row, err := r.db.NamedQuery(insertThreadQuery, thread)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	if row.Next() {
		err = row.Scan(&threadId)
		if err != nil {
			return nil, err
		}
	}
	return &threadId, nil

}

func (r *ThreadRepo) SaveUpvote(threadID int, userID int, upvote bool) error {

	tx, err := r.db.Beginx()

	if err != nil {
		return err
	}

	defer tx.Rollback()

	if upvote {
		_, err = tx.Exec(`
            WITH upsert AS (
                INSERT INTO upvotes (userID, threadID)
                VALUES ($1, $2)
                ON CONFLICT (userID, threadID) DO NOTHING
                RETURNING 1
            )
            UPDATE threads
            SET up_votes = up_votes + 
                (SELECT COUNT(*) FROM upsert)
            WHERE thread_id = $2
        `, userID, threadID)
	} else {
		_, err = tx.Exec(`
            WITH deleted AS (
                DELETE FROM upvotes
                WHERE userID = $1 AND threadID = $2
                RETURNING 1
            )
            UPDATE threads
            SET up_votes = up_votes - 
                (SELECT COUNT(*) FROM deleted)
            WHERE thread_id = $2
        `, userID, threadID)
	}

	if err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}
