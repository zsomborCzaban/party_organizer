package domains

import "github.com/zsomborCzaban/party_organizer/common/api"

type IPartyInviteService interface {
	Invite(invitedId, invitorId, partyId uint) api.IResponse
	Accept(invitedId, partyId uint) api.IResponse
	Decline(invitedId, partyId uint) api.IResponse

	GetPendingInvites(uint) api.IResponse

	Kick(kickedId, userId, partyId uint) api.IResponse
	JoinPublicParty(partyId, userId uint) api.IResponse
	JoinPrivateParty(partyId, userId uint, accessCode string) api.IResponse
}
