package usecases

import (
	"errors"
	"github.com/zsomborCzaban/party_organizer/db"
	"github.com/zsomborCzaban/party_organizer/services/creation/drink_requirement/domains"
)

type DrinkRequirementRepository struct {
	dbAccess db.IDatabaseAccess
}

func NewDrinkRequirementRepository(databaseAccessManager db.IDatabaseAccessManager) domains.IDrinkRequirementRepository {
	entityProvider := EntityProvider{}
	databaseAccess := databaseAccessManager.RegisterEntity("drinkRequirementProvider", entityProvider)

	return &DrinkRequirementRepository{
		dbAccess: databaseAccess,
	}
}

func (pr DrinkRequirementRepository) CreateDrinkRequirement(drinkRequirement *domains.DrinkRequirement) error {
	err := pr.dbAccess.Create(drinkRequirement)
	if err != nil {
		return err
	}
	return nil

}

func (pr DrinkRequirementRepository) GetDrinkRequirement(id uint) (*domains.DrinkRequirement, error) {
	drinkRequirement, err := pr.dbAccess.FindById(id)
	if err != nil {
		return nil, err
	}

	drinkRequirement2, err2 := drinkRequirement.(*domains.DrinkRequirement)
	if !err2 {
		return nil, errors.New("failed to convert database entity to drinkRequirement")
	}
	return drinkRequirement2, nil
}

func (pr DrinkRequirementRepository) UpdateDrinkRequirement(drinkRequirement *domains.DrinkRequirement) error {
	err := pr.dbAccess.Update(drinkRequirement)
	if err != nil {
		return err
	}
	return nil
}

func (pr DrinkRequirementRepository) DeleteDrinkRequirement(drinkRequirement *domains.DrinkRequirement) error {
	err := pr.dbAccess.Delete(drinkRequirement)
	if err != nil {
		return err
	}
	return nil
}

type EntityProvider struct {
}

func (e EntityProvider) Create() interface{} {
	return &domains.DrinkRequirement{}
}

func (e EntityProvider) CreateArray() interface{} {
	return &[]domains.DrinkRequirement{}
}
