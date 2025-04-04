package usecases

import (
	"github.com/stretchr/testify/mock"
	"github.com/zsomborCzaban/party_organizer/services/user/domains"
	"github.com/zsomborCzaban/party_organizer/utils/api"
	"mime/multipart"
)

type MockService struct {
	mock.Mock
}

func (m *MockService) Login(req domains.LoginRequest) api.IResponse {
	args := m.Called(req)
	return args.Get(0).(api.IResponse)
}

func (m *MockService) Register(req domains.RegisterRequest) api.IResponse {
	args := m.Called(req)
	return args.Get(0).(api.IResponse)
}

func (m *MockService) AddFriend(friendId, userId uint) api.IResponse {
	args := m.Called(friendId, userId)
	return args.Get(0).(api.IResponse)
}

func (m *MockService) GetFriends(userId uint) api.IResponse {
	args := m.Called(userId)
	return args.Get(0).(api.IResponse)
}

func (m *MockService) UploadProfilePicture(userId uint, file multipart.File, header *multipart.FileHeader) api.IResponse {
	args := m.Called(userId, file, header)
	return args.Get(0).(api.IResponse)
}
