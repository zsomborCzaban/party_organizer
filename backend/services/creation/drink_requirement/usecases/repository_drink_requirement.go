package usecases

import (
	"errors"
	"github.com/zsomborCzaban/party_organizer/db"
	"github.com/zsomborCzaban/party_organizer/services/creation/drink_requirement/domains"
)

type DrinkRequirementRepository struct {
	DbAccess db.IDatabaseAccess
}

func NewDrinkRequirementRepository(databaseAccessManager db.IDatabaseAccessManager) domains.IDrinkRequirementRepository {
	entityProvider := EntityProvider{}
	databaseAccess := databaseAccessManager.RegisterEntity("drinkRequirementProvider", entityProvider)

	return &DrinkRequirementRepository{
		DbAccess: databaseAccess,
	}
}

func (dr DrinkRequirementRepository) CreateDrinkRequirement(drinkRequirement *domains.DrinkRequirement) error {
	err := dr.DbAccess.Create(drinkRequirement)
	if err != nil {
		return err
	}
	return nil

}

func (dr DrinkRequirementRepository) FindById(id uint) (*domains.DrinkRequirement, error) {
	drinkRequirement, err := dr.DbAccess.FindById(id, "Party")
	if err != nil {
		return nil, err
	}

	drinkRequirement2, err2 := drinkRequirement.(*domains.DrinkRequirement)
	if !err2 {
		return nil, errors.New("failed to convert database entity to drinkRequirement")
	}
	return drinkRequirement2, nil
}

func (dr DrinkRequirementRepository) UpdateDrinkRequirement(drinkRequirement *domains.DrinkRequirement) error {
	err := dr.DbAccess.Update(drinkRequirement)
	if err != nil {
		return err
	}
	return nil
}

func (dr DrinkRequirementRepository) DeleteDrinkRequirement(drinkRequirement *domains.DrinkRequirement) error {
	err := dr.DbAccess.Delete(drinkRequirement)
	if err != nil {
		return err
	}
	return nil
}

func (dr DrinkRequirementRepository) GetByPartyId(id uint) (*[]domains.DrinkRequirement, error) {
	queryParams := []db.QueryParameter{
		{Field: "party_id", Operator: "=", Value: id},
	}

	fetchedDrinkReqs, fetchedError := dr.DbAccess.Query(queryParams)
	if fetchedError != nil {
		return nil, fetchedError
	}

	drinkReqs, err := fetchedDrinkReqs.(*[]domains.DrinkRequirement)
	if !err {
		return nil, errors.New("unable to parse fetched data to DrinkRequirements")
	}

	//not sure if parties can be nil after the db function call
	if drinkReqs == nil {
		return nil, errors.New("Error. DrinkRequirements were nil")
	}

	return drinkReqs, nil
}

type EntityProvider struct {
}

func (e EntityProvider) Create() interface{} {
	return &domains.DrinkRequirement{}
}

func (e EntityProvider) CreateArray() interface{} {
	return &[]domains.DrinkRequirement{}
}
