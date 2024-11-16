package domains

import (
	"github.com/zsomborCzaban/party_organizer/utils/api"
)

type IPartyService interface {
	GetParticipants(partyId, userId uint) api.IResponse
	GetPublicParties() api.IResponse
	GetPartiesByOrganizerId(uint) api.IResponse
	GetPartiesByParticipantId(uint) api.IResponse

	CreateParty(PartyDTO PartyDTO, userId uint) api.IResponse
	GetParty(partyId, userId uint) api.IResponse
	UpdateParty(partyDTO PartyDTO, userId uint) api.IResponse
	DeleteParty(uint) api.IResponse //auth this
}
