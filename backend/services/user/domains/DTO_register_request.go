package domains

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type RegisterRequest struct {
	Username string `json:"username" validate:"required,min=3"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required,min=3"`
}

func (rq *RegisterRequest) TransformToUser() (*User, error) {
	encodedPassword, err := bcrypt.GenerateFromPassword([]byte(rq.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("failed to encode provided password: " + rq.Password)
	}

	return &User{
		Username: rq.Username,
		Email:    rq.Email,
		Password: string(encodedPassword),
	}, nil
}
