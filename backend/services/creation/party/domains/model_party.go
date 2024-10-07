package domains

import (
	"github.com/zsomborCzaban/party_organizer/services/user/domains"
	"gorm.io/gorm"
	"time"
)

type Party struct {
	gorm.Model

	Place        string         `json:"place"`
	StartTime    time.Time      `json:"start_time"`
	Name         string         `json:"name"`
	OrganizerID  uint           `json:"organizer_id"`
	Participants []domains.User `gorm:"many2many:party_participants;"`
}

func (p *Party) TransformToPartyDTO() *PartyDTO {
	return &PartyDTO{
		ID:           p.ID,
		Place:        p.Place,
		StartTime:    p.StartTime,
		Name:         p.Name,
		OrganizerID:  p.OrganizerID,
		Participants: p.Participants,
	}
}
