package grpc

import (
	"context"

	"github.com/a5bbbbb/AITUmoment/api_gateway/internal/adapter/grpc/dto"
	"github.com/a5bbbbb/AITUmoment/api_gateway/internal/models"

	pb "github.com/a5bbbbb/AITUmoment/common/pkg/grpc/proto/core_service/gen/service/frontend/thread/v2"
)

type ThreadServiceClient struct {
	client pb.ThreadServiceClient
}

func NewThreadServiceClient(client pb.ThreadServiceClient) *ThreadServiceClient {
	return &ThreadServiceClient{
		client: client,
	}
}

func (tsc *ThreadServiceClient) GetParentThreads(userId *int) (*[]models.Thread, error) {
	resp, err := tsc.client.GetParentThreads(context.TODO(), &pb.GetParentThreadsRequest{
		UserId: int64(*userId),
	})

	if err != nil {
		return &[]models.Thread{}, err
	}

	return dto.FromRepeatedBaseThread(resp.Threads), nil
}

func (tsc *ThreadServiceClient) GetSubThreads(parentThread int, userId int) (*[]models.Thread, *models.Thread, error) {
	resp, err := tsc.client.GetSubThreads(context.TODO(), &pb.GetSubThreadsRequest{
		ParentThreadId: int64(parentThread),
		UserId:         int64(userId),
	})

	if err != nil {
		return &[]models.Thread{}, &models.Thread{}, err
	}

	return dto.FromRepeatedBaseThread(resp.SubThreads), dto.FromBaseThread(resp.ParentThread), nil
}

func (tsc *ThreadServiceClient) GetThread(threadId int, userId int) (*models.Thread, error) {
	resp, err := tsc.client.GetThread(context.TODO(), &pb.GetThreadRequest{
		ThreadId: int64(threadId),
		UserId:   int64(userId),
	})

	if err != nil {
		return &models.Thread{}, err
	}

	return dto.FromBaseThread(resp.Thread), nil
}

func (tsc *ThreadServiceClient) SaveThread(content string, creatorId int, parentThreadId *int) (*int, error) {
	req := &pb.SaveThreadRequest{
		Content:   content,
		CreatorId: int64(creatorId),
	}

	if parentThreadId != nil {
		req.ParentThreadId = int64(*parentThreadId)
	}

	resp, err := tsc.client.SaveThread(context.TODO(), req)

	if err != nil {
		return nil, err
	}

	threadId := int(resp.ThreadId)

	return &threadId, nil
}

func (tsc *ThreadServiceClient) SaveUpvote(threadId *int, userId *int, upvote *bool) error {
	_, err := tsc.client.SaveUpvote(context.TODO(), &pb.SaveUpvoteRequest{
		ThreadId: int64(*threadId),
		UserId:   int64(*userId),
		Upvote:   *upvote,
	})
	return err
}
