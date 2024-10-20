package api

type FieldError struct {
	Field string      `json:"field"`
	Value interface{} `json:"value"`
	Err   string      `json:"err"`
}

func NewValidationError(field string, err string, val interface{}) *FieldError {
	return &FieldError{
		Field: field,
		Err:   err,
		Value: val,
	}
}

type ValidationErrors struct {
	Errors []FieldError
}

func NewValidationErrors() *ValidationErrors {
	return &ValidationErrors{Errors: make([]FieldError, 0)}
}

func (ve *ValidationErrors) CollectValidationError(field string, err string, val interface{}) {
	valError := NewValidationError(field, err, val)

	validationErrors := append(ve.Errors, *valError)

	ve.Errors = validationErrors
}
