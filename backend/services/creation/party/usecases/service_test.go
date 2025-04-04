package usecases

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	drinkRequirementUsecases "github.com/zsomborCzaban/party_organizer/services/creation/drink_requirement/usecases"
	foodRequirementUsecases "github.com/zsomborCzaban/party_organizer/services/creation/food_requirement/usecases"
	"github.com/zsomborCzaban/party_organizer/services/creation/party/domains"
	drinkContributionUsecases "github.com/zsomborCzaban/party_organizer/services/interaction/drink_contributions/usecases"
	foodContributionUsecases "github.com/zsomborCzaban/party_organizer/services/interaction/food_contributions/usecases"
	partyInvitationUsecases "github.com/zsomborCzaban/party_organizer/services/managers/party_attendance_manager/usecases"
	userDomain "github.com/zsomborCzaban/party_organizer/services/user/domains"
	userUsecases "github.com/zsomborCzaban/party_organizer/services/user/usecases"
	"github.com/zsomborCzaban/party_organizer/utils/api"
	"github.com/zsomborCzaban/party_organizer/utils/repo"
	"gorm.io/gorm"
	"testing"
)

func setupDefaultPartyService() (domains.IPartyService, *api.MockValidator, *MockRepository, *userUsecases.MockRepository, *drinkRequirementUsecases.MockRepository, *drinkContributionUsecases.MockRepository, *foodRequirementUsecases.MockRepository, *foodContributionUsecases.MockRepository, *partyInvitationUsecases.MockRepository) {
	validator := new(api.MockValidator)
	partyRepo := new(MockRepository)
	userRepo := new(userUsecases.MockRepository)
	drinkReqRepo := new(drinkRequirementUsecases.MockRepository)
	drinkContribRepo := new(drinkContributionUsecases.MockRepository)
	foodReqRepo := new(foodRequirementUsecases.MockRepository)
	foodContribRepo := new(foodContributionUsecases.MockRepository)
	partyInviteRepo := new(partyInvitationUsecases.MockRepository)

	repoCollector := &repo.RepoCollector{
		PartyRepo:        partyRepo,
		UserRepo:         userRepo,
		DrinkReqRepo:     drinkReqRepo,
		DrinkContribRepo: drinkContribRepo,
		FoodReqRepo:      foodReqRepo,
		FoodContribRepo:  foodContribRepo,
		PartyInviteRepo:  partyInviteRepo,
	}

	service := NewPartyService(repoCollector, validator)

	return service, validator, partyRepo, userRepo, drinkReqRepo, drinkContribRepo, foodReqRepo, foodContribRepo, partyInviteRepo
}

func Test_PartyService_Create_Success(t *testing.T) {
	service, validator, partyRepo, userRepo, _, _, _, _, _ := setupDefaultPartyService()

	userId := uint(1)
	partyDTO := domains.PartyDTO{
		Name:              "Test Party",
		AccessCodeEnabled: true,
	}
	party := partyDTO.TransformToParty()
	party.OrganizerID = userId
	organizer := &userDomain.User{Model: gorm.Model{
		ID: userId,
	},
	}

	validator.On("Validate", partyDTO).Return(nil)
	userRepo.On("FindById", userId, mock.Anything).Return(organizer, nil)
	partyRepo.On("Create", mock.Anything).Return(nil)

	response := service.Create(partyDTO, userId)

	assert.False(t, response.GetIsError())
}

func Test_PartyService_Create_FailOnValidation(t *testing.T) {
	service, validator, _, _, _, _, _, _, _ := setupDefaultPartyService()

	userId := uint(1)
	partyDTO := domains.PartyDTO{
		Name:              "Test Party",
		AccessCodeEnabled: true,
	}
	validationError := api.NewValidationErrors()

	validator.On("Validate", partyDTO).Return(validationError)

	response := service.Create(partyDTO, userId)

	assert.Equal(t, api.ErrorValidation(validationError), response)
}

