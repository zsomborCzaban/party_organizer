package usecases

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zsomborCzaban/party_organizer/services/managers/friend_manager/domains"
	userDomain "github.com/zsomborCzaban/party_organizer/services/user/domains"
	userUsecases "github.com/zsomborCzaban/party_organizer/services/user/usecases"
	"github.com/zsomborCzaban/party_organizer/utils/api"
	"github.com/zsomborCzaban/party_organizer/utils/repo"
	"gorm.io/gorm"
	"testing"
)

func setupDefaultService() (FriendInviteService, *MockRepository, *userUsecases.MockRepository) {
	friendInviteRepo := new(MockRepository)
	userRepo := new(userUsecases.MockRepository)

	service := FriendInviteService{
		FriendInviteRepository: friendInviteRepo,
		UserRepository:         userRepo,
	}

	return service, friendInviteRepo, userRepo
}

func Test_FriendInviteService_Accept_Success(t *testing.T) {
	service, friendRepo, userRepo := setupDefaultService()

	invitorId := uint(1)
	userId := uint(2)
	invite := &domains.FriendInvite{
		InvitorId: invitorId,
		InvitedId: userId,
		State:     domains.PENDING,
	}
	invitor := &userDomain.User{Model: gorm.Model{ID: invitorId}}
	user := &userDomain.User{Model: gorm.Model{ID: userId}}

	friendRepo.On("FindByIds", invitorId, userId).Return(invite, nil)
	userRepo.On("FindById", invitorId, mock.Anything, mock.Anything).Return(invitor, nil)
	userRepo.On("FindById", userId, mock.Anything).Return(user, nil)
	friendRepo.On("Update", invite).Return(nil)
	userRepo.On("AddFriend", invitor, user).Return(nil)

	response := service.Accept(invitorId, userId)

	assert.False(t, response.GetIsError())
	friendRepo.AssertExpectations(t)
	userRepo.AssertExpectations(t)
}

func Test_FriendInviteService_Accept_FailOnFindInvite(t *testing.T) {
	service, friendRepo, _ := setupDefaultService()

	invitorId := uint(1)
	userId := uint(2)
	expectedErr := errors.New("not found")

	friendRepo.On("FindByIds", invitorId, userId).Return(&domains.FriendInvite{}, expectedErr)

	response := service.Accept(invitorId, userId)

	assert.Equal(t, api.ErrorInternalServerError(expectedErr.Error()), response)
}

func Test_FriendInviteService_Accept_AlreadyDeclined(t *testing.T) {
	service, friendRepo, _ := setupDefaultService()

	invitorId := uint(1)
	userId := uint(2)
	invite := &domains.FriendInvite{
		InvitorId: invitorId,
		InvitedId: userId,
		State:     domains.DECLINED,
	}

	friendRepo.On("FindByIds", invitorId, userId).Return(invite, nil)

	response := service.Accept(invitorId, userId)

	assert.Equal(t, api.ErrorBadRequest("Cannot accept already declined friends. Try inviting them"), response)
}

func Test_FriendInviteService_Accept_AlreadyAccepted(t *testing.T) {
	service, friendRepo, _ := setupDefaultService()

	invitorId := uint(1)
	userId := uint(2)
	invite := &domains.FriendInvite{
		InvitorId: invitorId,
		InvitedId: userId,
		State:     domains.ACCEPTED,
	}

	friendRepo.On("FindByIds", invitorId, userId).Return(invite, nil)

	response := service.Accept(invitorId, userId)

	assert.False(t, response.GetIsError())
}

func Test_FriendInviteService_Accept_FailOnFindInvitor(t *testing.T) {
	service, friendRepo, userRepo := setupDefaultService()

	invitorId := uint(1)
	userId := uint(2)
	invite := &domains.FriendInvite{
		InvitorId: invitorId,
		InvitedId: userId,
		State:     domains.PENDING,
	}
	expectedErr := errors.New("not found")

	friendRepo.On("FindByIds", invitorId, userId).Return(invite, nil)
	userRepo.On("FindById", invitorId, mock.Anything, mock.Anything).Return(&userDomain.User{}, expectedErr)

	response := service.Accept(invitorId, userId)

	assert.Equal(t, api.ErrorBadRequest(expectedErr.Error()), response)
}

