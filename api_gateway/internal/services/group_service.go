package services

import (
	"github.com/a5bbbbb/AITUmoment/api_gateway/internal/models"
)

type GroupService struct {
	groupServicePresenter GroupServicePresenter
}

func NewGroupService(presenter GroupServicePresenter) *GroupService {
	return &GroupService{groupServicePresenter: presenter}
}

func (s *GroupService) GetGroups(eduProg uint8) (*[]models.Group, error) {
	return s.groupServicePresenter.GetGroups(eduProg)
}

func (s *GroupService) GetGroup(groupID int) (*models.Group, error) {
	return s.groupServicePresenter.GetGroup(groupID)
}
