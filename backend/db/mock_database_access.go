package db

import "github.com/stretchr/testify/mock"

type MockDatabaseAccess struct {
	mock.Mock
}

func (m *MockDatabaseAccess) Create(entity interface{}) error {
	args := m.Called(entity)
	return args.Error(0)
}

func (m *MockDatabaseAccess) FindById(id interface{}, associations ...string) (interface{}, error) {
	args := m.Called(id, associations)
	return args.Get(0), args.Error(1)
}

func (m *MockDatabaseAccess) FindAll(associations ...string) (interface{}, error) {
	args := m.Called(associations)
	return args.Get(0), args.Error(1)
}

func (m *MockDatabaseAccess) Update(entity interface{}) error {
	args := m.Called(entity)
	return args.Error(0)
}

func (m *MockDatabaseAccess) Delete(value interface{}) error {
	args := m.Called(value)
	return args.Error(0)
}

func (m *MockDatabaseAccess) BatchDelete(conds []QueryParameter) error {
	args := m.Called(conds)
	return args.Error(0)
}

func (m *MockDatabaseAccess) Query(conds []QueryParameter, associations ...string) (interface{}, error) {
	args := m.Called(conds, associations)
	return args.Get(0), args.Error(1)
}

func (m *MockDatabaseAccess) Many2ManyQueryId(cond Many2ManyQueryParameter, associations ...string) (interface{}, error) {
	args := m.Called(cond, associations)
	return args.Get(0), args.Error(1)
}

func (m *MockDatabaseAccess) ReplaceAssociations(param AssociationParameter) error {
	args := m.Called(param)
	return args.Error(0)
}
func (m *MockDatabaseAccess) TransactionBegin() IDatabaseAccess {
	args := m.Called()
	return args.Get(0).(IDatabaseAccess)
}
func (m *MockDatabaseAccess) TransactionCommit() error {
	args := m.Called()
	return args.Error(0)
}
func (m *MockDatabaseAccess) TransactionRollback() error {
	args := m.Called()
	return args.Error(0)
}
