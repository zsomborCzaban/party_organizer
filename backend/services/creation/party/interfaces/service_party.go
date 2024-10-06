package interfaces

import (
	"github.com/zsomborCzaban/party_organizer/common/api"
	"github.com/zsomborCzaban/party_organizer/services/creation/party/domains"
	"gorm.io/gorm"
)

type PartyService struct {
	PartyRepository domains.IPartyRepository
	Validator       api.IValidator
}

func NewPartyService(repository domains.IPartyRepository, validator api.IValidator) domains.IPartyService {
	return &PartyService{
		PartyRepository: repository,
		Validator:       validator,
	}
}

func (ps PartyService) CreateParty(partyDTO domains.PartyDTO) api.IResponse {
	errors := ps.Validator.Validate(partyDTO)
	if errors != nil {
		return api.ErrorValidation(errors)
	}

	party := partyDTO.TransformToParty()

	err := ps.PartyRepository.CreateParty(party)
	if err != nil {
		return api.ErrorInternalServerError(err)
	}

	return api.Success("create_success")
}

func (ps PartyService) GetParty(id uint) api.IResponse {
	party, err := ps.PartyRepository.GetParty(id)

	if err != nil {
		return api.ErrorInternalServerError(err)
	}

	return api.Success(party.TransformToPartyDTO())
}

func (ps PartyService) UpdateParty(partyDTO domains.PartyDTO) api.IResponse {
	errors := ps.Validator.Validate(partyDTO)
	if errors != nil {
		return api.ErrorValidation(errors)
	}

	party := partyDTO.TransformToParty()

	err := ps.PartyRepository.UpdateParty(party)
	if err != nil {
		return api.ErrorInternalServerError(err)
	}

	return api.Success("update_success")
}

func (ps PartyService) DeleteParty(id uint) api.IResponse {
	//bc the repository layer only checks for id
	party := &domains.Party{
		Model: gorm.Model{ID: id},
	}

	err := ps.PartyRepository.DeleteParty(party)
	if err != nil {
		return api.ErrorInternalServerError(err)
	}
	return api.Success("delete_success")
}
