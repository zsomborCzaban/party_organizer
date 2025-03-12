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

func (m *MockDatabaseAccess) AddToAssociation(entity interface{}, association string, associatedEntities ...interface{}) error {
	args := m.Called(entity, association, associatedEntities)
	return args.Error(0)
}

func (m *MockDatabaseAccess) DeleteFromAssociation(entity interface{}, association string, associatedEntities ...interface{}) error {
	args := m.Called(entity, association, associatedEntities)
	return args.Error(0)
}

func (m *MockDatabaseAccess) ClearAssociation(entity interface{}, association string) error {
	args := m.Called(entity, association)
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
