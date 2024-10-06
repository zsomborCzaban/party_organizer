package interfaces

import (
	"github.com/zsomborCzaban/party_organizer/common/api"
	"github.com/zsomborCzaban/party_organizer/services/creation/food_requirement/domains"
	"gorm.io/gorm"
)

type FoodRequirementService struct {
	FoodRequirementRepository domains.IFoodRequirementRepository
	Validator                 api.IValidator
}

func NewFoodRequirementService(repository domains.IFoodRequirementRepository, validator api.IValidator) domains.IFoodRequirementService {
	return &FoodRequirementService{
		FoodRequirementRepository: repository,
		Validator:                 validator,
	}
}

func (ps FoodRequirementService) CreateFoodRequirement(foodRequirementDTO domains.FoodRequirementDTO) api.IResponse {
	errors := ps.Validator.Validate(foodRequirementDTO)
	if errors != nil {
		return api.ErrorValidation(errors)
	}

	foodRequirement := foodRequirementDTO.TransformToFoodRequirement()

	err := ps.FoodRequirementRepository.CreateFoodRequirement(foodRequirement)
	if err != nil {
		return api.ErrorInternalServerError(err)
	}

	return api.Success("create_success")
}

func (ps FoodRequirementService) GetFoodRequirement(id uint) api.IResponse {
	foodRequirement, err := ps.FoodRequirementRepository.GetFoodRequirement(id)

	if err != nil {
		return api.ErrorInternalServerError(err)
	}

	return api.Success(foodRequirement.TransformToFoodRequirementDTO())
}

func (ps FoodRequirementService) UpdateFoodRequirement(foodRequirementDTO domains.FoodRequirementDTO) api.IResponse {
	errors := ps.Validator.Validate(foodRequirementDTO)
	if errors != nil {
		return api.ErrorValidation(errors)
	}

	foodRequirement := foodRequirementDTO.TransformToFoodRequirement()

	err := ps.FoodRequirementRepository.UpdateFoodRequirement(foodRequirement)
	if err != nil {
		return api.ErrorInternalServerError(err)
	}

	return api.Success("update_success")
}

func (ps FoodRequirementService) DeleteFoodRequirement(id uint) api.IResponse {
	//bc the repository layer only checks for id
	foodRequirement := &domains.FoodRequirement{
		Model: gorm.Model{ID: id},
	}

	err := ps.FoodRequirementRepository.DeleteFoodRequirement(foodRequirement)
	if err != nil {
		return api.ErrorInternalServerError(err)
	}
	return api.Success("delete_success")
}
