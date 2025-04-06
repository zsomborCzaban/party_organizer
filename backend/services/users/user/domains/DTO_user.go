package domains

import (
	"gorm.io/gorm"
)

type UserDTO struct {
	ID                uint   `json:"id"`
	Username          string `json:"username" validate:"required,min=3"`
	Email             string `json:"email" validate:"required"`
	ProfilePictureUrl string `json:"profile_picture_url"`
	Friends           []User `json:"-"`
	//OrganizedParties []domains.Party `json:"organized_parties"`
}

func (u *UserDTO) TransformToUser() *User {
	return &User{
		Model:             gorm.Model{ID: u.ID},
		Username:          u.Username,
		Email:             u.Email,
		ProfilePictureUrl: u.ProfilePictureUrl,
		Friends:           u.Friends,
		//OrganizedParties: u.OrganizedParties,
	}
}
