package usecases

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zsomborCzaban/party_organizer/db"
	"github.com/zsomborCzaban/party_organizer/services/interaction/food_contributions/domains"
	"testing"
)

func setupDefaultRepository() (domains.IFoodContributionRepository, *db.MockDatabaseAccess) {
	dbAccess := new(db.MockDatabaseAccess)

	accessManager := new(db.MockDatabaseAccessManager)
	accessManager.On("RegisterEntity", mock.Anything, mock.Anything).Return(dbAccess)

	repository := NewFoodContributionRepository(accessManager)

	return repository, dbAccess
}

func Test_Create_Success(t *testing.T) {
	repository, dbAccess := setupDefaultRepository()

	dbAccess.On("Create", mock.Anything).Return(nil)

	err := repository.Create(nil)

	assert.Nil(t, err)
}

func Test_Create_Fail(t *testing.T) {
	repository, dbAccess := setupDefaultRepository()

	expectedError := errors.New("create failed")
	dbAccess.On("Create", mock.Anything).Return(expectedError)

	err := repository.Create(nil)

	assert.Equal(t, expectedError, err)
}

func Test_Update_Success(t *testing.T) {
	repository, dbAccess := setupDefaultRepository()

	dbAccess.On("Update", mock.Anything).Return(nil)

	err := repository.Update(nil)

	assert.Nil(t, err)
}

func Test_Update_Fail(t *testing.T) {
	repository, dbAccess := setupDefaultRepository()

	expectedError := errors.New("update failed")
	dbAccess.On("Update", mock.Anything).Return(expectedError)

	err := repository.Update(nil)

	assert.Equal(t, expectedError, err)
}

func Test_FindById_Success(t *testing.T) {
	repository, dbAccess := setupDefaultRepository()

	expectedFoodContribution := new(domains.FoodContribution)
	dbAccess.On("FindById", mock.Anything, mock.Anything).Return(expectedFoodContribution, nil)

	foodContribution, err := repository.FindById(1, "anyAssociation", "Organizer")

	assert.Nil(t, err)
	assert.Equal(t, expectedFoodContribution, foodContribution)
}

func Test_FindById_FailOnFind(t *testing.T) {
	repository, dbAccess := setupDefaultRepository()

	expectedError := errors.New("find failed")
	dbAccess.On("FindById", mock.Anything, mock.Anything).Return(nil, expectedError)

	foodContribution, err := repository.FindById(1, "anyAssociation", "Organizer")

	assert.Nil(t, foodContribution)
	assert.Equal(t, err, expectedError)
}

func Test_FindById_FailOnParse(t *testing.T) {
	repository, dbAccess := setupDefaultRepository()

	dbAccess.On("FindById", mock.Anything, mock.Anything).Return(2, nil)

	foodContribution, err := repository.FindById(1, "anyAssociation", "Organizer")

	assert.Nil(t, foodContribution)
	assert.NotNil(t, err)
}

func Test_Delete_Success(t *testing.T) {
	repository, dbAccess := setupDefaultRepository()

	dbAccess.On("Delete", mock.Anything).Return(nil)

	err := repository.Delete(nil)

	assert.Nil(t, err)
}

func Test_Delete_Fail(t *testing.T) {
	repository, dbAccess := setupDefaultRepository()

	expectedError := errors.New("delete failed")
	dbAccess.On("Delete", mock.Anything).Return(expectedError)

	err := repository.Delete(nil)

	assert.Equal(t, expectedError, err)
}

func Test_DeleteByPartyId_Success(t *testing.T) {
	repository, dbAccess := setupDefaultRepository()

	dbAccess.On("BatchDelete", mock.Anything).Return(nil)

	err := repository.DeleteByPartyId(2)

	assert.Nil(t, err)
}

func Test_DeleteByPartyId_Fail(t *testing.T) {
	repository, dbAccess := setupDefaultRepository()

	expectedError := errors.New("delete by party id failed")
	dbAccess.On("BatchDelete", mock.Anything).Return(expectedError)

	err := repository.DeleteByPartyId(2)

	assert.Equal(t, expectedError, err)
}

