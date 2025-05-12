package grpc

import (
	"context"

	"github.com/a5bbbbb/AITUmoment/api_gateway/internal/adapter/grpc/dto"
	"github.com/a5bbbbb/AITUmoment/api_gateway/internal/models"

	pb "github.com/a5bbbbb/AITUmoment/common/pkg/grpc/proto/core_service/gen/service/frontend/user/v2"
)

type AuthServiceClient struct {
	client pb.UserServiceClient
}

func NewAuthServiceClient(client pb.UserServiceClient) *AuthServiceClient {
	return &AuthServiceClient{
		client: client,
	}
}

func (asc *AuthServiceClient) Authorize(email string, passwd string) (*models.User, error) {
	resp, err := asc.client.Authorize(context.TODO(), &pb.AuthorizeRequest{
		Email:  email,
		Passwd: passwd,
	})

	if err != nil {
		return &models.User{}, err
	}

	return dto.FromBaseUser(resp.User), nil
}

func (asc *AuthServiceClient) CreateUser(user *models.User) (*models.User, error) {
	resp, err := asc.client.CreateUser(context.TODO(), &pb.CreateUserRequest{
		User: dto.ToBaseUser(user),
	})

	if err != nil {
		return &models.User{}, err
	}

	return dto.FromBaseUser(resp.User), nil
}

func (asc *AuthServiceClient) GetFullUserInfo(userID int) (*models.User, *[]models.EduProgram, *models.Group, error) {
	resp, err := asc.client.GetFullUserInfo(context.TODO(), &pb.GetFullUserInfoRequest{
		UserId: int64(userID),
	})

	if err != nil {
		return &models.User{}, &[]models.EduProgram{}, &models.Group{}, err
	}

	return dto.FromBaseUser(resp.User), dto.FromRepeatedBaseEduProgram(resp.EduPrograms), dto.FromBaseGroup(resp.Group), nil
}

func (asc *AuthServiceClient) UpdateUser(user *models.User) (*models.User, error) {
	resp, err := asc.client.UpdateUser(context.TODO(), &pb.UpdateUserRequest{
		User: dto.ToBaseUser(user),
	})

	if err != nil {
		return &models.User{}, err
	}

	return dto.FromBaseUser(resp.User), nil
}

func (asc *AuthServiceClient) VerifyEmail(token string) (*models.User, error) {
	resp, err := asc.client.VerifyEmail(context.TODO(), &pb.VerifyEmailRequest{
		Token: token,
	})

	if err != nil {
		return &models.User{}, err
	}

	return dto.FromBaseUser(resp.User), nil
}
