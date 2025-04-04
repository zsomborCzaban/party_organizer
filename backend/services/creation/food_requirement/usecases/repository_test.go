package usecases

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zsomborCzaban/party_organizer/db"
	"github.com/zsomborCzaban/party_organizer/services/creation/food_requirement/domains"
	"testing"
)

func setupDefaultRepository() (domains.IFoodRequirementRepository, *db.MockDatabaseAccess) {
	dbAccess := new(db.MockDatabaseAccess)

	accessManager := new(db.MockDatabaseAccessManager)
	accessManager.On("RegisterEntity", mock.Anything, mock.Anything).Return(dbAccess)

	repository := NewFoodRequirementRepository(accessManager)

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

func Test_FindById_Success(t *testing.T) {
	repository, dbAccess := setupDefaultRepository()

	expectedFoodRequirement := new(domains.FoodRequirement)
	dbAccess.On("FindById", mock.Anything, mock.Anything).Return(expectedFoodRequirement, nil)

	foodRequirement, err := repository.FindById(1, "anyAssociation", "Organizer")

	assert.Nil(t, err)
	assert.Equal(t, expectedFoodRequirement, foodRequirement)
}

func Test_FindById_FailOnFind(t *testing.T) {
	repository, dbAccess := setupDefaultRepository()

	expectedError := errors.New("find failed")
	dbAccess.On("FindById", mock.Anything, mock.Anything).Return(nil, expectedError)

	foodRequirement, err := repository.FindById(1, "anyAssociation", "Organizer")

	assert.Nil(t, foodRequirement)
	assert.Equal(t, err, expectedError)
}

func Test_FindById_FailOnParse(t *testing.T) {
	repository, dbAccess := setupDefaultRepository()

	dbAccess.On("FindById", mock.Anything, mock.Anything).Return(2, nil)

	foodRequirement, err := repository.FindById(1, "anyAssociation", "Organizer")

	assert.Nil(t, foodRequirement)
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

func Test_GetByPartyId_Success(t *testing.T) {
	repository, dbAccess := setupDefaultRepository()

	expectedFoodRequirements := &[]domains.FoodRequirement{}
	dbAccess.On("Query", mock.Anything, mock.Anything).Return(expectedFoodRequirements, nil)

	foodRequirements, err := repository.GetByPartyId(2)

	assert.Nil(t, err)
	assert.Equal(t, expectedFoodRequirements, foodRequirements)
}

func Test_GetByPartyId_FailOnQuery(t *testing.T) {
	repository, dbAccess := setupDefaultRepository()

	expectedError := errors.New("get by party id failed")
	dbAccess.On("Query", mock.Anything, mock.Anything).Return(nil, expectedError)

	foodRequirements, err := repository.GetByPartyId(2)

	assert.Nil(t, foodRequirements)
	assert.Equal(t, expectedError, err)
}

func Test_GetByPartyId_FailOnParse(t *testing.T) {
	repository, dbAccess := setupDefaultRepository()

	dbAccess.On("Query", mock.Anything, mock.Anything).Return(3, nil)

	foodRequirements, err := repository.GetByPartyId(2)

	assert.Nil(t, foodRequirements)
	assert.NotNil(t, err)
}

func Test_GetByPartyId_FailOnNilValue(t *testing.T) {
	repository, dbAccess := setupDefaultRepository()

	dbAccess.On("Query", mock.Anything, mock.Anything).Return(nil, nil)

	foodRequirements, err := repository.GetByPartyId(2)

	assert.Nil(t, foodRequirements)
	assert.NotNil(t, err)
}

func Test_EntityProvider_Create(t *testing.T) {
	entityProvide := EntityProvider{}
	entity := entityProvide.Create()

	_, ok := entity.(*domains.FoodRequirement)
	assert.True(t, ok)
}

func Test_EntityProvider_CreateArray(t *testing.T) {
	entityProvide := EntityProvider{}
	entity := entityProvide.CreateArray()

	_, ok := entity.(*[]domains.FoodRequirement)
	assert.True(t, ok)
}
