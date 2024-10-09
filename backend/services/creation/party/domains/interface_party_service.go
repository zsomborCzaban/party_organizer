package domains

import "github.com/zsomborCzaban/party_organizer/common/api"

type IPartyService interface {
	GetPartiesByOrganizerId(uint) api.IResponse
	GetPartiesByParticipantId(uint) api.IResponse
	AddUserToParty(partyId, userId uint) api.IResponse

	CreateParty(PartyDTO) api.IResponse
	GetParty(uint) api.IResponse        //if the user is organizer or participant
	UpdateParty(PartyDTO) api.IResponse //if the user os organizer
	DeleteParty(uint) api.IResponse
}
