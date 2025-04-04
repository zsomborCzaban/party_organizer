package usecases

import (
	"github.com/stretchr/testify/mock"
	"github.com/zsomborCzaban/party_organizer/services/interaction/food_contributions/domains"
	"github.com/zsomborCzaban/party_organizer/utils/api"
)

type MockService struct {
	mock.Mock
}

func (m *MockService) Create(contribution domains.FoodContribution, userId uint) api.IResponse {
	args := m.Called(contribution, userId)
	return args.Get(0).(api.IResponse)
}

func (m *MockService) Update(contribution domains.FoodContribution, userId uint) api.IResponse {
	args := m.Called(contribution, userId)
	return args.Get(0).(api.IResponse)
}

func (m *MockService) Delete(contributionId, userId uint) api.IResponse {
	args := m.Called(contributionId, userId)
	return args.Get(0).(api.IResponse)
}

func (m *MockService) GetByPartyIdAndContributorId(partyId, contributorId, userId uint) api.IResponse {
	args := m.Called(partyId, contributorId, userId)
	return args.Get(0).(api.IResponse)
}

func (m *MockService) GetByRequirementId(requirementId, userId uint) api.IResponse {
	args := m.Called(requirementId, userId)
	return args.Get(0).(api.IResponse)
}

func (m *MockService) GetByPartyId(partyId, userId uint) api.IResponse {
	args := m.Called(partyId, userId)
	return args.Get(0).(api.IResponse)
}
