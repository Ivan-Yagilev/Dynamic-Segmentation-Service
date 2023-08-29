package repository

import (
	"fmt"
	segmentation_service "segmentation-service"

	"github.com/jmoiron/sqlx"
)

type UserSegmentPostgres struct {
	db *sqlx.DB
}

func NewUserSegmentPostgres(db *sqlx.DB) *UserSegmentPostgres {
	return &UserSegmentPostgres{
		db: db,
	}
}

func (r *UserSegmentPostgres) CreateUserSegment(userSegment segmentation_service.UserSegment) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (user_id, segment_id) VALUES ($1, (SELECT id FROM segments WHERE segmentname=$2)) RETURNING id;", user_segmentTable)

	row := r.db.QueryRow(query, userSegment.UserId, userSegment.Segmentname)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}