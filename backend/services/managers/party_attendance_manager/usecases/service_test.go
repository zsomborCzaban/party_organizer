package usecases

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	partyDomains "github.com/zsomborCzaban/party_organizer/services/creation/party/domains"
	partyUsecases "github.com/zsomborCzaban/party_organizer/services/creation/party/usecases"
	drinkContributionUsecases "github.com/zsomborCzaban/party_organizer/services/interaction/drink_contributions/usecases"
	foodContributionUsecases "github.com/zsomborCzaban/party_organizer/services/interaction/food_contributions/usecases"
	"github.com/zsomborCzaban/party_organizer/services/managers/party_attendance_manager/domains"
	userDomain "github.com/zsomborCzaban/party_organizer/services/user/domains"
	userUsecases "github.com/zsomborCzaban/party_organizer/services/user/usecases"
	"github.com/zsomborCzaban/party_organizer/utils/api"
	"github.com/zsomborCzaban/party_organizer/utils/repo"
	"gorm.io/gorm"
	"testing"
)

func setupDefaultService() (PartyInviteService, *MockRepository, *userUsecases.MockRepository, *partyUsecases.MockRepository, *foodContributionUsecases.MockRepository, *drinkContributionUsecases.MockRepository) {
	partyInviteRepo := new(MockRepository)
	userRepo := new(userUsecases.MockRepository)
	partyRepo := new(partyUsecases.MockRepository)
	foodContribRepo := new(foodContributionUsecases.MockRepository)
	drinkContribRepo := new(drinkContributionUsecases.MockRepository)

	service := PartyInviteService{
		PartyInviteRepository:       partyInviteRepo,
		UserRepository:              userRepo,
		PartyRepository:             partyRepo,
		FoodContributionRepository:  foodContribRepo,
		DrinkContributionRepository: drinkContribRepo,
	}

	return service, partyInviteRepo, userRepo, partyRepo, foodContribRepo, drinkContribRepo
}

func Test_PartyInviteService_Accept_Success(t *testing.T) {
	service, inviteRepo, userRepo, partyRepo, _, _ := setupDefaultService()

	invitedId := uint(1)
	partyId := uint(2)
	invite := &domains.PartyInvite{
		InvitedId: invitedId,
		PartyId:   partyId,
		State:     domains.PENDING,
	}
	invitedUser := &userDomain.User{Model: gorm.Model{ID: invitedId}}
	party := &partyDomains.Party{Model: gorm.Model{ID: partyId}}

	inviteRepo.On("FindByIds", invitedId, partyId).Return(invite, nil)
	userRepo.On("FindById", mock.Anything, mock.Anything).Return(invitedUser, nil)
	partyRepo.On("FindById", mock.Anything, mock.Anything).Return(party, nil)
	inviteRepo.On("Update", invite).Return(nil)
	partyRepo.On("AddUserToParty", mock.Anything, mock.Anything).Return(nil)

	response := service.Accept(invitedId, partyId)

	assert.False(t, response.GetIsError())
	assert.Equal(t, domains.ACCEPTED, invite.State)
}

func Test_PartyInviteService_Accept_AlreadyDeclined(t *testing.T) {
	service, inviteRepo, _, _, _, _ := setupDefaultService()

	invitedId := uint(1)
	partyId := uint(2)
	invite := &domains.PartyInvite{
		InvitedId: invitedId,
		PartyId:   partyId,
		State:     domains.DECLINED,
	}

	inviteRepo.On("FindByIds", invitedId, partyId).Return(invite, nil)

	response := service.Accept(invitedId, partyId)

	assert.Equal(t, api.ErrorBadRequest("Cannot accept already declined parties. Try inviting them"), response)
}

func Test_PartyInviteService_Accept_AlreadyAccepted(t *testing.T) {
	service, inviteRepo, _, _, _, _ := setupDefaultService()

	invitedId := uint(1)
	partyId := uint(2)
	invite := &domains.PartyInvite{
		InvitedId: invitedId,
		PartyId:   partyId,
		State:     domains.ACCEPTED,
	}

	inviteRepo.On("FindByIds", invitedId, partyId).Return(invite, nil)

	response := service.Accept(invitedId, partyId)

	assert.False(t, response.GetIsError())
}

