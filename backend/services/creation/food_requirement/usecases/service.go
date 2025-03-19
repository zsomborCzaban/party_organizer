package usecases

import (
	"github.com/zsomborCzaban/party_organizer/services/creation/food_requirement/domains"
	partyDomains "github.com/zsomborCzaban/party_organizer/services/creation/party/domains"
	foodContributionDomains "github.com/zsomborCzaban/party_organizer/services/interaction/food_contributions/domains"
	"github.com/zsomborCzaban/party_organizer/utils/api"
	"github.com/zsomborCzaban/party_organizer/utils/repo"
)

type FoodRequirementService struct {
	Validator                  api.IValidator
	FoodRequirementRepository  domains.IFoodRequirementRepository
	PartyRepository            partyDomains.IPartyRepository
	FoodContributionRepository foodContributionDomains.IFoodContributionRepository
}

func NewFoodRequirementService(repoCollector *repo.RepoCollector, validator api.IValidator) domains.IFoodRequirementService {
	return &FoodRequirementService{
		Validator:                  validator,
		FoodRequirementRepository:  repoCollector.FoodReqRepo,
		PartyRepository:            repoCollector.PartyRepo,
		FoodContributionRepository: repoCollector.FoodContribRepo,
	}
}

func (fs FoodRequirementService) Create(foodRequirementDTO domains.FoodRequirementDTO, userId uint) api.IResponse {
	err := fs.Validator.Validate(foodRequirementDTO)
	if err != nil {
		return api.ErrorValidation(err)
	}

	foodRequirement := foodRequirementDTO.TransformToFoodRequirement()

	party, err2 := fs.PartyRepository.FindById(foodRequirement.PartyID, partyDomains.FullPartyPreload...)
	if err2 != nil {
		return api.ErrorBadRequest(domains.PartyNotFound)
	}

	if !party.CanBeOrganizedBy(userId) {
		return api.ErrorUnauthorized(domains.NoOrganizerAccess)
	}

	err3 := fs.FoodRequirementRepository.Create(foodRequirement)
	if err3 != nil {
		return api.ErrorInternalServerError(err3.Error())
	}

	return api.Success(foodRequirement)
}

func (fs FoodRequirementService) FindById(foodReqId, userId uint) api.IResponse {
	foodRequirement, err := fs.FoodRequirementRepository.FindById(foodReqId, partyDomains.FullPartyNestedPreload...)
	if err != nil {
		return api.ErrorBadRequest(domains.RequirementNotFound)
	}

	if !foodRequirement.Party.CanBeAccessedBy(userId) {
		return api.ErrorUnauthorized(domains.NoViewAccess)
	}

	return api.Success(foodRequirement)
}

func (fs FoodRequirementService) Delete(foodReqId, userId uint) api.IResponse {
	foodRequirement, err := fs.FoodRequirementRepository.FindById(foodReqId, "Party", "Party.Organizer")
	if err != nil {
		return api.ErrorBadRequest(domains.RequirementNotFound)
	}

	if !foodRequirement.Party.CanBeOrganizedBy(userId) {
		return api.ErrorUnauthorized(domains.NoOrganizerAccess)
	}

	//todo: put this in transaction
	if err2 := fs.FoodContributionRepository.DeleteByReqId(foodReqId); err2 != nil {
		return api.ErrorInternalServerError(err2.Error())
	}

	err3 := fs.FoodRequirementRepository.Delete(foodRequirement)
	if err3 != nil {
		return api.ErrorInternalServerError(err3.Error())
	}
	return api.Success("delete_success")
}

func (fs FoodRequirementService) GetByPartyId(partyId, userId uint) api.IResponse {
	party, err := fs.PartyRepository.FindById(partyId, partyDomains.FullPartyPreload...)
	if err != nil {
		return api.ErrorBadRequest(domains.PartyNotFound)
	}

	if !party.CanBeAccessedBy(userId) {
		return api.ErrorUnauthorized(domains.NoViewAccess)
	}

	foodReqs, err3 := fs.FoodRequirementRepository.GetByPartyId(partyId)
	if err3 != nil {
		return api.ErrorInternalServerError(err3.Error())
	}

	return api.Success(foodReqs)
}
