package usecases

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	drinkReqDomain "github.com/zsomborCzaban/party_organizer/services/creation/drink_requirement/domains"
	drinkReqUsecases "github.com/zsomborCzaban/party_organizer/services/creation/drink_requirement/usecases"
	partyDomains "github.com/zsomborCzaban/party_organizer/services/creation/party/domains"
	partyUsecases "github.com/zsomborCzaban/party_organizer/services/creation/party/usecases"
	"github.com/zsomborCzaban/party_organizer/services/interaction/drink_contributions/domains"
	userDomain "github.com/zsomborCzaban/party_organizer/services/user/domains"
	userUsecases "github.com/zsomborCzaban/party_organizer/services/user/usecases"
	"github.com/zsomborCzaban/party_organizer/utils/api"
	"github.com/zsomborCzaban/party_organizer/utils/repo"
	"gorm.io/gorm"
	"testing"
)

func setupDefaultService() (domains.IDrinkContributionService, *api.MockValidator, *MockRepository, *partyUsecases.MockRepository, *userUsecases.MockRepository, *drinkReqUsecases.MockRepository) {
	validator := new(api.MockValidator)
	contribRepo := new(MockRepository)
	partyRepo := new(partyUsecases.MockRepository)
	userRepo := new(userUsecases.MockRepository)
	drinkReqRepo := new(drinkReqUsecases.MockRepository)

	repoCollector := &repo.RepoCollector{
		DrinkContribRepo: contribRepo,
		PartyRepo:        partyRepo,
		UserRepo:         userRepo,
		DrinkReqRepo:     drinkReqRepo,
	}

	service := NewDrinkContributionService(repoCollector, validator)

	return service, validator, contribRepo, partyRepo, userRepo, drinkReqRepo
}

func Test_DrinkContributionService_Create_Success(t *testing.T) {
	service, validator, contribRepo, _, _, drinkReqRepo := setupDefaultService()

	userId := uint(1)
	contribution := domains.DrinkContribution{
		DrinkReqId: 1,
	}
	req := &drinkReqDomain.DrinkRequirement{
		Party: partyDomains.Party{
			OrganizerID: 2,
			Participants: []userDomain.User{
				{Model: gorm.Model{ID: userId}},
			},
		},
	}

	validator.On("Validate", contribution).Return(nil)
	drinkReqRepo.On("FindById", contribution.DrinkReqId, mock.Anything).Return(req, nil)
	contribRepo.On("Create", mock.Anything).Return(nil)

	response := service.Create(contribution, userId)

	assert.False(t, response.GetIsError())
}

func Test_DrinkContributionService_Create_FailOnValidation(t *testing.T) {
	service, validator, _, _, _, _ := setupDefaultService()

	userId := uint(1)
	contribution := domains.DrinkContribution{}
	validationError := api.NewValidationErrors()

	validator.On("Validate", contribution).Return(validationError)

	response := service.Create(contribution, userId)

	assert.Equal(t, api.ErrorValidation(validationError.Errors), response)
}

func Test_DrinkContributionService_Create_FailOnReqNotFound(t *testing.T) {
	service, validator, _, _, _, drinkReqRepo := setupDefaultService()

	userId := uint(1)
	contribution := domains.DrinkContribution{
		DrinkReqId: 1,
	}
	expectedErr := errors.New("requirement not found")

	validator.On("Validate", contribution).Return(nil)
	drinkReqRepo.On("FindById", contribution.DrinkReqId, mock.Anything).Return(&drinkReqDomain.DrinkRequirement{}, expectedErr)

	response := service.Create(contribution, userId)

	assert.Equal(t, api.ErrorBadRequest(expectedErr.Error()), response)
}

