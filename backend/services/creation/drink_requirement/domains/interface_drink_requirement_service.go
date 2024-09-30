package domains

type IDrinkRequirementService interface {
	CreateDrinkRequirement(DrinkRequirementDTO) IResponse
	GetDrinkRequirement(uint) IResponse
	UpdateDrinkRequirement(DrinkRequirementDTO) IResponse
	DeleteDrinkRequirement(uint) IResponse
}