func Test_PartyInviteService_Accept_FailOnFindInvite(t *testing.T) {
	service, inviteRepo, _, _, _, _ := setupDefaultService()

	invitedId := uint(1)
	partyId := uint(2)
	expectedErr := errors.New("not found")

	inviteRepo.On("FindByIds", invitedId, partyId).Return(&domains.PartyInvite{}, expectedErr)

	response := service.Accept(invitedId, partyId)

	assert.Equal(t, api.ErrorInternalServerError(expectedErr.Error()), response)
}

func Test_PartyInviteService_Accept_FailOnFindUser(t *testing.T) {
	service, inviteRepo, userRepo, _, _, _ := setupDefaultService()

	invitedId := uint(1)
	partyId := uint(2)
	invite := &domains.PartyInvite{
		InvitedId: invitedId,
		PartyId:   partyId,
		State:     domains.PENDING,
	}
	expectedErr := errors.New("user not found")

	inviteRepo.On("FindByIds", invitedId, partyId).Return(invite, nil)
	userRepo.On("FindById", mock.Anything, mock.Anything).Return(&userDomain.User{}, expectedErr)

	response := service.Accept(invitedId, partyId)

	assert.Equal(t, api.ErrorInternalServerError(expectedErr.Error()), response)
}

func Test_PartyInviteService_Accept_FailOnFindParty(t *testing.T) {
	service, inviteRepo, userRepo, partyRepo, _, _ := setupDefaultService()

	invitedId := uint(1)
	partyId := uint(2)
	invite := &domains.PartyInvite{
		InvitedId: invitedId,
		PartyId:   partyId,
		State:     domains.PENDING,
	}
	invitedUser := &userDomain.User{Model: gorm.Model{ID: invitedId}}
	expectedErr := errors.New("party not found")

	inviteRepo.On("FindByIds", invitedId, partyId).Return(invite, nil)
	userRepo.On("FindById", mock.Anything, mock.Anything).Return(invitedUser, nil)
	partyRepo.On("FindById", mock.Anything, mock.Anything).Return(&partyDomains.Party{}, expectedErr)

	response := service.Accept(invitedId, partyId)

	assert.Equal(t, api.ErrorBadRequest(expectedErr.Error()), response)
}

func Test_PartyInviteService_Accept_FailOnUpdate(t *testing.T) {
	service, inviteRepo, userRepo, partyRepo, _, _ := setupDefaultService()

	invitedId := uint(1)
	partyId := uint(2)
	invite := &domains.PartyInvite{
		InvitedId: invitedId,
		PartyId:   partyId,
		State:     domains.PENDING,
	}
	invitedUser := &userDomain.User{Model: gorm.Model{ID: invitedId}}
	party := &partyDomains.Party{Model: gorm.Model{ID: partyId}}
	expectedErr := errors.New("update failed")

	inviteRepo.On("FindByIds", invitedId, partyId).Return(invite, nil)
	userRepo.On("FindById", mock.Anything, mock.Anything).Return(invitedUser, nil)
	partyRepo.On("FindById", mock.Anything, mock.Anything).Return(party, nil)
	inviteRepo.On("Update", invite).Return(expectedErr)

	response := service.Accept(invitedId, partyId)

	assert.Equal(t, api.ErrorInternalServerError(expectedErr.Error()), response)
}

func Test_PartyInviteService_Accept_FailOnAddUser(t *testing.T) {
	service, inviteRepo, userRepo, partyRepo, _, _ := setupDefaultService()

	invitedId := uint(1)
	partyId := uint(2)
	invite := &domains.PartyInvite{
		InvitedId: invitedId,
		PartyId:   partyId,
		State:     domains.PENDING,
	}
	invitedUser := &userDomain.User{Model: gorm.Model{ID: invitedId}}
	party := &partyDomains.Party{Model: gorm.Model{ID: partyId}}
	expectedErr := errors.New("add failed")

	inviteRepo.On("FindByIds", invitedId, partyId).Return(invite, nil)
	userRepo.On("FindById", mock.Anything, mock.Anything).Return(invitedUser, nil)
	partyRepo.On("FindById", mock.Anything, mock.Anything).Return(party, nil)
	inviteRepo.On("Update", invite).Return(nil)
	partyRepo.On("AddUserToParty", mock.Anything, mock.Anything).Return(expectedErr)

	response := service.Accept(invitedId, partyId)

	assert.Equal(t, api.ErrorInternalServerError(expectedErr.Error()), response)
}