func Test_FriendInviteService_Accept_FailOnFindUser(t *testing.T) {
	service, friendRepo, userRepo := setupDefaultService()

	invitorId := uint(1)
	userId := uint(2)
	invite := &domains.FriendInvite{
		InvitorId: invitorId,
		InvitedId: userId,
		State:     domains.PENDING,
	}
	invitor := &userDomain.User{Model: gorm.Model{ID: invitorId}}
	expectedErr := errors.New("not found")

	friendRepo.On("FindByIds", invitorId, userId).Return(invite, nil)
	userRepo.On("FindById", invitorId, mock.Anything, mock.Anything).Return(invitor, nil)
	userRepo.On("FindById", userId, mock.Anything).Return(&userDomain.User{}, expectedErr)

	response := service.Accept(invitorId, userId)

	assert.Equal(t, api.ErrorBadRequest(expectedErr.Error()), response)
}

func Test_FriendInviteService_Accept_FailOnUpdate(t *testing.T) {
	service, friendRepo, userRepo := setupDefaultService()

	invitorId := uint(1)
	userId := uint(2)
	invite := &domains.FriendInvite{
		InvitorId: invitorId,
		InvitedId: userId,
		State:     domains.PENDING,
	}
	invitor := &userDomain.User{Model: gorm.Model{ID: invitorId}}
	user := &userDomain.User{Model: gorm.Model{ID: userId}}
	expectedErr := errors.New("update failed")

	friendRepo.On("FindByIds", invitorId, userId).Return(invite, nil)
	userRepo.On("FindById", invitorId, mock.Anything).Return(invitor, nil)
	userRepo.On("FindById", userId, mock.Anything).Return(user, nil)
	friendRepo.On("Update", invite).Return(expectedErr)

	response := service.Accept(invitorId, userId)

	assert.Equal(t, api.ErrorInternalServerError(expectedErr.Error()), response)
}

func Test_FriendInviteService_Accept_FailOnAddFriend(t *testing.T) {
	service, friendRepo, userRepo := setupDefaultService()

	invitorId := uint(1)
	userId := uint(2)
	invite := &domains.FriendInvite{
		InvitorId: invitorId,
		InvitedId: userId,
		State:     domains.PENDING,
	}
	invitor := &userDomain.User{Model: gorm.Model{ID: invitorId}}
	user := &userDomain.User{Model: gorm.Model{ID: userId}}
	expectedErr := errors.New("add friend failed")

	friendRepo.On("FindByIds", invitorId, userId).Return(invite, nil)
	userRepo.On("FindById", invitorId, mock.Anything).Return(invitor, nil)
	userRepo.On("FindById", userId, mock.Anything).Return(user, nil)
	friendRepo.On("Update", invite).Return(nil)
	userRepo.On("AddFriend", invitor, user).Return(expectedErr)

	response := service.Accept(invitorId, userId)

	assert.Equal(t, api.ErrorInternalServerError(expectedErr.Error()), response)
}

func Test_FriendInviteService_Decline_Success(t *testing.T) {
	service, friendRepo, _ := setupDefaultService()

	invitorId := uint(1)
	userId := uint(2)
	invite := &domains.FriendInvite{
		InvitorId: invitorId,
		InvitedId: userId,
		State:     domains.PENDING,
	}

	friendRepo.On("FindByIds", invitorId, userId).Return(invite, nil)
	friendRepo.On("Update", invite).Return(nil)

	response := service.Decline(invitorId, userId)

	assert.False(t, response.GetIsError())
}

func Test_FriendInviteService_Decline_AlreadyAccepted(t *testing.T) {
	service, friendRepo, _ := setupDefaultService()

	invitorId := uint(1)
	userId := uint(2)
	invite := &domains.FriendInvite{
		InvitorId: invitorId,
		InvitedId: userId,
		State:     domains.ACCEPTED,
	}

	friendRepo.On("FindByIds", invitorId, userId).Return(invite, nil)

	response := service.Decline(invitorId, userId)

	assert.Equal(t, api.ErrorBadRequest("Cannot decline already accepted friends. Try deleting them"), response)
}

func Test_FriendInviteService_Decline_AlreadyDeclined(t *testing.T) {
	service, friendRepo, _ := setupDefaultService()

	invitorId := uint(1)
	userId := uint(2)
	invite := &domains.FriendInvite{
		InvitorId: invitorId,
		InvitedId: userId,
		State:     domains.DECLINED,
	}

	friendRepo.On("FindByIds", invitorId, userId).Return(invite, nil)

	response := service.Decline(invitorId, userId)

	assert.False(t, response.GetIsError())
}

