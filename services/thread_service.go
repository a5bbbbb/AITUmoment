package services

import (
	"aitu-moment/db/repository"
	"aitu-moment/models"
)

type ThreadService struct {
	repo *repository.ThreadRepo
}

func NewThreadRepo() *ThreadService {
	return &ThreadService{repo: repository.NewThreadRepo()}
}

func (s *ThreadService) GetParentThreads(userID *int) (*[]models.Thread, error) {
	threads, err := s.repo.GetParentThreads(userID)
	return threads, err
}

func (s *ThreadService) GetThread(threadID int, userID int) (*models.Thread, error) {
	thread, err := s.repo.GetThread(threadID, userID)

	if err != nil {
		return nil, err
	}

	return thread, nil

}

func (s *ThreadService) GetSubThreads(parentThread int, userID int) (*[]models.Thread, *models.Thread, error) {
	subThreads, err := s.repo.GetSubThreads(userID, parentThread)

	if err != nil {
		return nil, nil, err
	}

	parent, err := s.repo.GetThread(parentThread, userID)

	if err != nil {
		return nil, nil, err
	}

	return subThreads, parent, nil

}

func (s *ThreadService) SaveThread(content string, creatorID int, parentThreadID *int) (*int, error) {
	thread := models.Thread{Content: content, CreatorID: creatorID, ParentThread: parentThreadID}

	newThreadID, err := s.repo.SaveThread(&thread)

	return newThreadID, err

}

func (s *ThreadService) SaveUpvote(threadID *int, userID *int, upvote *bool) error {
	err := s.repo.SaveUpvote(*threadID, *userID, *upvote)

	return err
}
