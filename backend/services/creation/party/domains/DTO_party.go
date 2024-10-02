package domains

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
	"time"
)

type PartyDTO struct {
	ID             uint          `json:"id,omitempty"`
	Place          string        `json:"place,omitempty" validate:"required,min=3"`
	StartTime      time.Time     `json:"start_time,omitempty" validate:"required"`
	Name           string        `json:"name,omitempty" validate:"required"`
	OrganizerID    uint          `json:"organizer_id,omitempty"`
	ParticipantIDs pq.Int64Array `json:"participant_ids,omitempty"`
}

func (p *PartyDTO) TransformToParty() *Party {
	return &Party{
		Model:          gorm.Model{ID: p.ID},
		Place:          p.Place,
		StartTime:      p.StartTime,
		Name:           p.Name,
		OrganizerID:    p.OrganizerID,
		ParticipantIDs: p.ParticipantIDs,
	}
}
