package domains

import (
	partyDomains "github.com/zsomborCzaban/party_organizer/services/creation/party/domains"
	userDomain "github.com/zsomborCzaban/party_organizer/services/user/domains"
	"gorm.io/gorm"
)

type PartyInvite struct {
	gorm.Model

	InvitorId uint               `json:"-"`
	Invitor   userDomain.User    `json:"invitor"`
	InvitedId uint               `json:"-"`
	Invited   userDomain.User    `json:"invited"`
	PartyId   uint               `json:"-"`
	Party     partyDomains.Party `json:"party"`
	State     string             `json:"state"`
}
