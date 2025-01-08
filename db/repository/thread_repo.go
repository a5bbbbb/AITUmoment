package repository

import (
	"aitu-moment/models"

	"github.com/jmoiron/sqlx"
)

type ThreadRepository struct {
	DB *sqlx.DB
}

func NewThreadRepository(db *sqlx.DB) *ThreadRepository {
	return &ThreadRepository{DB: db}
}

func (r *ThreadRepository) GetAllThreads() ([]models.Thread, error) {
	query := "SELECT * FROM threads"
	var threads []models.Thread
	err := r.DB.Select(&threads, query)
	return threads, err
}

func (r *ThreadRepository) CreateThread(thread models.Thread) (int64, error) {
	var insertedID int
	err := r.DB.QueryRow(
		"INSERT INTO threads (creator_id, \"content\", up_votes, parent_thread_id) VALUES ($1, $2, 0, $3) RETURNING thread_id",
		thread.CreatorId, thread.Content, thread.ParentThreadId,
	).Scan(&insertedID)
	return int64(insertedID), err
}