func Test_PartyInviteService_Decline_Success(t *testing.T) {
	service, inviteRepo, _, _, _, _ := setupDefaultService()

	invitedId := uint(1)
	partyId := uint(2)
	invite := &domains.PartyInvite{
		InvitedId: invitedId,
		PartyId:   partyId,
		State:     domains.PENDING,
	}

	inviteRepo.On("FindByIds", invitedId, partyId).Return(invite, nil)
	inviteRepo.On("Update", invite).Return(nil)

	response := service.Decline(invitedId, partyId)

	assert.False(t, response.GetIsError())
	assert.Equal(t, domains.DECLINED, invite.State)
}

func Test_PartyInviteService_Decline_AlreadyAccepted(t *testing.T) {
	service, inviteRepo, _, _, _, _ := setupDefaultService()

	invitedId := uint(1)
	partyId := uint(2)
	invite := &domains.PartyInvite{
		InvitedId: invitedId,
		PartyId:   partyId,
		State:     domains.ACCEPTED,
	}

	inviteRepo.On("FindByIds", invitedId, partyId).Return(invite, nil)

	response := service.Decline(invitedId, partyId)

	assert.Equal(t, api.ErrorBadRequest("Cannot decline already accepted parties. Try deleting them"), response)
}

func Test_PartyInviteService_Decline_AlreadyDeclined(t *testing.T) {
	service, inviteRepo, _, _, _, _ := setupDefaultService()

	invitedId := uint(1)
	partyId := uint(2)
	invite := &domains.PartyInvite{
		InvitedId: invitedId,
		PartyId:   partyId,
		State:     domains.DECLINED,
	}

	inviteRepo.On("FindByIds", invitedId, partyId).Return(invite, nil)

	response := service.Decline(invitedId, partyId)

	assert.False(t, response.GetIsError())
}

func Test_PartyInviteService_Decline_FailOnFind(t *testing.T) {
	service, inviteRepo, _, _, _, _ := setupDefaultService()

	invitedId := uint(1)
	partyId := uint(2)
	expectedErr := errors.New("not found")

	inviteRepo.On("FindByIds", invitedId, partyId).Return(&domains.PartyInvite{}, expectedErr)

	response := service.Decline(invitedId, partyId)

	assert.Equal(t, api.ErrorInternalServerError(expectedErr.Error()), response)
}

func Test_PartyInviteService_Decline_FailOnUpdate(t *testing.T) {
	service, inviteRepo, _, _, _, _ := setupDefaultService()

	invitedId := uint(1)
	partyId := uint(2)
	invite := &domains.PartyInvite{
		InvitedId: invitedId,
		PartyId:   partyId,
		State:     domains.PENDING,
	}
	expectedErr := errors.New("update failed")

	inviteRepo.On("FindByIds", invitedId, partyId).Return(invite, nil)
	inviteRepo.On("Update", invite).Return(expectedErr)

	response := service.Decline(invitedId, partyId)

	assert.Equal(t, api.ErrorInternalServerError(expectedErr.Error()), response)
}