func Test_PartyService_Create_FailOnUserNotFound(t *testing.T) {
	service, validator, _, userRepo, _, _, _, _, _ := setupDefaultPartyService()

	userId := uint(1)
	partyDTO := domains.PartyDTO{
		Name:              "Test Party",
		AccessCodeEnabled: true,
	}

	validator.On("Validate", partyDTO).Return(nil)
	userRepo.On("FindById", userId, mock.Anything).Return(&userDomain.User{}, errors.New("user not found"))

	response := service.Create(partyDTO, userId)

	assert.Equal(t, api.ErrorInternalServerError(domains.DeletedUser), response)
}

func Test_PartyService_Create_FailOnCreateParty(t *testing.T) {
	service, validator, partyRepo, userRepo, _, _, _, _, _ := setupDefaultPartyService()

	userId := uint(1)
	partyDTO := domains.PartyDTO{
		Name:              "Test Party",
		AccessCodeEnabled: true,
	}
	party := partyDTO.TransformToParty()
	party.OrganizerID = userId
	organizer := &userDomain.User{Model: gorm.Model{
		ID: userId,
	},
	}

	validator.On("Validate", partyDTO).Return(nil)
	userRepo.On("FindById", userId, mock.Anything).Return(organizer, nil)
	partyRepo.On("Create", mock.Anything).Return(errors.New("failed to create party"))

	response := service.Create(partyDTO, userId)

	assert.Equal(t, api.ErrorInternalServerError("failed to create party"), response)
}

func Test_PartyService_Get_Success(t *testing.T) {
	service, _, partyRepo, _, _, _, _, _, _ := setupDefaultPartyService()

	userId := uint(1)
	partyId := uint(1)
	party := &domains.Party{
		Model: gorm.Model{
			ID: partyId,
		},
		OrganizerID: 2,
		Participants: []userDomain.User{
			{Model: gorm.Model{ID: userId}},
		},
	}

	partyRepo.On("FindById", mock.Anything, mock.Anything).Return(party, nil)

	response := service.Get(partyId, userId)

	assert.False(t, response.GetIsError())
}

func Test_PartyService_Get_FailOnPartyNotFound(t *testing.T) {
	service, _, partyRepo, _, _, _, _, _, _ := setupDefaultPartyService()

	userId := uint(1)
	partyId := uint(1)

	partyRepo.On("FindById", mock.Anything, mock.Anything).Return(&domains.Party{}, errors.New("party not found"))

	response := service.Get(partyId, userId)

	assert.Equal(t, api.ErrorInternalServerError("party not found"), response)
}

func Test_PartyService_Get_FailOnUnauthorized(t *testing.T) {
	service, _, partyRepo, _, _, _, _, _, _ := setupDefaultPartyService()

	userId := uint(1)
	partyId := uint(1)
	party := &domains.Party{
		Model: gorm.Model{
			ID: partyId,
		},
		OrganizerID: 2, // Different organizer
	}

	partyRepo.On("FindById", mock.Anything, mock.Anything).Return(party, nil)

	response := service.Get(partyId, userId)

	assert.Equal(t, api.ErrorUnauthorized("you cannot access this party"), response)
}

func Test_PartyService_GetPublicParty_Success(t *testing.T) {
	service, _, partyRepo, _, _, _, _, _, _ := setupDefaultPartyService()

	partyId := uint(1)
	party := &domains.Party{
		Model: gorm.Model{
			ID: partyId,
		},
		Private: false,
	}

	partyRepo.On("FindById", mock.Anything, mock.Anything).Return(party, nil)

	response := service.GetPublicParty(partyId)

	assert.False(t, response.GetIsError())
}

func Test_PartyService_GetPublicParty_FailOnFindParty(t *testing.T) {
	service, _, partyRepo, _, _, _, _, _, _ := setupDefaultPartyService()

	partyId := uint(1)
	party := &domains.Party{
		Model: gorm.Model{
			ID: partyId,
		},
		Private: true,
	}
	expectedErr := errors.New("")

	partyRepo.On("FindById", mock.Anything, mock.Anything).Return(party, expectedErr)

	response := service.GetPublicParty(partyId)

	assert.Equal(t, api.ErrorBadRequest(expectedErr.Error()), response)
}

