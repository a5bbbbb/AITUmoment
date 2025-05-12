package dto

import (
	"github.com/a5bbbbb/AITUmoment/core_service/internal/models"

	base "github.com/a5bbbbb/AITUmoment/common/pkg/grpc/proto/core_service/gen/base/frontend/v2"
)

func ToBaseGroup(group *models.Group) *base.Group {
	return &base.Group{
		Id:                 int64(group.Id),
		EducationalProgram: uint32(group.EducationalProgram),
		Year:               int32(group.Year),
		Number:             uint32(group.Number),
		EduProgName:        group.EduProgName,
		GroupName:          group.GroupName,
	}
}

func ToRepeatedBaseGroup(group *[]models.Group) []*base.Group {
	res := make([]*base.Group, len(*group))

	for i, group := range *group {
		res[i] = ToBaseGroup(&group)
	}

	return res
}
