package domains

type IPartyRepository interface {
	GetPartiesByOrganizerId(uint) (*[]Party, error)
	GetPartiesByParticipantId(uint) (*[]Party, error)
	AddUserToParty(uint, uint) error

	CreateParty(*Party) error
	GetParty(uint) (*Party, error)
	UpdateParty(*Party) error
	DeleteParty(*Party) error
}
