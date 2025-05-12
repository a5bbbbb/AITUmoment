package dto

import (
	"github.com/a5bbbbb/AITUmoment/core_service/internal/models"

	base "github.com/a5bbbbb/AITUmoment/common/pkg/grpc/proto/core_service/gen/base/frontend/v2"
)

func ToRepeatedBaseEduProgram(programs *[]models.EduProgram) []*base.EduProgram {
	res := make([]*base.EduProgram, len(*programs))

	for i, program := range *programs {
		res[i] = &base.EduProgram{
			Id:   int64(program.Id),
			Name: program.Name,
		}
	}

	return res
}
