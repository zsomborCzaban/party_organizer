package domains

import (
	"github.com/stretchr/testify/assert"
	userDomain "github.com/zsomborCzaban/party_organizer/services/users/user/domains"
	"github.com/zsomborCzaban/party_organizer/utils/adminUser"
	"gorm.io/gorm"
	"testing"
)

func Test_HasParticipant_UserIsOrganizer(t *testing.T) {
	party := Party{
		OrganizerID:  1,
		Participants: make([]userDomain.User, 0),
	}

	hasParticipant := party.HasParticipant(1)
	assert.True(t, hasParticipant)
}

func Test_HasParticipant_UserIsParticipant(t *testing.T) {
	participants := []userDomain.User{{
		Model: gorm.Model{
			ID: 2,
		},
	}, {
		Model: gorm.Model{
			ID: 3,
		},
	}}

	party := Party{
		OrganizerID:  1,
		Participants: participants,
	}

	hasParticipant := party.HasParticipant(3)
	assert.True(t, hasParticipant)
}

func Test_HasParticipant_UserIsNotParticipant(t *testing.T) {
	party := Party{
		OrganizerID:  1,
		Participants: make([]userDomain.User, 0),
	}

	hasParticipant := party.HasParticipant(4)
	assert.False(t, hasParticipant)
}

func Test_CanBeOrganizedBy_UserIsOrganizer(t *testing.T) {
	party := Party{
		OrganizerID: 1,
	}

	canBeOrganizedBy := party.CanBeOrganizedBy(1)
	assert.True(t, canBeOrganizedBy)
}

func Test_CanBeOrganizedBy_UserIsAdmin(t *testing.T) {
	party := Party{
		OrganizerID: 1,
	}

	canBeOrganizedBy := party.CanBeOrganizedBy(adminUser.ADMIN_USER_ID)
	assert.True(t, canBeOrganizedBy)
}

func Test_CanBeAccessedBy_True(t *testing.T) {
	party := Party{
		OrganizerID: 1,
	}

	canBeAccessedBy := party.CanBeAccessedBy(1)
	assert.True(t, canBeAccessedBy)

}

func Test_CanBeAccessedBy_False(t *testing.T) {
	party := Party{
		OrganizerID: 1,
	}

	canBeOrganizedBy := party.CanBeAccessedBy(2)
	assert.False(t, canBeOrganizedBy)
}
