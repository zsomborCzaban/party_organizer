package interfaces

import (
	drinkReqDomain "github.com/zsomborCzaban/party_organizer/services/creation/drink_requirement/domains"
	partyDomains "github.com/zsomborCzaban/party_organizer/services/creation/party/domains"
	"github.com/zsomborCzaban/party_organizer/services/interaction/drink_contributions/domains"
	userDomain "github.com/zsomborCzaban/party_organizer/services/user/domains"
	"github.com/zsomborCzaban/party_organizer/utils/adminUser"
	"github.com/zsomborCzaban/party_organizer/utils/api"
	"github.com/zsomborCzaban/party_organizer/utils/repo"
)

type DrinkContributionService struct {
	Validator              api.IValidator
	ContributionRepository *domains.IDrinkContributionRepository
	UserRepository         *userDomain.IUserRepository
	PartyRepository        *partyDomains.IPartyRepository
	DrinkReqRepository     *drinkReqDomain.IDrinkRequirementRepository
}

func NewDrinkContributionService(repoCollector *repo.RepoCollector, vali api.IValidator) domains.IDrinkContributionService {
	return &DrinkContributionService{
		Validator:              vali,
		ContributionRepository: repoCollector.DrinkContribRepo,
		UserRepository:         repoCollector.UserRepo,
		PartyRepository:        repoCollector.PartyRepo,
		DrinkReqRepository:     repoCollector.DrinkReqRepo,
	}
}

func (ds DrinkContributionService) Create(contribution domains.DrinkContribution, userId uint) api.IResponse {
	err := ds.Validator.Validate(contribution)
	if err != nil {
		return api.ErrorValidation(err.Errors)
	}

	req, err3 := (*ds.DrinkReqRepository).FindById(contribution.DrinkReqId, partyDomains.FullPartyNestedPreload...)
	if err3 != nil {
		return api.ErrorBadRequest(err3.Error())
	}

	if !req.Party.CanBeAccessedBy(userId) {
		return api.ErrorUnauthorized(domains.NO_ACCESS_TO_PARTY)
	}

	contribution.ContributorId = userId
	contribution.PartyId = req.PartyID

	if err5 := (*ds.ContributionRepository).Create(&contribution); err5 != nil {
		return api.ErrorInternalServerError(err5.Error())
	}

	return api.Success(contribution)
}

func (ds DrinkContributionService) Update(contribution domains.DrinkContribution, userId uint) api.IResponse {
	err := ds.Validator.Validate(contribution)
	if err != nil {
		return api.ErrorValidation(err.Errors)
	}

	oldContribution, err2 := (*ds.ContributionRepository).FindById(contribution.ID, partyDomains.FullPartyNestedPreload...)
	if err2 != nil {
		return api.ErrorBadRequest(err2.Error())
	}

	if !oldContribution.Party.CanBeAccessedBy(userId) {
		return api.ErrorUnauthorized(domains.NO_ACCESS_TO_PARTY)
	}

	if oldContribution.ContributorId != userId && userId != adminUser.ADMIN_USER_ID {
		return api.ErrorBadRequest("cannot update other people's contribution")
	}

	if contribution.DrinkReqId != oldContribution.DrinkReqId {
		return api.ErrorBadRequest("cannot change drink requirement of the contribution")
	}

	contribution.ContributorId = oldContribution.ContributorId
	contribution.PartyId = oldContribution.PartyId

	if err6 := (*ds.ContributionRepository).Create(&contribution); err6 != nil {
		return api.ErrorInternalServerError(err6.Error())
	}

	return api.Success(contribution)
}

func (ds DrinkContributionService) Delete(contributionId, userId uint) api.IResponse {
	contribution, err := (*ds.ContributionRepository).FindById(contributionId, partyDomains.FullPartyNestedPreload...)
	if err != nil {
		return api.ErrorBadRequest(err.Error())
	}

	if userId != contribution.ContributorId && !contribution.Party.CanBeOrganizedBy(userId) {
		return api.ErrorUnauthorized("cannot delete other people's contribution")
	}

	if err3 := (*ds.ContributionRepository).Delete(contribution); err3 != nil {
		return api.ErrorInternalServerError(err3.Error())
	}

	return api.Success("contribution deleted successfully")
}

func (ds DrinkContributionService) GetByPartyIdAndContributorId(partyId, contributorId, userId uint) api.IResponse {
	party, err := (*ds.PartyRepository).FindById(partyId, partyDomains.FullPartyNestedPreload...)
	if err != nil {
		return api.ErrorBadRequest(err.Error())
	}

	if party.CanBeAccessedBy(userId) {
		return api.ErrorUnauthorized(domains.NO_ACCESS_TO_PARTY)
	}

	columnNames := []string{"party_id", "contributor_id"}
	values := []interface{}{partyId, contributorId}

	contributions, err := (*ds.ContributionRepository).FindAllBy(columnNames, values, "Contributor")
	if err != nil {
		return api.ErrorInternalServerError(err.Error())
	}

	return api.Success(contributions)
}

func (ds DrinkContributionService) GetByRequirementId(requirementId, userId uint) api.IResponse {
	requirement, err := (*ds.DrinkReqRepository).FindById(requirementId, partyDomains.FullPartyNestedPreload...)
	if err != nil {
		return api.ErrorBadRequest(err.Error())
	}

	if requirement.Party.CanBeAccessedBy(userId) {
		return api.ErrorUnauthorized(domains.NO_ACCESS_TO_PARTY)
	}

	columnNames := []string{"drink_req_id"}
	values := []interface{}{requirementId}

	contributions, err := (*ds.ContributionRepository).FindAllBy(columnNames, values, "Contributor")
	if err != nil {
		return api.ErrorInternalServerError(err.Error())
	}

	return api.Success(contributions)
}

func (ds DrinkContributionService) GetByPartyId(partyId, userId uint) api.IResponse {
	party, err := (*ds.PartyRepository).FindById(partyId, partyDomains.FullPartyPreload...)
	if err != nil {
		return api.ErrorBadRequest(err.Error())
	}

	if !party.CanBeAccessedBy(userId) {
		return api.ErrorUnauthorized(domains.NO_ACCESS_TO_PARTY)
	}

	columnNames := []string{"party_id"}
	values := []interface{}{partyId}

	contributions, err := (*ds.ContributionRepository).FindAllBy(columnNames, values, "Contributor")
	if err != nil {
		return api.ErrorInternalServerError(err.Error())
	}

	return api.Success(contributions)
}
