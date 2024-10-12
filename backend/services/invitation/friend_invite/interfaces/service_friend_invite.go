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
	invite, err := fs.FriendInviteRepository.FindByIds(invitorId, userId)
	if err != nil {
		return api.ErrorInternalServerError(err)
	}

	if invite.State != "Accepted" {
		invite.State = "Accepted"
		if err2 := fs.FriendInviteRepository.Update(invite); err2 != nil {
			return api.ErrorInternalServerError(err2)
		}
	}

	return api.Success(invite)
}

func (fs FriendInviteService) Decline(invitorId, userId uint) api.IResponse {
	invite, err := fs.FriendInviteRepository.FindByIds(invitorId, userId)
	if err != nil {
		return api.ErrorInternalServerError(err)
	}

	if invite.State != "Declined" {
		invite.State = "Declined"
		if err2 := fs.FriendInviteRepository.Update(invite); err2 != nil {
			return api.ErrorInternalServerError(err2)
		}
	}

	return api.Success(invite)
}

func (fs FriendInviteService) Invite(invitedId, userId uint) api.IResponse {
	invite, err := fs.FriendInviteRepository.FindByIds(userId, invitedId)
	if err != nil && err.Error() != "error, invite was nil" {
		return api.ErrorInternalServerError(err.Error())
	}

	if err != nil && err.Error() == "error, invite was nil" {
		invitation := &domains.FriendInvitation{
			InvitorId: userId,
			InvitedId: invitedId,
			State:     "Pending",
		}
		if errCreation := fs.FriendInviteRepository.Create(invitation); errCreation != nil {
			return api.ErrorInternalServerError(errCreation.Error())
		}
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
