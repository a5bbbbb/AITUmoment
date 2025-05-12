package services

import (
	"github.com/a5bbbbb/AITUmoment/api_gateway/internal/models"
)

type EduService struct {
	eduServicePresenter EduServicePresenter
}

func NewEduProgramService(presenter EduServicePresenter) *EduService {
	return &EduService{eduServicePresenter: presenter}
}

func (s *EduService) GetPrograms() (*[]models.EduProgram, error) {
	return s.eduServicePresenter.GetPrograms()
}
