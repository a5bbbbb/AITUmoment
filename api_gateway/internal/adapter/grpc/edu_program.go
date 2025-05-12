package grpc

import (
	"context"

	"github.com/a5bbbbb/AITUmoment/api_gateway/internal/adapter/grpc/dto"
	"github.com/a5bbbbb/AITUmoment/api_gateway/internal/models"

	pb "github.com/a5bbbbb/AITUmoment/common/pkg/grpc/proto/core_service/gen/service/frontend/edu_program/v2"
)

type EduProgramServiceClient struct {
	client pb.EduProgramServiceClient
}

func NewEduProgramServiceClient(client pb.EduProgramServiceClient) *EduProgramServiceClient {
	return &EduProgramServiceClient{
		client: client,
	}
}

func (epsc *EduProgramServiceClient) GetPrograms() (*[]models.EduProgram, error) {
	resp, err := epsc.client.GetEduPrograms(context.TODO(), &pb.GetEduProgramsRequest{})

	if err != nil {
		return &[]models.EduProgram{}, err
	}

	return dto.FromRepeatedBaseEduProgram(resp.EduPrograms), nil
}
