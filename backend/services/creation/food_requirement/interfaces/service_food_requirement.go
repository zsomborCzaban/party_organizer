package interfaces

import (
	"github.com/zsomborCzaban/party_organizer/common/api"
	"github.com/zsomborCzaban/party_organizer/services/creation/food_requirement/domains"
	partyDomains "github.com/zsomborCzaban/party_organizer/services/creation/party/domains"
	"gorm.io/gorm"
)

type FoodRequirementService struct {
	FoodRequirementRepository domains.IFoodRequirementRepository
	Validator                 api.IValidator
	PartyRepository           partyDomains.IPartyRepository
}

func NewFoodRequirementService(repository domains.IFoodRequirementRepository, validator api.IValidator, partyRepository partyDomains.IPartyRepository) domains.IFoodRequirementService {
	return &FoodRequirementService{
		FoodRequirementRepository: repository,
		Validator:                 validator,
		PartyRepository:           partyRepository,
	}
}

func (fs FoodRequirementService) CreateFoodRequirement(foodRequirementDTO domains.FoodRequirementDTO, userId uint) api.IResponse {
	err := fs.Validator.Validate(foodRequirementDTO)
	if err != nil {
		return api.ErrorValidation(err)
	}

	foodRequirement := foodRequirementDTO.TransformToFoodRequirement()

	party, err2 := fs.PartyRepository.GetParty(foodRequirement.PartyID)
	if err2 != nil {
		return api.ErrorBadRequest("partyId does not exists")
	}

	if foodRequirement.PartyID != userId && userId != 0 {
		return api.ErrorUnauthorized("cannot create foodRequirement for other people's party")
	}

	foodRequirement.Party = *party

	err3 := fs.FoodRequirementRepository.CreateFoodRequirement(foodRequirement)
	if err3 != nil {
		return api.ErrorInternalServerError(err)
	}

	return api.Success(foodRequirement.TransformToFoodRequirementDTO())
}

func (fs FoodRequirementService) GetFoodRequirement(id uint) api.IResponse {
	foodRequirement, err := fs.FoodRequirementRepository.GetFoodRequirement(id)

	if err != nil {
		return api.ErrorInternalServerError(err)
	}

	return api.Success(foodRequirement.TransformToFoodRequirementDTO())
}

func (fs FoodRequirementService) UpdateFoodRequirement(foodRequirementDTO domains.FoodRequirementDTO) api.IResponse {
	errors := fs.Validator.Validate(foodRequirementDTO)
	if errors != nil {
		return api.ErrorValidation(errors)
	}

	foodRequirement := foodRequirementDTO.TransformToFoodRequirement()

	err := fs.FoodRequirementRepository.UpdateFoodRequirement(foodRequirement)
	if err != nil {
		return api.ErrorInternalServerError(err)
	}

	return api.Success("update_success")
}

func (fs FoodRequirementService) DeleteFoodRequirement(id uint) api.IResponse {
	//bc the repository layer only checks for id
	foodRequirement := &domains.FoodRequirement{
		Model: gorm.Model{ID: id},
	}

	err := fs.FoodRequirementRepository.DeleteFoodRequirement(foodRequirement)
	if err != nil {
		return api.ErrorInternalServerError(err)
	}
	return api.Success("delete_success")
}

func (fs FoodRequirementService) GetByPartyId(partyId, userId uint) api.IResponse {
	party, err := fs.PartyRepository.GetParty(partyId)
	if err != nil {
		return api.ErrorBadRequest("party not found")
	}

	if party.OrganizerID != userId && userId != 0 {
		return api.ErrorUnauthorized("you are not the party organizer")
	}

	foodReqs, err3 := fs.FoodRequirementRepository.GetByPartyId(partyId)
	if err3 != nil {
		return api.ErrorInternalServerError(err3)
	}

	return api.Success(foodReqs)
}
