package domains

type IValidator interface {
	Validate(data interface{}) *ValidationErrors
}
