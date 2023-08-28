package service

import "segmentation-service/pkg/repository"

type SegmentService struct {
	repo repository.Segment
}

func NewSegmentService(repo repository.Segment) *SegmentService {
	return &SegmentService{
		repo: repo,
	}
}