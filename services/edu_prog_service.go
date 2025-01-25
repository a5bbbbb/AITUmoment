package services

import (
	"aitu-moment/db/repository"
	"aitu-moment/models"
)


type EduService struct{
    repo *repository.EduRepo
}


func NewEduService() *EduService{
    return &EduService{repo: repository.NewEduRepo()}
}

func (s *EduService) GetPrograms()([]models.EduProgram, error){
    programs,err := s.repo.GetEduPrograms()
    return programs,err
}


