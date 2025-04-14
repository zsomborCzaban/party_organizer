package usecases

import (
	"github.com/stretchr/testify/mock"
	domains2 "github.com/zsomborCzaban/party_organizer/services/users/user/domains"
	"github.com/zsomborCzaban/party_organizer/utils/api"
	"mime/multipart"
)

type MockService struct {
	mock.Mock
}

func (m *MockService) Login(req domains2.LoginRequest) api.IResponse {
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
