package domains

import (
	foodReqDomain "github.com/zsomborCzaban/party_organizer/services/creation/food_requirement/domains"
	userDomain "github.com/zsomborCzaban/party_organizer/services/user/domains"
	"gorm.io/gorm"
)

type FoodContribution struct {
	gorm.Model

	ContributorId uint                          `json:"contributor_id"`
	Contributor   userDomain.User               `json:"-"`
	FoodReqId     uint                          `json:"food_req_id"`
	FoodReq       foodReqDomain.FoodRequirement `json:"-"`
	PartyId       uint                          `json:"party_id"`
	Quantity      int                           `json:"quantity"`
	Description   string                        `json:"description"`
}
