package interfaces

import (
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

	user, err := ps.UserRepository.FindById(userId)
	if err != nil {
		return api.ErrorBadRequest(err.Error())
	}

	party := partyDTO.TransformToParty()
	party.OrganizerID = userId
	party.Organizer = *user

	err2 := ps.PartyRepository.CreateParty(party)
	if err != nil {
		return api.ErrorInternalServerError(err2.Error())
	}

	return api.Success(party)
}

func (ps PartyService) GetParty(partyId, userId uint) api.IResponse {
	party, err := ps.PartyRepository.FindById(partyId)
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

	//todo: cehck against original party
	if partyDTO.OrganizerID != userId && userId != 0 {
		return api.ErrorUnauthorized("cannot update other people's party")
	}

	party := partyDTO.TransformToParty()

	err2 := ps.PartyRepository.UpdateParty(party)
	if err2 != nil {
		return api.ErrorInternalServerError(err2.Error())
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

func (ps PartyService) AddUserToParty(partyId, userId uint) api.IResponse {
	party, err := ps.PartyRepository.FindById(partyId)
	if err != nil {
		return api.ErrorBadRequest(err.Error())
	}

	user, err2 := ps.UserRepository.FindById(userId)
	if err2 != nil {
		return api.ErrorInternalServerError(err2.Error())
	}

	if err3 := ps.PartyRepository.AddUserToParty(party, user); err3 != nil {
		return api.ErrorInternalServerError(err3.Error())
	}

	return api.Success("user added to party")
}
