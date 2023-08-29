package service

import (
	segmentation_service "segmentation-service"
	"segmentation-service/pkg/repository"
)

type User interface {
	GetAllUsers() ([]segmentation_service.UserResponse, error)
	CreateUser(user segmentation_service.User) (int, error)
	DeleteUser(userId int) error
}

type Segment interface {
	GetAllSegments() ([]segmentation_service.Segment, error)
	CreateSegment(segment segmentation_service.Segment) (int, error)
	UpdateSegment(id int, input segmentation_service.UpdateSegmentInput) error
	DeleteSegment(segment segmentation_service.Segment) error
}

type UserSegment interface {
	CreateUserSegment(userSegment segmentation_service.UserSegment) error
	DeleteUserSegment(userSegment segmentation_service.UserSegment) error
	GetAllSegments(userId int) ([]segmentation_service.Segment, error)
}

type Service struct {
	User
	Segment
	UserSegment
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		User:    NewUserService(repo.User),
		Segment: NewSegmentService(repo.Segment),
		UserSegment: NewUserSegmentService(repo.UserSegment),
	}
}
