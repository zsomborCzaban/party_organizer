package interfaces

import (
	"fmt"
	"github.com/zsomborCzaban/party_organizer/common/api"
	"github.com/zsomborCzaban/party_organizer/services/creation/party/domains"
	userDomain "github.com/zsomborCzaban/party_organizer/services/user/domains"
	"gorm.io/gorm"
)

type PartyService struct {
	PartyRepository domains.IPartyRepository
	UserRepository  userDomain.IUserRepository
	Validator       api.IValidator
}

func NewPartyService(repo domains.IPartyRepository, validator api.IValidator, userRepo userDomain.IUserRepository) domains.IPartyService {
	return &PartyService{
		PartyRepository: repo,
		UserRepository:  userRepo,
		Validator:       validator,
	}
}

func (ps PartyService) CreateParty(partyDTO domains.PartyDTO, userId uint) api.IResponse {
	errors := ps.Validator.Validate(partyDTO)
	if errors != nil {
		return api.ErrorValidation(errors)
	}

	party := partyDTO.TransformToParty()
	party.OrganizerID = userId
	if party.AccessCodeEnabled {
		party.AccessCode = fmt.Sprintf("%d_%s", party.ID, party.AccessCode)
	}

	err2 := ps.PartyRepository.CreateParty(party)
	if err2 != nil {
		return api.ErrorInternalServerError(err2.Error())
	}

	return api.Success(party)
}

func (ps PartyService) GetParty(partyId, userId uint) api.IResponse {
	party, err := ps.PartyRepository.FindById(partyId, domains.FullPartyPreload...)
	if err != nil {
		return api.ErrorInternalServerError(err.Error())
	}

	if !party.CanBeAccessedBy(userId) {
		return api.ErrorUnauthorized("you have to be in the party for private parties")
	}

	if !party.CanBeOrganizedBy(userId) {
		party.AccessCode = ""
	}

	return api.Success(party.TransformToPartyDTO())
}

func (ps PartyService) UpdateParty(partyDTO domains.PartyDTO, userId uint) api.IResponse {
	err := ps.Validator.Validate(partyDTO)
	if err != nil {
		return api.ErrorValidation(err)
	}

	originalParty, err2 := ps.PartyRepository.FindById(partyDTO.ID, "Organizer")
	if err2 != nil {
		return api.ErrorBadRequest(err2.Error())
	}

	if !originalParty.CanBeOrganizedBy(userId) {
		return api.ErrorUnauthorized("cannot update other people's party")
	}
	party := partyDTO.TransformToParty()
	party.OrganizerID = originalParty.OrganizerID

	err3 := ps.PartyRepository.UpdateParty(party)
	if err3 != nil {
		return api.ErrorInternalServerError(err3.Error())
	}

	return api.Success(party)
}

// todo: come back here and do this
func (ps PartyService) DeleteParty(id uint) api.IResponse {
	//bc the repository layer only checks for id
	party := &domains.Party{
		Model: gorm.Model{ID: id},
	}

	err := ps.PartyRepository.DeleteParty(party)
	if err != nil {
		return api.ErrorInternalServerError(err.Error())
	}
	return api.Success("delete_success")
}

func (ps PartyService) GetPublicParties() api.IResponse {
	parties, err := ps.PartyRepository.GetPublicParties()
	if err != nil {
		return api.ErrorInternalServerError(err.Error())
	}
	return api.Success(parties)
}

func (ps PartyService) GetPartiesByOrganizerId(id uint) api.IResponse {
	parties, err := ps.PartyRepository.GetPartiesByOrganizerId(id)

	if err != nil {
		return api.ErrorInternalServerError(err.Error())
	}

	//maybe transform to dto before
	return api.Success(parties)
}

func (ps PartyService) GetPartiesByParticipantId(id uint) api.IResponse {
	parties, err := ps.PartyRepository.GetPartiesByParticipantId(id)

	if err != nil {
		return api.ErrorInternalServerError(err.Error())
	}

	return api.Success(parties)
}

func (ps PartyService) GetParticipants(partyId, userId uint) api.IResponse {
	party, err := ps.PartyRepository.FindById(partyId, domains.FullPartyPreload...)
	if err != nil {
		return api.ErrorBadRequest(err.Error())
	}

	if !party.CanBeAccessedBy(userId) {
		return api.ErrorUnauthorized(err.Error())
	}

	return api.Success(append(party.Participants, party.Organizer))
}
