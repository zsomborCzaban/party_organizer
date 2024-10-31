package interfaces

import (
	"github.com/zsomborCzaban/party_organizer/common/adminUser"
	"github.com/zsomborCzaban/party_organizer/common/api"
	drinkReqDomain "github.com/zsomborCzaban/party_organizer/services/creation/drink_requirement/domains"
	partyDomains "github.com/zsomborCzaban/party_organizer/services/creation/party/domains"
	"github.com/zsomborCzaban/party_organizer/services/interaction/drink_contributions/domains"
	userDomain "github.com/zsomborCzaban/party_organizer/services/user/domains"
)

type DrinkContributionService struct {
	Validator              api.IValidator
	ContributionRepository domains.IDrinkContributionRepository
	UserRepository         userDomain.IUserRepository
	PartyRepository        partyDomains.IPartyRepository
	DrinkReqRepository     drinkReqDomain.IDrinkRequirementRepository
}

func NewDrinkContributionService(contributionRepo domains.IDrinkContributionRepository, vali api.IValidator, userRepo userDomain.IUserRepository, partyRepo partyDomains.IPartyRepository, drinkReqRepo drinkReqDomain.IDrinkRequirementRepository) domains.IDrinkContributionService {
	return &DrinkContributionService{
		Validator:              vali,
		ContributionRepository: contributionRepo,
		UserRepository:         userRepo,
		PartyRepository:        partyRepo,
		DrinkReqRepository:     drinkReqRepo,
	}
}

func (ds DrinkContributionService) Create(contribution domains.DrinkContribution, userId uint) api.IResponse {
	err := ds.Validator.Validate(contribution)
	if err != nil {
		return api.ErrorValidation(err.Errors)
	}

	contributor, err2 := ds.UserRepository.FindById(userId)
	if err2 != nil {
		return api.ErrorBadRequest(err2.Error())
	}

	drinkReq, err3 := ds.DrinkReqRepository.FindById(contribution.DrinkReqId)
	if err3 != nil {
		return api.ErrorBadRequest(err3.Error())
	}

	party, err4 := ds.PartyRepository.FindById(drinkReq.PartyID)
	if err4 != nil {
		return api.ErrorBadRequest(err4.Error())
	}

	if !party.CanBeAccessedBy(userId) {
		return api.ErrorUnauthorized(domains.NO_ACCESS_TO_PARTY)
	}

	contribution.ContributorId = userId
	contribution.Contributor = *contributor
	contribution.PartyId = party.ID
	contribution.DrinkReq = *drinkReq

	if err5 := ds.ContributionRepository.Create(&contribution); err5 != nil {
		return api.ErrorInternalServerError(err5.Error())
	}

	return api.Success(contribution)
}

func (ds DrinkContributionService) Update(contribution domains.DrinkContribution, userId uint) api.IResponse {
	err := ds.Validator.Validate(contribution)
	if err != nil {
		return api.ErrorValidation(err.Errors)
	}

	oldContribution, err2 := ds.ContributionRepository.FindById(contribution.ID)
	if err2 != nil {
		return api.ErrorBadRequest(err2.Error())
	}

	drinkReq, err3 := ds.DrinkReqRepository.FindById(contribution.DrinkReqId)
	if err3 != nil {
		return api.ErrorBadRequest(err3.Error())
	}
	party := drinkReq.Party

	if !party.CanBeAccessedBy(userId) {
		return api.ErrorUnauthorized(domains.NO_ACCESS_TO_PARTY)
	}

	if oldContribution.ContributorId != userId && userId != adminUser.ADMIN_USER_ID {
		return api.ErrorBadRequest("cannot update other people's contribution")
	}

	if drinkReq.PartyID != contribution.PartyId {
		return api.ErrorBadRequest("drink requirement doesnt belong to party")
	}

	contribution.ContributorId = oldContribution.ContributorId
	contribution.Contributor = oldContribution.Contributor
	contribution.PartyId = party.ID
	contribution.DrinkReq = *drinkReq

	if err6 := ds.ContributionRepository.Create(&contribution); err6 != nil {
		return api.ErrorInternalServerError(err6.Error())
	}

	return api.Success(contribution)
}

func (ds DrinkContributionService) Delete(contributionId, userId uint) api.IResponse {
	contribution, err := ds.ContributionRepository.FindById(contributionId)
	if err != nil {
		return api.ErrorBadRequest(err.Error())
	}

	if userId != contribution.ContributorId && userId != contribution.DrinkReq.Party.OrganizerID && userId != adminUser.ADMIN_USER_ID {
		return api.ErrorUnauthorized("cannot delete other people's contribution")
	}

	if err2 := ds.ContributionRepository.Delete(contribution); err2 != nil {
		return api.ErrorInternalServerError(err2.Error())
	}

	return api.Success("contribution deleted successfully")
}

func (ds DrinkContributionService) GetByPartyIdAndContributorId(partyId, contributorId, userId uint) api.IResponse {
	party, err := ds.PartyRepository.FindById(partyId)
	if err != nil {
		return api.ErrorBadRequest(err.Error())
	}

	if party.CanBeAccessedBy(userId) {
		return api.ErrorUnauthorized(domains.NO_ACCESS_TO_PARTY)
	}

	columnNames := []string{"party_id", "contributor_id"}
	values := []interface{}{partyId, contributorId}

	contributions, err := ds.ContributionRepository.FindAllBy(columnNames, values)
	if err != nil {
		return api.ErrorInternalServerError(err.Error())
	}

	return api.Success(contributions)
}

func (ds DrinkContributionService) GetByRequirementId(requirementId, userId uint) api.IResponse {
	requirement, err := ds.DrinkReqRepository.FindById(requirementId)
	if err != nil {
		return api.ErrorBadRequest(err.Error())
	}

	if requirement.Party.CanBeAccessedBy(userId) {
		return api.ErrorUnauthorized(domains.NO_ACCESS_TO_PARTY)
	}

	columnNames := []string{"drink_req_id"}
	values := []interface{}{requirementId}

	contributions, err := ds.ContributionRepository.FindAllBy(columnNames, values)
	if err != nil {
		return api.ErrorInternalServerError(err.Error())
	}

	return api.Success(contributions)
}

func (ds DrinkContributionService) GetByPartyId(partyId, userId uint) api.IResponse {
	party, err := ds.PartyRepository.FindById(partyId)
	if err != nil {
		return api.ErrorBadRequest(err.Error())
	}

	if !party.CanBeAccessedBy(userId) {
		return api.ErrorUnauthorized(domains.NO_ACCESS_TO_PARTY)
	}

	columnNames := []string{"party_id"}
	values := []interface{}{partyId}

	contributions, err := ds.ContributionRepository.FindAllBy(columnNames, values)
	if err != nil {
		return api.ErrorInternalServerError(err.Error())
	}

	return api.Success(contributions)
}
