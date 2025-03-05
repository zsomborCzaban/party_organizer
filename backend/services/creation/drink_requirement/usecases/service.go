package usecases

import (
	"github.com/zsomborCzaban/party_organizer/services/creation/drink_requirement/domains"
	partyDomains "github.com/zsomborCzaban/party_organizer/services/creation/party/domains"
	drinkContributionDomains "github.com/zsomborCzaban/party_organizer/services/interaction/drink_contributions/domains"
	"github.com/zsomborCzaban/party_organizer/utils/api"
	"github.com/zsomborCzaban/party_organizer/utils/repo"
)

type DrinkRequirementService struct {
	Validator                   api.IValidator
	DrinkRequirementRepository  domains.IDrinkRequirementRepository
	PartyRepository             partyDomains.IPartyRepository
	DrinkContributionRepository drinkContributionDomains.IDrinkContributionRepository
}

func NewDrinkRequirementService(repoCollector *repo.RepoCollector, validator api.IValidator) domains.IDrinkRequirementService {
	return &DrinkRequirementService{
		Validator:                   validator,
		DrinkRequirementRepository:  *repoCollector.DrinkReqRepo,
		PartyRepository:             *repoCollector.PartyRepo,
		DrinkContributionRepository: *repoCollector.DrinkContribRepo,
	}
}

func (ds DrinkRequirementService) CreateDrinkRequirement(drinkRequirementDTO domains.DrinkRequirementDTO, userId uint) api.IResponse {
	err := ds.Validator.Validate(drinkRequirementDTO)
	if err != nil {
		return api.ErrorValidation(err)
	}

	drinkRequirement := drinkRequirementDTO.TransformToDrinkRequirement()

	party, err2 := ds.PartyRepository.FindById(drinkRequirement.PartyID, partyDomains.FullPartyPreload...)
	if err2 != nil {
		return api.ErrorBadRequest("Party id doesnt exists")
	}

	if !party.CanBeOrganizedBy(userId) {
		return api.ErrorUnauthorized("cannot create drinkRequirements for other peoples party")
	}

	err3 := ds.DrinkRequirementRepository.CreateDrinkRequirement(drinkRequirement)
	if err3 != nil {
		return api.ErrorInternalServerError(err3)
	}

	return api.Success(drinkRequirement)
}

func (ds DrinkRequirementService) GetDrinkRequirement(drinkReqId, userId uint) api.IResponse {
	drinkRequirement, err := ds.DrinkRequirementRepository.FindById(drinkReqId, partyDomains.FullPartyNestedPreload...)
	if err != nil {
		return api.ErrorInternalServerError(err)
	}

	if !drinkRequirement.Party.CanBeAccessedBy(userId) {
		return api.ErrorUnauthorized("you are not in the party")
	}

	return api.Success(drinkRequirement)
}

func (ds DrinkRequirementService) DeleteDrinkRequirement(drinkReqId, userId uint) api.IResponse {
	drinkRequirement, err := ds.DrinkRequirementRepository.FindById(drinkReqId, partyDomains.FullPartyNestedPreload...)
	if err != nil {
		return api.ErrorBadRequest(err.Error())
	}

	if !drinkRequirement.Party.CanBeOrganizedBy(userId) {
		return api.ErrorUnauthorized(domains.UNAUTHORIZED)
	}

	//todo: put this in transaction
	if err2 := ds.DrinkContributionRepository.DeleteByReqId(drinkReqId); err2 != nil {
		return api.ErrorInternalServerError(err2.Error())
	}

	err3 := ds.DrinkRequirementRepository.DeleteDrinkRequirement(drinkRequirement)
	if err3 != nil {
		return api.ErrorInternalServerError(err3)
	}
	return api.Success("delete_success")
}

func (ds DrinkRequirementService) GetByPartyId(partyId, userId uint) api.IResponse {
	party, err := ds.PartyRepository.FindById(partyId, partyDomains.FullPartyPreload...)
	if err != nil {
		return api.ErrorBadRequest("party not found")
	}

	if !party.CanBeAccessedBy(userId) {
		return api.ErrorUnauthorized("you are not in the party")
	}

	drinkReqs, err3 := ds.DrinkRequirementRepository.GetByPartyId(partyId)
	if err3 != nil {
		return api.ErrorInternalServerError(err3)
	}

	return api.Success(drinkReqs)
}
