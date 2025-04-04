package usecases

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	foodReqDomain "github.com/zsomborCzaban/party_organizer/services/creation/food_requirement/domains"
	foodReqUsecases "github.com/zsomborCzaban/party_organizer/services/creation/food_requirement/usecases"
	partyDomains "github.com/zsomborCzaban/party_organizer/services/creation/party/domains"
	partyUsecases "github.com/zsomborCzaban/party_organizer/services/creation/party/usecases"
	"github.com/zsomborCzaban/party_organizer/services/interaction/food_contributions/domains"
	userDomain "github.com/zsomborCzaban/party_organizer/services/user/domains"
	userUsecases "github.com/zsomborCzaban/party_organizer/services/user/usecases"
	"github.com/zsomborCzaban/party_organizer/utils/api"
	"github.com/zsomborCzaban/party_organizer/utils/repo"
	"gorm.io/gorm"
	"testing"
)

func setupDefaultService() (domains.IFoodContributionService, *api.MockValidator, *MockRepository, *partyUsecases.MockRepository, *userUsecases.MockRepository, *foodReqUsecases.MockRepository) {
	validator := new(api.MockValidator)
	contribRepo := new(MockRepository)
	partyRepo := new(partyUsecases.MockRepository)
	userRepo := new(userUsecases.MockRepository)
	foodReqRepo := new(foodReqUsecases.MockRepository)

	repoCollector := &repo.RepoCollector{
		FoodContribRepo: contribRepo,
		PartyRepo:       partyRepo,
		UserRepo:        userRepo,
		FoodReqRepo:     foodReqRepo,
	}

	service := NewFoodContributionService(repoCollector, validator)

	return service, validator, contribRepo, partyRepo, userRepo, foodReqRepo
}

func Test_FoodContributionService_Create_Success(t *testing.T) {
	service, validator, contribRepo, _, _, foodReqRepo := setupDefaultService()

	userId := uint(1)
	contribution := domains.FoodContribution{
		FoodReqId: 1,
	}
	req := &foodReqDomain.FoodRequirement{
		Party: partyDomains.Party{
			OrganizerID: 2,
			Participants: []userDomain.User{
				{Model: gorm.Model{ID: userId}},
			},
		},
	}

	validator.On("Validate", contribution).Return(nil)
	foodReqRepo.On("FindById", contribution.FoodReqId, mock.Anything).Return(req, nil)
	contribRepo.On("Create", mock.Anything).Return(nil)

	response := service.Create(contribution, userId)

	assert.False(t, response.GetIsError())
}

func Test_FoodContributionService_Create_FailOnValidation(t *testing.T) {
	service, validator, _, _, _, _ := setupDefaultService()

	userId := uint(1)
	contribution := domains.FoodContribution{}
	validationError := api.NewValidationErrors()

	validator.On("Validate", contribution).Return(validationError)

	response := service.Create(contribution, userId)

	assert.Equal(t, api.ErrorValidation(validationError.Errors), response)
}

func Test_FoodContributionService_Create_FailOnReqNotFound(t *testing.T) {
	service, validator, _, _, _, foodReqRepo := setupDefaultService()

	userId := uint(1)
	contribution := domains.FoodContribution{
		FoodReqId: 1,
	}
	expectedErr := errors.New("requirement not found")

	validator.On("Validate", contribution).Return(nil)
	foodReqRepo.On("FindById", contribution.FoodReqId, mock.Anything).Return(&foodReqDomain.FoodRequirement{}, expectedErr)

	response := service.Create(contribution, userId)

	assert.Equal(t, api.ErrorBadRequest(expectedErr.Error()), response)
}

func Test_FoodContributionService_Create_FailOnNoAccess(t *testing.T) {
	service, validator, _, _, _, foodReqRepo := setupDefaultService()

	userId := uint(1)
	contribution := domains.FoodContribution{
		FoodReqId: 1,
	}
	req := &foodReqDomain.FoodRequirement{
		Party: partyDomains.Party{
			OrganizerID: 2, // Different organizer
			Private:     true,
		},
	}

	validator.On("Validate", contribution).Return(nil)
	foodReqRepo.On("FindById", contribution.FoodReqId, mock.Anything).Return(req, nil)

	response := service.Create(contribution, userId)

	assert.Equal(t, api.ErrorUnauthorized(domains.NO_ACCESS_TO_PARTY), response)
}