func Test_DrinkContributionService_Create_FailOnNoAccess(t *testing.T) {
	service, validator, _, _, _, drinkReqRepo := setupDefaultService()

	userId := uint(1)
	contribution := domains.DrinkContribution{
		DrinkReqId: 1,
	}
	req := &drinkReqDomain.DrinkRequirement{
		Party: partyDomains.Party{
			OrganizerID: 2, // Different organizer
			Private:     true,
		},
	}

	validator.On("Validate", contribution).Return(nil)
	drinkReqRepo.On("FindById", contribution.DrinkReqId, mock.Anything).Return(req, nil)

	response := service.Create(contribution, userId)

	assert.Equal(t, api.ErrorUnauthorized(domains.NO_ACCESS_TO_PARTY), response)
}

func Test_DrinkContributionService_Create_FailOnCreate(t *testing.T) {
	service, validator, contribRepo, _, _, drinkReqRepo := setupDefaultService()

	userId := uint(1)
	contribution := domains.DrinkContribution{
		DrinkReqId: 1,
	}
	req := &drinkReqDomain.DrinkRequirement{
		Party: partyDomains.Party{
			OrganizerID: 2,
			Participants: []userDomain.User{
				{Model: gorm.Model{ID: userId}},
			},
		},
	}
	expectedErr := errors.New("create failed")

	validator.On("Validate", contribution).Return(nil)
	drinkReqRepo.On("FindById", contribution.DrinkReqId, mock.Anything).Return(req, nil)
	contribRepo.On("Create", mock.Anything).Return(expectedErr)

	response := service.Create(contribution, userId)

	assert.Equal(t, api.ErrorInternalServerError(expectedErr.Error()), response)
}

func Test_DrinkContributionService_Update_Success(t *testing.T) {
	service, validator, contribRepo, _, _, _ := setupDefaultService()

	userId := uint(1)
	contribution := domains.DrinkContribution{
		Model:      gorm.Model{ID: 1},
		DrinkReqId: 1,
	}
	oldContribution := &domains.DrinkContribution{
		Model:         gorm.Model{ID: 1},
		DrinkReqId:    1,
		ContributorId: userId,
		Party: partyDomains.Party{
			OrganizerID: 2,
			Participants: []userDomain.User{
				{Model: gorm.Model{ID: userId}},
			},
		},
	}

	validator.On("Validate", contribution).Return(nil)
	contribRepo.On("FindById", contribution.ID, mock.Anything).Return(oldContribution, nil)
	contribRepo.On("Update", mock.Anything).Return(nil)

	response := service.Update(contribution, userId)

	assert.False(t, response.GetIsError())
}

func Test_DrinkContributionService_Update_FailOnValidation(t *testing.T) {
	service, validator, _, _, _, _ := setupDefaultService()

	userId := uint(1)
	contribution := domains.DrinkContribution{}
	validationError := api.NewValidationErrors()

	validator.On("Validate", contribution).Return(validationError)

	response := service.Update(contribution, userId)

	assert.Equal(t, api.ErrorValidation(validationError.Errors), response)
}

func Test_DrinkContributionService_Update_FailOnFind(t *testing.T) {
	service, validator, contribRepo, _, _, _ := setupDefaultService()

	userId := uint(1)
	contribution := domains.DrinkContribution{Model: gorm.Model{ID: 1}}
	expectedErr := errors.New("not found")

	validator.On("Validate", contribution).Return(nil)
	contribRepo.On("FindById", contribution.ID, mock.Anything).Return(&domains.DrinkContribution{}, expectedErr)

	response := service.Update(contribution, userId)

	assert.Equal(t, api.ErrorBadRequest(expectedErr.Error()), response)
}

func Test_DrinkContributionService_Update_FailOnNoAccess(t *testing.T) {
	service, validator, contribRepo, _, _, _ := setupDefaultService()

	userId := uint(1)
	contribution := domains.DrinkContribution{Model: gorm.Model{ID: 1}}
	oldContribution := &domains.DrinkContribution{
		Model: gorm.Model{ID: 1},
		Party: partyDomains.Party{
			OrganizerID: 2,
			Private:     true,
		},
	}

	validator.On("Validate", contribution).Return(nil)
	contribRepo.On("FindById", contribution.ID, mock.Anything).Return(oldContribution, nil)

	response := service.Update(contribution, userId)

	assert.Equal(t, api.ErrorUnauthorized(domains.NO_ACCESS_TO_PARTY), response)
}

