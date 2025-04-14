package usecases

import (
	"fmt"
	drinkReqDomain "github.com/zsomborCzaban/party_organizer/services/creation/drink_requirement/domains"
	foodReqDomain "github.com/zsomborCzaban/party_organizer/services/creation/food_requirement/domains"
	"github.com/zsomborCzaban/party_organizer/services/creation/party/domains"
	drinkContribDomain "github.com/zsomborCzaban/party_organizer/services/interaction/drink_contributions/domains"
	foodContributionDomains "github.com/zsomborCzaban/party_organizer/services/interaction/food_contributions/domains"
	partyInvitationDomains "github.com/zsomborCzaban/party_organizer/services/managers/party_attendance_manager/domains"
	userDomain "github.com/zsomborCzaban/party_organizer/services/users/user/domains"
	"github.com/zsomborCzaban/party_organizer/utils/api"
	"github.com/zsomborCzaban/party_organizer/utils/repo"
)

type PartyService struct {
	Validator              api.IValidator
	PartyRepository        domains.IPartyRepository
	UserRepository         userDomain.IUserRepository
	DrinkReqRepository     drinkReqDomain.IDrinkRequirementRepository
	DrinkContribRepository drinkContribDomain.IDrinkContributionRepository
	FoodReqRepository      foodReqDomain.IFoodRequirementRepository
	FoodContribRepository  foodContributionDomains.IFoodContributionRepository
	PartyInviteRepository  partyInvitationDomains.IPartyInviteRepository
}

func NewPartyService(repoCollector *repo.RepoCollector, validator api.IValidator) domains.IPartyService {
	return &PartyService{
		Validator:              validator,
		PartyRepository:        repoCollector.PartyRepo,
		UserRepository:         repoCollector.UserRepo,
		DrinkReqRepository:     repoCollector.DrinkReqRepo,
		DrinkContribRepository: repoCollector.DrinkContribRepo,
		FoodReqRepository:      repoCollector.FoodReqRepo,
		FoodContribRepository:  repoCollector.FoodContribRepo,
		PartyInviteRepository:  repoCollector.PartyInviteRepo,
	}
}

func (ps PartyService) Create(partyDTO domains.PartyDTO, userId uint) api.IResponse {
	errors := ps.Validator.Validate(partyDTO)
	if errors != nil {
		return api.ErrorValidation(errors)
	}

	party := partyDTO.TransformToParty()
	party.OrganizerID = userId
	if party.AccessCodeEnabled {
		party.AccessCode = fmt.Sprintf("%d_%s", party.ID, party.AccessCode)
	}

	organizer, err := ps.UserRepository.FindById(userId)
	if err != nil {
		return api.ErrorInternalServerError(domains.DeletedUser)
	}

	err2 := ps.PartyRepository.Create(party)
	if err2 != nil {
		return api.ErrorInternalServerError(err2.Error())
	}

	party.Organizer = *organizer
	return api.Success(party)
}

func (ps PartyService) Get(partyId, userId uint) api.IResponse {
	party, err := ps.PartyRepository.FindById(partyId, domains.FullPartyPreload...)
	if err != nil {
		return api.ErrorInternalServerError(err.Error())
	}

	if !party.CanBeAccessedBy(userId) {
		return api.ErrorUnauthorized("you cannot access this party")
	}

	if !party.CanBeOrganizedBy(userId) {
		party.AccessCode = ""
	}

	return api.Success(party)
}

func (ps PartyService) GetPublicParty(partyId uint) api.IResponse {
	party, err := ps.PartyRepository.FindById(partyId, domains.FullPartyPreload...)
	if err != nil {
		return api.ErrorBadRequest(err.Error())
	}

	if party.Private {
		return api.ErrorUnauthorized("this party is private")
	}

	party.AccessCode = ""

	return api.Success(party)
}

func (ps PartyService) Update(partyDTO domains.PartyDTO, userId uint) api.IResponse {
	err := ps.Validator.Validate(partyDTO)
	if err != nil {
		return api.ErrorValidation(err)
	}

	originalParty, err2 := ps.PartyRepository.FindById(partyDTO.ID, domains.FullPartyPreload...)
	if err2 != nil {
		return api.ErrorBadRequest(err2.Error())
	}

	if !originalParty.CanBeOrganizedBy(userId) {
		return api.ErrorUnauthorized("cannot update other people's party")
	}
	party := partyDTO.TransformToParty()
	party.OrganizerID = originalParty.OrganizerID
	if party.AccessCodeEnabled {
		party.AccessCode = fmt.Sprintf("%d_%s", party.ID, party.AccessCode)
	}

	err3 := ps.PartyRepository.Update(party)
	if err3 != nil {
		return api.ErrorInternalServerError(err3.Error())
	}

	return api.Success(party)
}

func (ps PartyService) Delete(partyId uint, userId uint) api.IResponse {
	//bc the repository layer only checks for id
	party, err := ps.PartyRepository.FindById(partyId, domains.FullPartyPreload...)
	if err != nil {
		return api.ErrorBadRequest(err.Error())
	}

	if !party.CanBeOrganizedBy(userId) {
		return api.ErrorUnauthorized("cannot delete other peoples party")
	}

	//todo: transaction maybe
	err2 := ps.DrinkContribRepository.DeleteByPartyId(partyId)
	err3 := ps.FoodContribRepository.DeleteByPartyId(partyId)
	if err2 != nil || err3 != nil {
		return api.ErrorInternalServerError("unexpected error while deleting the contributions of the party")
	}

	err4 := ps.DrinkReqRepository.DeleteByPartyId(partyId)
	err5 := ps.FoodReqRepository.DeleteByPartyId(partyId)
	if err4 != nil || err5 != nil {
		return api.ErrorInternalServerError("unexpected error while deleting the requirements of the party")
	}

	err6 := ps.PartyInviteRepository.DeleteByPartyId(partyId)
	if err6 != nil {
		return api.ErrorInternalServerError("unexpected error while deleting party invites of the party")
	}

	err7 := ps.PartyRepository.Delete(party)
	if err7 != nil {
		return api.ErrorInternalServerError(err7.Error())
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
		return api.ErrorUnauthorized(domains.NoAccessToParty)
	}

	return api.Success(append(party.Participants, party.Organizer))
}
