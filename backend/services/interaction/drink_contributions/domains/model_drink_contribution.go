package domains

import (
	drinkReqDomain "github.com/zsomborCzaban/party_organizer/services/creation/drink_requirement/domains"
	partyDomains "github.com/zsomborCzaban/party_organizer/services/creation/party/domains"
	userDomain "github.com/zsomborCzaban/party_organizer/services/user/domains"
	"gorm.io/gorm"
)

type DrinkContribution struct {
	gorm.Model

	ContributorId uint                            `json:"contributor_id"`
	Contributor   userDomain.User                 `json:"contributor"`
	DrinkReqId    uint                            `json:"drink_req_id"`
	DrinkReq      drinkReqDomain.DrinkRequirement `json:"drink_req"`
	PartyId       uint                            `json:"party_id"`
	Party         partyDomains.Party              `json:"party"`
	Quantity      int                             `json:"quantity"`
	Description   string                          `json:"description"`
}
