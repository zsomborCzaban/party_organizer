package domains

import "github.com/zsomborCzaban/party_organizer/utils/api"

type IFoodRequirementService interface {
	CreateFoodRequirement(foodReq FoodRequirementDTO, userId uint) api.IResponse
	GetFoodRequirement(foodReqId, userId uint) api.IResponse
	DeleteFoodRequirement(foodReqId, userId uint) api.IResponse
	GetByPartyId(partyId, userId uint) api.IResponse
}
