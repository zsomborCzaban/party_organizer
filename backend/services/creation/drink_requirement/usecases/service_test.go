package usecases

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zsomborCzaban/party_organizer/services/creation/drink_requirement/domains"
	domains2 "github.com/zsomborCzaban/party_organizer/services/creation/party/domains"
	partyUsecases "github.com/zsomborCzaban/party_organizer/services/creation/party/usecases"
	drinkContributionUsecases "github.com/zsomborCzaban/party_organizer/services/interaction/drink_contributions/usecases"
	"github.com/zsomborCzaban/party_organizer/utils/api"
	repo2 "github.com/zsomborCzaban/party_organizer/utils/repo"
	"testing"
)

func setupDefaultService() (domains.IDrinkRequirementService, *api.MockValidator, *MockRepository, *partyUsecases.MockRepository, *drinkContributionUsecases.MockRepository) {
	drinkReqRepo := new(MockRepository)
	partyRepo := new(partyUsecases.MockRepository)
	drinkContribRepo := new(drinkContributionUsecases.MockRepository)
	repo := repo2.RepoCollector{
		PartyRepo:        partyRepo,
		UserRepo:         nil,
		DrinkReqRepo:     drinkReqRepo,
		DrinkContribRepo: drinkContribRepo,
		FoodReqRepo:      nil,
		FoodContribRepo:  nil,
		PartyInviteRepo:  nil,
		FriendInviteRepo: nil,
	}
	validator := new(api.MockValidator)
	service := NewDrinkRequirementService(&repo, validator)

	return service, validator, drinkReqRepo, partyRepo, drinkContribRepo
}

func Test_ServiceCreate_Success(t *testing.T) {
	service, vali, drinkReqRepo, partyRepo, _ := setupDefaultService()

	userId := uint(1)
	drinkReqDto := domains.DrinkRequirementDTO{
		PartyID: 1,
	}
	party := domains2.Party{
		OrganizerID: userId,
	}

	vali.On("Validate", mock.Anything).Return(nil)
	partyRepo.On("FindById", mock.Anything, mock.Anything).Return(&party, nil)
	drinkReqRepo.On("Create", mock.Anything).Return(nil)

	service.Create(drinkReqDto, userId)
}

func Test_ServiceCreate_FailOnValidator(t *testing.T) {
	service, vali, _, _, _ := setupDefaultService()

	userId := uint(1)
	drinkReqDto := domains.DrinkRequirementDTO{
		PartyID: 1,
	}

	validationError := api.NewValidationErrors()
	expectedResponse := api.ErrorValidation(validationError)

	vali.On("Validate", mock.Anything).Return(validationError)

	response := service.Create(drinkReqDto, userId)

	assert.Equal(t, expectedResponse, response)
}

func Test_ServiceCreate_FailOnNoParty(t *testing.T) {
	service, vali, _, partyRepo, _ := setupDefaultService()

	userId := uint(1)
	drinkReqDto := domains.DrinkRequirementDTO{
		PartyID: 1,
	}

	expectedResponse := api.ErrorBadRequest(domains.PartyNotFound)

	vali.On("Validate", mock.Anything).Return(nil)
	partyRepo.On("FindById", mock.Anything, mock.Anything).Return(&domains2.Party{}, errors.New(""))

	response := service.Create(drinkReqDto, userId)

	assert.Equal(t, expectedResponse, response)
}

func Test_ServiceCreate_FailOnPartyAuth(t *testing.T) {
	service, vali, _, partyRepo, _ := setupDefaultService()

	userId := uint(1)
	drinkReqDto := domains.DrinkRequirementDTO{
		PartyID: 1,
	}
	party := domains2.Party{
		OrganizerID: 2,
	}
	expectedResponse := api.ErrorUnauthorized(domains.NoOrganizerAccess)

	vali.On("Validate", mock.Anything).Return(nil)
	partyRepo.On("FindById", mock.Anything, mock.Anything).Return(&party, nil)

	response := service.Create(drinkReqDto, userId)

	assert.Equal(t, expectedResponse, response)
}

func Test_ServiceCreate_FailOnCreate(t *testing.T) {
	service, vali, drinkReqRepo, partyRepo, _ := setupDefaultService()

	userId := uint(1)
	drinkReqDto := domains.DrinkRequirementDTO{
		PartyID: 1,
	}
	party := domains2.Party{
		OrganizerID: userId,
	}
	err := errors.New("error")
	expectedResponse := api.ErrorInternalServerError(err.Error())

	vali.On("Validate", mock.Anything).Return(nil)
	partyRepo.On("FindById", mock.Anything, mock.Anything).Return(&party, nil)
	drinkReqRepo.On("Create", mock.Anything).Return(err)

	response := service.Create(drinkReqDto, userId)

	assert.Equal(t, expectedResponse, response)
}

func Test_ServiceGet_Success(t *testing.T) {
	service, _, drinkReqRepo, _, _ := setupDefaultService()

	userId := uint(1)
	drinkReq := domains.DrinkRequirement{
		Party: domains2.Party{
			OrganizerID: userId,
		},
	}
	expectedResponse := api.Success(&drinkReq)

	drinkReqRepo.On("FindById", mock.Anything, mock.Anything).Return(&drinkReq, nil)

	response := service.FindById(0, userId)

	assert.Equal(t, expectedResponse, response)
}

