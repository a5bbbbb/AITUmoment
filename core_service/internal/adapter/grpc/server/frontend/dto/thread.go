package dto

import (
	"github.com/a5bbbbb/AITUmoment/core_service/internal/models"

	base "github.com/a5bbbbb/AITUmoment/common/pkg/grpc/proto/core_service/gen/base/frontend/v2"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToRepeatedBaseThread(threads *[]models.Thread) []*base.Thread {
	res := make([]*base.Thread, len(*threads))

	for i, thread := range *threads {
		res[i] = ToBaseThread(&thread)
	}

	return res
}

func ToBaseThread(thread *models.Thread) *base.Thread {
	res := &base.Thread{
		Id:          int64(thread.Id),
		CreatorId:   int64(thread.CreatorID),
		CreatorName: thread.CreatorName,
		Content:     thread.Content,
		CreateDate:  timestamppb.New(thread.CreateDate),
		UpVotes:     int64(thread.UpVotes),
		UserUpvoted: thread.UserUpvoted,
	}
	if thread.ParentThread != nil {
		res.ParentThread = int64(*thread.ParentThread)
	}
	return res
}