func Test_DrinkContributionService_Update_FailOnNotOwner(t *testing.T) {
	service, validator, contribRepo, _, _, _ := setupDefaultService()

	userId := uint(1)
	contribution := domains.DrinkContribution{Model: gorm.Model{ID: 1}}
	oldContribution := &domains.DrinkContribution{
		Model:         gorm.Model{ID: 1},
		ContributorId: 2, // Different contributor
		Party: partyDomains.Party{
			OrganizerID: 3,
			Participants: []userDomain.User{
				{Model: gorm.Model{ID: userId}},
			},
		},
	}

	validator.On("Validate", contribution).Return(nil)
	contribRepo.On("FindById", contribution.ID, mock.Anything).Return(oldContribution, nil)

	response := service.Update(contribution, userId)

	assert.Equal(t, api.ErrorBadRequest("cannot update other people's contribution"), response)
}

func Test_DrinkContributionService_Update_FailOnReqChange(t *testing.T) {
	service, validator, contribRepo, _, _, _ := setupDefaultService()

	userId := uint(1)
	contribution := domains.DrinkContribution{
		Model:      gorm.Model{ID: 1},
		DrinkReqId: 2, // Different from original
	}
	oldContribution := &domains.DrinkContribution{
		Model:         gorm.Model{ID: 1},
		DrinkReqId:    1,
		ContributorId: userId,
		Party: partyDomains.Party{
			OrganizerID: 2,
			Participants: []userDomain.User{
				{Model: gorm.Model{ID: userId}},
			},
		},
	}

	validator.On("Validate", contribution).Return(nil)
	contribRepo.On("FindById", contribution.ID, mock.Anything).Return(oldContribution, nil)

	response := service.Update(contribution, userId)

	assert.Equal(t, api.ErrorBadRequest("cannot change drink requirement of the contribution"), response)
}

func Test_DrinkContributionService_Update_FailOnUpdate(t *testing.T) {
	service, validator, contribRepo, _, _, _ := setupDefaultService()

	userId := uint(1)
	contribution := domains.DrinkContribution{
		Model:      gorm.Model{ID: 1},
		DrinkReqId: 1,
	}
	oldContribution := &domains.DrinkContribution{
		Model:         gorm.Model{ID: 1},
		DrinkReqId:    1,
		ContributorId: userId,
		Party: partyDomains.Party{
			OrganizerID: 2,
			Participants: []userDomain.User{
				{Model: gorm.Model{ID: userId}},
			},
		},
	}
	expectedErr := errors.New("update failed")

	validator.On("Validate", contribution).Return(nil)
	contribRepo.On("FindById", contribution.ID, mock.Anything).Return(oldContribution, nil)
	contribRepo.On("Update", mock.Anything).Return(expectedErr)

	response := service.Update(contribution, userId)

	assert.Equal(t, api.ErrorInternalServerError(expectedErr.Error()), response)
}

func Test_DrinkContributionService_Delete_Success(t *testing.T) {
	service, _, contribRepo, _, _, _ := setupDefaultService()

	userId := uint(1)
	contributionId := uint(1)
	contribution := &domains.DrinkContribution{
		Model:         gorm.Model{ID: contributionId},
		ContributorId: userId,
		Party: partyDomains.Party{
			OrganizerID: 2,
		},
	}

	contribRepo.On("FindById", contributionId, mock.Anything).Return(contribution, nil)
	contribRepo.On("Delete", contribution).Return(nil)

	response := service.Delete(contributionId, userId)

	assert.False(t, response.GetIsError())
}

func Test_DrinkContributionService_Delete_SuccessAsOrganizer(t *testing.T) {
	service, _, contribRepo, _, _, _ := setupDefaultService()

	userId := uint(1)
	contributionId := uint(1)
	contribution := &domains.DrinkContribution{
		Model:         gorm.Model{ID: contributionId},
		ContributorId: 2, // Different contributor
		Party: partyDomains.Party{
			OrganizerID: userId, // But user is organizer
		},
	}

	contribRepo.On("FindById", contributionId, mock.Anything).Return(contribution, nil)
	contribRepo.On("Delete", contribution).Return(nil)

	response := service.Delete(contributionId, userId)

	assert.False(t, response.GetIsError())
}

