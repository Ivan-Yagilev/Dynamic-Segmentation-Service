package repository

import (
	segmentation_service "segmentation-service"

	"github.com/jmoiron/sqlx"
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
}

type Repository struct {
	User
	Segment
	UserSegment
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User:    NewUserPostgres(db),
		Segment: NewSegmentPostgres(db),
		UserSegment: NewUserSegmentPostgres(db),
	}
}
