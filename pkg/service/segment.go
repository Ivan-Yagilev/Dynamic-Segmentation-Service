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

func (s *SegmentService) CreateSegment(segment segmentation_service.SegmentWOId) (int, error) {
	return s.repo.CreateSegment(segment)
}

func (s *SegmentService) DeleteSegment(segment segmentation_service.SegmentWOId) error {
	return s.repo.DeleteSegment(segment)
}

func (s *SegmentService) UpdateSegment(id int, input segmentation_service.UpdateSegmentInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.UpdateSegment(id, input)
}
