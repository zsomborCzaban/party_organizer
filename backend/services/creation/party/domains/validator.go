package domains

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type Validator struct {
	v *validator.Validate
}

func NewValidator(v *validator.Validate) *Validator { return &Validator{v: v} }

func (val *Validator) Validate(data interface{}) *ValidationErrors {
	ret := NewValidationErrors()

	err := val.v.Struct(data)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			ret.CollectValidationError(err.Field(), val.CustomErrorMessage(err), err.Value())
		}

		return ret
	}

	return nil
}

func (val *Validator) CustomErrorMessage(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", err.Field())
	case "min":
		return fmt.Sprintf("%s must be at least %s characters long", err.Field(), err.Param())
	default:
		return fmt.Sprintf("%s is not valid", err.Field())
	}
}
