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

	party := partyDTO.TransformToParty()
	party.OrganizerID = userId

	err := ps.PartyRepository.CreateParty(party)
	if err != nil {
		return api.ErrorInternalServerError(err.Error())
	}

	return api.Success("create_success")
}

func (ps PartyService) GetParty(partyId uint) api.IResponse {
	party, err := ps.PartyRepository.GetParty(partyId)

	if err != nil {
		return api.ErrorInternalServerError(err.Error())
	}

	return api.Success(party.TransformToPartyDTO())
}

func (ps PartyService) UpdateParty(partyDTO domains.PartyDTO, userId uint) api.IResponse {
	err := ps.Validator.Validate(partyDTO)
	if err != nil {
		return api.ErrorValidation(err)
	}

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
	user, err := ps.UserRepository.FindById(userId)
	if err != nil {
		return api.ErrorInternalServerError(err.Error())
	}

	if err2 := ps.PartyRepository.AddUserToParty(partyId, user); err2 != nil {
		return api.ErrorInternalServerError(err2.Error())
	}

	return api.Success("user added to party")
}
