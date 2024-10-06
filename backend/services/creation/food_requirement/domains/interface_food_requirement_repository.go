package domains

type IFoodRequirementRepository interface {
	CreateFoodRequirement(*FoodRequirement) error
	GetFoodRequirement(uint) (*FoodRequirement, error)
	UpdateFoodRequirement(*FoodRequirement) error
	DeleteFoodRequirement(*FoodRequirement) error
}
