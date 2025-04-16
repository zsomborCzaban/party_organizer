package domains

import (
	"github.com/zsomborCzaban/party_organizer/services/users/user/domains"
	"gorm.io/gorm"
	"time"
)

type PartyDTO struct {
	ID                uint           `json:"ID,omitempty"`
	Place             string         `json:"place,omitempty" validate:"required,min=3"`
	StartTime         time.Time      `json:"start_time,omitempty" validate:"required"`
	Name              string         `json:"name,omitempty" validate:"required"`
	GoogleMapsLink    string         `json:"google_maps_link" validate:"http_url"`
	FacebookLink      string         `json:"facebook_link" validate:"http_url"`
	WhatsappLink      string         `json:"whatsapp_link" validate:"http_url"`
	Private           bool           `json:"is_private"`
	AccessCodeEnabled bool           `json:"access_code_enabled" validate:"bool_allowed_by_bool=Private"`
	AccessCode        string         `json:"access_code" validate:"string_allowed_by_bool_and_min_3=AccessCodeEnabled,string_allowed_by_bool=AccessCodeEnabled"`
	OrganizerID       uint           `json:"organizer_id,omitempty"`
	Participants      []domains.User `json:"participants"`
}

// todo: refactor this to usecases, bc Organiter and accessCode is set dynamicly, after this function call. which is bad practice and counts as buisness logic
func (p *PartyDTO) TransformToParty() *Party {
	return &Party{
		Model:             gorm.Model{ID: p.ID},
		Place:             p.Place,
		StartTime:         p.StartTime,
		Name:              p.Name,
		GoogleMapsLink:    p.GoogleMapsLink,
		FacebookLink:      p.FacebookLink,
		WhatsappLink:      p.WhatsappLink,
		Private:           p.Private,
		AccessCodeEnabled: p.AccessCodeEnabled,
		AccessCode:        p.AccessCode,
		OrganizerID:       p.OrganizerID,
	}
}
