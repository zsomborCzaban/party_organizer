package domains

import "gorm.io/gorm"

type FoodRequirementDTO struct {
	ID             uint    `json:"id,omitempty"`
	PartyID        uint    `json:"party_id,omitempty" validate:"required"`
	Type           string  `json:"type,omitempty" validate:"required"`
	TargetQuantity float32 `json:"target_quantity,omitempty" validate:"required,min=1,gt=0"`
	QuantityMark   string  `json:"quantity_mark,omitempty" validate:"required,min=1"`
}

func (frDTO *FoodRequirementDTO) TransformToFoodRequirement() *FoodRequirement {
	return &FoodRequirement{
		Model:          gorm.Model{ID: frDTO.ID},
		PartyID:        frDTO.PartyID,
		Type:           frDTO.Type,
		TargetQuantity: frDTO.TargetQuantity,
		QuantityMark:   frDTO.QuantityMark,
	}
}
