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

var ErrUserAlreadyExists = errors.New("user with this email already exists")

type AuthInteractor struct {
	userRepo domain.UserRepository
}

func NewAuthInteractor(repo domain.UserRepository) *AuthInteractor {
	return &AuthInteractor{
		userRepo: repo,
	}
}

func (ai *AuthInteractor) Register(ctx context.Context, email, password string) (*domain.User, error) {
	user, err := ai.userRepo.GetByEmail(ctx, email)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	if user != nil {
		return nil, ErrUserAlreadyExists
	}

	passwordHashed, err := crypto.HashPassword(password)
	if err != nil {
		return nil, err
	}

	newUser := &domain.User{
		ID:           uuid.NewV4().String(),
		Email:        email,
		PasswordHash: passwordHashed,
		CreatedAt:    time.Now(),
	}
	ai.userRepo.Create(ctx, newUser)

	return newUser, nil
}
