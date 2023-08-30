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

func (s *UserSegmentService) CreateUserSegment(userSegment segmentation_service.UserSegment) error {
	return s.repo.CreateUserSegment(userSegment)
}

func (s *UserSegmentService) DeleteUserSegment(userSegment segmentation_service.UserSegment) error {
	return s.repo.DeleteUserSegment(userSegment)
}

func (s *UserSegmentService) GetAllSegments(userId int) ([]segmentation_service.Segment, error) {
	return s.repo.GetAllSegments(userId)
}
