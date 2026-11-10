package storage

import (
	"database/sql"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepositiry(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}
