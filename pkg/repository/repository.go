package repository

import (
	segmentation_service "segmentation-service"

	"github.com/jmoiron/sqlx"
)

type User interface {
	CreateUser(user segmentation_service.User) (int, error)
	DeleteUser(userId int) error
}

type Segment interface {
}

type Repository struct {
	User
	Segment
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User:    NewUserPostgres(db),
		Segment: NewSegmentPostgres(db),
	}
}
