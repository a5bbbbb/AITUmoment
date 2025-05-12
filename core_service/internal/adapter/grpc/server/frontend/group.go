package frontend

import (
	"context"

	"github.com/a5bbbbb/AITUmoment/core_service/internal/adapter/grpc/server/frontend/dto"
	"github.com/a5bbbbb/AITUmoment/core_service/internal/adapter/logger"

	pb "github.com/a5bbbbb/AITUmoment/common/pkg/grpc/proto/core_service/gen/service/frontend/group/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GroupServiceHandler struct {
	pb.UnimplementedGroupServiceServer

	gs GroupService
}

func NewGroupServiceHandler(gs GroupService) *GroupServiceHandler {
	logger.GetLogger().Trace("Creating group service handler")
	return &GroupServiceHandler{
		gs: gs,
	}
}

func (gsh *GroupServiceHandler) GetGroup(ctx context.Context, req *pb.GetGroupRequest) (*pb.GetGroupResponse, error) {
	groupId := int(req.GroupId)

	group, err := gsh.gs.GetGroup(groupId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &pb.GetGroupResponse{
		Group: dto.ToBaseGroup(group),
	}, nil
}

func (gsh *GroupServiceHandler) GetGroups(ctx context.Context, req *pb.GetGroupsRequest) (*pb.GetGroupsResponse, error) {
	eduProgramId := int8(req.EduProgramId)

	groups, err := gsh.gs.GetGroups(uint8(eduProgramId))
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &pb.GetGroupsResponse{
		Groups: dto.ToRepeatedBaseGroup(groups),
	}, nil
}
