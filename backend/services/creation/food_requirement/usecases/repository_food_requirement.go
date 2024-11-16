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

func (fr FoodRequirementRepository) CreateFoodRequirement(foodRequirement *domains.FoodRequirement) error {
	err := fr.DbAccess.Create(foodRequirement)
	if err != nil {
		return err
	}
	return nil

}

func (fr FoodRequirementRepository) FindById(id uint, associations ...string) (*domains.FoodRequirement, error) {
	foodRequirement, err := fr.DbAccess.FindById(id, associations...)
	if err != nil {
		return nil, err
	}

	foodRequirement2, err2 := foodRequirement.(*domains.FoodRequirement)
	if !err2 {
		return nil, errors.New("failed to convert database entity to foodRequirement")
	}
	return foodRequirement2, nil
}

func (fr FoodRequirementRepository) UpdateFoodRequirement(foodRequirement *domains.FoodRequirement) error {
	err := fr.DbAccess.Update(foodRequirement)
	if err != nil {
		return err
	}
	return nil
}

func (fr FoodRequirementRepository) DeleteFoodRequirement(foodRequirement *domains.FoodRequirement) error {
	err := fr.DbAccess.Delete(foodRequirement)
	if err != nil {
		return err
	}
	return nil
}

func (fr FoodRequirementRepository) DeleteByPartyId(partyId uint) error {
	conds := []db.QueryParameter{{
		Field:    "party_id",
		Operator: "=",
		Value:    partyId,
	},
	}

	if err := fr.DbAccess.BatchDelete(conds); err != nil {
		return err
	}
	return nil
}

func (fr FoodRequirementRepository) GetByPartyId(id uint) (*[]domains.FoodRequirement, error) {
	queryParams := []db.QueryParameter{
		{Field: "party_id", Operator: "=", Value: id},
	}

	fetchedFoodReqs, fetchedError := fr.DbAccess.Query(queryParams)
	if fetchedError != nil {
		return nil, fetchedError
	}

	foodReqs, err := fetchedFoodReqs.(*[]domains.FoodRequirement)
	if !err {
		return nil, errors.New("unable to parse fetched data to DrinkRequirements")
	}

	//not sure if parties can be nil after the db function call
	if foodReqs == nil {
		return nil, errors.New("Error. DrinkRequirements were nil")
	}

	return foodReqs, nil
}

type EntityProvider struct {
}

func (e EntityProvider) Create() interface{} {
	return &domains.FoodRequirement{}
}

func (e EntityProvider) CreateArray() interface{} {
	return &[]domains.FoodRequirement{}
}
