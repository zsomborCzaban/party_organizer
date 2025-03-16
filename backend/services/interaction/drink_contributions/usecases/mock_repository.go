package usecases

import (
	"github.com/stretchr/testify/mock"
	"github.com/zsomborCzaban/party_organizer/services/interaction/drink_contributions/domains"
)

type MockRepository struct {
	mock.Mock
}

func (mr *MockRepository) Create(requirement *domains.DrinkContribution) error {
	args := mr.Called(requirement)
	return args.Error(0)
}

func (mr *MockRepository) Update(requirement *domains.DrinkContribution) error {
	args := mr.Called(requirement)
	return args.Error(0)
}

func (mr *MockRepository) Delete(requirement *domains.DrinkContribution) error {
	args := mr.Called(requirement)
	return args.Error(0)
}

func (mr *MockRepository) DeleteByPartyId(id uint) error {
	args := mr.Called(id)
	return args.Error(0)
}

func (mr *MockRepository) DeleteByReqId(id uint) error {
	args := mr.Called(id)
	return args.Error(0)
}

func (mr *MockRepository) DeleteByContributorId(id uint) error {
	args := mr.Called(id)
	return args.Error(0)
}

func (mr *MockRepository) FindAllBy(columnNames []string, values []interface{}, associations ...string) (*[]domains.DrinkContribution, error) {
	args := mr.Called(columnNames, values, associations)
	return args.Get(0).(*[]domains.DrinkContribution), args.Error(1)
}

func (mr *MockRepository) FindById(id uint, associations ...string) (*domains.DrinkContribution, error) {
	args := mr.Called(id, associations)
	return args.Get(0).(*domains.DrinkContribution), args.Error(1)
}
