package frontend

import (
	"context"

	"github.com/a5bbbbb/AITUmoment/core_service/internal/adapter/grpc/server/frontend/dto"
	"github.com/a5bbbbb/AITUmoment/core_service/internal/adapter/logger"

	pb "github.com/a5bbbbb/AITUmoment/common/pkg/grpc/proto/core_service/gen/service/frontend/thread/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ThreadServiceHandler struct {
	pb.UnimplementedThreadServiceServer

	ts ThreadService
}

func NewThreadServiceHandler(ts ThreadService) *ThreadServiceHandler {
	logger.GetLogger().Trace("Creating thread service handler")
	return &ThreadServiceHandler{
		ts: ts,
	}
}

func (tsh *ThreadServiceHandler) GetParentThreads(ctx context.Context, req *pb.GetParentThreadsRequest) (*pb.GetParentThreadsResponse, error) {
	userId := int(req.UserId)

	threads, err := tsh.ts.GetParentThreads(&userId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &pb.GetParentThreadsResponse{
		Threads: dto.ToRepeatedBaseThread(threads),
	}, nil
}

func (tsh *ThreadServiceHandler) GetSubThreads(ctx context.Context, req *pb.GetSubThreadsRequest) (*pb.GetSubThreadsResponse, error) {
	parentTreadId := int(req.ParentThreadId)
	userId := int(req.UserId)

	subThreads, parentThread, err := tsh.ts.GetSubThreads(parentTreadId, userId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &pb.GetSubThreadsResponse{
		SubThreads:   dto.ToRepeatedBaseThread(subThreads),
		ParentThread: dto.ToBaseThread(parentThread),
	}, nil
}

func (tsh *ThreadServiceHandler) GetThread(ctx context.Context, req *pb.GetThreadRequest) (*pb.GetThreadResponse, error) {
	threadId := int(req.ThreadId)
	userId := int(req.UserId)

	thread, err := tsh.ts.GetThread(threadId, userId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &pb.GetThreadResponse{
		Thread: dto.ToBaseThread(thread),
	}, nil
}

func (tsh *ThreadServiceHandler) SaveThread(ctx context.Context, req *pb.SaveThreadRequest) (*pb.SaveThreadResponse, error) {
	content := req.Content
	creatorId := int(req.CreatorId)
	parentThreadId := new(int)

	if req.ParentThreadId == 0 {
		parentThreadId = nil
	} else {
		*parentThreadId = int(req.ParentThreadId)
	}

	threadId, err := tsh.ts.SaveThread(content, creatorId, parentThreadId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &pb.SaveThreadResponse{
		ThreadId: int64(*threadId),
	}, nil
}

func (tsh *ThreadServiceHandler) SaveUpvote(ctx context.Context, req *pb.SaveUpvoteRequest) (*pb.SaveUpvoteResponse, error) {
	threadId := int(req.ThreadId)
	userId := int(req.UserId)
	upvote := req.Upvote

	err := tsh.ts.SaveUpvote(&threadId, &userId, &upvote)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &pb.SaveUpvoteResponse{}, nil
}
