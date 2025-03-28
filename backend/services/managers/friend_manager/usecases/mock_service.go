package usecases

import (
	"github.com/stretchr/testify/mock"
	"github.com/zsomborCzaban/party_organizer/utils/api"
)

type MockService struct {
	mock.Mock
}

func (m *MockService) Invite(invitedUsername string, userId uint) api.IResponse {
	args := m.Called(invitedUsername, userId)
	return args.Get(0).(api.IResponse)
}

func (m *MockService) Accept(invitorId, userId uint) api.IResponse {
	args := m.Called(invitorId, userId)
	return args.Get(0).(api.IResponse)
}

func (m *MockService) Decline(invitorId, userId uint) api.IResponse {
	args := m.Called(invitorId, userId)
	return args.Get(0).(api.IResponse)
}

func (m *MockService) GetPendingInvites(userId uint) api.IResponse {
	args := m.Called(userId)
	return args.Get(0).(api.IResponse)
}

func (m *MockService) RemoveFriend(userId, friendId uint) api.IResponse {
	args := m.Called(userId, friendId)
	return args.Get(0).(api.IResponse)
}
