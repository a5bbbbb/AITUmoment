package dto

import (
	"github.com/a5bbbbb/AITUmoment/api_gateway/internal/models"

	base "github.com/a5bbbbb/AITUmoment/common/pkg/grpc/proto/core_service/gen/base/frontend/v2"
)

func FromBaseGroup(group *base.Group) *models.Group {
	return &models.Group{
		Id:                 int(group.Id),
		EducationalProgram: uint8(group.EducationalProgram),
		Year:               int16(group.Year),
		Number:             uint8(group.Number),
		EduProgName:        group.EduProgName,
		GroupName:          group.GroupName,
	}
}

func FromRepeatedBaseGroup(group []*base.Group) *[]models.Group {
	res := make([]models.Group, len(group))

	for i, group := range group {
		res[i] = *FromBaseGroup(group)
	}

	return &res
}
