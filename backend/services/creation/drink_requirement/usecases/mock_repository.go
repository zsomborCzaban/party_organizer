package usecases

import (
	"github.com/stretchr/testify/mock"
	"github.com/zsomborCzaban/party_organizer/services/creation/drink_requirement/domains"
)

type MockRepository struct {
	mock.Mock
}

func (mr *MockRepository) Create(requirement *domains.DrinkRequirement) error {
	args := mr.Called(requirement)
	return args.Error(0)
}

func (mr *MockRepository) FindById(id uint, associations ...string) (*domains.DrinkRequirement, error) {
	args := mr.Called(id, associations)
	return args.Get(0).(*domains.DrinkRequirement), args.Error(1)
}

func (mr *MockRepository) Delete(requirement *domains.DrinkRequirement) error {
	args := mr.Called(requirement)
	return args.Error(0)
}

func (mr *MockRepository) DeleteByPartyId(id uint) error {
	args := mr.Called(id)
	return args.Error(0)
}

func (mr *MockRepository) GetByPartyId(id uint) (*[]domains.DrinkRequirement, error) {
	args := mr.Called(id)
	return args.Get(0).(*[]domains.DrinkRequirement), args.Error(1)
}
