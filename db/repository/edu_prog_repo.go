package repository

import (
	"aitu-moment/db"
	"aitu-moment/models"

	"github.com/jmoiron/sqlx"
)



type EduRepo struct{
    db *sqlx.DB
}

func NewEduRepo() *EduRepo{
    return &EduRepo{db: db.GetDB()}
}

func (r *EduRepo) GetEduPrograms() ([]models.EduProgram, error){
    query := " SELECT * FROM educational_programs "
	var programs []models.EduProgram
	err := r.db.Select(&programs, query)
	return programs, err
}
