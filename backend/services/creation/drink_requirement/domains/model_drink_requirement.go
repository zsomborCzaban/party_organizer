package domains

import (
	"github.com/zsomborCzaban/party_organizer/services/creation/party/domains"
	"gorm.io/gorm"
)

type DrinkRequirement struct {
	gorm.Model

	PartyID        uint          `json:"party_id"`
	Party          domains.Party `json:"-"`
	Type           string        `json:"type" validate:"required"`
	TargetQuantity int           `json:"target_quantity" validate:"required,gt=0"`
	QuantityMark   string        `json:"quantity_mark" validate:"required"`
	Description    string        `json:"description"`
}

func (d *DrinkRequirement) TransformToDrinkRequirementDTO() *DrinkRequirementDTO {
	return &DrinkRequirementDTO{
		ID:             d.ID,
		PartyID:        d.PartyID,
		Type:           d.Type,
		TargetQuantity: d.TargetQuantity,
		QuantityMark:   d.QuantityMark,
		Description:    d.Description,
	}
}
