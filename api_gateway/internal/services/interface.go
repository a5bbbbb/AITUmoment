package services

import "github.com/a5bbbbb/AITUmoment/api_gateway/internal/models"

type EduServicePresenter interface {
	GetPrograms() (*[]models.EduProgram, error)
}

type GroupServicePresenter interface {
	GetGroup(groupID int) (*models.Group, error)
	GetGroups(eduProg uint8) (*[]models.Group, error)
}

type ThreadServicePresenter interface {
	GetParentThreads(userID *int) (*[]models.Thread, error)
	GetSubThreads(parentThread int, userID int) (*[]models.Thread, *models.Thread, error)
	GetThread(threadID int, userID int) (*models.Thread, error)
	SaveThread(content string, creatorID int, parentThreadID *int) (*int, error)
	SaveUpvote(threadID *int, userID *int, upvote *bool) error
}

type UserServicePresenter interface {
	Authorize(email string, passwd string) (*models.User, error)
	CreateUser(user *models.User) (*models.User, error)
	GetFullUserInfo(userID int) (*models.User, *[]models.EduProgram, *models.Group, error)
	UpdateUser(user *models.User) (*models.User, error)
	VerifyEmail(token string) (*models.User, error)
}
