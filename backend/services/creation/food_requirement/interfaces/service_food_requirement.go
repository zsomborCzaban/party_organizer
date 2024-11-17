package interfaces

import (
	"github.com/zsomborCzaban/party_organizer/services/creation/food_requirement/domains"
	partyDomains "github.com/zsomborCzaban/party_organizer/services/creation/party/domains"
	foodContributionDomains "github.com/zsomborCzaban/party_organizer/services/interaction/food_contributions/domains"
	"github.com/zsomborCzaban/party_organizer/utils/api"
	"github.com/zsomborCzaban/party_organizer/utils/repo"
)

type FoodRequirementService struct {
	Validator                  api.IValidator
	FoodRequirementRepository  *domains.IFoodRequirementRepository
	PartyRepository            *partyDomains.IPartyRepository
	FoodContributionRepository *foodContributionDomains.IFoodContributionRepository
}

func NewFoodRequirementService(repoCollector *repo.RepoCollector, validator api.IValidator) domains.IFoodRequirementService {
	return &FoodRequirementService{
		Validator:                  validator,
		FoodRequirementRepository:  repoCollector.FoodReqReqRepo,
		PartyRepository:            repoCollector.PartyRepo,
		FoodContributionRepository: repoCollector.FoodContribRepo,
	}
}

func (fs FoodRequirementService) CreateFoodRequirement(foodRequirementDTO domains.FoodRequirementDTO, userId uint) api.IResponse {
	err := fs.Validator.Validate(foodRequirementDTO)
	if err != nil {
		return api.ErrorValidation(err)
	}

	foodRequirement := foodRequirementDTO.TransformToFoodRequirement()

	party, err2 := (*fs.PartyRepository).FindById(foodRequirement.PartyID, partyDomains.FullPartyPreload...)
	if err2 != nil {
		return api.ErrorBadRequest("partyId does not exists")
	}

	if !party.CanBeOrganizedBy(userId) {
		return api.ErrorUnauthorized("cannot create foodRequirement for other people's party")
	}

	err3 := (*fs.FoodRequirementRepository).CreateFoodRequirement(foodRequirement)
	if err3 != nil {
		return api.ErrorInternalServerError(err)
	}

	return api.Success(foodRequirement.TransformToFoodRequirementDTO())
}

func (fs FoodRequirementService) GetFoodRequirement(foodReqId, userId uint) api.IResponse {
	foodRequirement, err := (*fs.FoodRequirementRepository).FindById(foodReqId, partyDomains.FullPartyNestedPreload...)
	if err != nil {
		return api.ErrorInternalServerError(err)
	}

	if foodRequirement.Party.CanBeAccessedBy(userId) {
		return api.ErrorUnauthorized("you are not in the party")
	}

	return api.Success(foodRequirement)
}

func (fs FoodRequirementService) DeleteFoodRequirement(foodReqId, userId uint) api.IResponse {
	foodRequirement, err := (*fs.FoodRequirementRepository).FindById(foodReqId, partyDomains.FullPartyNestedPreload...)
	if err != nil {
		return api.ErrorBadRequest(err.Error())
	}

	if !foodRequirement.Party.CanBeOrganizedBy(userId) {
		return api.ErrorUnauthorized(domains.UNAUTHORIZED)
	}

	//todo: put this in transaction
	if err2 := (*fs.FoodContributionRepository).DeleteByReqId(foodReqId); err2 != nil {
		return api.ErrorInternalServerError(err2.Error())
	}

	err3 := (*fs.FoodRequirementRepository).DeleteFoodRequirement(foodRequirement)
	if err3 != nil {
		return api.ErrorInternalServerError(err3)
	}
	return api.Success("delete_success")
}

func (fs FoodRequirementService) GetByPartyId(partyId, userId uint) api.IResponse {
	party, err := (*fs.PartyRepository).FindById(partyId, partyDomains.FullPartyPreload...)
	if err != nil {
		return api.ErrorBadRequest("party not found")
	}

	if !party.CanBeAccessedBy(userId) {
		return api.ErrorUnauthorized("you are not in the party")
	}

	foodReqs, err3 := (*fs.FoodRequirementRepository).GetByPartyId(partyId)
	if err3 != nil {
		return api.ErrorInternalServerError(err3)
	}

	return api.Success(foodReqs)
}
