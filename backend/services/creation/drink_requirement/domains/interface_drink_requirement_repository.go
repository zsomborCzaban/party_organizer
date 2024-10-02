package domains

type IDrinkRequirementRepository interface {
	CreateDrinkRequirement(*DrinkRequirement) error
	GetDrinkRequirement(uint) (*DrinkRequirement, error)
	UpdateDrinkRequirement(*DrinkRequirement) error
	DeleteDrinkRequirement(*DrinkRequirement) error
}
