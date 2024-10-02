package interfaces

import (
	"github.com/zsomborCzaban/party_organizer/services/creation/drink_requirement/domains"
	"gorm.io/gorm"
)

type DrinkRequirementService struct {
	DrinkRequirementRepository domains.IDrinkRequirementRepository
	Validator                  domains.IValidator
}

func NewDrinkRequirementService(repository domains.IDrinkRequirementRepository, validator domains.IValidator) domains.IDrinkRequirementService {
	return &DrinkRequirementService{
		DrinkRequirementRepository: repository,
		Validator:                  validator,
	}
}

func (ps DrinkRequirementService) CreateDrinkRequirement(drinkRequirementDTO domains.DrinkRequirementDTO) domains.IResponse {
	errors := ps.Validator.Validate(drinkRequirementDTO)
	if errors != nil {
		return domains.ErrorValidation(errors)
	}

	drinkRequirement := drinkRequirementDTO.TransformToDrinkRequirement()

	err := ps.DrinkRequirementRepository.CreateDrinkRequirement(drinkRequirement)
	if err != nil {
		return domains.ErrorInternalServerError(err)
	}

	return domains.Success("create_success")
}

func (ps DrinkRequirementService) GetDrinkRequirement(id uint) domains.IResponse {
	drinkRequirement, err := ps.DrinkRequirementRepository.GetDrinkRequirement(id)

	if err != nil {
		return domains.ErrorInternalServerError(err)
	}

	return domains.Success(drinkRequirement.TransformToDrinkRequirementDTO())
}

func (ps DrinkRequirementService) UpdateDrinkRequirement(drinkRequirementDTO domains.DrinkRequirementDTO) domains.IResponse {
	errors := ps.Validator.Validate(drinkRequirementDTO)
	if errors != nil {
		return domains.ErrorValidation(errors)
	}

	drinkRequirement := drinkRequirementDTO.TransformToDrinkRequirement()

	err := ps.DrinkRequirementRepository.UpdateDrinkRequirement(drinkRequirement)
	if err != nil {
		return domains.ErrorInternalServerError(err)
	}

	return domains.Success("update_success")
}

func (ps DrinkRequirementService) DeleteDrinkRequirement(id uint) domains.IResponse {
	//bc the repository layer only checks for id
	drinkRequirement := &domains.DrinkRequirement{
		Model: gorm.Model{ID: id},
	}

	err := ps.DrinkRequirementRepository.DeleteDrinkRequirement(drinkRequirement)
	if err != nil {
		return domains.ErrorInternalServerError(err)
	}
	return domains.Success("delete_success")
}
