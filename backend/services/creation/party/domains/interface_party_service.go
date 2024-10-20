package domains

import "github.com/zsomborCzaban/party_organizer/common/api"

type IPartyService interface {
	GetPublicParties() api.IResponse
	GetPartiesByOrganizerId(uint) api.IResponse
	GetPartiesByParticipantId(uint) api.IResponse
	AddUserToParty(partyId, userId uint) api.IResponse //this will only get called by party_invite

	CreateParty(PartyDTO PartyDTO, userId uint) api.IResponse
	GetParty(partyId, userId uint) api.IResponse
	UpdateParty(partyDTO PartyDTO, userId uint) api.IResponse
	DeleteParty(uint) api.IResponse //auth this
}
