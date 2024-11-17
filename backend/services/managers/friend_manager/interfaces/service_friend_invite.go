package interfaces

import (
	"fmt"
	"github.com/zsomborCzaban/party_organizer/services/managers/friend_manager/domains"
	userDomain "github.com/zsomborCzaban/party_organizer/services/user/domains"
	"github.com/zsomborCzaban/party_organizer/utils/api"
	"github.com/zsomborCzaban/party_organizer/utils/repo"
)

type FriendInviteService struct {
	FriendInviteRepository *domains.IFriendInviteRepository
	UserRepository         *userDomain.IUserRepository
}

func NewFriendInviteService(repoCollector *repo.RepoCollector) domains.IFriendInviteService {
	return &FriendInviteService{
		FriendInviteRepository: repoCollector.FriendInviteRepo,
		UserRepository:         repoCollector.UserRepo,
	}
}

func (fs FriendInviteService) Accept(invitorId, userId uint) api.IResponse {
	//should exists in database
	//if userId == invitorId {
	//	return api.ErrorBadRequest("cannot accept yourself")
	//}

	invite, err := (*fs.FriendInviteRepository).FindByIds(invitorId, userId)
	if err != nil {
		return api.ErrorInternalServerError(err.Error())
	}

	if invite.State == domains.DECLINED {
		return api.ErrorBadRequest("Cannot accept already declined friends. Try inviting them")
	}

	if invite.State == domains.ACCEPTED {
		return api.Success(invite)
	}

	invitor, err2 := (*fs.UserRepository).FindById(invitorId)
	if err2 != nil {
		return api.ErrorBadRequest(err2.Error())
	}

	user, err3 := (*fs.UserRepository).FindById(userId)
	if err3 != nil {
		return api.ErrorBadRequest(err3.Error())
	}

	//todo: put this in a transaction
	invite.State = domains.ACCEPTED
	if err4 := (*fs.FriendInviteRepository).Update(invite); err4 != nil {
		return api.ErrorInternalServerError(err4.Error())
	}

	if err5 := (*fs.UserRepository).AddFriend(invitor, user); err5 != nil {
		return api.ErrorInternalServerError(err5.Error())
	}

	return api.Success(invite)
}

func (fs FriendInviteService) Decline(invitorId, userId uint) api.IResponse {
	invite, err := (*fs.FriendInviteRepository).FindByIds(invitorId, userId)
	if err != nil {
		return api.ErrorInternalServerError(err.Error())
	}

	if invite.State == domains.ACCEPTED {
		return api.ErrorBadRequest("Cannot decline already accepted friends. Try deleting them")
	}

	if invite.State == domains.DECLINED {
		return api.Success(invite)
	}

	invite.State = domains.DECLINED
	if err2 := (*fs.FriendInviteRepository).Update(invite); err2 != nil {
		return api.ErrorInternalServerError(err2.Error())
	}

	return api.Success(invite)
}

func (fs FriendInviteService) Invite(invitedUsername string, userId uint) api.IResponse {
	invited, err := (*fs.UserRepository).FindByUsername(invitedUsername)
	if err != nil {
		return api.ErrorBadRequest(err.Error())
	}

	if userId == invited.ID {
		return api.ErrorBadRequest("cannot friend invite yourself")
	}

	//check reverse invite
	reverseInvite, _ := (*fs.FriendInviteRepository).FindByIds(invited.ID, userId)
	if reverseInvite != nil && reverseInvite.State == domains.DECLINED {
		return fs.ReverseInvite(reverseInvite)
	}
	if reverseInvite != nil {
		return api.ErrorBadRequest("friend request already exists, try accepting it")
	}
	//check reverse invite end

	invite, err2 := (*fs.FriendInviteRepository).FindByIds(userId, invited.ID)
	if err2 != nil && err2.Error() == domains.NOT_FOUND {
		return fs.CreateInvitation(invited.ID, userId)
	}
	if err2 != nil {
		return api.ErrorInternalServerError(err2.Error())
	}

	if invite.State == domains.ACCEPTED {
		return api.ErrorBadRequest("User already accepted the invite")
	}

	if invite.State == domains.PENDING {
		return api.Success(invite)
	}

	invite.State = domains.PENDING
	if err3 := (*fs.FriendInviteRepository).Update(invite); err3 != nil {
		return api.ErrorInternalServerError(err3.Error())
	}

	return api.Success(invite)
}

func (fs FriendInviteService) CreateInvitation(invitedId, userId uint) api.IResponse {
	invitor, err := (*fs.UserRepository).FindById(userId)
	if err != nil {
		return api.ErrorBadRequest(fmt.Sprintf("cannot find user with id: %d", userId))
	}

	invited, err2 := (*fs.UserRepository).FindById(invitedId)
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

	if errCreation := (*fs.FriendInviteRepository).Create(invitation); errCreation != nil {
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

	if err := (*fs.FriendInviteRepository).Update(invitation); err != nil {
		return api.ErrorInternalServerError(err.Error())
	}
	return api.Success(invitation)
}

func (fs FriendInviteService) GetPendingInvites(userId uint) api.IResponse {
	invites, err := (*fs.FriendInviteRepository).FindPendingByInvitedId(userId)
	if err != nil {
		return api.ErrorInternalServerError(err.Error())
	}

	return api.Success(invites)
}

func (fs FriendInviteService) RemoveFriend(userId, friendId uint) api.IResponse {
	user, err := (*fs.UserRepository).FindByIdWithFriends(userId)
	if err != nil {
		return api.ErrorBadRequest(err.Error())
	}

	friend, err2 := (*fs.UserRepository).FindByIdWithFriends(friendId)
	if err2 != nil {
		return api.ErrorBadRequest(err2.Error())
	}

	if !user.HasFriend(friendId) {
		return api.Success("removed friends successfully")
	}

	invite, _ := (*fs.FriendInviteRepository).FindByIds(userId, friendId)
	rInvite, _ := (*fs.FriendInviteRepository).FindByIds(friendId, userId)

	if invite != nil {
		return fs.RemoveFriendAndInvite(user, friend, invite)
	}

	if rInvite != nil {
		return fs.RemoveFriendAndInvite(user, friend, rInvite)
	}

	//should only get here if the database state is bad
	return api.ErrorInternalServerError("error while removing friend")
}

func (fs FriendInviteService) RemoveFriendAndInvite(user, friend *userDomain.User, invite *domains.FriendInvite) api.IResponse {
	if err := (*fs.UserRepository).RemoveFriend(user, friend); err != nil {
		return api.ErrorInternalServerError(err.Error())
	}

	//state is always ACCEPTED before, bc the user has the friend
	invite.State = domains.DECLINED
	if err2 := (*fs.FriendInviteRepository).Update(invite); err2 != nil {
		return api.ErrorInternalServerError(err2.Error())
	}

	return api.Success("removed friend successfully")
}