func Test_FriendInviteService_Decline_FailOnFind(t *testing.T) {
	service, friendRepo, _ := setupDefaultService()

	invitorId := uint(1)
	userId := uint(2)
	expectedErr := errors.New("not found")

	friendRepo.On("FindByIds", invitorId, userId).Return(&domains.FriendInvite{}, expectedErr)

	response := service.Decline(invitorId, userId)

	assert.Equal(t, api.ErrorInternalServerError(expectedErr.Error()), response)
}

func Test_FriendInviteService_Decline_FailOnUpdate(t *testing.T) {
	service, friendRepo, _ := setupDefaultService()

	invitorId := uint(1)
	userId := uint(2)
	invite := &domains.FriendInvite{
		InvitorId: invitorId,
		InvitedId: userId,
		State:     domains.PENDING,
	}
	expectedErr := errors.New("update failed")

	friendRepo.On("FindByIds", invitorId, userId).Return(invite, nil)
	friendRepo.On("Update", invite).Return(expectedErr)

	response := service.Decline(invitorId, userId)

	assert.Equal(t, api.ErrorInternalServerError(expectedErr.Error()), response)
}

func Test_FriendInviteService_Invite_Success(t *testing.T) {
	service, friendRepo, userRepo := setupDefaultService()

	invitedUsername := "friend"
	userId := uint(1)
	invited := &userDomain.User{
		Model:    gorm.Model{ID: 2},
		Username: invitedUsername,
	}

	userRepo.On("FindByUsername", invitedUsername).Return(invited, nil)
	friendRepo.On("FindByIds", userId, invited.ID).Return(&domains.FriendInvite{}, errors.New(domains.NOT_FOUND))
	friendRepo.On("FindByIds", invited.ID, userId).Return(&domains.FriendInvite{}, errors.New(domains.NOT_FOUND))
	userRepo.On("FindById", userId, mock.Anything).Return(&userDomain.User{Model: gorm.Model{ID: userId}}, nil)
	userRepo.On("FindById", invited.ID, mock.Anything).Return(invited, nil)
	friendRepo.On("Create", mock.AnythingOfType("*domains.FriendInvite")).Return(nil)

	response := service.Invite(invitedUsername, userId)

	assert.False(t, response.GetIsError())
}

func Test_FriendInviteService_Invite_SelfInvite(t *testing.T) {
	service, _, userRepo := setupDefaultService()

	invitedUsername := "myself"
	userId := uint(1)
	invited := &userDomain.User{
		Model:    gorm.Model{ID: userId},
		Username: invitedUsername,
	}

	userRepo.On("FindByUsername", invitedUsername).Return(invited, nil)

	response := service.Invite(invitedUsername, userId)

	assert.Equal(t, api.ErrorBadRequest("cannot friend invite yourself"), response)
}

func Test_FriendInviteService_Invite_ReverseInviteExists(t *testing.T) {
	service, friendRepo, userRepo := setupDefaultService()

	invitedUsername := "friend"
	userId := uint(1)
	invitedId := uint(2)
	invited := &userDomain.User{
		Model:    gorm.Model{ID: invitedId},
		Username: invitedUsername,
	}
	reverseInvite := &domains.FriendInvite{
		InvitorId: invitedId,
		InvitedId: userId,
		State:     domains.DECLINED,
	}

	userRepo.On("FindByUsername", invitedUsername).Return(invited, nil)
	friendRepo.On("FindByIds", invitedId, userId).Return(reverseInvite, nil)
	friendRepo.On("Update", reverseInvite).Return(nil)

	response := service.Invite(invitedUsername, userId)

	assert.False(t, response.GetIsError())
}

func Test_FriendInviteService_Invite_ReverseInvitePending(t *testing.T) {
	service, friendRepo, userRepo := setupDefaultService()

	invitedUsername := "friend"
	userId := uint(1)
	invitedId := uint(2)
	invited := &userDomain.User{
		Model:    gorm.Model{ID: invitedId},
		Username: invitedUsername,
	}
	reverseInvite := &domains.FriendInvite{
		InvitorId: invitedId,
		InvitedId: userId,
		State:     domains.PENDING,
	}

	userRepo.On("FindByUsername", invitedUsername).Return(invited, nil)
	friendRepo.On("FindByIds", invitedId, userId).Return(reverseInvite, nil)

	response := service.Invite(invitedUsername, userId)

	assert.Equal(t, api.ErrorBadRequest("friend request already exists, try accepting it"), response)
}

