package domains

import (
	"github.com/zsomborCzaban/party_organizer/services/user/domains"
	"gorm.io/gorm"
)

type FriendInvitation struct {
	gorm.Model

	InvitatorId uint
	Invitator   domains.User
	InvitedId   uint
	Invited     domains.User
	State       string //Pending, Accepted, Declined
}
