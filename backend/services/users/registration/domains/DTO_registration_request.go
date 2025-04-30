package domains

import (
	"errors"
	"github.com/zsomborCzaban/party_organizer/utils/random"
	"golang.org/x/crypto/bcrypt"
)

type DTORegistrationRequest struct {
	Username        string `json:"username" validate:"required,min=3"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,min=3,containsany=0123456789"`
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password"`
}

func (rq *DTORegistrationRequest) TransformToRegistrationRequest() (*RegistrationRequest, error) {
	encodedPassword, err := bcrypt.GenerateFromPassword([]byte(rq.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("failed to encode provided password: " + rq.Password)
	}

	confirmHash := random.GenerateRandomString(64)

	return &RegistrationRequest{
		Username:    rq.Username,
		Email:       rq.Email,
		Password:    string(encodedPassword),
		ConfirmHash: confirmHash,
	}, nil
}
