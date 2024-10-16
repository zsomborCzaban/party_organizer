package interfaces

import (
	"github.com/zsomborCzaban/party_organizer/common/api"
	"github.com/zsomborCzaban/party_organizer/services/creation/drink_requirement/domains"
	partyDomains "github.com/zsomborCzaban/party_organizer/services/creation/party/domains"
	drinkContributionDomains "github.com/zsomborCzaban/party_organizer/services/interaction/drink_contributions/domains"
)

type DrinkRequirementService struct {
	DrinkRequirementRepository  domains.IDrinkRequirementRepository
	Validator                   api.IValidator
	PartyRepository             partyDomains.IPartyRepository
	DrinkContributionRepository drinkContributionDomains.IDrinkContributionRepository
}

func NewDrinkRequirementService(repository domains.IDrinkRequirementRepository, validator api.IValidator, partyRepository partyDomains.IPartyRepository, drinkContributionRepository drinkContributionDomains.IDrinkContributionRepository) domains.IDrinkRequirementService {
	return &DrinkRequirementService{
		DrinkRequirementRepository:  repository,
		Validator:                   validator,
		PartyRepository:             partyRepository,
		DrinkContributionRepository: drinkContributionRepository,
	}
}

func (ds DrinkRequirementService) CreateDrinkRequirement(drinkRequirementDTO domains.DrinkRequirementDTO, userId uint) api.IResponse {
	err := ds.Validator.Validate(drinkRequirementDTO)
	if err != nil {
		return api.ErrorValidation(err)
	}

	drinkRequirement := drinkRequirementDTO.TransformToDrinkRequirement()

	party, err2 := ds.PartyRepository.FindById(drinkRequirement.PartyID)
	if err2 != nil {
		return api.ErrorBadRequest("Party id doesnt exists")
	}

	if party.CanBeOrganizedBy(userId) {
		return api.ErrorUnauthorized("cannot create drinkRequirements for other peoples party")
	}

	drinkRequirement.Party = *party

	err3 := ds.DrinkRequirementRepository.CreateDrinkRequirement(drinkRequirement)
	if err3 != nil {
		return api.ErrorInternalServerError(err3)
	}

	return api.Success(drinkRequirement.TransformToDrinkRequirementDTO())
}

func (ds DrinkRequirementService) GetDrinkRequirement(drinkReqId, userId uint) api.IResponse {
	drinkRequirement, err := ds.DrinkRequirementRepository.FindById(drinkReqId)
	if err != nil {
		return api.ErrorInternalServerError(err)
	}

	if !drinkRequirement.Party.CanBeAccessedBy(userId) {
		return api.ErrorUnauthorized("you are not in the party")
	}

	return api.Success(drinkRequirement.TransformToDrinkRequirementDTO())
}

func (ds DrinkRequirementService) UpdateDrinkRequirement(drinkRequirementDTO domains.DrinkRequirementDTO, userId uint) api.IResponse {
	//this wont be used!

	err := ds.Validator.Validate(drinkRequirementDTO)
	if err != nil {
		return api.ErrorValidation(err)
	}

	drinkRequirement := drinkRequirementDTO.TransformToDrinkRequirement()

	err2 := ds.DrinkRequirementRepository.UpdateDrinkRequirement(drinkRequirement)
	if err2 != nil {
		return api.ErrorInternalServerError(err2)
	}

	return api.Success("update_success")
}

func (ds DrinkRequirementService) DeleteDrinkRequirement(drinkReqId, userId uint) api.IResponse {
	drinkRequirement, err := ds.DrinkRequirementRepository.FindById(drinkReqId)
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
	party, err := ds.PartyRepository.FindById(partyId)
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