func Test_PartyInviteService_Invite_Success(t *testing.T) {
	service, inviteRepo, userRepo, partyRepo, _, _ := setupDefaultService()

	invitedUsername := "friend"
	invitorId := uint(1)
	partyId := uint(2)
	invitedUser := &userDomain.User{
		Model:    gorm.Model{ID: 3},
		Username: invitedUsername,
	}

	userRepo.On("FindByUsername", invitedUsername).Return(invitedUser, nil)
	inviteRepo.On("FindByIds", invitedUser.ID, partyId).Return(&domains.PartyInvite{}, errors.New(domains.NOT_FOUND))
	userRepo.On("FindById", invitorId, mock.Anything).Return(&userDomain.User{Model: gorm.Model{ID: invitorId}}, nil)
	partyRepo.On("FindById", mock.Anything, mock.Anything).Return(&partyDomains.Party{OrganizerID: invitorId}, nil)
	userRepo.On("FindById", invitedUser.ID, mock.Anything).Return(invitedUser, nil)
	inviteRepo.On("Create", mock.AnythingOfType("*domains.PartyInvite")).Return(nil)

	response := service.Invite(invitedUsername, invitorId, partyId)

	assert.False(t, response.GetIsError())
}

func Test_PartyInviteService_Invite_SelfInvite(t *testing.T) {
	service, _, userRepo, _, _, _ := setupDefaultService()

	invitedUsername := "myself"
	invitorId := uint(1)
	partyId := uint(2)
	invitedUser := &userDomain.User{
		Model:    gorm.Model{ID: invitorId},
		Username: invitedUsername,
	}

	userRepo.On("FindByUsername", invitedUsername).Return(invitedUser, nil)

	response := service.Invite(invitedUsername, invitorId, partyId)

	assert.Equal(t, api.ErrorBadRequest("cannot party invite yourself"), response)
}

func Test_PartyInviteService_Invite_ExistingPending(t *testing.T) {
	service, inviteRepo, userRepo, _, _, _ := setupDefaultService()

	invitedUsername := "friend"
	invitorId := uint(1)
	partyId := uint(2)
	invitedUser := &userDomain.User{
		Model:    gorm.Model{ID: 3},
		Username: invitedUsername,
	}
	existingInvite := &domains.PartyInvite{
		InvitedId: invitedUser.ID,
		PartyId:   partyId,
		State:     domains.PENDING,
	}

	userRepo.On("FindByUsername", invitedUsername).Return(invitedUser, nil)
	inviteRepo.On("FindByIds", invitedUser.ID, partyId).Return(existingInvite, nil)

	response := service.Invite(invitedUsername, invitorId, partyId)

	assert.False(t, response.GetIsError())
}

func Test_PartyInviteService_Invite_ExistingAccepted(t *testing.T) {
	service, inviteRepo, userRepo, _, _, _ := setupDefaultService()

	invitedUsername := "friend"
	invitorId := uint(1)
	partyId := uint(2)
	invitedUser := &userDomain.User{
		Model:    gorm.Model{ID: 3},
		Username: invitedUsername,
	}
	existingInvite := &domains.PartyInvite{
		InvitedId: invitedUser.ID,
		PartyId:   partyId,
		State:     domains.ACCEPTED,
	}

	userRepo.On("FindByUsername", invitedUsername).Return(invitedUser, nil)
	inviteRepo.On("FindByIds", invitedUser.ID, partyId).Return(existingInvite, nil)

	response := service.Invite(invitedUsername, invitorId, partyId)

	assert.Equal(t, api.ErrorBadRequest("User already accepted the invite"), response)
}

func Test_PartyInviteService_Invite_ExistingDeclined(t *testing.T) {
	service, inviteRepo, userRepo, _, _, _ := setupDefaultService()

	invitedUsername := "friend"
	invitorId := uint(1)
	partyId := uint(2)
	invitedUser := &userDomain.User{
		Model:    gorm.Model{ID: 3},
		Username: invitedUsername,
	}
	existingInvite := &domains.PartyInvite{
		InvitedId: invitedUser.ID,
		PartyId:   partyId,
		State:     domains.DECLINED,
	}

	userRepo.On("FindByUsername", invitedUsername).Return(invitedUser, nil)
	inviteRepo.On("FindByIds", invitedUser.ID, partyId).Return(existingInvite, nil)
	inviteRepo.On("Update", existingInvite).Return(nil)

	response := service.Invite(invitedUsername, invitorId, partyId)

	assert.False(t, response.GetIsError())
	assert.Equal(t, domains.PENDING, existingInvite.State)
}

