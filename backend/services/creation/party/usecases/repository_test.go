package usecases

import (
	"errors"
	userDomain "github.com/zsomborCzaban/party_organizer/services/users/user/domains"
	"gorm.io/gorm"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zsomborCzaban/party_organizer/db"
	"github.com/zsomborCzaban/party_organizer/services/creation/party/domains"
)

func setupPartyRepository() (Repository, *db.MockDatabaseAccess) {
	mockDb := new(db.MockDatabaseAccess)
	repo := Repository{
		DbAccess: mockDb,
	}
	return repo, mockDb
}

func TestPartyRepository_AddUserToParty_Success(t *testing.T) {
	repo, mockDb := setupPartyRepository()
	party := &domains.Party{}
	user := &userDomain.User{Model: gorm.Model{ID: 1}}

	mockDb.On("Update", party, party.ID).Return(nil)

	err := repo.AddUserToParty(party, user)

	assert.NoError(t, err)
	assert.Len(t, party.Participants, 1)
	mockDb.AssertExpectations(t)
}

func TestPartyRepository_RemoveUserFromParty_Success(t *testing.T) {
	repo, mockDb := setupPartyRepository()
	user := &userDomain.User{Model: gorm.Model{ID: 1}}
	party := &domains.Party{
		Participants: []userDomain.User{{Model: gorm.Model{ID: 1}}, {Model: gorm.Model{ID: 2}}},
	}

	mockDb.On("ReplaceAssociations", mock.Anything).Return(nil)

	err := repo.RemoveUserFromParty(party, user)

	assert.NoError(t, err)
	mockDb.AssertExpectations(t)
}

func TestPartyRepository_GetPublicParties_Success(t *testing.T) {
	repo, mockDb := setupPartyRepository()
	expectedParties := &[]domains.Party{{Model: gorm.Model{ID: 1}}}

	mockDb.On("Query", []db.QueryParameter{
		{Field: "private", Operator: "=", Value: false},
	}, mock.Anything).Return(expectedParties, nil)

	result, err := repo.GetPublicParties()

	assert.NoError(t, err)
	assert.Equal(t, expectedParties, result)
	mockDb.AssertExpectations(t)
}

func TestPartyRepository_GetPublicParties_QueryError(t *testing.T) {
	repo, mockDb := setupPartyRepository()
	expectedErr := errors.New("query error")

	mockDb.On("Query", mock.Anything, mock.Anything).Return(nil, expectedErr)

	_, err := repo.GetPublicParties()

	assert.Error(t, err)
	mockDb.AssertExpectations(t)
}

func TestPartyRepository_GetPublicParties_TypeAssertionError(t *testing.T) {
	repo, mockDb := setupPartyRepository()

	mockDb.On("Query", mock.Anything, mock.Anything).Return(&[]int{}, nil)

	_, err := repo.GetPublicParties()

	assert.Error(t, err)
	mockDb.AssertExpectations(t)
}

func TestPartyRepository_GetPartiesByOrganizerId_Success(t *testing.T) {
	repo, mockDb := setupPartyRepository()
	expectedParties := &[]domains.Party{{Model: gorm.Model{ID: 1}}}

	mockDb.On("Query", []db.QueryParameter{
		{Field: "organizer_id", Operator: "=", Value: uint(1)},
	}, mock.Anything).Return(expectedParties, nil)

	result, err := repo.GetPartiesByOrganizerId(1)

	assert.NoError(t, err)
	assert.Equal(t, expectedParties, result)
	mockDb.AssertExpectations(t)
}

func TestPartyRepository_GetPartiesByParticipantId_Success(t *testing.T) {
	repo, mockDb := setupPartyRepository()
	expectedParties := &[]domains.Party{{Model: gorm.Model{ID: 1}}}
	queryParam := db.Many2ManyQueryParameter{
		QueriedTable:            "parties",
		Many2ManyTable:          "party_participants",
		M2MQueriedColumnName:    "party_id",
		M2MConditionColumnName:  "user_id",
		M2MConditionColumnValue: uint(1),
	}

	mockDb.On("Many2ManyQueryId", queryParam, mock.Anything).Return(expectedParties, nil)

	result, err := repo.GetPartiesByParticipantId(1)

	assert.NoError(t, err)
	assert.Equal(t, expectedParties, result)
	mockDb.AssertExpectations(t)
}

func TestPartyRepository_CreateParty_Success(t *testing.T) {
	repo, mockDb := setupPartyRepository()
	party := &domains.Party{}

	mockDb.On("Create", party).Return(nil)

	err := repo.Create(party)

	assert.NoError(t, err)
	mockDb.AssertExpectations(t)
}

func TestPartyRepository_CreateParty_Error(t *testing.T) {
	repo, mockDb := setupPartyRepository()
	party := &domains.Party{}
	expectedErr := errors.New("create error")

	mockDb.On("Create", party).Return(expectedErr)

	err := repo.Create(party)

	assert.Equal(t, expectedErr, err)
	mockDb.AssertExpectations(t)
}

func TestPartyRepository_FindById_Success(t *testing.T) {
	repo, mockDb := setupPartyRepository()
	expectedParty := &domains.Party{Model: gorm.Model{ID: 1}}

	mockDb.On("FindById", uint(1), mock.Anything).Return(expectedParty, nil)

	result, err := repo.FindById(1)

	assert.NoError(t, err)
	assert.Equal(t, expectedParty, result)
	mockDb.AssertExpectations(t)
}

func TestPartyRepository_FindById_TypeAssertionError(t *testing.T) {
	repo, mockDb := setupPartyRepository()

	mockDb.On("FindById", mock.Anything, mock.Anything).Return(&userDomain.User{}, nil)

	_, err := repo.FindById(1)

	assert.Error(t, err)
	mockDb.AssertExpectations(t)
}

func TestPartyRepository_UpdateParty_Success(t *testing.T) {
	repo, mockDb := setupPartyRepository()
	party := &domains.Party{Model: gorm.Model{ID: 1}}

	mockDb.On("Update", mock.Anything, party.ID).Return(nil)

	err := repo.Update(party)

	assert.NoError(t, err)
	mockDb.AssertExpectations(t)
}

func TestPartyRepository_DeleteParty_Success(t *testing.T) {
	repo, mockDb := setupPartyRepository()
	party := &domains.Party{}

	mockDb.On("Delete", party).Return(nil)

	err := repo.Delete(party)

	assert.NoError(t, err)
	mockDb.AssertExpectations(t)
}
