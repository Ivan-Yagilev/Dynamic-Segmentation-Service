package service

import (
	segmentation_service "segmentation-service"
	"segmentation-service/pkg/repository"
)

type SegmentService struct {
	repo repository.Segment
}

func NewSegmentService(repo repository.Segment) *SegmentService {
	return &SegmentService{
		repo: repo,
	}
}

func (s *SegmentService) GetAllSegments() ([]segmentation_service.Segment, error) {
	return s.repo.GetAllSegments()
}

func (s *SegmentService) CreateSegment(segment segmentation_service.Segment) (int, error) {
	return s.repo.CreateSegment(segment)
}

func (s *SegmentService) DeleteSegment(segment segmentation_service.Segment) error {
	return s.repo.DeleteSegment(segment)
}