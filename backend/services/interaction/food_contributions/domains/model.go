package domains

import (
	foodReqDomain "github.com/zsomborCzaban/party_organizer/services/creation/food_requirement/domains"
	partyDomains "github.com/zsomborCzaban/party_organizer/services/creation/party/domains"
	userDomain "github.com/zsomborCzaban/party_organizer/services/users/user/domains"
	"gorm.io/gorm"
)

type FoodContribution struct {
	gorm.Model

	ContributorId uint                          `json:"contributor_id"`
	Contributor   userDomain.User               `json:"contributor"`
	FoodReqId     uint                          `json:"requirement_id" validate:"required"`
	FoodReq       foodReqDomain.FoodRequirement `json:"requirement"`
	PartyId       uint                          `json:"-"`
	Party         partyDomains.Party            `json:"-"`
	Quantity      float32                       `json:"quantity" validate:"required,gt=0"`
	Description   string                        `json:"description"`
}
