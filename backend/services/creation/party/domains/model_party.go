package domains

import (
	"gorm.io/gorm"
	"time"
)

type Party struct {
	gorm.Model

	Place          string    `json:"place"`
	StartTime      time.Time `json:"start_time"`
	OrganizerID    uint      `json:"organizer_id"`
	ParticipantIDs []uint    `json:"Participant_ids"`
}

func (p *Party) TransformToPartyDTO() *PartyDTO {
	return &PartyDTO{
		ID:             p.ID,
		Place:          p.Place,
		StartTime:      p.StartTime,
		OrganizerID:    p.OrganizerID,
		ParticipantIDs: p.ParticipantIDs,
	}
}
