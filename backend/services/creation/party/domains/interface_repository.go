package domains

import userDomain "github.com/zsomborCzaban/party_organizer/services/user/domains"

type IPartyRepository interface {
	GetPartiesByOrganizerId(uint) (*[]Party, error)
	GetPartiesByParticipantId(uint) (*[]Party, error)
	GetPublicParties() (*[]Party, error)
	AddUserToParty(*Party, *userDomain.User) error
	RemoveUserFromParty(*Party, *userDomain.User) error

	Create(*Party) error
	FindById(id uint, associations ...string) (*Party, error)
	Update(*Party) error
	Delete(*Party) error
}
