package domains

type IDrinkRequirementRepository interface {
	CreateDrinkRequirement(*DrinkRequirement) error
	FindById(uint) (*DrinkRequirement, error)
	UpdateDrinkRequirement(*DrinkRequirement) error
	DeleteDrinkRequirement(*DrinkRequirement) error
	GetByPartyId(uint) (*[]DrinkRequirement, error)
}
