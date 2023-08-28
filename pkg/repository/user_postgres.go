package repository

import (
	"fmt"
	segmentation_service "segmentation-service"

	"github.com/jmoiron/sqlx"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{
		db: db,
	}
}

func (r *UserPostgres) CreateUser(user segmentation_service.User) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (username, password) VALUES ($1, $2) RETURNING id;", usersTable)

	row := r.db.QueryRow(query, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *UserPostgres) DeleteUser(userId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1;", usersTable)

	_, err := r.db.Exec(query, userId)

	return err
}
