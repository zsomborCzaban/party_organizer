package domains

import "github.com/zsomborCzaban/party_organizer/utils/api"

type IDrinkRequirementService interface {
	Create(dr DrinkRequirementDTO, userId uint) api.IResponse
	FindById(drinkReqId, userId uint) api.IResponse
	Delete(drinkReqId, userId uint) api.IResponse
	GetByPartyId(partyId, userId uint) api.IResponse
}
