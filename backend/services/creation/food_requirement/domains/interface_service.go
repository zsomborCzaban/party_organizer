package domains

import "github.com/zsomborCzaban/party_organizer/utils/api"

type IFoodRequirementService interface {
	Create(foodReq FoodRequirementDTO, userId uint) api.IResponse
	FindById(foodReqId, userId uint) api.IResponse
	Delete(foodReqId, userId uint) api.IResponse
	GetByPartyId(partyId, userId uint) api.IResponse
}
