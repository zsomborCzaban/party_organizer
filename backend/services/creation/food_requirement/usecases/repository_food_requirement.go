package usecases

import (
	"errors"
	"github.com/zsomborCzaban/party_organizer/db"
	"github.com/zsomborCzaban/party_organizer/services/creation/food_requirement/domains"
)

type FoodRequirementRepository struct {
	DbAccess db.IDatabaseAccess
}

func NewFoodRequirementRepository(databaseAccessManager db.IDatabaseAccessManager) domains.IFoodRequirementRepository {
	entityProvider := EntityProvider{}
	databaseAccess := databaseAccessManager.RegisterEntity("foodRequirementProvider", entityProvider)

	return &FoodRequirementRepository{
		DbAccess: databaseAccess,
	}
}

func (pr FoodRequirementRepository) CreateFoodRequirement(foodRequirement *domains.FoodRequirement) error {
	err := pr.DbAccess.Create(foodRequirement)
	if err != nil {
		return err
	}
	return nil

}

func (pr FoodRequirementRepository) GetFoodRequirement(id uint) (*domains.FoodRequirement, error) {
	foodRequirement, err := pr.DbAccess.FindById(id)
	if err != nil {
		return nil, err
	}

	foodRequirement2, err2 := foodRequirement.(*domains.FoodRequirement)
	if !err2 {
		return nil, errors.New("failed to convert database entity to foodRequirement")
	}
	return foodRequirement2, nil
}

func (pr FoodRequirementRepository) UpdateFoodRequirement(foodRequirement *domains.FoodRequirement) error {
	err := pr.DbAccess.Update(foodRequirement)
	if err != nil {
		return err
	}
	return nil
}

func (pr FoodRequirementRepository) DeleteFoodRequirement(foodRequirement *domains.FoodRequirement) error {
	err := pr.DbAccess.Delete(foodRequirement)
	if err != nil {
		return err
	}
	return nil
}

type EntityProvider struct {
}

func (e EntityProvider) Create() interface{} {
	return &domains.FoodRequirement{}
}

func (e EntityProvider) CreateArray() interface{} {
	return &[]domains.FoodRequirement{}
}
