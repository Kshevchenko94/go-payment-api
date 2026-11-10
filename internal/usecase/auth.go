package usecase

import "Kshevchenko94/go-payment-api/internal/domain"

type AuthInteractor struct {
	userRepo domain.UserRepository
}

func NewAuthInteractor(repo domain.UserRepository) *AuthInteractor {
	return &AuthInteractor{
		userRepo: repo,
	}
}
