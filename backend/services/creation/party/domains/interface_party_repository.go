package domains

import userDomain "github.com/zsomborCzaban/party_organizer/services/user/domains"

type IPartyRepository interface {
	GetPartiesByOrganizerId(uint) (*[]Party, error)
	GetPartiesByParticipantId(uint) (*[]Party, error)
	AddUserToParty(uint, *userDomain.User) error

	CreateParty(*Party) error
	GetParty(uint) (*Party, error)
	UpdateParty(*Party) error
	DeleteParty(*Party) error
}
