package services

import (
	"github.com/a5bbbbb/AITUmoment/api_gateway/internal/models"
)

type UserService struct {
	userServicePresenter UserServicePresenter
}

func NewUserService(presenter UserServicePresenter) *UserService {
	return &UserService{userServicePresenter: presenter}
}

func (us *UserService) Authorize(email string, passwd string) (*models.User, error) {
	return us.userServicePresenter.Authorize(email, passwd)
}

func (us *UserService) UpdateUser(user *models.User) (*models.User, error) {
	return us.userServicePresenter.UpdateUser(user)
}

func (us *UserService) CreateUser(user *models.User) (*models.User, error) {
	return us.userServicePresenter.CreateUser(user)
}

func (us *UserService) GetFullUserInfo(userID int) (*models.User, *[]models.EduProgram, *models.Group, error) {
	return us.userServicePresenter.GetFullUserInfo(userID)
}

func (us *UserService) VerifyEmail(token string) (*models.User, error) {
	return us.userServicePresenter.VerifyEmail(token)
}