func Test_PartyInviteService_Invite_FailOnFindUser(t *testing.T) {
	service, _, userRepo, _, _, _ := setupDefaultService()

	invitedUsername := "friend"
	invitorId := uint(1)
	partyId := uint(2)
	expectedErr := errors.New("not found")

	userRepo.On("FindByUsername", invitedUsername).Return(&userDomain.User{}, expectedErr)

	response := service.Invite(invitedUsername, invitorId, partyId)

	assert.Equal(t, api.ErrorBadRequest(expectedErr.Error()), response)
}

func Test_PartyInviteService_CreateInvitation_Success(t *testing.T) {
	service, inviteRepo, userRepo, partyRepo, _, _ := setupDefaultService()

	invitedId := uint(2)
	invitorId := uint(1)
	partyId := uint(3)
	invitor := &userDomain.User{Model: gorm.Model{ID: invitorId}}
	invited := &userDomain.User{Model: gorm.Model{ID: invitedId}}
	party := &partyDomains.Party{
		Model:       gorm.Model{ID: partyId},
		OrganizerID: invitorId,
	}

	userRepo.On("FindById", invitorId, mock.Anything).Return(invitor, nil)
	partyRepo.On("FindById", mock.Anything, mock.Anything).Return(party, nil)
	userRepo.On("FindById", mock.Anything, mock.Anything).Return(invited, nil)
	inviteRepo.On("Create", mock.AnythingOfType("*domains.PartyInvite")).Return(nil)

	response := service.CreateInvitation(invitedId, invitorId, partyId)

	assert.False(t, response.GetIsError())
}

func Test_PartyInviteService_CreateInvitation_Unauthorized(t *testing.T) {
	service, _, userRepo, partyRepo, _, _ := setupDefaultService()

	invitedId := uint(2)
	invitorId := uint(1)
	partyId := uint(3)
	invitor := &userDomain.User{Model: gorm.Model{ID: invitorId}}
	party := &partyDomains.Party{
		Model:       gorm.Model{ID: partyId},
		OrganizerID: 999, // Different organizer
	}

	userRepo.On("FindById", mock.Anything, mock.Anything).Return(invitor, nil)
	partyRepo.On("FindById", mock.Anything, mock.Anything).Return(party, nil)

	response := service.CreateInvitation(invitedId, invitorId, partyId)

	assert.Equal(t, api.ErrorUnauthorized("cannot invite users for other people's party"), response)
}

func Test_PartyInviteService_GetUserPendingInvites_Success(t *testing.T) {
	service, inviteRepo, _, _, _, _ := setupDefaultService()

	userId := uint(1)
	invites := []domains.PartyInvite{
		{InvitedId: userId, State: domains.PENDING},
	}

	inviteRepo.On("FindPendingByInvitedId", userId).Return(&invites, nil)

	response := service.GetUserPendingInvites(userId)

	assert.False(t, response.GetIsError())
}

func Test_PartyInviteService_GetUserPendingInvites_Fail(t *testing.T) {
	service, inviteRepo, _, _, _, _ := setupDefaultService()

	userId := uint(1)
	expectedErr := errors.New("not found")

	inviteRepo.On("FindPendingByInvitedId", userId).Return(&[]domains.PartyInvite{}, expectedErr)

	response := service.GetUserPendingInvites(userId)

	assert.Equal(t, api.ErrorInternalServerError(expectedErr.Error()), response)
}

func Test_PartyInviteService_GetPartyPendingInvites_Success(t *testing.T) {
	service, inviteRepo, _, partyRepo, _, _ := setupDefaultService()

	userId := uint(1)
	partyId := uint(2)
	party := &partyDomains.Party{
		Model:       gorm.Model{ID: partyId},
		OrganizerID: userId,
	}
	invites := []domains.PartyInvite{
		{PartyId: partyId, State: domains.PENDING},
	}

	partyRepo.On("FindById", mock.Anything, mock.Anything).Return(party, nil)
	inviteRepo.On("FindPendingByPartyId", mock.Anything, mock.Anything).Return(&invites, nil)

	response := service.GetPartyPendingInvites(partyId, userId)

	assert.False(t, response.GetIsError())
}

