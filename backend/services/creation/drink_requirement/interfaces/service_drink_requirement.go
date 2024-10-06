package interfaces

import (
	"github.com/zsomborCzaban/party_organizer/common/api"
	"github.com/zsomborCzaban/party_organizer/services/creation/drink_requirement/domains"
	"gorm.io/gorm"
)

type DrinkRequirementService struct {
	DrinkRequirementRepository domains.IDrinkRequirementRepository
	Validator                  api.IValidator
}

func NewDrinkRequirementService(repository domains.IDrinkRequirementRepository, validator api.IValidator) domains.IDrinkRequirementService {
	return &DrinkRequirementService{
		DrinkRequirementRepository: repository,
		Validator:                  validator,
	}
}

func (ps DrinkRequirementService) CreateDrinkRequirement(drinkRequirementDTO domains.DrinkRequirementDTO) api.IResponse {
	errors := ps.Validator.Validate(drinkRequirementDTO)
	if errors != nil {
		return api.ErrorValidation(errors)
	}

	drinkRequirement := drinkRequirementDTO.TransformToDrinkRequirement()

	err := ps.DrinkRequirementRepository.CreateDrinkRequirement(drinkRequirement)
	if err != nil {
		return api.ErrorInternalServerError(err)
	}

	return api.Success("create_success")
}

func (ps DrinkRequirementService) GetDrinkRequirement(id uint) api.IResponse {
	drinkRequirement, err := ps.DrinkRequirementRepository.GetDrinkRequirement(id)

	if err != nil {
		return api.ErrorInternalServerError(err)
	}

	return api.Success(drinkRequirement.TransformToDrinkRequirementDTO())
}

func (ps DrinkRequirementService) UpdateDrinkRequirement(drinkRequirementDTO domains.DrinkRequirementDTO) api.IResponse {
	errors := ps.Validator.Validate(drinkRequirementDTO)
	if errors != nil {
		return api.ErrorValidation(errors)
	}

	drinkRequirement := drinkRequirementDTO.TransformToDrinkRequirement()

	err := ps.DrinkRequirementRepository.UpdateDrinkRequirement(drinkRequirement)
	if err != nil {
		return api.ErrorInternalServerError(err)
	}

	return api.Success("update_success")
}

func (ps DrinkRequirementService) DeleteDrinkRequirement(id uint) api.IResponse {
	//bc the repository layer only checks for id
	drinkRequirement := &domains.DrinkRequirement{
		Model: gorm.Model{ID: id},
	}

	err := ps.DrinkRequirementRepository.DeleteDrinkRequirement(drinkRequirement)
	if err != nil {
		return api.ErrorInternalServerError(err)
	}
	return api.Success("delete_success")
}
