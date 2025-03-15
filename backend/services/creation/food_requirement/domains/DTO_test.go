package domains

import (
	"github.com/go-playground/assert/v2"
	"gorm.io/gorm"
	"testing"
)

func TestTransformToFoodRequirement(t *testing.T) {
	dto := FoodRequirementDTO{
		ID:             1,
		PartyID:        2,
		Type:           "3",
		TargetQuantity: 4,
		QuantityMark:   "5",
	}
	foodRequirement := dto.TransformToFoodRequirement()

	expectedFoodRequirement := FoodRequirement{
		Model: gorm.Model{
			ID: 1,
		},
		PartyID:        2,
		Type:           "3",
		TargetQuantity: 4,
		QuantityMark:   "5",
	}

	assert.Equal(t, expectedFoodRequirement.ID, foodRequirement.ID)
	assert.Equal(t, expectedFoodRequirement.PartyID, foodRequirement.PartyID)
	assert.Equal(t, expectedFoodRequirement.Type, foodRequirement.Type)
	assert.Equal(t, expectedFoodRequirement.TargetQuantity, foodRequirement.TargetQuantity)
	assert.Equal(t, expectedFoodRequirement.QuantityMark, foodRequirement.QuantityMark)
}
