package domains

import (
	"github.com/zsomborCzaban/party_organizer/common/adminUser"
	"github.com/zsomborCzaban/party_organizer/services/user/domains"
	"gorm.io/gorm"
	"time"
)

type Party struct {
	gorm.Model

	Place              string    `json:"place"`
	StartTime          time.Time `json:"start_time"`
	Name               string    `json:"name"`
	GoogleMapsPlusCode string
	FacebookLink       string
	WhatsappLink       string
	Private            bool           `json:"is_private"`
	AccessCodeEnabled  bool           `json:"access_code_enabled"`
	AccessCode         string         `json:"access_code"` //starts with the id of the party and after that has a '_' character. The "partyId_" part is appended to the code in the buisness logic. the user only sends the code part of the code
	OrganizerID        uint           `json:"organizer_id"`
	Organizer          domains.User   `json:"organizer"`
	Participants       []domains.User `json:"-" gorm:"many2many:party_participants;"`
}

func (p *Party) TransformToPartyDTO() *PartyDTO {
	return &PartyDTO{
		ID:                 p.ID,
		Place:              p.Place,
		StartTime:          p.StartTime,
		Name:               p.Name,
		GoogleMapsPlusCode: p.GoogleMapsPlusCode,
		FacebookLink:       p.FacebookLink,
		WhatsappLink:       p.WhatsappLink,
		Private:            p.Private,
		AccessCodeEnabled:  p.AccessCodeEnabled,
		AccessCode:         p.AccessCode,
		OrganizerID:        p.OrganizerID,
	}
}

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
