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
	DrinkReqId    uint                            `json:"requirement_id" validate:"required"`
	DrinkReq      drinkReqDomain.DrinkRequirement `json:"requirement"`
	PartyId       uint                            `json:"-"`
	Party         partyDomains.Party              `json:"-"`
	Quantity      float32                         `json:"quantity" validate:"required,gt=0"`
	Description   string                          `json:"description"`
}