func Test_PartyInviteService_GetPartyPendingInvites_Unauthorized(t *testing.T) {
	service, _, _, partyRepo, _, _ := setupDefaultService()

	userId := uint(1)
	partyId := uint(2)
	party := &partyDomains.Party{
		Model:       gorm.Model{ID: partyId},
		OrganizerID: 999, // Different organizer
	}

	partyRepo.On("FindById", mock.Anything, mock.Anything).Return(party, nil)

	response := service.GetPartyPendingInvites(partyId, userId)

	assert.Equal(t, api.ErrorUnauthorized("cannot organize this party"), response)
}

func Test_PartyInviteService_JoinPublicParty_Success(t *testing.T) {
	service, inviteRepo, userRepo, partyRepo, _, _ := setupDefaultService()

	userId := uint(1)
	partyId := uint(2)
	user := &userDomain.User{Model: gorm.Model{ID: userId}}
	party := &partyDomains.Party{
		Model:       gorm.Model{ID: partyId},
		OrganizerID: 3,
		Private:     false,
	}

	partyRepo.On("FindById", mock.Anything, mock.Anything).Return(party, nil)
	userRepo.On("FindById", userId, mock.Anything).Return(user, nil)
	inviteRepo.On("FindByIds", userId, partyId).Return(&domains.PartyInvite{}, errors.New(domains.NOT_FOUND))
	inviteRepo.On("Create", mock.AnythingOfType("*domains.PartyInvite")).Return(nil)
	partyRepo.On("AddUserToParty", mock.Anything, mock.Anything).Return(nil)

	response := service.JoinPublicParty(partyId, userId)

	assert.False(t, response.GetIsError())
}

func Test_PartyInviteService_JoinPublicParty_AlreadyParticipant(t *testing.T) {
	service, _, userRepo, partyRepo, _, _ := setupDefaultService()

	userId := uint(1)
	partyId := uint(2)
	user := &userDomain.User{Model: gorm.Model{ID: userId}}
	party := &partyDomains.Party{
		Model:       gorm.Model{ID: partyId},
		OrganizerID: 3,
		Private:     false,
		Participants: []userDomain.User{
			{Model: gorm.Model{ID: userId}},
		},
	}

	partyRepo.On("FindById", mock.Anything, mock.Anything).Return(party, nil)
	userRepo.On("FindById", userId, mock.Anything).Return(user, nil)

	response := service.JoinPublicParty(partyId, userId)

	assert.False(t, response.GetIsError())
}

func Test_PartyInviteService_JoinPublicParty_ExistingInvite(t *testing.T) {
	service, inviteRepo, userRepo, partyRepo, _, _ := setupDefaultService()

	userId := uint(1)
	partyId := uint(2)
	user := &userDomain.User{Model: gorm.Model{ID: userId}}
	party := &partyDomains.Party{
		Model:       gorm.Model{ID: partyId},
		OrganizerID: 3,
		Private:     false,
	}
	invite := &domains.PartyInvite{
		InvitedId: userId,
		PartyId:   partyId,
		State:     domains.PENDING,
	}

	partyRepo.On("FindById", mock.Anything, mock.Anything).Return(party, nil)
	userRepo.On("FindById", userId, mock.Anything).Return(user, nil)
	inviteRepo.On("FindByIds", userId, partyId).Return(invite, nil)
	inviteRepo.On("Update", invite).Return(nil)
	partyRepo.On("AddUserToParty", mock.Anything, mock.Anything).Return(nil)

	response := service.JoinPublicParty(partyId, userId)

	assert.False(t, response.GetIsError())
	assert.Equal(t, domains.ACCEPTED, invite.State)
}

