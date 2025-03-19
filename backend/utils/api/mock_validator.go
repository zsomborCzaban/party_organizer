package api

import "github.com/stretchr/testify/mock"

type MockValidator struct {
	mock.Mock
}

func (v *MockValidator) Validate(entity interface{}) *ValidationErrors {
	args := v.Called(entity)

	if args.Get(0) == nil {
		return nil
	}

	return args.Get(0).(*ValidationErrors)
}
