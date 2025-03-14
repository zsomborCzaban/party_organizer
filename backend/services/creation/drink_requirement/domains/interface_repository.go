package domains

type IDrinkRequirementRepository interface {
	CreateDrinkRequirement(*DrinkRequirement) error
	FindById(id uint, associations ...string) (*DrinkRequirement, error)
	DeleteDrinkRequirement(*DrinkRequirement) error
	DeleteByPartyId(uint) error
	GetByPartyId(uint) (*[]DrinkRequirement, error)
}
