package services

import (
	"github.com/a5bbbbb/AITUmoment/core_service/internal/adapter/logger"
	"github.com/a5bbbbb/AITUmoment/core_service/internal/models"
)

type EduService struct {
	eduProgramRepo EduProgramRepository
}

func NewEduService(eduProgramRepo EduProgramRepository) *EduService {
	logger.GetLogger().Trace("Creating edu service")
	return &EduService{eduProgramRepo: eduProgramRepo}
}

func (s *EduService) GetPrograms() (*[]models.EduProgram, error) {
	programs, err := s.eduProgramRepo.GetEduPrograms()
	return programs, err
}
