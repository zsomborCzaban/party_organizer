package domains

type IDrinkContributionRepository interface {
	Create(*DrinkContribution) error
	Update(*DrinkContribution) error
	Delete(*DrinkContribution) error

	FindAllBy(columnNames []string, values []interface{}) (*[]DrinkContribution, error)
	FindById(uint) (*DrinkContribution, error)
}
