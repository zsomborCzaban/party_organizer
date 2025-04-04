package usecases

import (
	"github.com/stretchr/testify/mock"
	"github.com/zsomborCzaban/party_organizer/services/managers/friend_manager/domains"
)

type MockRepository struct {
	mock.Mock
}

func (mr *MockRepository) FindByIds(invitorId, invitedId uint) (*domains.FriendInvite, error) {
	args := mr.Called(invitorId, invitedId)
	return args.Get(0).(*domains.FriendInvite), args.Error(1)
}

func (mr *MockRepository) Update(requirement *domains.FriendInvite) error {
	args := mr.Called(requirement)
	return args.Error(0)
}

func (mr *MockRepository) Create(requirement *domains.FriendInvite) error {
	args := mr.Called(requirement)
	return args.Error(0)
}

func (mr *MockRepository) FindPendingByInvitedId(id uint) (*[]domains.FriendInvite, error) {
	args := mr.Called(id)
	return args.Get(0).(*[]domains.FriendInvite), args.Error(1)
}
