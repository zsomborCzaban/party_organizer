package usecases

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zsomborCzaban/party_organizer/db"
	"github.com/zsomborCzaban/party_organizer/services/managers/friend_manager/domains"
	"gorm.io/gorm"
	"testing"
)

func setupDefaultRepository() (domains.IFriendInviteRepository, *db.MockDatabaseAccess) {
	dbAccess := new(db.MockDatabaseAccess)

	accessManager := new(db.MockDatabaseAccessManager)
	accessManager.On("RegisterEntity", mock.Anything, mock.Anything).Return(dbAccess)

	repository := NewFriendInviteRepository(accessManager)

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

	dbAccess.On("Update", mock.Anything, mock.Anything).Return(nil)

	friendInvite := domains.FriendInvite{Model: gorm.Model{ID: 1}}
	err := repository.Update(&friendInvite)

	assert.Nil(t, err)
}

func Test_Update_Fail(t *testing.T) {
	repository, dbAccess := setupDefaultRepository()

	expectedError := errors.New("update failed")
	dbAccess.On("Update", mock.Anything, mock.Anything).Return(expectedError)

	friendInvite := domains.FriendInvite{Model: gorm.Model{ID: 1}}
	err := repository.Update(&friendInvite)

	assert.Equal(t, expectedError, err)
}

func Test_FindById_Success(t *testing.T) {
	repository, dbAccess := setupDefaultRepository()

	expectedFriendInvite := &domains.FriendInvite{}
	expectedFriendInvites := []domains.FriendInvite{*expectedFriendInvite}
	dbAccess.On("Query", mock.Anything, mock.Anything).Return(&expectedFriendInvites, nil)

	invitorId := uint(1)
	invitedId := uint(2)
	friendInvite, err := repository.FindByIds(invitorId, invitedId)

	queryParams := []db.QueryParameter{
		{Field: "invitor_id", Operator: "=", Value: invitorId},
		{Field: "invited_id", Operator: "=", Value: invitedId},
	}

	dbAccess.AssertCalled(t, "Query", queryParams, []string{"Invited", "Invitor"})
	assert.Nil(t, err)
	assert.Equal(t, expectedFriendInvite, friendInvite)
}

func Test_FindById_FailOnFind(t *testing.T) {
	repository, dbAccess := setupDefaultRepository()

	expectedError := errors.New(domains.NOT_FOUND)
	dbAccess.On("Query", mock.Anything, mock.Anything).Return(nil, expectedError)

	friendInvite, err := repository.FindByIds(1, 2)

	assert.Nil(t, friendInvite)
	assert.Equal(t, err, expectedError)
}

func Test_FindById_FailOnParse(t *testing.T) {
	repository, dbAccess := setupDefaultRepository()

	expectedError := errors.New(domains.FAILED_PARSE_TO_ARRAY)
	dbAccess.On("Query", mock.Anything, mock.Anything).Return(2, nil)

	friendInvite, err := repository.FindByIds(1, 2)

	assert.Nil(t, friendInvite)
	assert.Equal(t, err, expectedError)
}

func Test_FindById_FailOnLengthGreaterThanOne(t *testing.T) {
	repository, dbAccess := setupDefaultRepository()

	expectedError := errors.New(domains.INVALID_DATABASE_STATE)
	dbAccess.On("Query", mock.Anything, mock.Anything).Return(&[]domains.FriendInvite{{}, {}}, nil)

	friendInvite, err := repository.FindByIds(1, 2)

	assert.Nil(t, friendInvite)
	assert.Equal(t, err, expectedError)
}

func Test_FindById_FailOnLengthEqualsZero(t *testing.T) {
	repository, dbAccess := setupDefaultRepository()

	expectedError := errors.New(domains.NOT_FOUND)
	dbAccess.On("Query", mock.Anything, mock.Anything).Return(&[]domains.FriendInvite{}, nil)

	friendInvite, err := repository.FindByIds(1, 2)

	assert.Nil(t, friendInvite)
	assert.Equal(t, err, expectedError)
}

func Test_FindPendingById_Success(t *testing.T) {
	repository, dbAccess := setupDefaultRepository()

	expectedFriendInvites := &[]domains.FriendInvite{{}}
	dbAccess.On("Query", mock.Anything, mock.Anything).Return(expectedFriendInvites, nil)

	invitedId := uint(1)
	queryParams := []db.QueryParameter{
		{Field: "state", Operator: "=", Value: domains.PENDING},
		{Field: "invited_id", Operator: "=", Value: invitedId},
	}
	pendingInvites, err := repository.FindPendingByInvitedId(invitedId)

	dbAccess.AssertCalled(t, "Query", queryParams, []string{"Invited", "Invitor"})
	assert.Nil(t, err)
	assert.Equal(t, expectedFriendInvites, pendingInvites)
}

func Test_FindPendingById_FailOnFind(t *testing.T) {
	repository, dbAccess := setupDefaultRepository()

	expectedError := errors.New("find failed")
	dbAccess.On("Query", mock.Anything, mock.Anything).Return(nil, expectedError)

	pendingInvites, err := repository.FindPendingByInvitedId(1)

	assert.Nil(t, pendingInvites)
	assert.Equal(t, err, expectedError)
}

func Test_FindPendingById_FailOnParse(t *testing.T) {
	repository, dbAccess := setupDefaultRepository()

	expectedError := errors.New(domains.FAILED_PARSE_TO_ARRAY)
	dbAccess.On("Query", mock.Anything, mock.Anything).Return(2, nil)

	friendInvite, err := repository.FindPendingByInvitedId(1)

	assert.Nil(t, friendInvite)
	assert.Equal(t, err, expectedError)
}

func Test_EntityProvider_Create(t *testing.T) {
	entityProvide := EntityProvider{}
	entity := entityProvide.Create()

	_, ok := entity.(*domains.FriendInvite)
	assert.True(t, ok)
}

func Test_EntityProvider_CreateArray(t *testing.T) {
	entityProvide := EntityProvider{}
	entity := entityProvide.CreateArray()

	_, ok := entity.(*[]domains.FriendInvite)
	assert.True(t, ok)
}
