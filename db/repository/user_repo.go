package repository

import (
	"aitu-moment/db"
	"aitu-moment/models"
	"github.com/jmoiron/sqlx"
)


type UserRepository struct {
    db *sqlx.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{db: db.GetDB()}
}

func (r *UserRepository) GetAllUsers() ([]models.User, error) {
	query := " SELECT u.*, ep.name as program_name FROM users u JOIN educational_programs ep ON u.educational_program = ep.id"
	var users []models.User
	err := r.db.Select(&users, query)
	return users, err
}

func (r *UserRepository) GetUser(userId int64) (*models.User, error) {
	var user models.User
	query := ` 
        SELECT u.*, ep.name as program_name 
        FROM users u 
        JOIN educational_programs ep ON u.educational_program = ep.id 
        WHERE u.id = $1`

	err := r.db.Get(&user, query, userId)
	return &user, err
}

func (r *UserRepository) CreateUser(user *models.User) (int64, error) {
	var insertedID int
	err := r.db.QueryRow(
		"INSERT INTO users (username, educational_program, email, password, public_name, bio, group_id) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id",
		user.Name, user.EducationalProgram, user.Email, user.Passwd, user.PublicName, "", user.Group,
	).Scan(&insertedID)
	return int64(insertedID), err

}

func (r *UserRepository) UpdateUser(user *models.User) (int64,error){
    query := 
    `
    UPDATE users
        SET username = :username,
        educational_program = :educational_program,
        public_name = :public_name,
        email = :email,
        password = :password,
        bio = :bio,
        group_id = :group_id
    WHERE id = :id
    `
    _,err := r.db.NamedExec(query, user)

    return int64(user.Id),err

}




func (r *UserRepository) VerifyUser(email string) (error){

    query :=
    `
        UPDATE users
            SET verified = TRUE
        WHERE email = $1

    `
    _,err := r.db.Exec(query, email)

    return err

}

func (r *UserRepository) GetUserIdByCred(email string, passwd string ) (*int, error){
    var userID int
    query := `  SELECT id 
                FROM users WHERE email = $1 AND password = $2 AND verified = TRUE
             `
    err := r.db.Get(&userID, query, email, passwd)

    if err != nil {
        return nil,err
    }
    return &userID,err
}

