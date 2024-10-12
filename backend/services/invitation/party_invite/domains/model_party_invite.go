package domains

import (
	partyDomains "github.com/zsomborCzaban/party_organizer/services/creation/party/domains"
	userDomain "github.com/zsomborCzaban/party_organizer/services/user/domains"
	"gorm.io/gorm"
)

type PartyInvite struct {
	gorm.Model

	InvitorId uint
	Invitor   userDomain.User
	InvitedId uint
	Invited   userDomain.User
	PartyId   uint
	Party     partyDomains.Party
	State     string
}
