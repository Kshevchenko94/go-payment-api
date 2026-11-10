package usecase

import (
	"Kshevchenko94/go-payment-api/internal/domain"
	"Kshevchenko94/go-payment-api/internal/pkg/crypto"
	"context"
	"database/sql"
	"errors"
	"time"
	"uuid"
)

type AuthInteractor struct {
	userRepo domain.UserRepository
}

func NewAuthInteractor(repo domain.UserRepository) *AuthInteractor {
	return &AuthInteractor{
		userRepo: repo,
	}
}

func (ai *AuthInteractor) Register(ctx context.Context, email, password string) error {
	_, err := ai.userRepo.GetByEmail(ctx, email)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return err
	}

	passwordHashed, err := crypto.HashPassword(password)
	if err != nil {
		return err
	}

	newUser := &domain.User{
		ID:           uuid.NewV4().String(),
		Email:        email,
		PasswordHash: passwordHashed,
		CreatedAt:    time.Now(),
	}
	ai.userRepo.Create(ctx, newUser)

	return nil
}