func Test_DeleteByReqId_Success(t *testing.T) {
	repository, dbAccess := setupDefaultRepository()

	dbAccess.On("BatchDelete", mock.Anything).Return(nil)

	err := repository.DeleteByReqId(2)

	assert.Nil(t, err)
}

func Test_DeleteByReqId_Fail(t *testing.T) {
	repository, dbAccess := setupDefaultRepository()

	expectedError := errors.New("delete by req id failed")
	dbAccess.On("BatchDelete", mock.Anything).Return(expectedError)

	err := repository.DeleteByReqId(2)

	assert.Equal(t, expectedError, err)
}

func Test_DeleteByContributorId_Success(t *testing.T) {
	repository, dbAccess := setupDefaultRepository()

	dbAccess.On("BatchDelete", mock.Anything).Return(nil)

	err := repository.DeleteByContributorId(2)

	assert.Nil(t, err)
}

func Test_DeleteByContributorId_Fail(t *testing.T) {
	repository, dbAccess := setupDefaultRepository()

	expectedError := errors.New("delete by contributor id failed")
	dbAccess.On("BatchDelete", mock.Anything).Return(expectedError)

	err := repository.DeleteByContributorId(2)

	assert.Equal(t, expectedError, err)
}

func Test_FindAllBy_Success(t *testing.T) {
	repository, dbAccess := setupDefaultRepository()

	expectedFoodContributions := new([]domains.FoodContribution)
	dbAccess.On("Query", mock.Anything, mock.Anything).Return(expectedFoodContributions, nil)

	columns := []string{"column", "names"}
	values := []interface{}{12, "alma"}
	foodContributions, err := repository.FindAllBy(columns, values, "anyAssociation")

	queryParams := make([]db.QueryParameter, len(columns))

	for i, _ := range columns {
		queryParams[i] = db.QueryParameter{
			Field:    columns[i],
			Operator: "=",
			Value:    values[i],
		}
	}

	dbAccess.AssertCalled(t, "Query", queryParams, []string{"anyAssociation"})
	assert.Nil(t, err)
	assert.Equal(t, expectedFoodContributions, foodContributions)
}

func Test_FindAllBy_FailOnIncorrectParams(t *testing.T) {
	repository, _ := setupDefaultRepository()

	expectedError := errors.New(domains.FIND_ALL_BY_INCORRECT_PARAMS)
	columns := []string{"column"}
	values := []interface{}{12, "alma"}
	foodContributions, err := repository.FindAllBy(columns, values, "anyAssociation", "Organizer")

	assert.Nil(t, foodContributions)
	assert.Equal(t, expectedError, err)
}

func Test_FindAllBy_FailOnFind(t *testing.T) {
	repository, dbAccess := setupDefaultRepository()

	expectedError := errors.New(domains.FETCH_ERROR)
	dbAccess.On("Query", mock.Anything, mock.Anything).Return(nil, errors.New(""))

	columns := []string{"column", "names"}
	values := []interface{}{12, "alma"}
	foodContributions, err := repository.FindAllBy(columns, values, "anyAssociation", "Organizer")

	assert.Nil(t, foodContributions)
	assert.Equal(t, err, expectedError)
}

func Test_FindAllBy_FailOnParse(t *testing.T) {
	repository, dbAccess := setupDefaultRepository()

	expectedError := errors.New(domains.FAILED_PARSE_TO_ARRAY)
	dbAccess.On("Query", mock.Anything, mock.Anything).Return(2, nil)

	columns := []string{"column", "names"}
	values := []interface{}{12, "alma"}
	foodContributions, err := repository.FindAllBy(columns, values, "anyAssociation", "Organizer")

	assert.Nil(t, foodContributions)
	assert.Equal(t, err, expectedError)
}

func Test_EntityProvider_Create(t *testing.T) {
	entityProvide := EntityProvider{}
	entity := entityProvide.Create()

	_, ok := entity.(*domains.FoodContribution)
	assert.True(t, ok)
}

func Test_EntityProvider_CreateArray(t *testing.T) {
	entityProvide := EntityProvider{}
	entity := entityProvide.CreateArray()

	_, ok := entity.(*[]domains.FoodContribution)
	assert.True(t, ok)
}
