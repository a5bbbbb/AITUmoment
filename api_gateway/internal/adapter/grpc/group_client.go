package grpc

import (
	"context"

	"github.com/a5bbbbb/AITUmoment/api_gateway/internal/adapter/grpc/dto"
	"github.com/a5bbbbb/AITUmoment/api_gateway/internal/models"

	pb "github.com/a5bbbbb/AITUmoment/common/pkg/grpc/proto/core_service/gen/service/frontend/group/v2"
)

type GroupServiceClient struct {
	client pb.GroupServiceClient
}

func NewGroupServiceClient(client pb.GroupServiceClient) *GroupServiceClient {
	return &GroupServiceClient{
		client: client,
	}
}

func (gsc *GroupServiceClient) GetGroup(groupId int) (*models.Group, error) {
	resp, err := gsc.client.GetGroup(context.TODO(), &pb.GetGroupRequest{
		GroupId: int64(groupId),
	})

	if err != nil {
		return &models.Group{}, err
	}

	return dto.FromBaseGroup(resp.Group), nil
}
func (gsc *GroupServiceClient) GetGroups(eduProgram uint8) (*[]models.Group, error) {
	resp, err := gsc.client.GetGroups(context.TODO(), &pb.GetGroupsRequest{
		EduProgramId: uint32(eduProgram),
	})

	if err != nil {
		return &[]models.Group{}, err
	}

	return dto.FromRepeatedBaseGroup(resp.Groups), nil
}