func Test_PartyService_GetPublicParty_FailOnPrivateParty(t *testing.T) {
	service, _, partyRepo, _, _, _, _, _, _ := setupDefaultPartyService()

	partyId := uint(1)
	party := &domains.Party{
		Model: gorm.Model{
			ID: partyId,
		},
		Private: true,
	}

	partyRepo.On("FindById", mock.Anything, mock.Anything).Return(party, nil)

	response := service.GetPublicParty(partyId)

	assert.Equal(t, api.ErrorUnauthorized("this party is private"), response)
}

func Test_PartyService_Update_Success(t *testing.T) {
	service, validator, partyRepo, _, _, _, _, _, _ := setupDefaultPartyService()

	userId := uint(1)
	partyDTO := domains.PartyDTO{
		ID:                1,
		Name:              "Updated Party",
		AccessCodeEnabled: true,
	}
	originalParty := &domains.Party{
		Model: gorm.Model{
			ID: partyDTO.ID,
		},
		OrganizerID: userId,
	}
	updatedParty := partyDTO.TransformToParty()
	updatedParty.OrganizerID = userId

	validator.On("Validate", partyDTO).Return(nil)
	partyRepo.On("FindById", mock.Anything, mock.Anything).Return(originalParty, nil)
	partyRepo.On("Update", mock.Anything).Return(nil)

	response := service.Update(partyDTO, userId)

	assert.False(t, response.GetIsError())
}

func Test_PartyService_Update_FailOnValidation(t *testing.T) {
	service, validator, _, _, _, _, _, _, _ := setupDefaultPartyService()

	userId := uint(1)
	partyDTO := domains.PartyDTO{
		ID:                1,
		Name:              "Updated Party",
		AccessCodeEnabled: true,
	}
	validationError := api.NewValidationErrors()

	validator.On("Validate", partyDTO).Return(validationError)

	response := service.Update(partyDTO, userId)

	assert.Equal(t, api.ErrorValidation(validationError), response)
}

func Test_PartyService_Update_FailOnPartyNotFound(t *testing.T) {
	service, validator, partyRepo, _, _, _, _, _, _ := setupDefaultPartyService()

	userId := uint(1)
	partyDTO := domains.PartyDTO{
		ID:                1,
		Name:              "Updated Party",
		AccessCodeEnabled: true,
	}

	validator.On("Validate", partyDTO).Return(nil)
	partyRepo.On("FindById", mock.Anything, mock.Anything).Return(&domains.Party{}, errors.New("party not found"))

	response := service.Update(partyDTO, userId)

	assert.Equal(t, api.ErrorBadRequest("party not found"), response)
}

func Test_PartyService_Update_FailOnUnauthorized(t *testing.T) {
	service, validator, partyRepo, _, _, _, _, _, _ := setupDefaultPartyService()

	userId := uint(1)
	partyDTO := domains.PartyDTO{
		ID:                1,
		Name:              "Updated Party",
		AccessCodeEnabled: true,
	}
	originalParty := &domains.Party{
		Model: gorm.Model{
			ID: partyDTO.ID,
		},
		OrganizerID: 2, // Different organizer
	}

	validator.On("Validate", partyDTO).Return(nil)
	partyRepo.On("FindById", mock.Anything, mock.Anything).Return(originalParty, nil)

	response := service.Update(partyDTO, userId)

	assert.Equal(t, api.ErrorUnauthorized("cannot update other people's party"), response)
}
func Test_PartyService_Update_FailOnUpdate(t *testing.T) {
	service, validator, partyRepo, _, _, _, _, _, _ := setupDefaultPartyService()

	userId := uint(1)
	partyDTO := domains.PartyDTO{
		ID:                1,
		Name:              "Updated Party",
		AccessCodeEnabled: true,
	}
	originalParty := &domains.Party{
		Model: gorm.Model{
			ID: partyDTO.ID,
		},
		OrganizerID: userId,
	}
	updatedParty := partyDTO.TransformToParty()
	updatedParty.OrganizerID = userId
	expectedErr := errors.New("some error")

	validator.On("Validate", partyDTO).Return(nil)
	partyRepo.On("FindById", mock.Anything, mock.Anything).Return(originalParty, nil)
	partyRepo.On("Update", mock.Anything).Return(expectedErr)

	response := service.Update(partyDTO, userId)

	assert.Equal(t, api.ErrorInternalServerError(expectedErr.Error()), response)
}

