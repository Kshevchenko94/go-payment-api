package delivery

import (
	"errors"
	"strings"
)

var (
	ErrEmptyEmail    = errors.New("Email is empty.")
	ErrInvalidEmail  = errors.New("Email is invalid.")
	ErrShortPassword = errors.New("Password is short.")
)

const MIN_LEN_PASS int = 8

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *RegisterRequest) Validate() error {
	r.Email = strings.TrimSpace(r.Email)

	if r.Email == "" {
		return ErrEmptyEmail
	}

	if !strings.Contains(r.Email, "@") || !strings.Contains(r.Email, ".") {
		return ErrInvalidEmail
	}

	if len(r.Password) < MIN_LEN_PASS {
		return ErrShortPassword
	}

	return nil
}
