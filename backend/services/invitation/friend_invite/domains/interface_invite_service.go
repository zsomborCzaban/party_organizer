package domains

import "github.com/zsomborCzaban/party_organizer/common/api"

type IFriendInviteService interface {
	Invite(invitedId, userId uint) api.IResponse
	Accept(invitorId, userId uint) api.IResponse
	Decline(invitorId, userId uint) api.IResponse
}
