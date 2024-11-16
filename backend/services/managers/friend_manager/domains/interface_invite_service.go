package domains

import "github.com/zsomborCzaban/party_organizer/utils/api"

type IFriendInviteService interface {
	Invite(invitedUsername string, userId uint) api.IResponse
	Accept(invitorId, userId uint) api.IResponse
	Decline(invitorId, userId uint) api.IResponse

	GetPendingInvites(uint) api.IResponse
	RemoveFriend(userId, friendId uint) api.IResponse
}