func Test_FoodContributionService_Create_FailOnCreate(t *testing.T) {
	service, validator, contribRepo, _, _, foodReqRepo := setupDefaultService()

	userId := uint(1)
	contribution := domains.FoodContribution{
		FoodReqId: 1,
	}
	req := &foodReqDomain.FoodRequirement{
		Party: partyDomains.Party{
			OrganizerID: 2,
			Participants: []userDomain.User{
				{Model: gorm.Model{ID: userId}},
			},
		},
	}
	expectedErr := errors.New("create failed")

	validator.On("Validate", contribution).Return(nil)
	foodReqRepo.On("FindById", contribution.FoodReqId, mock.Anything).Return(req, nil)
	contribRepo.On("Create", mock.Anything).Return(expectedErr)

	response := service.Create(contribution, userId)

	assert.Equal(t, api.ErrorInternalServerError(expectedErr.Error()), response)
}

func Test_FoodContributionService_Update_Success(t *testing.T) {
	service, validator, contribRepo, _, _, _ := setupDefaultService()

	userId := uint(1)
	contribution := domains.FoodContribution{
		Model:     gorm.Model{ID: 1},
		FoodReqId: 1,
	}
	oldContribution := &domains.FoodContribution{
		Model:         gorm.Model{ID: 1},
		FoodReqId:     1,
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

func Test_FoodContributionService_Update_FailOnValidation(t *testing.T) {
	service, validator, _, _, _, _ := setupDefaultService()

	userId := uint(1)
	contribution := domains.FoodContribution{}
	validationError := api.NewValidationErrors()

	validator.On("Validate", contribution).Return(validationError)

	response := service.Update(contribution, userId)

	assert.Equal(t, api.ErrorValidation(validationError.Errors), response)
}

func Test_FoodContributionService_Update_FailOnFind(t *testing.T) {
	service, validator, contribRepo, _, _, _ := setupDefaultService()

	userId := uint(1)
	contribution := domains.FoodContribution{Model: gorm.Model{ID: 1}}
	expectedErr := errors.New("not found")

	validator.On("Validate", contribution).Return(nil)
	contribRepo.On("FindById", contribution.ID, mock.Anything).Return(&domains.FoodContribution{}, expectedErr)

	response := service.Update(contribution, userId)

	assert.Equal(t, api.ErrorBadRequest(expectedErr.Error()), response)
}

func Test_FoodContributionService_Update_FailOnNoAccess(t *testing.T) {
	service, validator, contribRepo, _, _, _ := setupDefaultService()

	userId := uint(1)
	contribution := domains.FoodContribution{Model: gorm.Model{ID: 1}}
	oldContribution := &domains.FoodContribution{
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

func Test_FoodContributionService_Update_FailOnNotOwner(t *testing.T) {
	service, validator, contribRepo, _, _, _ := setupDefaultService()

	userId := uint(1)
	contribution := domains.FoodContribution{Model: gorm.Model{ID: 1}}
	oldContribution := &domains.FoodContribution{
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

func Test_FoodContributionService_Update_FailOnReqChange(t *testing.T) {
	service, validator, contribRepo, _, _, _ := setupDefaultService()

	userId := uint(1)
	contribution := domains.FoodContribution{
		Model:     gorm.Model{ID: 1},
		FoodReqId: 2, // Different from original
	}
	oldContribution := &domains.FoodContribution{
		Model:         gorm.Model{ID: 1},
		FoodReqId:     1,
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

	assert.Equal(t, api.ErrorBadRequest("cannot change food requirement of the contribution"), response)
}

func Test_FoodContributionService_Update_FailOnUpdate(t *testing.T) {
	service, validator, contribRepo, _, _, _ := setupDefaultService()

	userId := uint(1)
	contribution := domains.FoodContribution{
		Model:     gorm.Model{ID: 1},
		FoodReqId: 1,
	}
	oldContribution := &domains.FoodContribution{
		Model:         gorm.Model{ID: 1},
		FoodReqId:     1,
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

func Test_FoodContributionService_Delete_Success(t *testing.T) {
	service, _, contribRepo, _, _, _ := setupDefaultService()

	userId := uint(1)
	contributionId := uint(1)
	contribution := &domains.FoodContribution{
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

func Test_FoodContributionService_Delete_SuccessAsOrganizer(t *testing.T) {
	service, _, contribRepo, _, _, _ := setupDefaultService()

	userId := uint(1)
	contributionId := uint(1)
	contribution := &domains.FoodContribution{
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

func Test_FoodContributionService_Delete_FailOnFind(t *testing.T) {
	service, _, contribRepo, _, _, _ := setupDefaultService()

	userId := uint(1)
	contributionId := uint(1)
	expectedErr := errors.New("not found")

	contribRepo.On("FindById", contributionId, mock.Anything).Return(&domains.FoodContribution{}, expectedErr)

	response := service.Delete(contributionId, userId)

	assert.Equal(t, api.ErrorBadRequest(expectedErr.Error()), response)
}

func Test_FoodContributionService_Delete_FailOnUnauthorized(t *testing.T) {
	service, _, contribRepo, _, _, _ := setupDefaultService()

	userId := uint(1)
	contributionId := uint(1)
	contribution := &domains.FoodContribution{
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

func Test_FoodContributionService_Delete_FailOnDelete(t *testing.T) {
	service, _, contribRepo, _, _, _ := setupDefaultService()

	userId := uint(1)
	contributionId := uint(1)
	contribution := &domains.FoodContribution{
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

func Test_FoodContributionService_GetByPartyIdAndContributorId_Success(t *testing.T) {
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
	contributions := []domains.FoodContribution{
		{Model: gorm.Model{ID: 1}},
	}

	partyRepo.On("FindById", mock.Anything, mock.Anything).Return(party, nil)
	contribRepo.On("FindAllBy", []string{"party_id", "contributor_id"}, []interface{}{partyId, contributorId}, []string{"Contributor", "FoodReq"}).Return(&contributions, nil)

	response := service.GetByPartyIdAndContributorId(partyId, contributorId, userId)

	assert.False(t, response.GetIsError())
}

func Test_FoodContributionService_GetByPartyIdAndContributorId_FailOnPartyFind(t *testing.T) {
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

func Test_FoodContributionService_GetByPartyIdAndContributorId_FailOnPartyAuth(t *testing.T) {
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

func Test_FoodContributionService_GetByPartyIdAndContributorId_FailOnContribFind(t *testing.T) {
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
	contributions := []domains.FoodContribution{
		{Model: gorm.Model{ID: 1}},
	}
	expectedErr := errors.New("")

	partyRepo.On("FindById", mock.Anything, mock.Anything).Return(party, nil)
	contribRepo.On("FindAllBy", []string{"party_id", "contributor_id"}, []interface{}{partyId, contributorId}, []string{"Contributor", "FoodReq"}).Return(&contributions, expectedErr)

	response := service.GetByPartyIdAndContributorId(partyId, contributorId, userId)

	assert.Equal(t, response, api.ErrorInternalServerError(expectedErr.Error()))
}

func Test_FoodContributionService_GetByRequirementId_Success(t *testing.T) {
	service, _, contribRepo, _, _, foodReqRepo := setupDefaultService()

	userId := uint(1)
	requirementId := uint(1)
	req := &foodReqDomain.FoodRequirement{
		Party: partyDomains.Party{
			OrganizerID: 2,
			Participants: []userDomain.User{
				{Model: gorm.Model{ID: userId}},
			},
		},
	}
	contributions := []domains.FoodContribution{
		{Model: gorm.Model{ID: 1}},
	}

	foodReqRepo.On("FindById", requirementId, mock.Anything).Return(req, nil)
	contribRepo.On("FindAllBy", []string{"food_req_id"}, []interface{}{requirementId}, []string{"Contributor", "FoodReq"}).Return(&contributions, nil)

	response := service.GetByRequirementId(requirementId, userId)

	assert.False(t, response.GetIsError())
}

func Test_FoodContributionService_GetByRequirementId_FailOnPartyFind(t *testing.T) {
	service, _, _, _, _, foodReqRepo := setupDefaultService()

	userId := uint(1)
	requirementId := uint(1)
	req := &foodReqDomain.FoodRequirement{
		Party: partyDomains.Party{
			OrganizerID: 2,
			Participants: []userDomain.User{
				{Model: gorm.Model{ID: userId}},
			},
		},
	}
	expectedErr := errors.New("")

	foodReqRepo.On("FindById", requirementId, mock.Anything).Return(req, expectedErr)

	response := service.GetByRequirementId(requirementId, userId)

	assert.Equal(t, response, api.ErrorBadRequest(expectedErr.Error()))
}

func Test_FoodContributionService_GetByRequirementId_FailOnPartyAuth(t *testing.T) {
	service, _, _, _, _, foodReqRepo := setupDefaultService()

	userId := uint(1)
	requirementId := uint(1)
	req := &foodReqDomain.FoodRequirement{
		Party: partyDomains.Party{
			OrganizerID: 2,
		},
	}

	foodReqRepo.On("FindById", requirementId, mock.Anything).Return(req, nil)

	response := service.GetByRequirementId(requirementId, userId)

	assert.Equal(t, response, api.ErrorUnauthorized(domains.NO_ACCESS_TO_PARTY))
}

func Test_FoodContributionService_GetByRequirementId_FailOnContribFind(t *testing.T) {
	service, _, contribRepo, _, _, foodReqRepo := setupDefaultService()

	userId := uint(1)
	requirementId := uint(1)
	req := &foodReqDomain.FoodRequirement{
		Party: partyDomains.Party{
			OrganizerID: 2,
			Participants: []userDomain.User{
				{Model: gorm.Model{ID: userId}},
			},
		},
	}
	contributions := []domains.FoodContribution{
		{Model: gorm.Model{ID: 1}},
	}
	expectedErr := errors.New("")

	foodReqRepo.On("FindById", requirementId, mock.Anything).Return(req, nil)
	contribRepo.On("FindAllBy", []string{"food_req_id"}, []interface{}{requirementId}, []string{"Contributor", "FoodReq"}).Return(&contributions, expectedErr)

	response := service.GetByRequirementId(requirementId, userId)

	assert.Equal(t, response, api.ErrorInternalServerError(expectedErr.Error()))
}

func Test_FoodContributionService_GetByPartyId_Success(t *testing.T) {
	service, _, contribRepo, partyRepo, _, _ := setupDefaultService()

	userId := uint(1)
	partyId := uint(1)
	party := &partyDomains.Party{
		OrganizerID: 2,
		Participants: []userDomain.User{
			{Model: gorm.Model{ID: userId}},
		},
	}
	contributions := []domains.FoodContribution{
		{Model: gorm.Model{ID: 1}},
	}

	partyRepo.On("FindById", mock.Anything, mock.Anything).Return(party, nil)
	contribRepo.On("FindAllBy", []string{"party_id"}, []interface{}{partyId}, []string{"Contributor", "FoodReq"}).Return(&contributions, nil)

	response := service.GetByPartyId(partyId, userId)

	assert.False(t, response.GetIsError())
}

func Test_FoodContributionService_GetByPartyId_FailOnPartyFind(t *testing.T) {
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

func Test_FoodContributionService_GetByPartyId_FailOnPartyAuth(t *testing.T) {
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

func Test_FoodContributionService_GetByPartyId_FailOnContribFind(t *testing.T) {
	service, _, contribRepo, partyRepo, _, _ := setupDefaultService()

	userId := uint(1)
	partyId := uint(1)
	party := &partyDomains.Party{
		OrganizerID: 2,
		Participants: []userDomain.User{
			{Model: gorm.Model{ID: userId}},
		},
	}
	contributions := []domains.FoodContribution{
		{Model: gorm.Model{ID: 1}},
	}
	expectedErr := errors.New("")

	partyRepo.On("FindById", mock.Anything, mock.Anything).Return(party, nil)
	contribRepo.On("FindAllBy", []string{"party_id"}, []interface{}{partyId}, []string{"Contributor", "FoodReq"}).Return(&contributions, expectedErr)

	response := service.GetByPartyId(partyId, userId)

	assert.Equal(t, response, api.ErrorInternalServerError(expectedErr.Error()))
}
