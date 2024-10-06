package domains

import (
	"gorm.io/gorm"
)

type FoodRequirement struct {
	gorm.Model

	PartyID        uint   `json:"party_id"`
	Type           string `json:"type"`
	TargetQuantity int    `json:"target_quantity"`
	QuantityMark   string `json:"quantity_mark"`
	Description    string `json:"description"`
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
