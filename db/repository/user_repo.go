package repository

import (
	"aitu-moment/models"
	"github.com/jmoiron/sqlx"
)

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
