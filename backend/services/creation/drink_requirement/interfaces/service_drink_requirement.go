package interfaces

import (
	"github.com/zsomborCzaban/party_organizer/common/api"
	"github.com/zsomborCzaban/party_organizer/services/creation/drink_requirement/domains"
	partyDomains "github.com/zsomborCzaban/party_organizer/services/creation/party/domains"
	"gorm.io/gorm"
)

type DrinkRequirementService struct {
	DrinkRequirementRepository domains.IDrinkRequirementRepository
	Validator                  api.IValidator
	PartyRepository            partyDomains.IPartyRepository
}

func NewDrinkRequirementService(repository domains.IDrinkRequirementRepository, validator api.IValidator, partyRepository partyDomains.IPartyRepository) domains.IDrinkRequirementService {
	return &DrinkRequirementService{
		DrinkRequirementRepository: repository,
		Validator:                  validator,
		PartyRepository:            partyRepository,
	}
}

func (ds DrinkRequirementService) CreateDrinkRequirement(drinkRequirementDTO domains.DrinkRequirementDTO, userId uint) api.IResponse {
	err := ds.Validator.Validate(drinkRequirementDTO)
	if err != nil {
		return api.ErrorValidation(err)
	}

	drinkRequirement := drinkRequirementDTO.TransformToDrinkRequirement()

	party, err2 := ds.PartyRepository.GetParty(drinkRequirement.PartyID)
	if err2 != nil {
		return api.ErrorBadRequest("Party id doesnt exists")
	}

	if party.OrganizerID != userId && userId != 0 {
		return api.ErrorUnauthorized("cannot create drinkRequirements for other peoples party")
	}

	drinkRequirement.Party = *party

	err3 := ds.DrinkRequirementRepository.CreateDrinkRequirement(drinkRequirement)
	if err3 != nil {
		return api.ErrorInternalServerError(err3)
	}

	return api.Success(drinkRequirement.TransformToDrinkRequirementDTO())
}

func (ds DrinkRequirementService) GetDrinkRequirement(id uint) api.IResponse {
	drinkRequirement, err := ds.DrinkRequirementRepository.GetDrinkRequirement(id)

	if err != nil {
		return api.ErrorInternalServerError(err)
	}

	return api.Success(drinkRequirement.TransformToDrinkRequirementDTO())
}

func (ds DrinkRequirementService) UpdateDrinkRequirement(drinkRequirementDTO domains.DrinkRequirementDTO) api.IResponse {
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

func (ds DrinkRequirementService) DeleteDrinkRequirement(id uint) api.IResponse {
	//bc the repository layer only checks for id
	drinkRequirement := &domains.DrinkRequirement{
		Model: gorm.Model{ID: id},
	}

	err := ds.DrinkRequirementRepository.DeleteDrinkRequirement(drinkRequirement)
	if err != nil {
		return api.ErrorInternalServerError(err)
	}
	return api.Success("delete_success")
}

func (ds DrinkRequirementService) GetByPartyId(partyId, userId uint) api.IResponse {
	party, err := ds.PartyRepository.GetParty(partyId)
	if err != nil {
		return api.ErrorBadRequest("party not found")
	}

	if party.OrganizerID != userId && userId != 0 {
		return api.ErrorUnauthorized("you are not the party organizer")
	}

	drinkReqs, err3 := ds.DrinkRequirementRepository.GetByPartyId(partyId)
	if err3 != nil {
		return api.ErrorInternalServerError(err3)
	}

	return api.Success(drinkReqs)
}
