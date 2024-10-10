package domains

import "github.com/zsomborCzaban/party_organizer/common/api"

type IFoodRequirementService interface {
	CreateFoodRequirement(foodReq FoodRequirementDTO, userId uint) api.IResponse
	GetFoodRequirement(uint) api.IResponse
	UpdateFoodRequirement(FoodRequirementDTO) api.IResponse
	DeleteFoodRequirement(uint) api.IResponse
	GetByPartyId(partyId, userId uint) api.IResponse
}
