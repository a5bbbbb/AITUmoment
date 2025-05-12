package services

import (
	"github.com/a5bbbbb/AITUmoment/api_gateway/internal/models"
)

type ThreadService struct {
	threadServicePresenter ThreadServicePresenter
}

func NewThreadService(presenter ThreadServicePresenter) *ThreadService {
	return &ThreadService{threadServicePresenter: presenter}
}

func (ts *ThreadService) GetParentThreads(userID *int) (*[]models.Thread, error) {
	return ts.threadServicePresenter.GetParentThreads(userID)
}

func (ts *ThreadService) GetThread(threadID int, userID int) (*models.Thread, error) {
	return ts.threadServicePresenter.GetThread(threadID, userID)
}

func (ts *ThreadService) GetSubThreads(parentThread int, userID int) (*[]models.Thread, *models.Thread, error) {
	return ts.threadServicePresenter.GetSubThreads(parentThread, userID)
}

func (ts *ThreadService) SaveThread(content string, creatorID int, parentThreadID *int) (*int, error) {
	return ts.threadServicePresenter.SaveThread(content, creatorID, parentThreadID)
}

func (ts *ThreadService) SaveUpvote(threadID *int, userID *int, upvote *bool) error {
	return ts.SaveUpvote(threadID, userID, upvote)
}