func Test_PartyService_Delete_Success(t *testing.T) {
	service, _, partyRepo, _, drinkReqRepo, drinkContribRepo, foodReqRepo, foodContribRepo, partyInviteRepo := setupDefaultPartyService()

	userId := uint(1)
	partyId := uint(1)
	party := &domains.Party{
		Model: gorm.Model{
			ID: partyId,
		},
		OrganizerID: userId,
	}

	partyRepo.On("FindById", mock.Anything, mock.Anything).Return(party, nil)
	drinkContribRepo.On("DeleteByPartyId", mock.Anything).Return(nil)
	drinkReqRepo.On("DeleteByPartyId", mock.Anything).Return(nil)
	foodContribRepo.On("DeleteByPartyId", mock.Anything).Return(nil)
	foodReqRepo.On("DeleteByPartyId", mock.Anything).Return(nil)
	partyRepo.On("Delete", mock.Anything).Return(nil)
	partyInviteRepo.On("DeleteByPartyId", mock.Anything).Return(nil)

	response := service.Delete(partyId, userId)

	assert.False(t, response.GetIsError())
}

func Test_PartyService_Delete_FailOnPartyNotFound(t *testing.T) {
	service, _, partyRepo, _, _, _, _, _, _ := setupDefaultPartyService()

	userId := uint(1)
	partyId := uint(1)

	partyRepo.On("FindById", mock.Anything, mock.Anything).Return(&domains.Party{}, errors.New("party not found"))

	response := service.Delete(partyId, userId)

	assert.Equal(t, api.ErrorBadRequest("party not found"), response)
}

func Test_PartyService_Delete_FailOnUnauthorized(t *testing.T) {
	service, _, partyRepo, _, _, _, _, _, _ := setupDefaultPartyService()

	userId := uint(1)
	partyId := uint(1)
	party := &domains.Party{
		Model: gorm.Model{
			ID: partyId,
		},
		OrganizerID: 2, // Different organizer
	}

	partyRepo.On("FindById", mock.Anything, mock.Anything).Return(party, nil)

	response := service.Delete(partyId, userId)

	assert.Equal(t, api.ErrorUnauthorized("cannot delete other peoples party"), response)
}

func Test_PartyService_Delete_FailOnDeleteContributions(t *testing.T) {
	service, _, partyRepo, _, _, drinkContribRepo, _, foodContribRepo, _ := setupDefaultPartyService()

	userId := uint(1)
	partyId := uint(1)
	party := &domains.Party{
		Model: gorm.Model{
			ID: partyId,
		},
		OrganizerID: userId,
	}

	partyRepo.On("FindById", mock.Anything, mock.Anything).Return(party, nil)
	drinkContribRepo.On("DeleteByPartyId", mock.Anything).Return(errors.New("failed to delete drink contributions"))
	foodContribRepo.On("DeleteByPartyId", mock.Anything).Return(nil)

	response := service.Delete(partyId, userId)

	assert.Equal(t, api.ErrorInternalServerError("unexpected error while deleting the contributions of the party"), response)
}

func Test_PartyService_Delete_FailOnDeleteRequirements(t *testing.T) {
	service, _, partyRepo, _, drinkReqRepo, drinkContribRepo, foodReqRepo, foodContribRepo, partyInviteRepo := setupDefaultPartyService()

	userId := uint(1)
	partyId := uint(1)
	party := &domains.Party{
		Model: gorm.Model{
			ID: partyId,
		},
		OrganizerID: userId,
	}

	partyRepo.On("FindById", mock.Anything, mock.Anything).Return(party, nil)
	drinkContribRepo.On("DeleteByPartyId", mock.Anything).Return(nil)
	drinkReqRepo.On("DeleteByPartyId", mock.Anything).Return(nil)
	foodContribRepo.On("DeleteByPartyId", mock.Anything).Return(nil)
	foodReqRepo.On("DeleteByPartyId", mock.Anything).Return(errors.New(""))
	partyRepo.On("Delete", mock.Anything).Return(nil)
	partyInviteRepo.On("DeleteByPartyId", mock.Anything).Return(nil)

	response := service.Delete(partyId, userId)

	assert.Equal(t, api.ErrorInternalServerError("unexpected error while deleting the requirements of the party"), response)
}

