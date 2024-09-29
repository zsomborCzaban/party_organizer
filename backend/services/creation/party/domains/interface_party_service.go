package domains

type IPartyService interface {
	CreateParty(PartyDTO) IResponse
	GetParty(uint) IResponse
	UpdateParty(PartyDTO) IResponse
	DeleteParty(uint) IResponse
}
