package dto

import (
	"github.com/a5bbbbb/AITUmoment/api_gateway/internal/models"

	base "github.com/a5bbbbb/AITUmoment/common/pkg/grpc/proto/core_service/gen/base/frontend/v2"
)

func FromBaseUser(user *base.User) *models.User {
	return &models.User{
		Id:                 int(user.Id),
		Name:               user.Name,
		EducationalProgram: uint8(user.EducationalProgram),
		ProgramName:        user.ProgramName,
		PublicName:         user.PublicName,
		Email:              user.Email,
		Passwd:             user.Passwd,
		Bio:                user.Bio,
		Group:              uint8(user.Group),
	}
}

func ToBaseUser(user *models.User) *base.User {
	return &base.User{
		Id:                 int64(user.Id),
		Name:               user.Name,
		EducationalProgram: uint32(user.EducationalProgram),
		ProgramName:        user.ProgramName,
		PublicName:         user.PublicName,
		Email:              user.Email,
		Passwd:             user.Passwd,
		Bio:                user.Bio,
		Group:              uint32(user.Group),
	}
}
