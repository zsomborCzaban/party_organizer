package domains

import (
	"time"
)

type PartyDTO struct {
	ID             uint      `json:"id,omitempty"`
	Place          string    `json:"place" validate:"required"`
	StartTime      time.Time `json:"start_time" validate:"required"`
	OrganizerID    uint      `json:"organizer_id,omitempty"`
	ParticipantIDs []uint    `json:"Participant_ids,omitempty"`
}
