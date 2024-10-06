package domains

import "github.com/zsomborCzaban/party_organizer/common/api"

type IPartyService interface {
	CreateParty(PartyDTO) api.IResponse
	GetParty(uint) api.IResponse
	UpdateParty(PartyDTO) api.IResponse
	DeleteParty(uint) api.IResponse
}
