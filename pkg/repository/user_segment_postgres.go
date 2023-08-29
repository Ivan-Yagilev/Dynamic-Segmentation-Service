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

func (r *UserSegmentPostgres) DeleteUserSegment(userSegment segmentation_service.UserSegment) error {
	var segmentListStr string

	for i := 0; i < len(userSegment.Segmentlist); i++ {
		if i != len(userSegment.Segmentlist) - 1{
			segmentListStr += "'" + userSegment.Segmentlist[i] + "', "
		} else {
			segmentListStr += "'" + userSegment.Segmentlist[i] + "'"
		}
	}

	query := fmt.Sprintf("DELETE FROM %s WHERE user_id=$1 AND segment_id IN (SELECT id FROM segments WHERE segmentname IN (%s));", user_segmentTable, segmentListStr)

	_, err := r.db.Exec(query, userSegment.UserId)

	return err
}

func (r *UserSegmentPostgres) GetAllSegments(userId int) ([]segmentation_service.Segment, error) {
	var lists []segmentation_service.Segment

	query := fmt.Sprintf("SELECT s.id AS id, s.segmentname AS segmentname FROM %s u LEFT JOIN %s us ON us.user_id = u.id LEFT JOIN %s s ON s.id = us.segment_id WHERE u.id=$1;", 
		usersTable, user_segmentTable, segmentsTable)
	err := r.db.Select(&lists, query, userId)

	if err != nil {
		return nil, nil
	}

	return lists, err
}