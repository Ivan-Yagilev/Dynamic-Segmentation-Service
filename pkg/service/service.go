package service

import (
	segmentation_service "segmentation-service"
	"segmentation-service/pkg/repository"
)

type User interface {
	CreateUser(user segmentation_service.User) (int, error)
	DeleteUser(userId int) error
}

type Segment interface {
	GetAllSegments() ([]segmentation_service.Segment, error)
	CreateSegment(segment segmentation_service.Segment) (int, error)
	DeleteSegment(segment segmentation_service.Segment) error
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
