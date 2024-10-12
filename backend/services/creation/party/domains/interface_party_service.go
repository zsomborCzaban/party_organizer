package domains

import "github.com/zsomborCzaban/party_organizer/common/api"

type IPartyService interface {
	GetPartiesByOrganizerId(uint) api.IResponse
	GetPartiesByParticipantId(uint) api.IResponse
	AddUserToParty(partyId, userId uint) api.IResponse //auth this. only party organizer can add to party

	CreateParty(PartyDTO PartyDTO, userId uint) api.IResponse
	GetParty(uint) api.IResponse //authenticate this
	UpdateParty(partyDTO PartyDTO, userId uint) api.IResponse
	DeleteParty(uint) api.IResponse //auth this
}
