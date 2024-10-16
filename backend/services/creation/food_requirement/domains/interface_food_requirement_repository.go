package domains

type IFoodRequirementRepository interface {
	CreateFoodRequirement(*FoodRequirement) error
	FindById(uint) (*FoodRequirement, error)
	UpdateFoodRequirement(*FoodRequirement) error
	DeleteFoodRequirement(*FoodRequirement) error
	GetByPartyId(uint) (*[]FoodRequirement, error)
}
