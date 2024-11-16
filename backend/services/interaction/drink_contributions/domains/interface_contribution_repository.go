package domains

type IDrinkContributionRepository interface {
	Create(*DrinkContribution) error
	Update(*DrinkContribution) error
	Delete(*DrinkContribution) error

	DeleteByReqId(foodReqId uint) error
	DeleteByContributorId(userId uint) error
	FindAllBy(columnNames []string, values []interface{}, associations ...string) (*[]DrinkContribution, error)
	FindById(id uint, associations ...string) (*DrinkContribution, error)
}
