package domains

import (
	"github.com/go-playground/assert/v2"
	"gorm.io/gorm"
	"testing"
)

func Test_TransformToDrinkRequirement(t *testing.T) {
	dto := DrinkRequirementDTO{
		ID:             1,
		PartyID:        2,
		Type:           "3",
		TargetQuantity: 4,
		QuantityMark:   "5",
	}
	drinkRequirement := dto.TransformToDrinkRequirement()

	expectedDrinkRequirement := DrinkRequirement{
		Model: gorm.Model{
			ID: 1,
		},
		PartyID:        2,
		Type:           "3",
		TargetQuantity: 4,
		QuantityMark:   "5",
	}

	assert.Equal(t, expectedDrinkRequirement.ID, drinkRequirement.ID)
	assert.Equal(t, expectedDrinkRequirement.PartyID, drinkRequirement.PartyID)
	assert.Equal(t, expectedDrinkRequirement.Type, drinkRequirement.Type)
	assert.Equal(t, expectedDrinkRequirement.TargetQuantity, drinkRequirement.TargetQuantity)
	assert.Equal(t, expectedDrinkRequirement.QuantityMark, drinkRequirement.QuantityMark)
}
