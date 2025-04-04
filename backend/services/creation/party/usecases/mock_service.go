package usecases

import (
	"github.com/stretchr/testify/mock"
	"github.com/zsomborCzaban/party_organizer/services/creation/party/domains"
	"github.com/zsomborCzaban/party_organizer/utils/api"
)

type MockService struct {
	mock.Mock
}

func (m *MockService) GetParticipants(partyId, userId uint) api.IResponse {
	args := m.Called(partyId, userId)
	return args.Get(0).(api.IResponse)
}

func (m *MockService) GetPublicParties() api.IResponse {
	args := m.Called()
	return args.Get(0).(api.IResponse)
}

func (m *MockService) GetPublicParty(partyId uint) api.IResponse {
	args := m.Called(partyId)
	return args.Get(0).(api.IResponse)
}

func (m *MockService) GetPartiesByOrganizerId(userId uint) api.IResponse {
	args := m.Called(userId)
	return args.Get(0).(api.IResponse)
}

func (m *MockService) GetPartiesByParticipantId(userId uint) api.IResponse {
	args := m.Called(userId)
	return args.Get(0).(api.IResponse)
}

func (m *MockService) Create(partyDTO domains.PartyDTO, userId uint) api.IResponse {
	args := m.Called(partyDTO, userId)
	return args.Get(0).(api.IResponse)
}

func (m *MockService) Update(partyDTO domains.PartyDTO, userId uint) api.IResponse {
	args := m.Called(partyDTO, userId)
	return args.Get(0).(api.IResponse)
}

func (m *MockService) Get(partyId, userId uint) api.IResponse {
	args := m.Called(partyId, userId)
	return args.Get(0).(api.IResponse)
}

func (m *MockService) Delete(partyId, userId uint) api.IResponse {
	args := m.Called(partyId, userId)
	return args.Get(0).(api.IResponse)
}
