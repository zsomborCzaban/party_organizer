package domains

import (
	"github.com/zsomborCzaban/party_organizer/services/user/domains"
	"github.com/zsomborCzaban/party_organizer/utils/adminUser"
	"gorm.io/gorm"
	"time"
)

type Party struct {
	gorm.Model

	Place             string         `json:"place"`
	StartTime         time.Time      `json:"start_time"`
	Name              string         `json:"name"`
	GoogleMapsLink    string         `json:"google_maps_link"`
	FacebookLink      string         `json:"facebook_link"`
	WhatsappLink      string         `json:"whatsapp_link"`
	Private           bool           `json:"is_private"`
	AccessCodeEnabled bool           `json:"access_code_enabled"`
	AccessCode        string         `json:"access_code"` //starts with the id of the party and after that has a '_' character. The "{partyId}_" part is appended to the code in the buisness logic. the user only sends the code part of the code
	OrganizerID       uint           `json:"organizer_id"`
	Organizer         domains.User   `json:"organizer"`
	Participants      []domains.User `json:"-" gorm:"many2many:party_participants;"`
}

// todo: refactor these methods into the usecases folder
func (p *Party) CanBeAccessedBy(userId uint) bool {
	return p.HasParticipant(userId) || userId == adminUser.ADMIN_USER_ID //we dont check for private bc it only means anyone cna join them
}

func (p *Party) CanBeOrganizedBy(userId uint) bool {
	return p.OrganizerID == userId || userId == adminUser.ADMIN_USER_ID
}

func (p *Party) HasParticipant(userId uint) bool {
	if p.OrganizerID == userId {
		return true
	}

	for _, participant := range p.Participants {
		if participant.ID == userId {
			return true
		}
	}
	return false
}
