package usecases

import (
	"github.com/stretchr/testify/mock"
	"github.com/zsomborCzaban/party_organizer/services/users/user/domains"
)

type MockRepository struct {
	mock.Mock
}

func (mr *MockRepository) FindByUsername(name string) (*domains.User, error) {
	args := mr.Called(name)
	return args.Get(0).(*domains.User), args.Error(1)
}

func (mr *MockRepository) FindById(id uint, associations ...string) (*domains.User, error) {
	args := mr.Called(id, associations)
	return args.Get(0).(*domains.User), args.Error(1)
}
func (mr *MockRepository) CreateUser(user *domains.User) error {
	args := mr.Called(user)
	return args.Error(0)
}
func (mr *MockRepository) UpdateUser(user *domains.User) error {
	args := mr.Called(user)
	return args.Error(0)
}
func (mr *MockRepository) AddFriend(user *domains.User, friend *domains.User) error {
	args := mr.Called(user, friend)
	return args.Error(0)
}
func (mr *MockRepository) RemoveFriend(user *domains.User, friend *domains.User) error {
	args := mr.Called(user, friend)
	return args.Error(0)
}
func (mr *MockRepository) GetFriends(userId uint) (*[]domains.User, error) {
	args := mr.Called(userId)
	return args.Get(0).(*[]domains.User), args.Error(1)
}
