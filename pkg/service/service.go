package service

import (
	"segmentation-service/pkg/repository"
)

type User interface {
}

type Segment interface {
}

type Service struct {
	User
	Segment
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		User:    NewUserService(repo.User),
		Segment: NewSegmentService(repo.Segment),
	}
}
