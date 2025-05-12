package dto

import (
	"github.com/a5bbbbb/AITUmoment/api_gateway/internal/models"

	base "github.com/a5bbbbb/AITUmoment/common/pkg/grpc/proto/core_service/gen/base/frontend/v2"
)

func FromRepeatedBaseEduProgram(eduPrograms []*base.EduProgram) *[]models.EduProgram {
	res := make([]models.EduProgram, len(eduPrograms))

	for i, program := range eduPrograms {
		res[i] = models.EduProgram{
			Id:   int(program.Id),
			Name: program.Name,
		}
	}

	return &res
}