func Test_DrinkContributionService_Delete_FailOnFind(t *testing.T) {
	service, _, contribRepo, _, _, _ := setupDefaultService()

	userId := uint(1)
	contributionId := uint(1)
	expectedErr := errors.New("not found")

	contribRepo.On("FindById", contributionId, mock.Anything).Return(&domains.DrinkContribution{}, expectedErr)

	response := service.Delete(contributionId, userId)

	assert.Equal(t, api.ErrorBadRequest(expectedErr.Error()), response)
}

func Test_DrinkContributionService_Delete_FailOnUnauthorized(t *testing.T) {
	service, _, contribRepo, _, _, _ := setupDefaultService()

	userId := uint(1)
	contributionId := uint(1)
	contribution := &domains.DrinkContribution{
		Model:         gorm.Model{ID: contributionId},
		ContributorId: 2, // Different contributor
		Party: partyDomains.Party{
			OrganizerID: 3, // And not organizer
			Private:     true,
		},
	}

	contribRepo.On("FindById", contributionId, mock.Anything).Return(contribution, nil)

	response := service.Delete(contributionId, userId)

	assert.Equal(t, api.ErrorUnauthorized("cannot delete other people's contribution"), response)
}

func Test_DrinkContributionService_Delete_FailOnDelete(t *testing.T) {
	service, _, contribRepo, _, _, _ := setupDefaultService()

	userId := uint(1)
	contributionId := uint(1)
	contribution := &domains.DrinkContribution{
		Model:         gorm.Model{ID: contributionId},
		ContributorId: userId,
		Party: partyDomains.Party{
			OrganizerID: 2,
		},
	}
	expectedErr := errors.New("delete failed")

	contribRepo.On("FindById", contributionId, mock.Anything).Return(contribution, nil)
	contribRepo.On("Delete", contribution).Return(expectedErr)

	response := service.Delete(contributionId, userId)

	assert.Equal(t, api.ErrorInternalServerError(expectedErr.Error()), response)
}

func Test_DrinkContributionService_GetByPartyIdAndContributorId_Success(t *testing.T) {
	service, _, contribRepo, partyRepo, _, _ := setupDefaultService()

	userId := uint(1)
	partyId := uint(1)
	contributorId := uint(2)
	party := &partyDomains.Party{
		OrganizerID: 3,
		Participants: []userDomain.User{
			{Model: gorm.Model{ID: userId}},
		},
	}
	contributions := []domains.DrinkContribution{
		{Model: gorm.Model{ID: 1}},
	}

	partyRepo.On("FindById", mock.Anything, mock.Anything).Return(party, nil)
	contribRepo.On("FindAllBy", []string{"party_id", "contributor_id"}, []interface{}{partyId, contributorId}, []string{"Contributor", "DrinkReq"}).Return(&contributions, nil)

	response := service.GetByPartyIdAndContributorId(partyId, contributorId, userId)

	assert.False(t, response.GetIsError())
}

func Test_DrinkContributionService_GetByPartyIdAndContributorId_FailOnPartyFind(t *testing.T) {
	service, _, _, partyRepo, _, _ := setupDefaultService()

	userId := uint(1)
	partyId := uint(1)
	contributorId := uint(2)
	party := &partyDomains.Party{
		OrganizerID: 3,
		Participants: []userDomain.User{
			{Model: gorm.Model{ID: userId}},
		},
	}
	expectedErr := errors.New("party not found")

	partyRepo.On("FindById", mock.Anything, mock.Anything).Return(party, expectedErr)

	response := service.GetByPartyIdAndContributorId(partyId, contributorId, userId)

	assert.Equal(t, response, api.ErrorBadRequest(expectedErr.Error()))
}

