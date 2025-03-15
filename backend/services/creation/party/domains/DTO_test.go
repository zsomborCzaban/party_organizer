package domains

import (
	"github.com/go-playground/assert/v2"
	"gorm.io/gorm"
	"testing"
	"time"
)

func Test_TransformToParty(t *testing.T) {
	dto := PartyDTO{
		ID:                1,
		Place:             "ligma",
		StartTime:         time.Time{},
		Name:              "updog",
		GoogleMapsLink:    "dragon",
		FacebookLink:      "sucondeese",
		WhatsappLink:      "out of ideas",
		Private:           false,
		AccessCodeEnabled: false,
		AccessCode:        "",
		OrganizerID:       3,
		Participants:      nil,
	}
	party := dto.TransformToParty()

	expectedParty := Party{
		Model: gorm.Model{
			ID: 1,
		},
		Place:             "ligma",
		StartTime:         time.Time{},
		Name:              "updog",
		GoogleMapsLink:    "dragon",
		FacebookLink:      "sucondeese",
		WhatsappLink:      "out of ideas",
		Private:           false,
		AccessCodeEnabled: false,
		AccessCode:        "",
		OrganizerID:       3,
		Participants:      nil,
	}

	assert.Equal(t, expectedParty.ID, party.ID)
	assert.Equal(t, expectedParty.Place, party.Place)
	assert.Equal(t, expectedParty.StartTime, party.StartTime)
	assert.Equal(t, expectedParty.Name, party.Name)
	assert.Equal(t, expectedParty.GoogleMapsLink, party.GoogleMapsLink)
	assert.Equal(t, expectedParty.FacebookLink, party.FacebookLink)
	assert.Equal(t, expectedParty.WhatsappLink, party.WhatsappLink)
	assert.Equal(t, expectedParty.Private, party.Private)
	assert.Equal(t, expectedParty.AccessCodeEnabled, party.AccessCodeEnabled)
	assert.Equal(t, expectedParty.AccessCode, party.AccessCode)
	assert.Equal(t, expectedParty.OrganizerID, party.OrganizerID)
	assert.Equal(t, expectedParty.Participants, party.Participants)

}
