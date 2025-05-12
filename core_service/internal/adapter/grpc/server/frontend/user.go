package frontend

import (
	"context"

	"github.com/a5bbbbb/AITUmoment/core_service/internal/adapter/grpc/server/frontend/dto"
	"github.com/a5bbbbb/AITUmoment/core_service/internal/adapter/logger"

	pb "github.com/a5bbbbb/AITUmoment/common/pkg/grpc/proto/core_service/gen/service/frontend/user/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserServiceHandler struct {
	pb.UnimplementedUserServiceServer

	us UserService
}

func NewUserServiceHandler(us UserService) *UserServiceHandler {
	logger.GetLogger().Trace("Creating user service handler")
	return &UserServiceHandler{
		us: us,
	}
}

func (ush *UserServiceHandler) Authorize(ctx context.Context, req *pb.AuthorizeRequest) (*pb.AuthorizeResponse, error) {
	email := req.Email
	passwd := req.Passwd

	user, err := ush.us.Authorize(email, passwd)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &pb.AuthorizeResponse{
		User: dto.ToBaseUser(user),
	}, nil
}

func (ush *UserServiceHandler) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	user := dto.FromBaseUser(req.User)

	user, err := ush.us.CreateUser(user)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &pb.CreateUserResponse{
		User: dto.ToBaseUser(user),
	}, nil
}

func (ush *UserServiceHandler) GetFullUserInfo(ctx context.Context, req *pb.GetFullUserInfoRequest) (*pb.GetFullUserInfoResponse, error) {
	userID := req.UserId

	user, programs, group, err := ush.us.GetFullUserInfo(int(userID))
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &pb.GetFullUserInfoResponse{
		User:        dto.ToBaseUser(user),
		EduPrograms: dto.ToRepeatedBaseEduProgram(programs),
		Group:       dto.ToBaseGroup(group),
	}, nil
}

func (ush *UserServiceHandler) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	user := dto.FromBaseUser(req.User)

	user, err := ush.us.UpdateUser(user)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &pb.UpdateUserResponse{
		User: dto.ToBaseUser(user),
	}, nil
}

func (ush *UserServiceHandler) VerifyEmail(ctx context.Context, req *pb.VerifyEmailRequest) (*pb.VerifyEmailResponse, error) {
	token := req.Token

	user, err := ush.us.VerifyEmail(token)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &pb.VerifyEmailResponse{
		User: dto.ToBaseUser(user),
	}, nil
}
