package domains

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
	"time"
)

type Party struct {
	gorm.Model

	Place          string        `json:"place"`
	StartTime      time.Time     `json:"start_time"`
	Name           string        `json:"name"`
	OrganizerID    uint          `json:"organizer_id"`
	ParticipantIDs pq.Int64Array `json:"participant_ids" gorm:"type:integer[]"`
}

func (p *Party) TransformToPartyDTO() *PartyDTO {
	return &PartyDTO{
		ID:             p.ID,
		Place:          p.Place,
		StartTime:      p.StartTime,
		Name:           p.Name,
		OrganizerID:    p.OrganizerID,
		ParticipantIDs: p.ParticipantIDs,
	}
}
