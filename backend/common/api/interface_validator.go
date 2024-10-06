package api

type IValidator interface {
	Validate(data interface{}) *ValidationErrors
}
