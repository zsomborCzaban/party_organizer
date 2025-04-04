package usecases

import (
	"github.com/stretchr/testify/mock"
	"github.com/zsomborCzaban/party_organizer/services/creation/drink_requirement/domains"
	"github.com/zsomborCzaban/party_organizer/utils/api"
)

type MockService struct {
	mock.Mock
}

func (ms *MockService) Create(dr domains.DrinkRequirementDTO, userId uint) api.IResponse {
	args := ms.Called(dr, userId)
	return args.Get(0).(api.IResponse)
}
func (ms *MockService) FindById(drinkReqId, userId uint) api.IResponse {
	args := ms.Called(drinkReqId, userId)
	return args.Get(0).(api.IResponse)
}
func (ms *MockService) Delete(drinkReqId, userId uint) api.IResponse {
	args := ms.Called(drinkReqId, userId)
	return args.Get(0).(api.IResponse)
}
func (ms *MockService) GetByPartyId(partyId, userId uint) api.IResponse {
	args := ms.Called(partyId, userId)
	return args.Get(0).(api.IResponse)
}
