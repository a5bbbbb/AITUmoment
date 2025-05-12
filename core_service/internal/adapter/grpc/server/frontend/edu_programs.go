package frontend

import (
	"context"

	"github.com/a5bbbbb/AITUmoment/core_service/internal/adapter/grpc/server/frontend/dto"
	"github.com/a5bbbbb/AITUmoment/core_service/internal/adapter/logger"

	pb "github.com/a5bbbbb/AITUmoment/common/pkg/grpc/proto/core_service/gen/service/frontend/edu_program/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type EduProgramsServiceHandler struct {
	pb.UnimplementedEduProgramServiceServer

	eps EduProgramService
}

func NewEduProgramServiceHandler(eps EduProgramService) *EduProgramsServiceHandler {
	logger.GetLogger().Trace("Creating edu programs service handler")
	return &EduProgramsServiceHandler{
		eps: eps,
	}
}

func (epsh *EduProgramsServiceHandler) GetEduPrograms(ctx context.Context, req *pb.GetEduProgramsRequest) (*pb.GetEduProgramsResponse, error) {
	eduPrograms, err := epsh.eps.GetPrograms()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.GetEduProgramsResponse{
		EduPrograms: dto.ToRepeatedBaseEduProgram(eduPrograms),
	}, nil
}
