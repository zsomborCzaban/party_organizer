package domains

type IFoodRequirementRepository interface {
	CreateFoodRequirement(*FoodRequirement) error
	FindById(id uint, associations ...string) (*FoodRequirement, error)
	UpdateFoodRequirement(*FoodRequirement) error
	DeleteFoodRequirement(*FoodRequirement) error
	DeleteByPartyId(uint) error
	GetByPartyId(uint) (*[]FoodRequirement, error)
}