func Test_ServiceGet_FailOnFind(t *testing.T) {
	service, _, drinkReqRepo, _, _ := setupDefaultService()

	err := errors.New("error")
	expectedResponse := api.ErrorBadRequest(domains.RequirementNotFound)

	drinkReqRepo.On("FindById", mock.Anything, mock.Anything).Return(&domains.DrinkRequirement{}, err)

	response := service.FindById(0, 0)

	assert.Equal(t, expectedResponse, response)
}

func Test_ServiceGet_FailOnPartyAuth(t *testing.T) {
	service, _, drinkReqRepo, _, _ := setupDefaultService()

	userId := uint(1)
	drinkReq := domains.DrinkRequirement{
		Party: domains2.Party{
			OrganizerID: 2,
		},
	}
	expectedResponse := api.ErrorUnauthorized(domains.NoViewAccess)

	drinkReqRepo.On("FindById", mock.Anything, mock.Anything).Return(&drinkReq, nil)

	response := service.FindById(0, userId)

	assert.Equal(t, expectedResponse, response)
}

func Test_ServiceDelete_Success(t *testing.T) {
	service, _, drinkReqRepo, _, drinkContribRepo := setupDefaultService()

	userId := uint(1)
	drinkReq := domains.DrinkRequirement{
		Party: domains2.Party{
			OrganizerID: userId,
		},
	}
	expectedResponse := api.Success("delete_success")

	drinkReqRepo.On("FindById", mock.Anything, mock.Anything).Return(&drinkReq, nil)
	drinkContribRepo.On("DeleteByReqId", mock.Anything).Return(nil)
	drinkReqRepo.On("Delete", mock.Anything).Return(nil)

	response := service.Delete(0, userId)

	assert.Equal(t, expectedResponse, response)
}

func Test_ServiceDelete_FailOnFind(t *testing.T) {
	service, _, drinkReqRepo, _, _ := setupDefaultService()

	userId := uint(1)
	drinkReq := domains.DrinkRequirement{}
	expectedResponse := api.ErrorBadRequest(domains.RequirementNotFound)

	drinkReqRepo.On("FindById", mock.Anything, mock.Anything).Return(&drinkReq, errors.New(""))

	response := service.Delete(0, userId)

	assert.Equal(t, expectedResponse, response)
}

func Test_ServiceDelete_FailOnPartyAuth(t *testing.T) {
	service, _, drinkReqRepo, _, _ := setupDefaultService()

	userId := uint(1)
	drinkReq := domains.DrinkRequirement{
		Party: domains2.Party{
			OrganizerID: 2,
		},
	}
	expectedResponse := api.ErrorUnauthorized(domains.NoOrganizerAccess)

	drinkReqRepo.On("FindById", mock.Anything, mock.Anything).Return(&drinkReq, nil)

	response := service.Delete(0, userId)

	assert.Equal(t, expectedResponse, response)
}

func Test_ServiceDelete_FailOnContributionDelete(t *testing.T) {
	service, _, drinkReqRepo, _, drinkContribRepo := setupDefaultService()

	userId := uint(1)
	drinkReq := domains.DrinkRequirement{
		Party: domains2.Party{
			OrganizerID: userId,
		},
	}
	err := errors.New("")
	expectedResponse := api.ErrorInternalServerError(err.Error())

	drinkReqRepo.On("FindById", mock.Anything, mock.Anything).Return(&drinkReq, nil)
	drinkContribRepo.On("DeleteByReqId", mock.Anything).Return(err)

	response := service.Delete(0, userId)

	assert.Equal(t, expectedResponse, response)
}

func Test_ServiceDelete_FailOnDelete(t *testing.T) {
	service, _, drinkReqRepo, _, drinkContribRepo := setupDefaultService()

	userId := uint(1)
	drinkReq := domains.DrinkRequirement{
		Party: domains2.Party{
			OrganizerID: userId,
		},
	}
	err := errors.New("")
	expectedResponse := api.ErrorInternalServerError(err.Error())

	drinkReqRepo.On("FindById", mock.Anything, mock.Anything).Return(&drinkReq, nil)
	drinkContribRepo.On("DeleteByReqId", mock.Anything).Return(nil)
	drinkReqRepo.On("Delete", mock.Anything).Return(err)

	response := service.Delete(0, userId)

	assert.Equal(t, expectedResponse, response)
}

func Test_ServiceGetByPartyId_Success(t *testing.T) {
	service, _, drinkReqRepo, partyRepo, _ := setupDefaultService()

	userId := uint(1)
	party := domains2.Party{
		OrganizerID: userId,
	}
	drinkReqs := []domains.DrinkRequirement{{
		Party: domains2.Party{
			OrganizerID: userId,
		},
	}}
	expectedResponse := api.Success(&drinkReqs)

	partyRepo.On("FindById", mock.Anything, mock.Anything).Return(&party, nil)
	drinkReqRepo.On("GetByPartyId", mock.Anything, mock.Anything).Return(&drinkReqs, nil)

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
	service, _, drinkReqRepo, partyRepo, _ := setupDefaultService()

	userId := uint(1)
	party := domains2.Party{
		OrganizerID: userId,
	}
	drinkReqs := []domains.DrinkRequirement{{
		Party: domains2.Party{
			OrganizerID: userId,
		},
	}}
	err := errors.New("")
	expectedResponse := api.ErrorInternalServerError(err.Error())

	partyRepo.On("FindById", mock.Anything, mock.Anything).Return(&party, nil)
	drinkReqRepo.On("GetByPartyId", mock.Anything, mock.Anything).Return(&drinkReqs, err)

	response := service.GetByPartyId(0, userId)

	assert.Equal(t, expectedResponse, response)
}
