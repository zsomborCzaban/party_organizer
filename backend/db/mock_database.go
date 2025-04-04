package db

import (
	"github.com/stretchr/testify/mock"
)

type MockDatabase struct {
	mock.Mock
}

func (m *MockDatabase) AutoMigrate(dest ...interface{}) error {
	args := m.Called(dest)
	return args.Error(0)
}

func (m *MockDatabase) NewSession() {
	m.Called()
}

func (m *MockDatabase) Create(value interface{}) error {
	args := m.Called(value)
	return args.Error(0)
}

func (m *MockDatabase) First(dest interface{}, preloadColumns []string, conds ...interface{}) error {
	args := m.Called(dest, preloadColumns, conds)
	return args.Error(0)
}

func (m *MockDatabase) Find(dest interface{}, preloadColumns []string, conds ...interface{}) error {
	args := m.Called(dest, preloadColumns, conds)
	return args.Error(0)
}

func (m *MockDatabase) Update(entity interface{}) error {
	args := m.Called(entity)
	return args.Error(0)
}

func (m *MockDatabase) Delete(value interface{}, conds ...interface{}) error {
	args := m.Called(value, conds)
	return args.Error(0)
}

func (m *MockDatabase) ProcessWhereStatements(conds []QueryParameter) {
	m.Called(conds)
}

func (m *MockDatabase) Many2ManyQueryId(dest interface{}, associations []string, cond Many2ManyQueryParameter) error {
	args := m.Called(dest, associations, cond)
	return args.Error(0)
}

func (m *MockDatabase) ReplaceAssociations(param AssociationParameter) error {
	args := m.Called(param)
	return args.Error(0)
}

func (m *MockDatabase) TransactionBegin() IDatabase {
	args := m.Called()
	return args.Get(0).(IDatabase)
}
func (m *MockDatabase) TransactionCommit() error {
	args := m.Called()
	return args.Error(0)
}
func (m *MockDatabase) TransactionRollback() error {
	args := m.Called()
	return args.Error(0)
}
