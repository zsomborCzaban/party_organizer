package domains

import (
	"github.com/zsomborCzaban/party_organizer/services/user/domains"
	"gorm.io/gorm"
)

type FriendInvite struct {
	gorm.Model

	InvitorId uint         `json:"-"`
	Invitor   domains.User `json:"invitor"` //if no json, there will be a typescript error when trying to access nested objects
	InvitedId uint         `json:"-"`
	Invited   domains.User `json:"invited"`
	State     string       `json:"state"`
}
