package api

import (
	"github.com/stretchr/testify/mock"
	"net/http"
)

type MockResponseWriter struct {
	mock.Mock
}

func (m *MockResponseWriter) Header() http.Header {
	args := m.Called()
	if args.Get(0) == nil {
		return make(http.Header)
	}
	return args.Get(0).(http.Header)
}

func (m *MockResponseWriter) Write(b []byte) (int, error) {
	args := m.Called(b)
	return args.Int(0), args.Error(1)
}

func (m *MockResponseWriter) WriteHeader(statusCode int) {
	m.Called(statusCode)
	return
}