func Test_FriendInviteService_Invite_ExistingInvitePending(t *testing.T) {
	service, friendRepo, userRepo := setupDefaultService()

	invitedUsername := "friend"
	userId := uint(1)
	invitedId := uint(2)
	invited := &userDomain.User{
		Model:    gorm.Model{ID: invitedId},
		Username: invitedUsername,
	}
	existingInvite := &domains.FriendInvite{
		InvitorId: userId,
		InvitedId: invitedId,
		State:     domains.PENDING,
	}

	userRepo.On("FindByUsername", invitedUsername).Return(invited, nil)
	friendRepo.On("FindByIds", invitedId, userId).Return(&domains.FriendInvite{}, errors.New(domains.NOT_FOUND))
	friendRepo.On("FindByIds", userId, invitedId).Return(existingInvite, nil)

	response := service.Invite(invitedUsername, userId)

	assert.False(t, response.GetIsError())
}

func Test_FriendInviteService_Invite_ExistingInviteAccepted(t *testing.T) {
	service, friendRepo, userRepo := setupDefaultService()

	invitedUsername := "friend"
	userId := uint(1)
	invitedId := uint(2)
	invited := &userDomain.User{
		Model:    gorm.Model{ID: invitedId},
		Username: invitedUsername,
	}
	existingInvite := &domains.FriendInvite{
		InvitorId: userId,
		InvitedId: invitedId,
		State:     domains.ACCEPTED,
	}

	userRepo.On("FindByUsername", invitedUsername).Return(invited, nil)
	friendRepo.On("FindByIds", invitedId, userId).Return(&domains.FriendInvite{}, errors.New(domains.NOT_FOUND))
	friendRepo.On("FindByIds", userId, invitedId).Return(existingInvite, nil)

	response := service.Invite(invitedUsername, userId)

	assert.Equal(t, api.ErrorBadRequest("User already accepted the invite"), response)
}

func Test_FriendInviteService_Invite_ExistingInviteDeclined(t *testing.T) {
	service, friendRepo, userRepo := setupDefaultService()

	invitedUsername := "friend"
	userId := uint(1)
	invitedId := uint(2)
	invited := &userDomain.User{
		Model:    gorm.Model{ID: invitedId},
		Username: invitedUsername,
	}
	existingInvite := &domains.FriendInvite{
		InvitorId: userId,
		InvitedId: invitedId,
		State:     domains.DECLINED,
	}

	userRepo.On("FindByUsername", invitedUsername).Return(invited, nil)
	friendRepo.On("FindByIds", invitedId, userId).Return(&domains.FriendInvite{}, errors.New(domains.NOT_FOUND))
	friendRepo.On("FindByIds", userId, invitedId).Return(existingInvite, nil)
	friendRepo.On("Update", existingInvite).Return(nil)

	response := service.Invite(invitedUsername, userId)

	assert.False(t, response.GetIsError())
}

func Test_FriendInviteService_Invite_FailOnFindUser(t *testing.T) {
	service, _, userRepo := setupDefaultService()

	invitedUsername := "friend"
	userId := uint(1)
	expectedErr := errors.New("not found")

	userRepo.On("FindByUsername", invitedUsername).Return(&userDomain.User{}, expectedErr)

	response := service.Invite(invitedUsername, userId)

	assert.Equal(t, api.ErrorBadRequest(expectedErr.Error()), response)
}

func Test_FriendInviteService_CreateInvitation_Success(t *testing.T) {
	service, friendRepo, userRepo := setupDefaultService()

	invitedId := uint(2)
	userId := uint(1)
	invitor := &userDomain.User{Model: gorm.Model{ID: userId}}
	invited := &userDomain.User{Model: gorm.Model{ID: invitedId}}

	userRepo.On("FindById", userId, mock.Anything).Return(invitor, nil)
	userRepo.On("FindById", invitedId, mock.Anything, mock.Anything).Return(invited, nil)
	friendRepo.On("Create", mock.AnythingOfType("*domains.FriendInvite")).Return(nil)

	response := service.CreateInvitation(invitedId, userId)

	assert.False(t, response.GetIsError())
}

func Test_FriendInviteService_CreateInvitation_FailOnFindInvitor(t *testing.T) {
	service, _, userRepo := setupDefaultService()

	invitedId := uint(2)
	userId := uint(1)
	expectedErr := errors.New("not found")

	userRepo.On("FindById", userId, mock.Anything).Return(&userDomain.User{}, expectedErr)

	response := service.CreateInvitation(invitedId, userId)

	assert.Equal(t, api.ErrorBadRequest(fmt.Sprintf("cannot find user with id: %d", userId)), response)
}

