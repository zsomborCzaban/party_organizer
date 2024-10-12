package interfaces

import (
	"fmt"
	"github.com/zsomborCzaban/party_organizer/common/api"
	"github.com/zsomborCzaban/party_organizer/services/invitation/friend_invite/domains"
	userDomain "github.com/zsomborCzaban/party_organizer/services/user/domains"
)

type FriendInviteService struct {
	FriendInviteRepository domains.IFriendInviteRepository
	UserRepository         userDomain.IUserRepository
}

func NewFriendInviteService(repo domains.IFriendInviteRepository, userRepo userDomain.IUserRepository) domains.IFriendInviteService {
	return &FriendInviteService{
		FriendInviteRepository: repo,
		UserRepository:         userRepo,
	}
}

func (fs FriendInviteService) Accept(invitorId, userId uint) api.IResponse {
	if userId == invitorId {
		return api.ErrorBadRequest("cannot accept yourself")
	}

	invite, err := fs.FriendInviteRepository.FindByIds(invitorId, userId)
	if err != nil {
		return api.ErrorInternalServerError(err.Error())
	}

	if invite.State == domains.DECLINED {
		return api.ErrorBadRequest("Cannot accept already declined friends. Try inviting them")
	}

	//todo: put this in a transaction
	if invite.State != domains.ACCEPTED {
		invite.State = domains.ACCEPTED
		if err2 := fs.FriendInviteRepository.Update(invite); err2 != nil {
			return api.ErrorInternalServerError(err2.Error())
		}
	}

	if err2 := fs.UserRepository.AddFriend(invitorId, userId); err2 != nil {
		return api.ErrorInternalServerError(err2.Error())
	}

	return api.Success(invite)
}

func (fs FriendInviteService) Decline(invitorId, userId uint) api.IResponse {
	if userId == invitorId {
		return api.ErrorBadRequest("cannot decline yourself")
	}

	invite, err := fs.FriendInviteRepository.FindByIds(invitorId, userId)
	if err != nil {
		return api.ErrorInternalServerError(err.Error())
	}

	if invite.State == domains.ACCEPTED {
		return api.ErrorBadRequest("Cannot decline already accepted friends. Try deleting them")
	}

	if invite.State != domains.DECLINED {
		invite.State = domains.DECLINED
		if err2 := fs.FriendInviteRepository.Update(invite); err2 != nil {
			return api.ErrorInternalServerError(err2.Error())
		}
	}

	return api.Success(invite)
}

func (fs FriendInviteService) Invite(invitedId, userId uint) api.IResponse {
	if userId == invitedId {
		return api.ErrorBadRequest("cannot friend invite yourself")
	}

	//check reverse invite
	reverseInvite, err := fs.FriendInviteRepository.FindByIds(invitedId, userId)
	if reverseInvite != nil && reverseInvite.State == domains.DECLINED {
		return fs.ReverseInvite(reverseInvite)
	}
	if reverseInvite != nil {
		return api.ErrorBadRequest("friend request already exists, try accepting it")
	}
	//check reverse invite end

	invite, err := fs.FriendInviteRepository.FindByIds(userId, invitedId)
	if err != nil && err.Error() == domains.NOT_FOUND {
		return fs.CreateInvitation(invitedId, userId)
	}
	if err != nil {
		return api.ErrorInternalServerError(err.Error())
	}

	if invite.State == domains.ACCEPTED {
		return api.ErrorBadRequest("User already accepted the invite")
	}

	if invite.State == domains.DECLINED {
		invite.State = domains.PENDING
		if err2 := fs.FriendInviteRepository.Update(invite); err2 != nil {
			return api.ErrorInternalServerError(err2.Error())
		}
	}

	return api.Success(invite)
}

func (fs FriendInviteService) CreateInvitation(invitedId, userId uint) api.IResponse {
	invitor, err := fs.UserRepository.FindById(userId)
	if err != nil {
		return api.ErrorBadRequest(fmt.Sprintf("cannot find user with id: %d", userId))
	}

	invited, err2 := fs.UserRepository.FindById(invitedId)
	if err2 != nil {
		return api.ErrorBadRequest(fmt.Sprintf("cannot find user with id: %d", userId))
	}

	invitation := &domains.FriendInvite{
		InvitorId: userId,
		Invited:   *invited,
		InvitedId: invitedId,
		Invitor:   *invitor,
		State:     domains.PENDING,
	}

	if errCreation := fs.FriendInviteRepository.Create(invitation); errCreation != nil {
		return api.ErrorInternalServerError(errCreation.Error())
	}
	return api.Success(invitation)
}

func (fs FriendInviteService) ReverseInvite(invitation *domains.FriendInvite) api.IResponse {
	invitedId := invitation.InvitedId
	invited := invitation.Invited
	invitatorId := invitation.InvitorId
	invitator := invitation.Invitor

	invitation.InvitedId = invitatorId
	invitation.Invited = invitator
	invitation.InvitorId = invitedId
	invitation.Invitor = invited
	invitation.State = domains.PENDING

	if err := fs.FriendInviteRepository.Update(invitation); err != nil {
		return api.ErrorInternalServerError(err.Error())
	}
	return api.Success(invitation)
}

func (fs FriendInviteService) GetPendingInvites(userId uint) api.IResponse {
	invites, err := fs.FriendInviteRepository.FindPendingByInvitedId(userId)
	if err != nil {
		return api.ErrorInternalServerError(err.Error())
	}

	return api.Success(invites)
}
