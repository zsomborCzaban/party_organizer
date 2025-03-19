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
		DrinkRequirementRepository:  repoCollector.DrinkReqRepo,
		PartyRepository:             repoCollector.PartyRepo,
		DrinkContributionRepository: repoCollector.DrinkContribRepo,
	}
}

func (ds DrinkRequirementService) Create(drinkRequirementDTO domains.DrinkRequirementDTO, userId uint) api.IResponse {
	err := ds.Validator.Validate(drinkRequirementDTO)
	if err != nil {
		return api.ErrorValidation(err)
	}

	drinkRequirement := drinkRequirementDTO.TransformToDrinkRequirement()

	party, err2 := ds.PartyRepository.FindById(drinkRequirement.PartyID, partyDomains.FullPartyPreload...)
	if err2 != nil {
		return api.ErrorBadRequest(domains.PartyNotFound)
	}

	if !party.CanBeOrganizedBy(userId) {
		return api.ErrorUnauthorized(domains.NoOrganizerAccess)
	}

	err3 := ds.DrinkRequirementRepository.Create(drinkRequirement)
	if err3 != nil {
		return api.ErrorInternalServerError(err3)
	}

	return api.Success(drinkRequirement)
}

func (ds DrinkRequirementService) FindById(drinkReqId, userId uint) api.IResponse {
	drinkRequirement, err := ds.DrinkRequirementRepository.FindById(drinkReqId, partyDomains.FullPartyNestedPreload...)
	if err != nil {
		return api.ErrorBadRequest(domains.PartyNotFound)
	}

	if !drinkRequirement.Party.CanBeAccessedBy(userId) {
		return api.ErrorUnauthorized(domains.NoViewAccess)
	}

	return api.Success(drinkRequirement)
}

func (ds DrinkRequirementService) Delete(drinkReqId, userId uint) api.IResponse {
	drinkRequirement, err := ds.DrinkRequirementRepository.FindById(drinkReqId, partyDomains.FullPartyNestedPreload...)
	if err != nil {
		return api.ErrorBadRequest(domains.RequirementNotFound)
	}

	if !drinkRequirement.Party.CanBeOrganizedBy(userId) {
		return api.ErrorUnauthorized(domains.NoOrganizerAccess)
	}

	//todo: put this in transaction
	if err2 := ds.DrinkContributionRepository.DeleteByReqId(drinkReqId); err2 != nil {
		return api.ErrorInternalServerError(err2.Error())
	}

	err3 := ds.DrinkRequirementRepository.Delete(drinkRequirement)
	if err3 != nil {
		return api.ErrorInternalServerError(err3.Error())
	}
	return api.Success("delete_success")
}

func (ds DrinkRequirementService) GetByPartyId(partyId, userId uint) api.IResponse {
	party, err := ds.PartyRepository.FindById(partyId, partyDomains.FullPartyPreload...)
	if err != nil {
		return api.ErrorBadRequest(domains.PartyNotFound)
	}

	if !party.CanBeAccessedBy(userId) {
		return api.ErrorUnauthorized(domains.NoViewAccess)
	}

	drinkReqs, err3 := ds.DrinkRequirementRepository.GetByPartyId(partyId)
	if err3 != nil {
		return api.ErrorInternalServerError(err3.Error())
	}

	return api.Success(drinkReqs)
}
