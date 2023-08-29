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

func (r *UserSegmentPostgres) CreateUserSegment(userSegment segmentation_service.UserSegment) error {
	var segmentListStr string

	for i := 0; i < len(userSegment.Segmentlist); i++ {
		if i != len(userSegment.Segmentlist) - 1{
			segmentListStr += "'" + userSegment.Segmentlist[i] + "', "
		} else {
			segmentListStr += "'" + userSegment.Segmentlist[i] + "'"
		}
	}

	query := fmt.Sprintf("INSERT INTO %s (user_id, segment_id) SELECT $1, id FROM segments WHERE segmentname IN (%s);", user_segmentTable, segmentListStr)

	_, err := r.db.Exec(query, userSegment.UserId)

	return err
}
