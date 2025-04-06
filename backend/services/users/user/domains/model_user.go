package domains

import (
	"github.com/zsomborCzaban/party_organizer/utils/jwt"
	"gorm.io/gorm"
	"strconv"
)

type User struct {
	gorm.Model

	Username          string `json:"username"`
	Password          string `json:"-"` //it is raw []bytes and not encoded to human-readable format.
	Email             string `json:"email"`
	ProfilePictureUrl string `json:"profile_picture_url"`
	Friends           []User `gorm:"many2many:user_friends;"`
	//OrganizedParties []domains.Party `json:"organized_parties"`
}

func (u *User) TransformToUserDTO() *UserDTO {
	return &UserDTO{
		ID:                u.Model.ID,
		Username:          u.Username,
		Email:             u.Email,
		ProfilePictureUrl: u.ProfilePictureUrl,
		Friends:           u.Friends,
		//OrganizedParties: u.OrganizedParties,
	}
}

func (u *User) HasFriend(friendId uint) bool {
	for _, friend := range u.Friends {
		if friend.ID == friendId {
			return true
		}
	}
	return false
}

func (u *User) GenerateJWT() (*string, error) {
	idString := strconv.FormatUint(uint64(u.ID), 10)

	return jwt.WithClaims(idString, map[string]string{
		"email":             u.Email,
		"username":          u.Username,
		"id":                idString,
		"profilePictureUrl": u.ProfilePictureUrl,
	})
}
