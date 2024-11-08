package domains

import "gorm.io/gorm"

type DrinkRequirementDTO struct {
	ID             uint    `json:"id,omitempty"`
	PartyID        uint    `json:"party_id,omitempty" validate:"required"`
	Type           string  `json:"type,omitempty" validate:"required"`
	TargetQuantity float32 `json:"target_quantity,omitempty" validate:"required,min=1,gt=0"`
	QuantityMark   string  `json:"quantity_mark,omitempty" validate:"required,min=1"`
}

func (drDTO *DrinkRequirementDTO) TransformToDrinkRequirement() *DrinkRequirement {
	return &DrinkRequirement{
		Model:          gorm.Model{ID: drDTO.ID},
		PartyID:        drDTO.PartyID,
		Type:           drDTO.Type,
		TargetQuantity: drDTO.TargetQuantity,
		QuantityMark:   drDTO.QuantityMark,
	}
}
