package services

import (
	"context"

	"github.com/a5bbbbb/AITUmoment/core_service/internal/models"
)

type EduProgramRepository interface {
	GetEduPrograms() (*[]models.EduProgram, error)
}

type GroupRepository interface {
	GetGroupByID(groupID int) (*models.Group, error)
	GetGroups(eduProg uint8) (*[]models.Group, error)
}

type ThreadRepository interface {
	GetParentThreads(userID *int) (*[]models.Thread, error)
	GetSubThreads(userID int, parentThread int) (*[]models.Thread, error)
	GetThread(threadID int, userID int) (*models.Thread, error)
	SaveThread(thread *models.Thread) (*int, error)
	SaveUpvote(threadID int, userID int, upvote bool) error
}

type UserRepository interface {
	CreateUser(user *models.User) (int64, error)
	GetAllUsers() ([]models.User, error)
	GetUser(userId int64) (*models.User, error)
	GetUserIdByCred(email string, passwd string) (*int, error)
	UpdateUser(user *models.User) (int64, error)
	UpdateUserVerified(email string, verified bool) (*int, error)
}

type EmailVerificationProducer interface {
	Push(context.Context, models.EmailVerification) error
}

type EmailVerificationManager interface {
	Generate(models.User) (*models.EmailVerification, error)
	GetEmail(token string) (string, error)
}

type RedisCache interface {
	Get(ctx context.Context, userId int) (*models.User, error)
	Set(ctx context.Context, user *models.User) error
	SetMany(ctx context.Context, users []models.User) error
	Delete(ctx context.Context, userId int) error
}
