package domains

type IPartyRepository interface {
	CreateParty(*PartyDTO) (*Party, error)
	UpdateParty(uint) (*Party, error)
	GetParty(*PartyDTO) (*Party, error)
	DeleteParty(uint) (*Party, error)
}
