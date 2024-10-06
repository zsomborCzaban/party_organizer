package domains

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Username              string        `json:"username"`
	Password              string        `json:"password"` //it is raw []bytes and not encoded to human-readable format.
	Email                 string        `json:"email"`
	FriendIDs             pq.Int64Array `json:"friend_ids" gorm:"type:integer[]"`
	OrganizingPartyIDs    pq.Int64Array `json:"organizing_party_ids" gorm:"type:integer[]"`
	ParticipatingPartyIDs pq.Int64Array `json:"participating_party_ids" gorm:"type:integer[]"`
}

func (u *User) TransformToUserDTO() *UserDTO {
	return &UserDTO{
		ID:                    u.Model.ID,
		Username:              u.Username,
		Email:                 u.Email,
		FriendIDs:             u.FriendIDs,
		OrganizingPartyIDs:    u.OrganizingPartyIDs,
		ParticipatingPartyIDs: u.ParticipatingPartyIDs,
	}
}
