package domains

import (
	"errors"
	userDomain "github.com/zsomborCzaban/party_organizer/services/users/user/domains"
	"golang.org/x/crypto/bcrypt"
)

type RegisterRequest struct {
	Username        string `json:"username" validate:"required,min=3"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,min=3,containsany=0123456789"`
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password"`
	ConfirmHash     string
}

func (rq *RegisterRequest) TransformToUser() (*userDomain.User, error) {
	encodedPassword, err := bcrypt.GenerateFromPassword([]byte(rq.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("failed to encode provided password: " + rq.Password)
	}

	return &userDomain.User{
		Username: rq.Username,
		Email:    rq.Email,
		Password: string(encodedPassword),
	}, nil
}