func Test_DrinkContributionService_GetByPartyIdAndContributorId_FailOnPartyAuth(t *testing.T) {
	service, _, _, partyRepo, _, _ := setupDefaultService()

	userId := uint(1)
	partyId := uint(1)
	contributorId := uint(2)
	party := &partyDomains.Party{
		OrganizerID: 3,
	}

	partyRepo.On("FindById", mock.Anything, mock.Anything).Return(party, nil)

	response := service.GetByPartyIdAndContributorId(partyId, contributorId, userId)

	assert.Equal(t, response, api.ErrorUnauthorized(domains.NO_ACCESS_TO_PARTY))
}

func Test_DrinkContributionService_GetByPartyIdAndContributorId_FailOnContribFind(t *testing.T) {
	service, _, contribRepo, partyRepo, _, _ := setupDefaultService()

	userId := uint(1)
	partyId := uint(1)
	contributorId := uint(2)
	party := &partyDomains.Party{
		OrganizerID: 3,
		Participants: []userDomain.User{
			{Model: gorm.Model{ID: userId}},
		},
	}
	contributions := []domains.DrinkContribution{
		{Model: gorm.Model{ID: 1}},
	}
	expectedErr := errors.New("")

	partyRepo.On("FindById", mock.Anything, mock.Anything).Return(party, nil)
	contribRepo.On("FindAllBy", []string{"party_id", "contributor_id"}, []interface{}{partyId, contributorId}, []string{"Contributor", "DrinkReq"}).Return(&contributions, expectedErr)

	response := service.GetByPartyIdAndContributorId(partyId, contributorId, userId)

	assert.Equal(t, response, api.ErrorInternalServerError(expectedErr.Error()))
}

func Test_DrinkContributionService_GetByRequirementId_Success(t *testing.T) {
	service, _, contribRepo, _, _, drinkReqRepo := setupDefaultService()

	userId := uint(1)
	requirementId := uint(1)
	req := &drinkReqDomain.DrinkRequirement{
		Party: partyDomains.Party{
			OrganizerID: 2,
			Participants: []userDomain.User{
				{Model: gorm.Model{ID: userId}},
			},
		},
	}
	contributions := []domains.DrinkContribution{
		{Model: gorm.Model{ID: 1}},
	}

	drinkReqRepo.On("FindById", requirementId, mock.Anything).Return(req, nil)
	contribRepo.On("FindAllBy", []string{"drink_req_id"}, []interface{}{requirementId}, []string{"Contributor", "DrinkReq"}).Return(&contributions, nil)

	response := service.GetByRequirementId(requirementId, userId)

	assert.False(t, response.GetIsError())
}

func Test_DrinkContributionService_GetByRequirementId_FailOnPartyFind(t *testing.T) {
	service, _, _, _, _, drinkReqRepo := setupDefaultService()

	userId := uint(1)
	requirementId := uint(1)
	req := &drinkReqDomain.DrinkRequirement{
		Party: partyDomains.Party{
			OrganizerID: 2,
			Participants: []userDomain.User{
				{Model: gorm.Model{ID: userId}},
			},
		},
	}
	expectedErr := errors.New("")

	drinkReqRepo.On("FindById", requirementId, mock.Anything).Return(req, expectedErr)

	response := service.GetByRequirementId(requirementId, userId)

	assert.Equal(t, response, api.ErrorBadRequest(expectedErr.Error()))
}

func Test_DrinkContributionService_GetByRequirementId_FailOnPartyAuth(t *testing.T) {
	service, _, _, _, _, drinkReqRepo := setupDefaultService()

	userId := uint(1)
	requirementId := uint(1)
	req := &drinkReqDomain.DrinkRequirement{
		Party: partyDomains.Party{
			OrganizerID: 2,
		},
	}

	drinkReqRepo.On("FindById", requirementId, mock.Anything).Return(req, nil)

	response := service.GetByRequirementId(requirementId, userId)

	assert.Equal(t, response, api.ErrorUnauthorized(domains.NO_ACCESS_TO_PARTY))
}

