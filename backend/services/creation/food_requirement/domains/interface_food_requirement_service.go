package domains

import "github.com/zsomborCzaban/party_organizer/common/api"

type IFoodRequirementService interface {
	CreateFoodRequirement(FoodRequirementDTO) api.IResponse
	GetFoodRequirement(uint) api.IResponse
	UpdateFoodRequirement(FoodRequirementDTO) api.IResponse
	DeleteFoodRequirement(uint) api.IResponse
}
