package repository

import (
	"fmt"
	segmentation_service "segmentation-service"

	"github.com/jmoiron/sqlx"
)

type SegmentPostgres struct {
	db *sqlx.DB
}

func NewSegmentPostgres(db *sqlx.DB) *SegmentPostgres {
	return &SegmentPostgres{
		db: db,
	}
}

func (r *SegmentPostgres) GetAllSegments() ([]segmentation_service.Segment, error) {
	var lists []segmentation_service.Segment

	query := fmt.Sprintf("SELECT s.id, s.segmentname FROM %s s", segmentsTable)
	err := r.db.Select(&lists, query)

	return lists, err
}

func (r *SegmentPostgres) CreateSegment(segment segmentation_service.Segment) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (segmentname) VALUES ($1) RETURNING id;", segmentsTable)

	row := r.db.QueryRow(query, segment.Segmentname)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *SegmentPostgres) DeleteSegment(segment segmentation_service.Segment) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE segmentname = $1;", segmentsTable)

	_, err := r.db.Exec(query, segment.Segmentname)

	return err
}