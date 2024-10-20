package domains

import (
	"github.com/zsomborCzaban/party_organizer/services/creation/party/domains"
	"gorm.io/gorm"
)

type FoodRequirement struct {
	gorm.Model

	PartyID        uint          `json:"party_id"`
	Party          domains.Party `json:"-"`
	Type           string        `json:"type" validate:"required"`
	TargetQuantity int           `json:"target_quantity" validate:"required,gt=0"`
	QuantityMark   string        `json:"quantity_mark" validate:"required"`
	Description    string        `json:"description"`
}

func (fr *FoodRequirement) TransformToFoodRequirementDTO() *FoodRequirementDTO {
	return &FoodRequirementDTO{
		ID:             fr.ID,
		PartyID:        fr.PartyID,
		Type:           fr.Type,
		TargetQuantity: fr.TargetQuantity,
		QuantityMark:   fr.QuantityMark,
		Description:    fr.Description,
	}
}
