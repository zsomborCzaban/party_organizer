package domains

import (
	"github.com/zsomborCzaban/party_organizer/common"
	userDomain "github.com/zsomborCzaban/party_organizer/services/users/user/domains"
	"gorm.io/gorm"
)

type RegistrationRequest struct {
	gorm.Model
	Username    string `json:"username"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	ConfirmHash string
}

func (rq *RegistrationRequest) TransformToUser() (*userDomain.User, error) {
	return &userDomain.User{
		Username:          rq.Username,
		Email:             rq.Email,
		ProfilePictureUrl: common.DEFAULT_PROFILE_PICTURE_URL,
		Password:          rq.Password,
	}, nil
}
