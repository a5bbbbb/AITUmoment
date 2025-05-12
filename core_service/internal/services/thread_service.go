package services

import (
	"github.com/a5bbbbb/AITUmoment/core_service/internal/adapter/logger"
	"github.com/a5bbbbb/AITUmoment/core_service/internal/models"
)

type ThreadService struct {
	threadRepo ThreadRepository
}

func NewThreadService(threadRepo ThreadRepository) *ThreadService {
	logger.GetLogger().Trace("Creating thread service")
	return &ThreadService{threadRepo: threadRepo}
}

func (s *ThreadService) GetParentThreads(userID *int) (*[]models.Thread, error) {
	threads, err := s.threadRepo.GetParentThreads(userID)
	return threads, err
}

// Return the thread with threadID and return whether user with userId has an upvote on it inside the Thread.
func (s *ThreadService) GetThread(threadID int, userID int) (*models.Thread, error) {
	thread, err := s.threadRepo.GetThread(threadID, userID)

	if err != nil {
		return nil, err
	}

	return thread, nil

}

// Return sub threads of parentThread and to every sub threads return whether user with userId has an upvote on it inside the Thread.
func (s *ThreadService) GetSubThreads(parentThread int, userID int) (*[]models.Thread, *models.Thread, error) {
	subThreads, err := s.threadRepo.GetSubThreads(userID, parentThread)

	if err != nil {
		return nil, nil, err
	}

	parent, err := s.threadRepo.GetThread(parentThread, userID)

	if err != nil {
		return nil, nil, err
	}

	return subThreads, parent, nil

}

func (s *ThreadService) SaveThread(content string, creatorID int, parentThreadID *int) (*int, error) {
	thread := models.Thread{Content: content, CreatorID: creatorID, ParentThread: parentThreadID}

	newThreadID, err := s.threadRepo.SaveThread(&thread)

	return newThreadID, err

}

func (s *ThreadService) SaveUpvote(threadID *int, userID *int, upvote *bool) error {
	err := s.threadRepo.SaveUpvote(*threadID, *userID, *upvote)

	return err
}
