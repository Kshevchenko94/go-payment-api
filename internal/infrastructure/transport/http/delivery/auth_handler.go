package delivery

import (
	"context"
	"encoding/json"
	"net/http"
)

type AuthuseCase interface {
	Register(ctx context.Context, email, password string) error
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
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid json schema", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if err = req.Validate(); err != nil {
		http.Error(w, "invalid json schema", http.StatusBadRequest)
		return
	}
	err = a.AuthuseCase.Register(r.Context(), req.Email, req.Password)
	if err != nil {
		http.Error(w, "invalid json schema", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
