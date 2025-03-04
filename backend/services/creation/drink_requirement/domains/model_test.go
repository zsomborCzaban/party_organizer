package domains

import (
	"github.com/go-playground/assert/v2"
	"gorm.io/gorm"
	"testing"
)

func TestTransformToDrinkRequirementDTO(t *testing.T) {
	expectedDto := DrinkRequirementDTO{
		ID:             1,
		PartyID:        2,
		Type:           "3",
		TargetQuantity: 4,
		QuantityMark:   "5",
	}

	drinkRequirement := DrinkRequirement{
		Model: gorm.Model{
			ID: 1,
		},
		PartyID:        2,
		Type:           "3",
		TargetQuantity: 4,
		QuantityMark:   "5",
	}
	dto := drinkRequirement.TransformToDrinkRequirementDTO()

	assert.Equal(t, expectedDto.ID, dto.ID)
	assert.Equal(t, expectedDto.PartyID, dto.PartyID)
	assert.Equal(t, expectedDto.Type, dto.Type)
	assert.Equal(t, expectedDto.TargetQuantity, dto.TargetQuantity)
	assert.Equal(t, expectedDto.QuantityMark, dto.QuantityMark)
}
