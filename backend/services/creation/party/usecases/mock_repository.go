package usecases

import (
	"github.com/stretchr/testify/mock"
	"github.com/zsomborCzaban/party_organizer/services/creation/party/domains"
	userDomain "github.com/zsomborCzaban/party_organizer/services/users/user/domains"
)

type MockRepository struct {
	mock.Mock
}

func (mr *MockRepository) GetPartiesByOrganizerId(uint) (*[]domains.Party, error) {
	args := mr.Called()
	return args.Get(0).(*[]domains.Party), args.Error(1)
}
func (mr *MockRepository) GetPartiesByParticipantId(uint) (*[]domains.Party, error) {
	args := mr.Called()
	return args.Get(0).(*[]domains.Party), args.Error(1)
}
func (mr *MockRepository) GetPublicParties() (*[]domains.Party, error) {
	args := mr.Called()
	return args.Get(0).(*[]domains.Party), args.Error(1)
}
func (mr *MockRepository) AddUserToParty(*domains.Party, *userDomain.User) error {
	args := mr.Called()
	return args.Error(0)
}
func (mr *MockRepository) RemoveUserFromParty(*domains.Party, *userDomain.User) error {
	args := mr.Called()
	return args.Error(0)
}
func (mr *MockRepository) Create(*domains.Party) error {
	args := mr.Called()
	return args.Error(0)
}
func (mr *MockRepository) FindById(uint, ...string) (*domains.Party, error) {
	args := mr.Called()
	return args.Get(0).(*domains.Party), args.Error(1)
}
func (mr *MockRepository) Update(*domains.Party) error {
	args := mr.Called()
	return args.Error(0)
}
func (mr *MockRepository) Delete(*domains.Party) error {
	args := mr.Called()
	return args.Error(0)
}
