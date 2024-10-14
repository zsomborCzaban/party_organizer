package domains

import (
	"github.com/zsomborCzaban/party_organizer/services/user/domains"
	"gorm.io/gorm"
	"time"
)

type Party struct {
	gorm.Model

	Place             string         `json:"place"`
	StartTime         time.Time      `json:"start_time"`
	Name              string         `json:"name"`
	Private           bool           `json:"is_private"`
	AccessCodeEnabled bool           `json:"access_code_enabled"`
	AccessCode        string         `json:"access_code"`
	OrganizerID       uint           `json:"organizer_id"`
	Organizer         domains.User   `json:"-"`
	Participants      []domains.User `json:"-" gorm:"many2many:party_participants;"`
}

func (p *Party) TransformToPartyDTO() *PartyDTO {
	return &PartyDTO{
		ID:                p.ID,
		Place:             p.Place,
		StartTime:         p.StartTime,
		Name:              p.Name,
		Private:           p.Private,
		AccessCodeEnabled: p.AccessCodeEnabled,
		AccessCode:        p.AccessCode,
		OrganizerID:       p.OrganizerID,
	}
}
