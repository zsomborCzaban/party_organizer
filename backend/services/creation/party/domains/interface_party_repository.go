package domains

type IPartyRepository interface {
	CreateParty(*PartyDTO) (*Party, error)
	GetParty(uint) (*Party, error)
	UpdateParty(*PartyDTO) (*Party, error)
	DeleteParty(uint) (*Party, error)
}
