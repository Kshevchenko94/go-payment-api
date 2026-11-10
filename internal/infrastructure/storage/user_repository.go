package storage

import (
	"Kshevchenko94/go-payment-api/internal/domain"
	"context"
	"database/sql"
	"fmt"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) Create(ctx context.Context, u *domain.User) (*domain.User, error) {
	query := "INSERT INTO users (id, email, password_hash, created_at) VALUES ($1, $2, $3, $4)"
	_, err := ur.db.ExecContext(ctx, query, u.ID, u.Email, u.PasswordHash, u.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to insert user: %w", err)
	}
	return u, nil
}

func (ur *UserRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	query := "SELECT id, email, password_hash, created_at FROM users WHERE email = $1"
	var u domain.User
	err := ur.db.QueryRowContext(ctx, query, email).Scan(&u.ID, &u.Email, &u.PasswordHash, &u.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, sql.ErrNoRows
	}
	if err != nil {
		return nil, fmt.Errorf("failed to select user: %w", err)
	}
	return &u, nil
}