func Test_PartyInviteService_JoinPrivateParty_Success(t *testing.T) {
	service, inviteRepo, userRepo, partyRepo, _, _ := setupDefaultService()

	userId := uint(1)
	partyId := uint(2)
	accessCode := "2_secret"
	user := &userDomain.User{Model: gorm.Model{ID: userId}}
	party := &partyDomains.Party{
		Model:             gorm.Model{ID: partyId},
		OrganizerID:       3,
		AccessCode:        accessCode,
		AccessCodeEnabled: true,
	}

	partyRepo.On("FindById", mock.Anything, mock.Anything).Return(party, nil)
	userRepo.On("FindById", mock.Anything, mock.Anything).Return(user, nil)
	inviteRepo.On("FindByIds", userId, uint(partyId)).Return(&domains.PartyInvite{}, errors.New(domains.NOT_FOUND))
	inviteRepo.On("Create", mock.AnythingOfType("*domains.PartyInvite")).Return(nil)
	partyRepo.On("AddUserToParty", mock.Anything, mock.Anything).Return(nil)

	response := service.JoinPrivateParty(userId, accessCode)

	assert.False(t, response.GetIsError())
}

func Test_PartyInviteService_JoinPrivateParty_InvalidCode(t *testing.T) {
	service, _, _, _, _, _ := setupDefaultService()

	userId := uint(1)
	invalidCode := "invalid"

	response := service.JoinPrivateParty(userId, invalidCode)

	assert.Equal(t, api.ErrorBadRequest(domains.INVALID_ACCESS_CODE), response)
}

func Test_PartyInviteService_JoinPrivateParty_CodeMismatch(t *testing.T) {
	service, _, _, partyRepo, _, _ := setupDefaultService()

	userId := uint(1)
	partyId := uint(2)
	accessCode := "2_secret"
	wrongCode := "2_wrong"
	party := &partyDomains.Party{
		Model:             gorm.Model{ID: partyId},
		AccessCode:        accessCode,
		AccessCodeEnabled: true,
	}

	partyRepo.On("FindById", mock.Anything, mock.Anything).Return(party, nil)

	response := service.JoinPrivateParty(userId, wrongCode)

	assert.Equal(t, api.ErrorUnauthorized(domains.INVALID_ACCESS_CODE), response)
}

func Test_PartyInviteService_Kick_Success(t *testing.T) {
	service, inviteRepo, userRepo, partyRepo, foodRepo, drinkRepo := setupDefaultService()

	kickedId := uint(2)
	userId := uint(1)
	partyId := uint(3)
	kickedUser := &userDomain.User{Model: gorm.Model{ID: kickedId}}
	party := &partyDomains.Party{
		Model:       gorm.Model{ID: partyId},
		OrganizerID: userId,
		Participants: []userDomain.User{
			{Model: gorm.Model{ID: kickedId}},
		},
	}
	invite := &domains.PartyInvite{
		InvitedId: kickedId,
		PartyId:   partyId,
		State:     domains.ACCEPTED,
	}

	userRepo.On("FindById", kickedId, mock.Anything).Return(kickedUser, nil)
	partyRepo.On("FindById", mock.Anything, mock.Anything).Return(party, nil)
	inviteRepo.On("FindByIds", kickedId, partyId).Return(invite, nil)
	foodRepo.On("DeleteByContributorId", kickedId).Return(nil)
	drinkRepo.On("DeleteByContributorId", kickedId).Return(nil)
	partyRepo.On("RemoveUserFromParty", mock.Anything, mock.Anything).Return(nil)
	inviteRepo.On("Update", invite).Return(nil)

	response := service.Kick(kickedId, userId, partyId)

	assert.False(t, response.GetIsError())
	assert.Equal(t, domains.DECLINED, invite.State)
}

func Test_PartyInviteService_Kick_LeaveParty(t *testing.T) {
	service, inviteRepo, userRepo, partyRepo, foodRepo, drinkRepo := setupDefaultService()

	kickedId := uint(1)
	userId := uint(1) // Same as kickedId (leaving)
	partyId := uint(3)
	kickedUser := &userDomain.User{Model: gorm.Model{ID: kickedId}}
	party := &partyDomains.Party{
		Model:       gorm.Model{ID: partyId},
		OrganizerID: 2, // Different organizer
		Participants: []userDomain.User{
			{Model: gorm.Model{ID: kickedId}},
		},
	}
	invite := &domains.PartyInvite{
		InvitedId: kickedId,
		PartyId:   partyId,
		State:     domains.ACCEPTED,
	}

	userRepo.On("FindById", mock.Anything, mock.Anything).Return(kickedUser, nil)
	partyRepo.On("FindById", mock.Anything, mock.Anything).Return(party, nil)
	inviteRepo.On("FindByIds", kickedId, partyId).Return(invite, nil)
	foodRepo.On("DeleteByContributorId", kickedId).Return(nil)
	drinkRepo.On("DeleteByContributorId", kickedId).Return(nil)
	partyRepo.On("RemoveUserFromParty", mock.Anything, mock.Anything).Return(nil)
	inviteRepo.On("Update", invite).Return(nil)

	response := service.Kick(kickedId, userId, partyId)

	assert.False(t, response.GetIsError())
}

