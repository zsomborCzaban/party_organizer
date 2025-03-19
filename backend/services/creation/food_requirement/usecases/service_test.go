package usecases

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zsomborCzaban/party_organizer/services/creation/food_requirement/domains"
	domains2 "github.com/zsomborCzaban/party_organizer/services/creation/party/domains"
	partyUsecases "github.com/zsomborCzaban/party_organizer/services/creation/party/usecases"
	foodContributionUsecases "github.com/zsomborCzaban/party_organizer/services/interaction/food_contributions/usecases"
	"github.com/zsomborCzaban/party_organizer/utils/api"
	repo2 "github.com/zsomborCzaban/party_organizer/utils/repo"
	"testing"
)

func setupDefaultService() (domains.IFoodRequirementService, *api.MockValidator, *MockRepository, *partyUsecases.MockRepository, *foodContributionUsecases.MockRepository) {
	foodReqRepo := new(MockRepository)
	partyRepo := new(partyUsecases.MockRepository)
	foodContribRepo := new(foodContributionUsecases.MockRepository)
	repo := repo2.RepoCollector{
		PartyRepo:        partyRepo,
		UserRepo:         nil,
		DrinkReqRepo:     nil,
		DrinkContribRepo: nil,
		FoodReqRepo:      foodReqRepo,
		FoodContribRepo:  foodContribRepo,
		PartyInviteRepo:  nil,
		FriendInviteRepo: nil,
	}
	validator := new(api.MockValidator)
	service := NewFoodRequirementService(&repo, validator)

	return service, validator, foodReqRepo, partyRepo, foodContribRepo
}

func Test_ServiceCreate_Success(t *testing.T) {
	service, vali, foodReqRepo, partyRepo, _ := setupDefaultService()

	userId := uint(1)
	foodReqDto := domains.FoodRequirementDTO{
		PartyID: 1,
	}
	party := domains2.Party{
		OrganizerID: userId,
	}

	vali.On("Validate", mock.Anything).Return(nil)
	partyRepo.On("FindById", mock.Anything, mock.Anything).Return(&party, nil)
	foodReqRepo.On("Create", mock.Anything).Return(nil)

	service.Create(foodReqDto, userId)
}

func Test_ServiceCreate_FailOnValidator(t *testing.T) {
	service, vali, _, _, _ := setupDefaultService()

	userId := uint(1)
	foodReqDto := domains.FoodRequirementDTO{
		PartyID: 1,
	}

	validationError := api.NewValidationErrors()
	expectedResponse := api.ErrorValidation(validationError)

	vali.On("Validate", mock.Anything).Return(validationError)

	response := service.Create(foodReqDto, userId)

	assert.Equal(t, expectedResponse, response)
}

func Test_ServiceCreate_FailOnNoParty(t *testing.T) {
	service, vali, _, partyRepo, _ := setupDefaultService()

	userId := uint(1)
	foodReqDto := domains.FoodRequirementDTO{
		PartyID: 1,
	}

	expectedResponse := api.ErrorBadRequest(domains.PartyNotFound)

	vali.On("Validate", mock.Anything).Return(nil)
	partyRepo.On("FindById", mock.Anything, mock.Anything).Return(&domains2.Party{}, errors.New(""))

	response := service.Create(foodReqDto, userId)

	assert.Equal(t, expectedResponse, response)
}

func Test_ServiceCreate_FailOnPartyAuth(t *testing.T) {
	service, vali, _, partyRepo, _ := setupDefaultService()

	userId := uint(1)
	foodReqDto := domains.FoodRequirementDTO{
		PartyID: 1,
	}
	party := domains2.Party{
		OrganizerID: 2,
	}
	expectedResponse := api.ErrorUnauthorized(domains.NoOrganizerAccess)

	vali.On("Validate", mock.Anything).Return(nil)
	partyRepo.On("FindById", mock.Anything, mock.Anything).Return(&party, nil)

	response := service.Create(foodReqDto, userId)

	assert.Equal(t, expectedResponse, response)
}

