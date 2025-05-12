package dto

import (
	"github.com/a5bbbbb/AITUmoment/api_gateway/internal/models"

	base "github.com/a5bbbbb/AITUmoment/common/pkg/grpc/proto/core_service/gen/base/frontend/v2"
)

func FromRepeatedBaseThread(threads []*base.Thread) *[]models.Thread {
	res := make([]models.Thread, len(threads))

	for i, thread := range threads {
		res[i] = *FromBaseThread(thread)
	}

	return &res
}

func FromBaseThread(thread *base.Thread) *models.Thread {
	res := &models.Thread{
		Id:           int(thread.Id),
		CreatorID:    int(thread.CreatorId),
		CreatorName:  thread.CreatorName,
		Content:      thread.Content,
		CreateDate:   thread.CreateDate.AsTime(),
		UpVotes:      int(thread.UpVotes),
		UserUpvoted:  thread.UserUpvoted,
		ParentThread: nil,
	}

	if thread.ParentThread != 0 {
		parentThread := int(thread.ParentThread)
		res.ParentThread = &parentThread
	}
	return res
}
