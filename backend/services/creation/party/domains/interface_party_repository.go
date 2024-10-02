package domains

type IPartyRepository interface {
	CreateParty(*Party) error
	GetParty(uint) (*Party, error)
	UpdateParty(*Party) error
	DeleteParty(*Party) error
}
