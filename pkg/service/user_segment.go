package service

import (
	segmentation_service "segmentation-service"
	"segmentation-service/pkg/repository"
)

type UserSegmentService struct {
	repo repository.UserSegment
}

func NewUserSegmentService(repo repository.UserSegment) *UserSegmentService {
	return &UserSegmentService{
		repo: repo,
	}
}

func (s *UserSegmentService) CreateUserSegment(userSegment segmentation_service.UserSegment) (int, error) {
	return s.repo.CreateUserSegment(userSegment)
}