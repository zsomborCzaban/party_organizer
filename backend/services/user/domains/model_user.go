package domains

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Username          string `json:"username"`
	Password          string `json:"-"` //it is raw []bytes and not encoded to human-readable format.
	Email             string `json:"email"`
	ProfilePictureKey string `json:"-"`
	Friends           []User `gorm:"many2many:user_friends;"`
	//OrganizedParties []domains.Party `json:"organized_parties"`
}

func (u *User) TransformToUserDTO() *UserDTO {
	return &UserDTO{
		ID:       u.Model.ID,
		Username: u.Username,
		Email:    u.Email,
		Friends:  u.Friends,
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
