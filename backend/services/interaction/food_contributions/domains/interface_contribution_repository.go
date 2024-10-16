package domains

type IFoodContributionRepository interface {
	Create(*FoodContribution) error
	Update(*FoodContribution) error
	Delete(*FoodContribution) error

	DeleteByReqId(foodReqId uint) error
	FindAllBy(columnNames []string, values []interface{}) (*[]FoodContribution, error)
	FindById(uint) (*FoodContribution, error)
}
