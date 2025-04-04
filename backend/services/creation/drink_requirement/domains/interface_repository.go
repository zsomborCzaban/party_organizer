package domains

type IDrinkRequirementRepository interface {
	Create(*DrinkRequirement) error
	FindById(id uint, associations ...string) (*DrinkRequirement, error)
	Delete(*DrinkRequirement) error
	DeleteByPartyId(uint) error
	GetByPartyId(uint) (*[]DrinkRequirement, error)
}
