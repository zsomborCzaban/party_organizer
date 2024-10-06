package domains

import "github.com/zsomborCzaban/party_organizer/common/api"

type IDrinkRequirementService interface {
	CreateDrinkRequirement(DrinkRequirementDTO) api.IResponse
	GetDrinkRequirement(uint) api.IResponse
	UpdateDrinkRequirement(DrinkRequirementDTO) api.IResponse
	DeleteDrinkRequirement(uint) api.IResponse
}
