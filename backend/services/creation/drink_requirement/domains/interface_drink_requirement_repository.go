package domains

type IDrinkRequirementRepository interface {
	CreateDrinkRequirement(*DrinkRequirementDTO) (*DrinkRequirement, error)
	GetDrinkRequirement(uint) (*DrinkRequirement, error)
	UpdateDrinkRequirement(*DrinkRequirementDTO) (*DrinkRequirement, error)
	DeleteDrinkRequirement(uint) (*DrinkRequirement, error)
}
