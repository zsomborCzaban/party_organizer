package api

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"time"
)

type Validator struct {
	v *validator.Validate
}

func NewValidator(v *validator.Validate) *Validator {
	registerCustomValidations(v)
	return &Validator{v: v}

}

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
	//case "containsany":
	//	return fmt.Sprintf("%s must contain one of these characters: %s", err.Field(), err.Param())
	case "containsany":
		return fmt.Sprintf("%s must contain a number", err.Field())
	case "eqfield":
		return fmt.Sprintf("%s must match %s", err.Field(), err.Param())
	case "gt":
		return fmt.Sprintf("%s must be greater than %s", err.Field(), err.Param())
	case "email":
		return fmt.Sprintf("%s must be a valid email", err.Field())
	case "http_url":
		return fmt.Sprintf("%s must be a valid http url", err.Field())
	case "bool_allowed_by_bool":
		return fmt.Sprintf("%s cannot be true if %s is false", err.Field(), err.Param())
	case "string_allowed_by_bool_and_min_3":
		return fmt.Sprintf("%s has to be longer than 2 characters, if %s is true", err.Field(), err.Param())
	case "string_allowed_by_bool":
		return fmt.Sprintf("%s has to be empty, if %s is false", err.Field(), err.Param())
	case "after_24_hours":
		return fmt.Sprintf("%s must be 1 day after the current time", err.Field())
	default:
		return fmt.Sprintf("%s is not valid", err.Field())
	}
}

func registerCustomValidations(v *validator.Validate) {
	v.RegisterValidation("bool_allowed_by_bool", func(fl validator.FieldLevel) bool {
		field := fl.Field()
		param := fl.Param()
		otherField := fl.Parent().FieldByName(param)

		if !otherField.Bool() && field.Bool() {
			return false
		}

		return true
	})

	v.RegisterValidation("string_allowed_by_bool", func(fl validator.FieldLevel) bool {
		field := fl.Field()
		param := fl.Param()
		otherField := fl.Parent().FieldByName(param)

		if !otherField.Bool() && field.String() != "" {
			return false
		}

		return true
	})

	v.RegisterValidation("string_allowed_by_bool_and_min_3", func(fl validator.FieldLevel) bool {
		field := fl.Field()
		param := fl.Param()
		otherField := fl.Parent().FieldByName(param)

		if otherField.Bool() {
			return len(field.String()) >= 3
		}

		return true
	})

	v.RegisterValidation("after_24_hours", func(fl validator.FieldLevel) bool {
		field := fl.Field()

		fieldTime, ok := field.Interface().(time.Time)
		if !ok {
			return false
		}

		return fieldTime.After(time.Now().Add(24 * time.Hour))
	})
}
