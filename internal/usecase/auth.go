package usecase

import (
	"Kshevchenko94/go-payment-api/internal/domain"
	"context"
	"errors"
	"time"
	"uuid"
)

var (
	ErrUserExist = errors.New("User does exist.")
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
	user, err := ai.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return err
	}

	if user != nil {
		return ErrUserExist
	}

	passwordHashed := ""

	newUser := &domain.User{
		ID:           uuid.NewV4().String(),
		Email:        email,
		PasswordHash: passwordHashed,
		CreatedAt:    time.Now(),
	}
	ai.userRepo.Create(ctx, newUser)

	return nil
}
