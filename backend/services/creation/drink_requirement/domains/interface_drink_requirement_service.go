package domains

import "github.com/zsomborCzaban/party_organizer/utils/api"

type IDrinkRequirementService interface {
	CreateDrinkRequirement(dr DrinkRequirementDTO, userId uint) api.IResponse
	GetDrinkRequirement(drinkReqId, userId uint) api.IResponse
	DeleteDrinkRequirement(foodReqId, userId uint) api.IResponse
	GetByPartyId(partyId, userId uint) api.IResponse
}
