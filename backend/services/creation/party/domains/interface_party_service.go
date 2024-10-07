package domains

import "github.com/zsomborCzaban/party_organizer/common/api"

type IPartyService interface {
	GetPartiesByOrganizerId(uint) api.IResponse
	GetPartiesByParticipantId(uint) api.IResponse

	CreateParty(PartyDTO) api.IResponse
	GetParty(uint) api.IResponse
	UpdateParty(PartyDTO) api.IResponse
	DeleteParty(uint) api.IResponse
}