func Test_FriendInviteService_CreateInvitation_FailOnFindInvited(t *testing.T) {
	service, _, userRepo := setupDefaultService()

	invitedId := uint(2)
	userId := uint(1)
	invitor := &userDomain.User{Model: gorm.Model{ID: userId}}
	expectedErr := errors.New("not found")

	userRepo.On("FindById", userId, mock.Anything).Return(invitor, nil)
	userRepo.On("FindById", invitedId, mock.Anything).Return(&userDomain.User{}, expectedErr)

	response := service.CreateInvitation(invitedId, userId)

	assert.Equal(t, api.ErrorBadRequest(fmt.Sprintf("cannot find user with id: %d", userId)), response)
}

func Test_FriendInviteService_CreateInvitation_FailOnCreate(t *testing.T) {
	service, friendRepo, userRepo := setupDefaultService()

	invitedId := uint(2)
	userId := uint(1)
	invitor := &userDomain.User{Model: gorm.Model{ID: userId}}
	invited := &userDomain.User{Model: gorm.Model{ID: invitedId}}
	expectedErr := errors.New("create failed")

	userRepo.On("FindById", userId, mock.Anything).Return(invitor, nil)
	userRepo.On("FindById", invitedId, mock.Anything).Return(invited, nil)
	friendRepo.On("Create", mock.AnythingOfType("*domains.FriendInvite")).Return(expectedErr)

	response := service.CreateInvitation(invitedId, userId)

	assert.Equal(t, api.ErrorInternalServerError(expectedErr.Error()), response)
}

func Test_FriendInviteService_ReverseInvite_Success(t *testing.T) {
	service, friendRepo, _ := setupDefaultService()

	invitation := &domains.FriendInvite{
		InvitorId: 1,
		InvitedId: 2,
		State:     domains.DECLINED,
	}

	friendRepo.On("Update", invitation).Return(nil)

	response := service.ReverseInvite(invitation)

	assert.False(t, response.GetIsError())
	assert.Equal(t, uint(2), invitation.InvitorId)
	assert.Equal(t, uint(1), invitation.InvitedId)
	assert.Equal(t, domains.PENDING, invitation.State)
}

func Test_FriendInviteService_ReverseInvite_FailOnUpdate(t *testing.T) {
	service, friendRepo, _ := setupDefaultService()

	invitation := &domains.FriendInvite{
		InvitorId: 1,
		InvitedId: 2,
		State:     domains.DECLINED,
	}
	expectedErr := errors.New("update failed")

	friendRepo.On("Update", invitation).Return(expectedErr)

	response := service.ReverseInvite(invitation)

	assert.Equal(t, api.ErrorInternalServerError(expectedErr.Error()), response)
}

func Test_FriendInviteService_GetPendingInvites_Success(t *testing.T) {
	service, friendRepo, _ := setupDefaultService()

	userId := uint(1)
	invites := []domains.FriendInvite{
		{InvitorId: 2, InvitedId: userId, State: domains.PENDING},
	}

	friendRepo.On("FindPendingByInvitedId", userId).Return(&invites, nil)

	response := service.GetPendingInvites(userId)

	assert.False(t, response.GetIsError())
}

func Test_FriendInviteService_GetPendingInvites_Fail(t *testing.T) {
	service, friendRepo, _ := setupDefaultService()

	userId := uint(1)
	expectedErr := errors.New("not found")

	friendRepo.On("FindPendingByInvitedId", userId).Return(&[]domains.FriendInvite{}, expectedErr)

	response := service.GetPendingInvites(userId)

	assert.Equal(t, api.ErrorInternalServerError(expectedErr.Error()), response)
}

func Test_FriendInviteService_RemoveFriend_Success(t *testing.T) {
	service, friendRepo, userRepo := setupDefaultService()

	userId := uint(1)
	friendId := uint(2)
	user := &userDomain.User{
		Model: gorm.Model{ID: userId},
	}
	friend := &userDomain.User{
		Model: gorm.Model{ID: friendId},
	}
	invite := &domains.FriendInvite{
		InvitorId: userId,
		InvitedId: friendId,
		State:     domains.ACCEPTED,
	}

	userRepo.On("FindById", userId, mock.Anything).Return(user, nil)
	userRepo.On("FindById", friendId, mock.Anything).Return(friend, nil)
	friendRepo.On("FindByIds", userId, friendId).Return(invite, nil)
	userRepo.On("RemoveFriend", user, friend).Return(nil)
	friendRepo.On("Update", invite).Return(nil)

	response := service.RemoveFriend(userId, friendId)

	assert.False(t, response.GetIsError())
}

