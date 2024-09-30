package interfaces

import (
	"github.com/zsomborCzaban/party_organizer/services/creation/drink_requirement/domains"
)

type DrinkRequirementService struct {
	DrinkRequirementRepository domains.IDrinkRequirementRepository
	Validator                  domains.Validator
}

func NewDrinkRequirementService(repository domains.IDrinkRequirementRepository, validator domains.Validator) domains.IDrinkRequirementService {
	return &DrinkRequirementService{
		DrinkRequirementRepository: repository,
		Validator:                  validator,
	}
}

func (ds DrinkRequirementService) CreateDrinkRequirement(partyDTO domains.DrinkRequirementDTO) domains.IResponse {
	errors := ds.Validator.Validate(partyDTO)
	if errors != nil {
		return domains.ErrorValidation(errors)
	}

	party, err := ds.DrinkRequirementRepository.CreateDrinkRequirement(&partyDTO)
	if err != nil {
		return domains.ErrorInternalServerError(err)
	}

	return domains.Success(party.TransformToDrinkRequirementDTO())
}

func (ds DrinkRequirementService) GetDrinkRequirement(id uint) domains.IResponse {
	party, err := ds.DrinkRequirementRepository.GetDrinkRequirement(id)

	if err != nil {
		return domains.ErrorInternalServerError(err)
	}

	return domains.Success(party.TransformToDrinkRequirementDTO())
}

func (ds DrinkRequirementService) UpdateDrinkRequirement(partyDTO domains.DrinkRequirementDTO) domains.IResponse {
	errors := ds.Validator.Validate(partyDTO)
	if errors != nil {
		return domains.ErrorValidation(errors)
	}

	party, err := ds.DrinkRequirementRepository.UpdateDrinkRequirement(&partyDTO)
	if err != nil {
		return domains.ErrorInternalServerError(err)
	}

	return domains.Success(party.TransformToDrinkRequirementDTO())
}

func (ds DrinkRequirementService) DeleteDrinkRequirement(id uint) domains.IResponse {
	party, err := ds.DrinkRequirementRepository.DeleteDrinkRequirement(id)
	if err != nil {
		return domains.ErrorInternalServerError(err)
	}
	return domains.Success(party.TransformToDrinkRequirementDTO())
}
