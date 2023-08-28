package repository

import "github.com/jmoiron/sqlx"

type SegmentPostgres struct {
	db *sqlx.DB
}

func NewSegmentPostgres(db *sqlx.DB) *SegmentPostgres {
	return &SegmentPostgres{
		db: db,
	}
}
