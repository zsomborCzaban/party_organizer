package domains

import (
	"github.com/zsomborCzaban/party_organizer/common/jwt"
	"gorm.io/gorm"
	"strconv"
)

type UserDTO struct {
	ID       uint   `json:"id"`
	Username string `json:"username" validate:"required,min=3"`
	Email    string `json:"email" validate:"required"`
	Friends  []User `json:"-"`
	//OrganizedParties []domains.Party `json:"organized_parties"`
}

func (u *UserDTO) TransformToUser() *User {
	return &User{
		Model:    gorm.Model{ID: u.ID},
		Username: u.Username,
		Email:    u.Email,
		Friends:  u.Friends,
		//OrganizedParties: u.OrganizedParties,
	}
}

func (u *UserDTO) GenerateJWT() (*string, error) {
	idString := strconv.FormatUint(uint64(u.ID), 10)

	return jwt.WithClaims(idString, map[string]string{
		"username": u.Username,
		"id":       idString,
	})
}
