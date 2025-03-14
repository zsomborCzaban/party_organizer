package db

import (
	"github.com/stretchr/testify/mock"
)

type MockDatabaseAccessManager struct {
	mock.Mock
}

func (m *MockDatabaseAccessManager) RegisterEntity(name string, dbEntityProvider IEntityProvider) IDatabaseAccess {
	args := m.Called(name, dbEntityProvider)
	return args.Get(0).(IDatabaseAccess)
}

func (m *MockDatabaseAccessManager) GetRegisteredDBAccess(name string) IDatabaseAccess {
	args := m.Called(name)
	return args.Get(0).(IDatabaseAccess)
}

func (m *MockDatabaseAccessManager) Close() error {
	args := m.Called()
	return args.Error(0)
}