func Test_ServiceCreate_FailOnCreate(t *testing.T) {
	service, vali, foodReqRepo, partyRepo, _ := setupDefaultService()

	userId := uint(1)
	foodReqDto := domains.FoodRequirementDTO{
		PartyID: 1,
	}
	party := domains2.Party{
		OrganizerID: userId,
	}
	err := errors.New("error")
	expectedResponse := api.ErrorInternalServerError(err.Error())

	vali.On("Validate", mock.Anything).Return(nil)
	partyRepo.On("FindById", mock.Anything, mock.Anything).Return(&party, nil)
	foodReqRepo.On("Create", mock.Anything).Return(err)

	response := service.Create(foodReqDto, userId)

	assert.Equal(t, expectedResponse, response)
}

func Test_ServiceGet_Success(t *testing.T) {
	service, _, foodReqRepo, _, _ := setupDefaultService()

	userId := uint(1)
	foodReq := domains.FoodRequirement{
		Party: domains2.Party{
			OrganizerID: userId,
		},
	}
	expectedResponse := api.Success(&foodReq)

	foodReqRepo.On("FindById", mock.Anything, mock.Anything).Return(&foodReq, nil)

	response := service.FindById(0, userId)

	assert.Equal(t, expectedResponse, response)
}

func Test_ServiceGet_FailOnFind(t *testing.T) {
	service, _, foodReqRepo, _, _ := setupDefaultService()

	err := errors.New("error")
	expectedResponse := api.ErrorBadRequest(domains.RequirementNotFound)

	foodReqRepo.On("FindById", mock.Anything, mock.Anything).Return(&domains.FoodRequirement{}, err)

	response := service.FindById(0, 0)

	assert.Equal(t, expectedResponse, response)
}

func Test_ServiceGet_FailOnPartyAuth(t *testing.T) {
	service, _, foodReqRepo, _, _ := setupDefaultService()

	userId := uint(1)
	foodReq := domains.FoodRequirement{
		Party: domains2.Party{
			OrganizerID: 2,
		},
	}
	expectedResponse := api.ErrorUnauthorized(domains.NoViewAccess)

	foodReqRepo.On("FindById", mock.Anything, mock.Anything).Return(&foodReq, nil)

	response := service.FindById(0, userId)

	assert.Equal(t, expectedResponse, response)
}

func Test_ServiceDelete_Success(t *testing.T) {
	service, _, foodReqRepo, _, foodContribRepo := setupDefaultService()

	userId := uint(1)
	foodReq := domains.FoodRequirement{
		Party: domains2.Party{
			OrganizerID: userId,
		},
	}
	expectedResponse := api.Success("delete_success")

	foodReqRepo.On("FindById", mock.Anything, mock.Anything).Return(&foodReq, nil)
	foodContribRepo.On("DeleteByReqId", mock.Anything).Return(nil)
	foodReqRepo.On("Delete", mock.Anything).Return(nil)

	response := service.Delete(0, userId)

	assert.Equal(t, expectedResponse, response)
}

func Test_ServiceDelete_FailOnFind(t *testing.T) {
	service, _, foodReqRepo, _, _ := setupDefaultService()

	userId := uint(1)
	foodReq := domains.FoodRequirement{}
	expectedResponse := api.ErrorBadRequest(domains.RequirementNotFound)

	foodReqRepo.On("FindById", mock.Anything, mock.Anything).Return(&foodReq, errors.New(""))

	response := service.Delete(0, userId)

	assert.Equal(t, expectedResponse, response)
}

func Test_ServiceDelete_FailOnPartyAuth(t *testing.T) {
	service, _, foodReqRepo, _, _ := setupDefaultService()

	userId := uint(1)
	foodReq := domains.FoodRequirement{
		Party: domains2.Party{
			OrganizerID: 2,
		},
	}
	expectedResponse := api.ErrorUnauthorized(domains.NoOrganizerAccess)

	foodReqRepo.On("FindById", mock.Anything, mock.Anything).Return(&foodReq, nil)

	response := service.Delete(0, userId)

	assert.Equal(t, expectedResponse, response)
}