func Test_FriendInviteService_RemoveFriend_AlreadyNotFriends(t *testing.T) {
	service, _, userRepo := setupDefaultService()

	userId := uint(1)
	friendId := uint(2)
	user := &userDomain.User{
		Model: gorm.Model{ID: userId},
	}
	friend := &userDomain.User{
		Model: gorm.Model{ID: friendId},
	}

	userRepo.On("FindById", userId, mock.Anything).Return(user, nil)
	userRepo.On("FindById", friendId, mock.Anything).Return(friend, nil)

	response := service.RemoveFriend(userId, friendId)

	assert.False(t, response.GetIsError())
}

func Test_FriendInviteService_RemoveFriend_FailOnFindUser(t *testing.T) {
	service, _, userRepo := setupDefaultService()

	userId := uint(1)
	friendId := uint(2)
	expectedErr := errors.New("not found")

	userRepo.On("FindById", userId, mock.Anything).Return(&userDomain.User{}, expectedErr)

	response := service.RemoveFriend(userId, friendId)

	assert.Equal(t, api.ErrorBadRequest(expectedErr.Error()), response)
}

func Test_FriendInviteService_RemoveFriend_FailOnFindFriend(t *testing.T) {
	service, _, userRepo := setupDefaultService()

	userId := uint(1)
	friendId := uint(2)
	user := &userDomain.User{
		Model: gorm.Model{ID: userId},
	}
	expectedErr := errors.New("not found")

	userRepo.On("FindById", userId, mock.Anything).Return(user, nil)
	userRepo.On("FindById", friendId, mock.Anything).Return(&userDomain.User{}, expectedErr)

	response := service.RemoveFriend(userId, friendId)

	assert.Equal(t, api.ErrorBadRequest(expectedErr.Error()), response)
}

func Test_FriendInviteService_RemoveFriendAndInvite_Success(t *testing.T) {
	service, friendRepo, userRepo := setupDefaultService()

	user := &userDomain.User{Model: gorm.Model{ID: 1}}
	friend := &userDomain.User{Model: gorm.Model{ID: 2}}
	invite := &domains.FriendInvite{
		State: domains.ACCEPTED,
	}

	userRepo.On("RemoveFriend", user, friend).Return(nil)
	friendRepo.On("Update", invite).Return(nil)

	response := service.RemoveFriendAndInvite(user, friend, invite)

	assert.False(t, response.GetIsError())
	assert.Equal(t, domains.DECLINED, invite.State)
}

func Test_FriendInviteService_RemoveFriendAndInvite_FailOnRemoveFriend(t *testing.T) {
	service, _, userRepo := setupDefaultService()

	user := &userDomain.User{Model: gorm.Model{ID: 1}}
	friend := &userDomain.User{Model: gorm.Model{ID: 2}}
	invite := &domains.FriendInvite{}
	expectedErr := errors.New("remove failed")

	userRepo.On("RemoveFriend", user, friend).Return(expectedErr)

	response := service.RemoveFriendAndInvite(user, friend, invite)

	assert.Equal(t, api.ErrorInternalServerError(expectedErr.Error()), response)
}

func Test_FriendInviteService_RemoveFriendAndInvite_FailOnUpdateInvite(t *testing.T) {
	service, friendRepo, userRepo := setupDefaultService()

	user := &userDomain.User{Model: gorm.Model{ID: 1}}
	friend := &userDomain.User{Model: gorm.Model{ID: 2}}
	invite := &domains.FriendInvite{}
	expectedErr := errors.New("update failed")

	userRepo.On("RemoveFriend", user, friend).Return(nil)
	friendRepo.On("Update", invite).Return(expectedErr)

	response := service.RemoveFriendAndInvite(user, friend, invite)

	assert.Equal(t, api.ErrorInternalServerError(expectedErr.Error()), response)
}

func Test_NewService(t *testing.T) {
	friendInviteRepo := new(MockRepository)
	userRepo := new(userUsecases.MockRepository)
	repoCollector := repo.RepoCollector{
		UserRepo:         userRepo,
		FriendInviteRepo: friendInviteRepo,
	}

	serviceInterface := NewFriendInviteService(&repoCollector)
	service := serviceInterface.(*FriendInviteService)

	assert.Equal(t, service.FriendInviteRepository, friendInviteRepo)
	assert.Equal(t, service.UserRepository, userRepo)
}
