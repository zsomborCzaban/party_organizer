package domains

type IFoodRequirementRepository interface {
	Create(*FoodRequirement) error
	FindById(id uint, associations ...string) (*FoodRequirement, error)
	Delete(*FoodRequirement) error
	DeleteByPartyId(uint) error
	GetByPartyId(uint) (*[]FoodRequirement, error)
}
