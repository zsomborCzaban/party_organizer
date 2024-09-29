package interfaces

import (
	"github.com/zsomborCzaban/party_organizer/services/creation/party/domains"
)

type PartyService struct {
	PartyRepository domains.IPartyRepository
	Validator       domains.Validator
}

func NewHospitalService(repository domains.IPartyRepository, validator domains.Validator) domains.IPartyService {
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
