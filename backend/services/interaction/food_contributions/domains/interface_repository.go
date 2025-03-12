package domains

type IFoodContributionRepository interface {
	Create(*FoodContribution) error
	Update(*FoodContribution) error
	Delete(*FoodContribution) error

	DeleteByPartyId(partyId uint) error
	DeleteByReqId(foodReqId uint) error
	DeleteByContributorId(userId uint) error
	FindAllBy(columnNames []string, values []interface{}, associations ...string) (*[]FoodContribution, error)
	FindById(id uint, associations ...string) (*FoodContribution, error)
}
