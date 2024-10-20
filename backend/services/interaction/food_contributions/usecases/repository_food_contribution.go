package usecases

import (
	"errors"
	"fmt"
	"github.com/zsomborCzaban/party_organizer/db"
	"github.com/zsomborCzaban/party_organizer/services/interaction/food_contributions/domains"
)

type FoodContributionRepository struct {
	DbAccess db.IDatabaseAccess
}

func NewFoodContributionRepository(dbAccessManager db.IDatabaseAccessManager) domains.IFoodContributionRepository {
	entityProvider := EntityProvider{}
	databaseAccess := dbAccessManager.RegisterEntity("foodContributionProvider", entityProvider)

	return &FoodContributionRepository{
		DbAccess: databaseAccess,
	}
}

func (dr FoodContributionRepository) Create(contribution *domains.FoodContribution) error {
	if err := dr.DbAccess.Create(contribution); err != nil {
		return err
	}
	return nil
}

func (dr FoodContributionRepository) Update(contribution *domains.FoodContribution) error {
	if err := dr.DbAccess.Update(contribution); err != nil {
		return err
	}
	return nil
}

func (dr FoodContributionRepository) Delete(contribution *domains.FoodContribution) error {
	if err := dr.DbAccess.Delete(contribution); err != nil {
		return err
	}
	return nil
}

func (dr FoodContributionRepository) DeleteByReqId(foodReqId uint) error {
	conds := []db.QueryParameter{{
		Field:    "food_req_id",
		Operator: "=",
		Value:    foodReqId,
	},
	}

	if err := dr.DbAccess.BatchDelete(conds); err != nil {
		return err
	}
	return nil
}

func (dr FoodContributionRepository) DeleteByContributorId(contributorId uint) error {
	conds := []db.QueryParameter{{
		Field:    "contributor_id",
		Operator: "=",
		Value:    contributorId,
	},
	}

	if err := dr.DbAccess.BatchDelete(conds); err != nil {
		return err
	}
	return nil
}

func (dr FoodContributionRepository) FindById(id uint) (*domains.FoodContribution, error) {
	fetchedContribution, fetchedErr := dr.DbAccess.FindById(id)
	if fetchedErr != nil {
		return nil, fetchedErr
	}

	contribution, err := fetchedContribution.(*domains.FoodContribution)
	if !err {
		return nil, errors.New(domains.FAILED_PARSE)
	}
	return contribution, nil
}

// FindAllBy culd also get the []db.QueryParameter as param, but then maybe move QueryParamter to utils package
func (dr FoodContributionRepository) FindAllBy(columnNames []string, values []interface{}) (*[]domains.FoodContribution, error) {
	if len(columnNames) != len(values) || len(columnNames) == 0 {
		return nil, errors.New("incorrect use of FindAllBy")
	}

	queryParams := make([]db.QueryParameter, len(columnNames))

	for i, _ := range columnNames {
		queryParams[i] = db.QueryParameter{
			Field:    columnNames[i],
			Operator: "=",
			Value:    values[i],
		}
	}

	fetchedContributions, fetchedError := dr.DbAccess.Query(queryParams)
	if fetchedError != nil {
		//we should return errors from the database layer
		return nil, errors.New(fmt.Sprintf("Error while fetching contributions"))
	}

	contributions, err := fetchedContributions.(*[]domains.FoodContribution)
	if !err {
		return nil, errors.New(domains.FAILED_PARSE_TO_ARRAY)
	}

	//not sure if contributions can be nil after the db function call
	if contributions == nil {
		return nil, errors.New(domains.NOT_FOUND)
	}

	return contributions, nil
}

type EntityProvider struct {
}

func (e EntityProvider) Create() interface{} {
	return &domains.FoodContribution{}
}

func (e EntityProvider) CreateArray() interface{} {
	return &[]domains.FoodContribution{}
}
