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
	FoodReqId     uint                          `json:"food_req_id" validate:"required"`
	FoodReq       foodReqDomain.FoodRequirement `json:"-"`
	PartyId       uint                          `json:"-"`
	Quantity      int                           `json:"quantity" validate:"required,gt=0"`
	Description   string                        `json:"description"`
}
