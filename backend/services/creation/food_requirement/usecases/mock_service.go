package usecases

import (
	"github.com/stretchr/testify/mock"
	"github.com/zsomborCzaban/party_organizer/services/creation/food_requirement/domains"
	"github.com/zsomborCzaban/party_organizer/utils/api"
)

type MockService struct {
	mock.Mock
}

func (ms *MockService) Create(dr domains.FoodRequirementDTO, userId uint) api.IResponse {
	args := ms.Called(dr, userId)
	return args.Get(0).(api.IResponse)
}
func (ms *MockService) FindById(foodReqId, userId uint) api.IResponse {
	args := ms.Called(foodReqId, userId)
	return args.Get(0).(api.IResponse)
}
func (ms *MockService) Delete(foodReqId, userId uint) api.IResponse {
	args := ms.Called(foodReqId, userId)
	return args.Get(0).(api.IResponse)
}
func (ms *MockService) GetByPartyId(partyId, userId uint) api.IResponse {
	args := ms.Called(partyId, userId)
	return args.Get(0).(api.IResponse)
}