func Test_PartyService_Delete_FailOnDeletePartyInvites(t *testing.T) {
	service, _, partyRepo, _, drinkReqRepo, drinkContribRepo, foodReqRepo, foodContribRepo, partyInviteRepo := setupDefaultPartyService()

	userId := uint(1)
	partyId := uint(1)
	party := &domains.Party{
		Model: gorm.Model{
			ID: partyId,
		},
		OrganizerID: userId,
	}

	partyRepo.On("FindById", mock.Anything, mock.Anything).Return(party, nil)
	drinkContribRepo.On("DeleteByPartyId", mock.Anything).Return(nil)
	drinkReqRepo.On("DeleteByPartyId", mock.Anything).Return(nil)
	foodContribRepo.On("DeleteByPartyId", mock.Anything).Return(nil)
	foodReqRepo.On("DeleteByPartyId", mock.Anything).Return(nil)
	partyRepo.On("Delete", mock.Anything).Return(nil)
	partyInviteRepo.On("DeleteByPartyId", mock.Anything).Return(errors.New(""))

	response := service.Delete(partyId, userId)

	assert.Equal(t, api.ErrorInternalServerError("unexpected error while deleting party invites of the party"), response)
}

func Test_PartyService_Delete_FailOnDeleteParty(t *testing.T) {
	service, _, partyRepo, _, drinkReqRepo, drinkContribRepo, foodReqRepo, foodContribRepo, partyInviteRepo := setupDefaultPartyService()

	userId := uint(1)
	partyId := uint(1)
	party := &domains.Party{
		Model: gorm.Model{
			ID: partyId,
		},
		OrganizerID: userId,
	}

	partyRepo.On("FindById", mock.Anything, mock.Anything).Return(party, nil)
	drinkContribRepo.On("DeleteByPartyId", mock.Anything).Return(nil)
	drinkReqRepo.On("DeleteByPartyId", mock.Anything).Return(nil)
	foodContribRepo.On("DeleteByPartyId", mock.Anything).Return(nil)
	foodReqRepo.On("DeleteByPartyId", mock.Anything).Return(nil)
	partyRepo.On("Delete", mock.Anything).Return(errors.New("failed to delete party"))
	partyInviteRepo.On("DeleteByPartyId", mock.Anything).Return(nil)

	response := service.Delete(partyId, userId)

	assert.Equal(t, api.ErrorInternalServerError("failed to delete party"), response)
}

func Test_PartyService_GetPublicParties_Success(t *testing.T) {
	service, _, partyRepo, _, _, _, _, _, _ := setupDefaultPartyService()

	parties := []domains.Party{
		{Model: gorm.Model{
			ID: 1,
		},
			Name: "Public Party 1", Private: false},
		{Model: gorm.Model{
			ID: 1,
		},
			Name: "Public Party 2", Private: false},
	}

	partyRepo.On("GetPublicParties").Return(&parties, nil)

	response := service.GetPublicParties()

	assert.False(t, response.GetIsError())
}

func Test_PartyService_GetPublicParties_Fail(t *testing.T) {
	service, _, partyRepo, _, _, _, _, _, _ := setupDefaultPartyService()

	partyRepo.On("GetPublicParties").Return(&[]domains.Party{}, errors.New("failed to get public parties"))

	response := service.GetPublicParties()

	assert.Equal(t, api.ErrorInternalServerError("failed to get public parties"), response)
}

