package usecases

import (
	"errors"
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

func (dr FoodContributionRepository) DeleteByPartyId(partyId uint) error {
	conds := []db.QueryParameter{{
		Field:    "party_id",
		Operator: "=",
		Value:    partyId,
	},
	}

	if err := dr.DbAccess.BatchDelete(conds); err != nil {
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

func (dr FoodContributionRepository) FindById(id uint, associations ...string) (*domains.FoodContribution, error) {
	fetchedContribution, fetchedErr := dr.DbAccess.FindById(id, associations...)
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
func (dr FoodContributionRepository) FindAllBy(columnNames []string, values []interface{}, associations ...string) (*[]domains.FoodContribution, error) {
	if len(columnNames) != len(values) || len(columnNames) == 0 {
		return nil, errors.New(domains.FIND_ALL_BY_INCORRECT_PARAMS)
	}

	queryParams := make([]db.QueryParameter, len(columnNames))

	for i, _ := range columnNames {
		queryParams[i] = db.QueryParameter{
			Field:    columnNames[i],
			Operator: "=",
			Value:    values[i],
		}
	}

	fetchedContributions, fetchedError := dr.DbAccess.Query(queryParams, associations...)
	if fetchedError != nil {
		//we should return errors from the database layer
		return nil, errors.New(domains.FETCH_ERROR)
	}

	contributions, err := fetchedContributions.(*[]domains.FoodContribution)
	if !err {
		return nil, errors.New(domains.FAILED_PARSE_TO_ARRAY)
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