func Test_DrinkContributionService_GetByRequirementId_FailOnContribFind(t *testing.T) {
	service, _, contribRepo, _, _, drinkReqRepo := setupDefaultService()

	userId := uint(1)
	requirementId := uint(1)
	req := &drinkReqDomain.DrinkRequirement{
		Party: partyDomains.Party{
			OrganizerID: 2,
			Participants: []userDomain.User{
				{Model: gorm.Model{ID: userId}},
			},
		},
	}
	contributions := []domains.DrinkContribution{
		{Model: gorm.Model{ID: 1}},
	}
	expectedErr := errors.New("")

	drinkReqRepo.On("FindById", requirementId, mock.Anything).Return(req, nil)
	contribRepo.On("FindAllBy", []string{"drink_req_id"}, []interface{}{requirementId}, []string{"Contributor", "DrinkReq"}).Return(&contributions, expectedErr)

	response := service.GetByRequirementId(requirementId, userId)

	assert.Equal(t, response, api.ErrorInternalServerError(expectedErr.Error()))
}

func Test_DrinkContributionService_GetByPartyId_Success(t *testing.T) {
	service, _, contribRepo, partyRepo, _, _ := setupDefaultService()

	userId := uint(1)
	partyId := uint(1)
	party := &partyDomains.Party{
		OrganizerID: 2,
		Participants: []userDomain.User{
			{Model: gorm.Model{ID: userId}},
		},
	}
	contributions := []domains.DrinkContribution{
		{Model: gorm.Model{ID: 1}},
	}

	partyRepo.On("FindById", mock.Anything, mock.Anything).Return(party, nil)
	contribRepo.On("FindAllBy", []string{"party_id"}, []interface{}{partyId}, []string{"Contributor", "DrinkReq"}).Return(&contributions, nil)

	response := service.GetByPartyId(partyId, userId)

	assert.False(t, response.GetIsError())
}

func Test_DrinkContributionService_GetByPartyId_FailOnPartyFind(t *testing.T) {
	service, _, _, partyRepo, _, _ := setupDefaultService()

	userId := uint(1)
	partyId := uint(1)
	party := &partyDomains.Party{
		OrganizerID: 2,
		Participants: []userDomain.User{
			{Model: gorm.Model{ID: userId}},
		},
	}
	expectedErr := errors.New("")

	partyRepo.On("FindById", mock.Anything, mock.Anything).Return(party, expectedErr)

	response := service.GetByPartyId(partyId, userId)

	assert.Equal(t, response, api.ErrorBadRequest(expectedErr.Error()))
}

func Test_DrinkContributionService_GetByPartyId_FailOnPartyAuth(t *testing.T) {
	service, _, _, partyRepo, _, _ := setupDefaultService()

	userId := uint(1)
	partyId := uint(1)
	party := &partyDomains.Party{
		OrganizerID: 2,
	}

	partyRepo.On("FindById", mock.Anything, mock.Anything).Return(party, nil)

	response := service.GetByPartyId(partyId, userId)

	assert.Equal(t, response, api.ErrorUnauthorized(domains.NO_ACCESS_TO_PARTY))
}

func Test_DrinkContributionService_GetByPartyId_FailOnContribFind(t *testing.T) {
	service, _, contribRepo, partyRepo, _, _ := setupDefaultService()

	userId := uint(1)
	partyId := uint(1)
	party := &partyDomains.Party{
		OrganizerID: 2,
		Participants: []userDomain.User{
			{Model: gorm.Model{ID: userId}},
		},
	}
	contributions := []domains.DrinkContribution{
		{Model: gorm.Model{ID: 1}},
	}
	expectedErr := errors.New("")

	partyRepo.On("FindById", mock.Anything, mock.Anything).Return(party, nil)
	contribRepo.On("FindAllBy", []string{"party_id"}, []interface{}{partyId}, []string{"Contributor", "DrinkReq"}).Return(&contributions, expectedErr)

	response := service.GetByPartyId(partyId, userId)

	assert.Equal(t, response, api.ErrorInternalServerError(expectedErr.Error()))
}
