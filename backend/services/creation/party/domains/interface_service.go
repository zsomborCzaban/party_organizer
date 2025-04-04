package domains

import (
	"github.com/zsomborCzaban/party_organizer/utils/api"
)

type IPartyService interface {
	GetParticipants(partyId, userId uint) api.IResponse
	GetPublicParties() api.IResponse
	GetPublicParty(partyId uint) api.IResponse
	GetPartiesByOrganizerId(uint) api.IResponse
	GetPartiesByParticipantId(uint) api.IResponse

	Create(PartyDTO PartyDTO, userId uint) api.IResponse
	Update(partyDTO PartyDTO, userId uint) api.IResponse
	Get(partyId, userId uint) api.IResponse
	Delete(partyId, userId uint) api.IResponse //auth this
}