func Test_PartyInviteService_Kick_Unauthorized(t *testing.T) {
	service, _, userRepo, partyRepo, _, _ := setupDefaultService()

	kickedId := uint(2)
	userId := uint(1)
	partyId := uint(3)
	kickedUser := &userDomain.User{Model: gorm.Model{ID: kickedId}}
	party := &partyDomains.Party{
		Model:       gorm.Model{ID: partyId},
		OrganizerID: 999, // Different organizer
	}

	userRepo.On("FindById", mock.Anything, mock.Anything).Return(kickedUser, nil)
	partyRepo.On("FindById", mock.Anything, mock.Anything).Return(party, nil)

	response := service.Kick(kickedId, userId, partyId)

	assert.Equal(t, api.ErrorUnauthorized(domains.UNAUTHORIZED), response)
}

func Test_PartyInviteService_Kick_OrganizerCannotLeave(t *testing.T) {
	service, _, userRepo, partyRepo, _, _ := setupDefaultService()

	kickedId := uint(1)
	userId := uint(1)
	partyId := uint(3)
	kickedUser := &userDomain.User{Model: gorm.Model{ID: kickedId}}
	party := &partyDomains.Party{
		Model:       gorm.Model{ID: partyId},
		OrganizerID: kickedId,
	}

	userRepo.On("FindById", mock.Anything, mock.Anything).Return(kickedUser, nil)
	partyRepo.On("FindById", mock.Anything, mock.Anything).Return(party, nil)

	response := service.Kick(kickedId, userId, partyId)

	assert.Equal(t, api.ErrorUnauthorized("The organizer cannot leave the party."), response)
}

func Test_PartyInviteService_Kick_NotParticipant(t *testing.T) {
	service, _, userRepo, partyRepo, _, _ := setupDefaultService()

	kickedId := uint(2)
	userId := uint(1)
	partyId := uint(3)
	kickedUser := &userDomain.User{Model: gorm.Model{ID: kickedId}}
	party := &partyDomains.Party{
		Model:       gorm.Model{ID: partyId},
		OrganizerID: userId,
	}

	userRepo.On("FindById", mock.Anything, mock.Anything).Return(kickedUser, nil)
	partyRepo.On("FindById", mock.Anything, mock.Anything).Return(party, nil)

	response := service.Kick(kickedId, userId, partyId)

	assert.False(t, response.GetIsError())
}

func Test_NewService(t *testing.T) {
	partyInviteRepo := new(MockRepository)
	userRepo := new(userUsecases.MockRepository)
	partyRepo := new(partyUsecases.MockRepository)
	foodContribRepo := new(foodContributionUsecases.MockRepository)
	drinkContribRepo := new(drinkContributionUsecases.MockRepository)
	repoCollector := repo.RepoCollector{
		PartyRepo:        partyRepo,
		UserRepo:         userRepo,
		DrinkContribRepo: drinkContribRepo,
		FoodContribRepo:  foodContribRepo,
		PartyInviteRepo:  partyInviteRepo,
	}

	serviceInterface := NewPartyInviteService(&repoCollector)
	service := serviceInterface.(*PartyInviteService)

	assert.Equal(t, service.PartyRepository, partyRepo)
	assert.Equal(t, service.UserRepository, userRepo)
	assert.Equal(t, service.DrinkContributionRepository, drinkContribRepo)
	assert.Equal(t, service.FoodContributionRepository, foodContribRepo)
	assert.Equal(t, service.PartyInviteRepository, partyInviteRepo)
}
