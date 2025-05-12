package frontend

import "github.com/a5bbbbb/AITUmoment/core_service/internal/models"

type UserService interface {
	Authorize(email, passwd string) (*models.User, error)
	UpdateUser(user *models.User) (*models.User, error)
	CreateUser(user *models.User) (*models.User, error)
	GetFullUserInfo(userId int) (*models.User, *[]models.EduProgram, *models.Group, error)
	VerifyEmail(token string) (*models.User, error)
}

type ThreadService interface {
	GetParentThreads(userID *int) (*[]models.Thread, error)
	GetSubThreads(parentThread int, userID int) (*[]models.Thread, *models.Thread, error)
	GetThread(threadID int, userID int) (*models.Thread, error)
	SaveThread(content string, creatorID int, parentThreadID *int) (*int, error)
	SaveUpvote(threadID *int, userID *int, upvote *bool) error
}

type GroupService interface {
	GetGroup(groupID int) (*models.Group, error)
	GetGroups(eduProg uint8) (*[]models.Group, error)
}

type EduProgramService interface {
	GetPrograms() (*[]models.EduProgram, error)
}
