package domains

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Username string `json:"username"`
	Password string `json:"password"` //it is raw []bytes and not encoded to human-readable format.
	Email    string `json:"email"`
	Friends  []User `gorm:"many2many:user_friends;"`
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
