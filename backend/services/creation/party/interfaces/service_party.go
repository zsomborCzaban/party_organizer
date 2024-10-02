package interfaces

import (
	"github.com/zsomborCzaban/party_organizer/services/creation/party/domains"
	"gorm.io/gorm"
)

type PartyService struct {
	PartyRepository domains.IPartyRepository
	Validator       domains.IValidator
}

func NewPartyService(repository domains.IPartyRepository, validator domains.IValidator) domains.IPartyService {
	return &PartyService{
		PartyRepository: repository,
		Validator:       validator,
	}
}

func (ps PartyService) CreateParty(partyDTO domains.PartyDTO) domains.IResponse {
	errors := ps.Validator.Validate(partyDTO)
	if errors != nil {
		return domains.ErrorValidation(errors)
	}

	party := partyDTO.TransformToParty()

	err := ps.PartyRepository.CreateParty(party)
	if err != nil {
		return domains.ErrorInternalServerError(err)
	}

	return domains.Success("create_success")
}

func (ps PartyService) GetParty(id uint) domains.IResponse {
	party, err := ps.PartyRepository.GetParty(id)

	if err != nil {
		return domains.ErrorInternalServerError(err)
	}

	return domains.Success(party.TransformToPartyDTO())
}

func (ps PartyService) UpdateParty(partyDTO domains.PartyDTO) domains.IResponse {
	errors := ps.Validator.Validate(partyDTO)
	if errors != nil {
		return domains.ErrorValidation(errors)
	}

	party := partyDTO.TransformToParty()

	err := ps.PartyRepository.UpdateParty(party)
	if err != nil {
		return domains.ErrorInternalServerError(err)
	}

	return domains.Success("update_success")
}

func (ps PartyService) DeleteParty(id uint) domains.IResponse {
	//bc the repository layer only checks for id
	party := &domains.Party{
		Model: gorm.Model{ID: id},
	}

	err := ps.PartyRepository.DeleteParty(party)
	if err != nil {
		return domains.ErrorInternalServerError(err)
	}
	return domains.Success("delete_success")
}
