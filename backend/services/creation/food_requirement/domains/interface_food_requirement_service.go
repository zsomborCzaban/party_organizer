package domains

import "github.com/zsomborCzaban/party_organizer/common/api"

type IFoodRequirementService interface {
	CreateFoodRequirement(foodReq FoodRequirementDTO, userId uint) api.IResponse
	GetFoodRequirement(foodReqId, userId uint) api.IResponse
	UpdateFoodRequirement(foodReqDTO FoodRequirementDTO, userId uint) api.IResponse
	DeleteFoodRequirement(uint) api.IResponse
	GetByPartyId(partyId, userId uint) api.IResponse
}
