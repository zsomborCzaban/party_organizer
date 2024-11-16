package domains

type IFoodRequirementRepository interface {
	CreateFoodRequirement(*FoodRequirement) error
	FindById(id uint, associations ...string) (*FoodRequirement, error)
	UpdateFoodRequirement(*FoodRequirement) error
	DeleteFoodRequirement(*FoodRequirement) error
	//DeleteFoodReqContribution(*FoodRequirement) error
	GetByPartyId(uint) (*[]FoodRequirement, error)
}
