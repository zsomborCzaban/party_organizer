package usecases

import (
	"github.com/zsomborCzaban/party_organizer/services/creation/drink_requirement/domains"
)

type DrinkRequirementRepository struct{}

func NewDrinkRequirementRepository() domains.IDrinkRequirementRepository {
	return &DrinkRequirementRepository{}
}

func (d DrinkRequirementRepository) CreateDrinkRequirement(drinkRequirementDTO *domains.DrinkRequirementDTO) (*domains.DrinkRequirement, error) {

	return &domains.DrinkRequirement{
			Type:           "vodka",
			TargetQuantity: 4,
			QuantityMark:   "L",
		},
		nil
}

func (d DrinkRequirementRepository) GetDrinkRequirement(id uint) (*domains.DrinkRequirement, error) {
	//TODO implement me
	return &domains.DrinkRequirement{
			Type:           "vodka",
			TargetQuantity: 4,
			QuantityMark:   "L",
		},
		nil
}

func (d DrinkRequirementRepository) UpdateDrinkRequirement(drinkRequirementDTO *domains.DrinkRequirementDTO) (*domains.DrinkRequirement, error) {
	//TODO implement me
	return &domains.DrinkRequirement{
			Type:           "vodka",
			TargetQuantity: 4,
			QuantityMark:   "L",
		},
		nil
}

func (d DrinkRequirementRepository) DeleteDrinkRequirement(id uint) (*domains.DrinkRequirement, error) {
	//TODO implement me
	return &domains.DrinkRequirement{
			Type:           "vodka",
			TargetQuantity: 4,
			QuantityMark:   "L",
		},
		nil
}
