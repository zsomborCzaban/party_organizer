package domains

import userDomain "github.com/zsomborCzaban/party_organizer/services/user/domains"

type IPartyRepository interface {
	GetPartiesByOrganizerId(uint) (*[]Party, error)
	GetPartiesByParticipantId(uint) (*[]Party, error)
	GetPublicParties() (*[]Party, error)
	AddUserToParty(*Party, *userDomain.User) error

	CreateParty(*Party) error
	FindById(uint) (*Party, error)
	UpdateParty(*Party) error
	DeleteParty(*Party) error
}
