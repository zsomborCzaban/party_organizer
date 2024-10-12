package interfaces

import (
	"github.com/zsomborCzaban/party_organizer/common/api"
	"github.com/zsomborCzaban/party_organizer/services/invitation/friend_invite/domains"
)

type FriendInviteService struct {
	FriendInviteRepository domains.IFriendInviteRepository
}

func NewFriendInviteService(repo domains.IFriendInviteRepository) domains.IFriendInviteService {
	return &FriendInviteService{
		FriendInviteRepository: repo,
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

	if invite.State != "Accepted" {
		invite.State = "Accepted"
		if err2 := fs.FriendInviteRepository.Update(invite); err2 != nil {
			return api.ErrorInternalServerError(err2.Error())
		}
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

	if invite.State != "Declined" {
		invite.State = "Declined"
		if err2 := fs.FriendInviteRepository.Update(invite); err2 != nil {
			return api.ErrorInternalServerError(err2.Error())
		}
	}

	return api.Success(invite)
}

// todo: refactor this
func (fs FriendInviteService) Invite(invitedId, userId uint) api.IResponse {
	if userId == invitedId {
		return api.ErrorBadRequest("cannot friend invite yourself")
	}

	invite, err := fs.FriendInviteRepository.FindByIds(userId, invitedId)
	if err != nil && err.Error() != "error, fetched invites cannot be transformed to *[]FriendInvite" {
		return api.ErrorInternalServerError(err.Error())
	}

	if err != nil && err.Error() == "error, fetched invites cannot be transformed to *[]FriendInvite" {
		invitation := &domains.FriendInvite{
			InvitorId: userId,
			InvitedId: invitedId,
			State:     "Pending",
		}
		if errCreation := fs.FriendInviteRepository.Create(invitation); errCreation != nil {
			return api.ErrorInternalServerError(errCreation.Error())
		}
		return api.Success(invitation)
	}

	if invite.State == "Accepted" {
		return api.ErrorBadRequest("User already accepted the invite")
	}

	if invite.State == "Declined" {
		invite.State = "Pending"
		if err2 := fs.FriendInviteRepository.Update(invite); err2 != nil {
			return api.ErrorInternalServerError(err2.Error())
		}
	}

	return api.Success(invite)
}