func Test_PartyService_GetPartiesByOrganizerId_Success(t *testing.T) {
	service, _, partyRepo, _, _, _, _, _, _ := setupDefaultPartyService()

	organizerId := uint(1)
	parties := []domains.Party{
		{Model: gorm.Model{
			ID: 1,
		},
			Name: "Organizer Party 1", OrganizerID: organizerId},
		{Model: gorm.Model{
			ID: 2,
		},
			Name: "Organizer Party 2", OrganizerID: organizerId},
	}

	partyRepo.On("GetPartiesByOrganizerId", mock.Anything).Return(&parties, nil)

	response := service.GetPartiesByOrganizerId(organizerId)

	assert.False(t, response.GetIsError())
}

func Test_PartyService_GetPartiesByOrganizerId_Fail(t *testing.T) {
	service, _, partyRepo, _, _, _, _, _, _ := setupDefaultPartyService()

	organizerId := uint(1)
	partyRepo.On("GetPartiesByOrganizerId", mock.Anything).Return(&[]domains.Party{}, errors.New("failed to get parties by organizer"))

	response := service.GetPartiesByOrganizerId(organizerId)

	assert.Equal(t, api.ErrorInternalServerError("failed to get parties by organizer"), response)
}

func Test_PartyService_GetPartiesByParticipantId_Success(t *testing.T) {
	service, _, partyRepo, _, _, _, _, _, _ := setupDefaultPartyService()

	participantId := uint(1)
	parties := []domains.Party{
		{Model: gorm.Model{
			ID: 1,
		},
			Name: "Participant Party 1"},
		{Model: gorm.Model{
			ID: 2,
		},
			Name: "Participant Party 2"},
	}

	partyRepo.On("GetPartiesByParticipantId", mock.Anything).Return(&parties, nil)

	response := service.GetPartiesByParticipantId(participantId)

	assert.False(t, response.GetIsError())
}

func Test_PartyService_GetPartiesByParticipantId_Fail(t *testing.T) {
	service, _, partyRepo, _, _, _, _, _, _ := setupDefaultPartyService()

	participantId := uint(1)
	partyRepo.On("GetPartiesByParticipantId", mock.Anything).Return(&[]domains.Party{}, errors.New("failed to get parties by participant"))

	response := service.GetPartiesByParticipantId(participantId)

	assert.Equal(t, api.ErrorInternalServerError("failed to get parties by participant"), response)
}

func Test_PartyService_GetParticipants_Success(t *testing.T) {
	service, _, partyRepo, _, _, _, _, _, _ := setupDefaultPartyService()

	partyId := uint(1)
	userId := uint(1)
	party := &domains.Party{
		Model: gorm.Model{
			ID: partyId,
		},
		OrganizerID: userId,
		Participants: []userDomain.User{
			{Model: gorm.Model{
				ID: 2,
			},
				Username: "Participant 1",
			},
			{Model: gorm.Model{
				ID: 3,
			},
				Username: "Participant 2",
			},
		},
	}

	partyRepo.On("FindById", mock.Anything, mock.Anything).Return(party, nil)

	response := service.GetParticipants(partyId, userId)

	assert.False(t, response.GetIsError())
}

func Test_PartyService_GetParticipants_FailOnPartyNotFound(t *testing.T) {
	service, _, partyRepo, _, _, _, _, _, _ := setupDefaultPartyService()

	partyId := uint(1)
	userId := uint(1)

	partyRepo.On("FindById", mock.Anything, mock.Anything).Return(&domains.Party{}, errors.New("party not found"))

	response := service.GetParticipants(partyId, userId)

	assert.Equal(t, api.ErrorBadRequest("party not found"), response)
}

func Test_PartyService_GetParticipants_FailOnUnauthorized(t *testing.T) {
	service, _, partyRepo, _, _, _, _, _, _ := setupDefaultPartyService()

	partyId := uint(1)
	userId := uint(1)
	party := &domains.Party{
		Model: gorm.Model{
			ID: partyId,
		},
		OrganizerID: 2, // Different organizer
	}

	partyRepo.On("FindById", mock.Anything, mock.Anything).Return(party, nil)

	response := service.GetParticipants(partyId, userId)

	assert.Equal(t, api.ErrorUnauthorized("you cannot access this party"), response)
}
