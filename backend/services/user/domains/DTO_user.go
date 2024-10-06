package domains

import (
	"github.com/lib/pq"
	"github.com/zsomborCzaban/party_organizer/common/jwt"
	"gorm.io/gorm"
	"strconv"
)

type UserDTO struct {
	ID                    uint          `json:"id"`
	Username              string        `json:"username" validate:"required,min=3"`
	Email                 string        `json:"email" validate:"required"`
	FriendIDs             pq.Int64Array `json:"friend_ids,omitempty"`
	OrganizingPartyIDs    pq.Int64Array `json:"organizing_party_ids,omitempty"`
	ParticipatingPartyIDs pq.Int64Array `json:"participating_party_ids,omitempty"`
}

func (u *UserDTO) TransformToUser() *User {
	return &User{
		Model:                 gorm.Model{ID: u.ID},
		Username:              u.Username,
		Email:                 u.Email,
		FriendIDs:             u.FriendIDs,
		OrganizingPartyIDs:    u.OrganizingPartyIDs,
		ParticipatingPartyIDs: u.ParticipatingPartyIDs,
	}
}

func (u *UserDTO) GenerateJWT() (*string, error) {
	idString := strconv.FormatUint(uint64(u.ID), 10)

	return jwt.WithClaims(idString, map[string]string{
		"username": u.Username,
		"id":       idString,
	})
}
