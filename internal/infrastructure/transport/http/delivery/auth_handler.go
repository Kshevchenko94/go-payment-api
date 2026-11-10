package delivery

import (
	"Kshevchenko94/go-payment-api/internal/domain"
	"Kshevchenko94/go-payment-api/internal/usecase"
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type AuthuseCase interface {
	Register(ctx context.Context, email, password string) (*domain.User, error)
}

type AuthHandler struct {
	AuthuseCase AuthuseCase
}

func NewAuthHandler(ac AuthuseCase) *AuthHandler {
	return &AuthHandler{
		AuthuseCase: ac,
	}
}

func (a *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	var res RegisterResponse
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid json schema", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if err = req.Validate(); err != nil {
		http.Error(w, "invalid validation", http.StatusBadRequest)
		return
	}
	newUser, err := a.AuthuseCase.Register(r.Context(), req.Email, req.Password)
	if errors.Is(err, usecase.ErrUserAlreadyExists) {
		http.Error(w, "this email is already registered", http.StatusConflict)
		return
	}
	if err != nil {
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	res.ID = newUser.ID
	res.Email = newUser.Email
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}
