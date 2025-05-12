package repository

import (
	"github.com/a5bbbbb/AITUmoment/core_service/internal/adapter/logger"
	"github.com/a5bbbbb/AITUmoment/core_service/internal/models"
	db "github.com/a5bbbbb/AITUmoment/core_service/pkg/sqlx"

	"github.com/jmoiron/sqlx"
)

type EduRepo struct {
	db *sqlx.DB
}

func NewEduProgramRepo() *EduRepo {
	logger.GetLogger().Trace("Creating edu repo")
	return &EduRepo{db: db.GetDB()}
}

func (r *EduRepo) GetEduPrograms() (*[]models.EduProgram, error) {
	query := " SELECT * FROM educational_programs "
	var programs []models.EduProgram
	err := r.db.Select(&programs, query)
	return &programs, err
}