func Test_ServiceDelete_FailOnContributionDelete(t *testing.T) {
	service, _, foodReqRepo, _, foodContribRepo := setupDefaultService()

	userId := uint(1)
	foodReq := domains.FoodRequirement{
		Party: domains2.Party{
			OrganizerID: userId,
		},
	}
	err := errors.New("")
	expectedResponse := api.ErrorInternalServerError(err.Error())

	foodReqRepo.On("FindById", mock.Anything, mock.Anything).Return(&foodReq, nil)
	foodContribRepo.On("DeleteByReqId", mock.Anything).Return(err)

	response := service.Delete(0, userId)

	assert.Equal(t, expectedResponse, response)
}

func Test_ServiceDelete_FailOnDelete(t *testing.T) {
	service, _, foodReqRepo, _, foodContribRepo := setupDefaultService()

	userId := uint(1)
	foodReq := domains.FoodRequirement{
		Party: domains2.Party{
			OrganizerID: userId,
		},
	}
	err := errors.New("")
	expectedResponse := api.ErrorInternalServerError(err.Error())

	foodReqRepo.On("FindById", mock.Anything, mock.Anything).Return(&foodReq, nil)
	foodContribRepo.On("DeleteByReqId", mock.Anything).Return(nil)
	foodReqRepo.On("Delete", mock.Anything).Return(err)

	response := service.Delete(0, userId)

	assert.Equal(t, expectedResponse, response)
}

func Test_ServiceGetByPartyId_Success(t *testing.T) {
	service, _, foodReqRepo, partyRepo, _ := setupDefaultService()

	userId := uint(1)
	party := domains2.Party{
		OrganizerID: userId,
	}
	foodReqs := []domains.FoodRequirement{{
		Party: domains2.Party{
			OrganizerID: userId,
		},
	}}
	expectedResponse := api.Success(&foodReqs)

	partyRepo.On("FindById", mock.Anything, mock.Anything).Return(&party, nil)
	foodReqRepo.On("GetByPartyId", mock.Anything, mock.Anything).Return(&foodReqs, nil)

	response := service.GetByPartyId(0, userId)

	assert.Equal(t, expectedResponse, response)
}

func Test_ServiceGetByPartyId_FailOnPartyFind(t *testing.T) {
	service, _, _, partyRepo, _ := setupDefaultService()

	userId := uint(1)
	party := domains2.Party{
		OrganizerID: userId,
	}
	expectedResponse := api.ErrorBadRequest(domains.PartyNotFound)

	partyRepo.On("FindById", mock.Anything, mock.Anything).Return(&party, errors.New(""))

	response := service.GetByPartyId(0, userId)

	assert.Equal(t, expectedResponse, response)
}

func Test_ServiceGetByPartyId_FailOnPartyAuth(t *testing.T) {
	service, _, _, partyRepo, _ := setupDefaultService()

	userId := uint(1)
	party := domains2.Party{
		OrganizerID: 2,
	}
	expectedResponse := api.ErrorUnauthorized(domains.NoViewAccess)

	partyRepo.On("FindById", mock.Anything, mock.Anything).Return(&party, nil)

	response := service.GetByPartyId(0, userId)

	assert.Equal(t, expectedResponse, response)
}

func Test_ServiceGetByPartyId_FailOnFind(t *testing.T) {
	service, _, foodReqRepo, partyRepo, _ := setupDefaultService()

	userId := uint(1)
	party := domains2.Party{
		OrganizerID: userId,
	}
	foodReqs := []domains.FoodRequirement{{
		Party: domains2.Party{
			OrganizerID: userId,
		},
	}}
	err := errors.New("")
	expectedResponse := api.ErrorInternalServerError(err.Error())

	partyRepo.On("FindById", mock.Anything, mock.Anything).Return(&party, nil)
	foodReqRepo.On("GetByPartyId", mock.Anything, mock.Anything).Return(&foodReqs, err)

	response := service.GetByPartyId(0, userId)

	assert.Equal(t, expectedResponse, response)
}
