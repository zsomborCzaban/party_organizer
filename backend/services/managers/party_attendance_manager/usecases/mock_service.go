package usecases

import (
	"github.com/stretchr/testify/mock"
	"github.com/zsomborCzaban/party_organizer/utils/api"
)

type MockService struct {
	mock.Mock
}

func (m *MockService) Invite(invitedId string, invitorId, partyId uint) api.IResponse {
	args := m.Called(invitedId, invitorId, partyId)
	return args.Get(0).(api.IResponse)
}

func (m *MockService) Accept(invitedId, partyId uint) api.IResponse {
	args := m.Called(invitedId, partyId)
	return args.Get(0).(api.IResponse)
}

func (m *MockService) Decline(invitedId, partyId uint) api.IResponse {
	args := m.Called(invitedId, partyId)
	return args.Get(0).(api.IResponse)
}

func (m *MockService) GetUserPendingInvites(userId uint) api.IResponse {
	args := m.Called(userId)
	return args.Get(0).(api.IResponse)
}

func (m *MockService) GetPartyPendingInvites(partyId, userId uint) api.IResponse {
	args := m.Called(partyId, userId)
	return args.Get(0).(api.IResponse)
}

func (m *MockService) Kick(kickedId, userId, partyId uint) api.IResponse {
	args := m.Called(kickedId, userId, partyId)
	return args.Get(0).(api.IResponse)
}

func (m *MockService) JoinPublicParty(partyId, userId uint) api.IResponse {
	args := m.Called(partyId, userId)
	return args.Get(0).(api.IResponse)
}

func (m *MockService) JoinPrivateParty(userId uint, accessCode string) api.IResponse {
	args := m.Called(userId, accessCode)
	return args.Get(0).(api.IResponse)
}
