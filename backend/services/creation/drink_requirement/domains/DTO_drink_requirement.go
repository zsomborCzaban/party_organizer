package domains

type DrinkRequirementDTO struct {
	ID             uint   `json:"id,omitempty"`
	PartyID        uint   `json:"party_id,omitempty"`
	Type           string `json:"type,omitempty" validate:"required"`
	TargetQuantity int    `json:"target_quantity,omitempty" validate:"required,min=1"`
	QuantityMark   string `json:"quantity_mark,omitempty" validate:"required,min=1"`
	Description    string `json:"description,omitempty" `
}
