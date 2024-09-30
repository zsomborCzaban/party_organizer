package interfaces

import (
	"github.com/zsomborCzaban/party_organizer/services/creation/party/domains"
)

type PartyService struct {
	PartyRepository domains.IPartyRepository
	Validator       domains.Validator
}

func NewPartyService(repository domains.IPartyRepository, validator domains.Validator) domains.IPartyService {
	return &PartyService{
		PartyRepository: repository,
		Validator:       validator,
	}
}

func (ps PartyService) CreateParty(partyDTO domains.PartyDTO) domains.IResponse {
	errors := ps.Validator.Validate(partyDTO)
	if errors != nil {
		return domains.ErrorValidation(partyDTO)
	}

	party, err := ps.PartyRepository.CreateParty(&partyDTO)
	if err != nil {
		return domains.ErrorInternalServerError(err)
	}

	return domains.Success(party)
}

func (ps PartyService) GetParty(id uint) domains.IResponse {
	party, err := ps.PartyRepository.GetParty(id)

	if err != nil {
		return domains.ErrorInternalServerError(err)
	}

	return domains.Success(party)
}

func (ps PartyService) UpdateParty(partyDTO domains.PartyDTO) domains.IResponse {
	errors := ps.Validator.Validate(partyDTO)
	if errors != nil {
		return domains.ErrorValidation(errors)
	}

	party, err := ps.PartyRepository.UpdateParty(&partyDTO)
	if err != nil {
		return domains.ErrorInternalServerError(err)
	}

	return domains.Success(party)
}

func (ps PartyService) DeleteParty(id uint) domains.IResponse {
	party, err := ps.PartyRepository.DeleteParty(id)
	if err != nil {
		return domains.ErrorInternalServerError(err)
	}
	return domains.Success(party)
}
