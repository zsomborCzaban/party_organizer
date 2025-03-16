package usecases

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zsomborCzaban/party_organizer/db"
	"github.com/zsomborCzaban/party_organizer/services/managers/party_attendance_manager/domains"
	"testing"
)

func setupDefaultRepository() (domains.IPartyInviteRepository, *db.MockDatabaseAccess) {
	dbAccess := new(db.MockDatabaseAccess)

	accessManager := new(db.MockDatabaseAccessManager)
	accessManager.On("RegisterEntity", mock.Anything, mock.Anything).Return(dbAccess)

	repository := NewPartyInviteRepository(accessManager)

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

func Test_Delete_Success(t *testing.T) {
	repository, dbAccess := setupDefaultRepository()

	dbAccess.On("BatchDelete", mock.Anything).Return(nil)

	partyId := uint(1)
	queryParams := []db.QueryParameter{
		{Field: "party_id", Operator: "=", Value: partyId},
	}
	err := repository.DeleteByPartyId(partyId)

	dbAccess.AssertCalled(t, "BatchDelete", queryParams)
	assert.Nil(t, err)
}

func Test_Delete_Fail(t *testing.T) {
	repository, dbAccess := setupDefaultRepository()

	expectedError := errors.New("delete failed")
	dbAccess.On("BatchDelete", mock.Anything).Return(expectedError)

	err := repository.DeleteByPartyId(1)

	assert.Equal(t, expectedError, err)
}

func Test_FindById_Success(t *testing.T) {
	repository, dbAccess := setupDefaultRepository()

	expectedPartyInvite := &domains.PartyInvite{}
	expectedPartyInvites := []domains.PartyInvite{*expectedPartyInvite}
	dbAccess.On("Query", mock.Anything, mock.Anything).Return(&expectedPartyInvites, nil)

	invitedId := uint(1)
	partyId := uint(2)
	partyInvite, err := repository.FindByIds(invitedId, partyId)

	queryParams := []db.QueryParameter{
		{Field: "invited_id", Operator: "=", Value: invitedId},
		{Field: "party_id", Operator: "=", Value: partyId},
	}

	dbAccess.AssertCalled(t, "Query", queryParams, domains.FullPartyInvitePreload)
	assert.Nil(t, err)
	assert.Equal(t, expectedPartyInvite, partyInvite)
}

func Test_FindById_FailOnFind(t *testing.T) {
	repository, dbAccess := setupDefaultRepository()

	expectedError := errors.New(domains.FETCH_ERROR)
	dbAccess.On("Query", mock.Anything, mock.Anything).Return(nil, expectedError)

	partyInvite, err := repository.FindByIds(1, 2)

	assert.Nil(t, partyInvite)
	assert.Equal(t, err, expectedError)
}

func Test_FindById_FailOnParse(t *testing.T) {
	repository, dbAccess := setupDefaultRepository()

	expectedError := errors.New(domains.FAILED_PARSE_TO_ARRAY)
	dbAccess.On("Query", mock.Anything, mock.Anything).Return(2, nil)

	partyInvite, err := repository.FindByIds(1, 2)

	assert.Nil(t, partyInvite)
	assert.Equal(t, err, expectedError)
}

func Test_FindById_FailOnLengthGreaterThanOne(t *testing.T) {
	repository, dbAccess := setupDefaultRepository()

	expectedError := errors.New(domains.INVALID_DATABASE_STATE)
	dbAccess.On("Query", mock.Anything, mock.Anything).Return(&[]domains.PartyInvite{{}, {}}, nil)

	partyInvite, err := repository.FindByIds(1, 2)

	assert.Nil(t, partyInvite)
	assert.Equal(t, err, expectedError)
}

func Test_FindById_FailOnLengthEqualsZero(t *testing.T) {
	repository, dbAccess := setupDefaultRepository()

	expectedError := errors.New(domains.NOT_FOUND)
	dbAccess.On("Query", mock.Anything, mock.Anything).Return(&[]domains.PartyInvite{}, nil)

	partyInvite, err := repository.FindByIds(1, 2)

	assert.Nil(t, partyInvite)
	assert.Equal(t, err, expectedError)
}

func Test_FindPendingByInvitedId_Success(t *testing.T) {
	repository, dbAccess := setupDefaultRepository()

	expectedPartyInvites := &[]domains.PartyInvite{{}}
	dbAccess.On("Query", mock.Anything, mock.Anything).Return(expectedPartyInvites, nil)

	invitedId := uint(1)
	queryParams := []db.QueryParameter{
		{Field: "state", Operator: "=", Value: domains.PENDING},
		{Field: "invited_id", Operator: "=", Value: invitedId},
	}
	pendingInvites, err := repository.FindPendingByInvitedId(invitedId)

	dbAccess.AssertCalled(t, "Query", queryParams, domains.FullPartyInvitePreload)
	assert.Nil(t, err)
	assert.Equal(t, expectedPartyInvites, pendingInvites)
}

func Test_FindPendingByInvitedId_FailOnFind(t *testing.T) {
	repository, dbAccess := setupDefaultRepository()

	expectedError := errors.New("find failed")
	dbAccess.On("Query", mock.Anything, mock.Anything).Return(nil, expectedError)

	pendingInvites, err := repository.FindPendingByInvitedId(1)

	assert.Nil(t, pendingInvites)
	assert.Equal(t, err, expectedError)
}

func Test_FindPendingByInvitedId_FailOnParse(t *testing.T) {
	repository, dbAccess := setupDefaultRepository()

	expectedError := errors.New(domains.FAILED_PARSE_TO_ARRAY)
	dbAccess.On("Query", mock.Anything, mock.Anything).Return(2, nil)

	friendInvite, err := repository.FindPendingByInvitedId(1)

	assert.Nil(t, friendInvite)
	assert.Equal(t, err, expectedError)
}

func Test_FindPendingByPartyId_Success(t *testing.T) {
	repository, dbAccess := setupDefaultRepository()

	expectedPartyInvites := &[]domains.PartyInvite{{}}
	dbAccess.On("Query", mock.Anything, mock.Anything).Return(expectedPartyInvites, nil)

	party_id := uint(1)
	queryParams := []db.QueryParameter{
		{Field: "state", Operator: "=", Value: domains.PENDING},
		{Field: "party_id", Operator: "=", Value: party_id},
	}
	pendingInvites, err := repository.FindPendingByPartyId(party_id)

	dbAccess.AssertCalled(t, "Query", queryParams, domains.FullPartyInvitePreload)
	assert.Nil(t, err)
	assert.Equal(t, expectedPartyInvites, pendingInvites)
}

func Test_FindPendingByPartyId_FailOnFind(t *testing.T) {
	repository, dbAccess := setupDefaultRepository()

	expectedError := errors.New("find failed")
	dbAccess.On("Query", mock.Anything, mock.Anything).Return(nil, expectedError)

	pendingInvites, err := repository.FindPendingByPartyId(1)

	assert.Nil(t, pendingInvites)
	assert.Equal(t, err, expectedError)
}

func Test_FindPendingByPartyId_FailOnParse(t *testing.T) {
	repository, dbAccess := setupDefaultRepository()

	expectedError := errors.New(domains.FAILED_PARSE_TO_ARRAY)
	dbAccess.On("Query", mock.Anything, mock.Anything).Return(2, nil)

	friendInvite, err := repository.FindPendingByPartyId(1)

	assert.Nil(t, friendInvite)
	assert.Equal(t, err, expectedError)
}

func Test_EntityProvider_Create(t *testing.T) {
	entityProvide := EntityProvider{}
	entity := entityProvide.Create()

	_, ok := entity.(*domains.PartyInvite)
	assert.True(t, ok)
}

func Test_EntityProvider_CreateArray(t *testing.T) {
	entityProvide := EntityProvider{}
	entity := entityProvide.CreateArray()

	_, ok := entity.(*[]domains.PartyInvite)
	assert.True(t, ok)
}
