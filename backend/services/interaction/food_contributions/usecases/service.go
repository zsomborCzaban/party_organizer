package usecases

import (
	foodReqDomain "github.com/zsomborCzaban/party_organizer/services/creation/food_requirement/domains"
	partyDomains "github.com/zsomborCzaban/party_organizer/services/creation/party/domains"
	"github.com/zsomborCzaban/party_organizer/services/interaction/food_contributions/domains"
	userDomain "github.com/zsomborCzaban/party_organizer/services/users/user/domains"
	"github.com/zsomborCzaban/party_organizer/utils/adminUser"
	"github.com/zsomborCzaban/party_organizer/utils/api"
	"github.com/zsomborCzaban/party_organizer/utils/repo"
)

type FoodContributionService struct {
	Validator              api.IValidator
	ContributionRepository domains.IFoodContributionRepository
	UserRepository         userDomain.IUserRepository
	PartyRepository        partyDomains.IPartyRepository
	FoodReqRepository      foodReqDomain.IFoodRequirementRepository
}

func NewFoodContributionService(repoCollector *repo.RepoCollector, vali api.IValidator) domains.IFoodContributionService {
	return &FoodContributionService{
		Validator:              vali,
		ContributionRepository: repoCollector.FoodContribRepo,
		UserRepository:         repoCollector.UserRepo,
		PartyRepository:        repoCollector.PartyRepo,
		FoodReqRepository:      repoCollector.FoodReqRepo,
	}
}

func (ds FoodContributionService) Create(contribution domains.FoodContribution, userId uint) api.IResponse {
	err := ds.Validator.Validate(contribution)
	if err != nil {
		return api.ErrorValidation(err.Errors)
	}

	req, err3 := ds.FoodReqRepository.FindById(contribution.FoodReqId, partyDomains.FullPartyNestedPreload...)
	if err3 != nil {
		return api.ErrorBadRequest(err3.Error())
	}

	if !req.Party.CanBeAccessedBy(userId) {
		return api.ErrorUnauthorized(domains.NO_ACCESS_TO_PARTY)
	}

	contribution.ContributorId = userId
	contribution.PartyId = req.PartyID

	if err5 := ds.ContributionRepository.Create(&contribution); err5 != nil {
		return api.ErrorInternalServerError(err5.Error())
	}

	return api.Success(contribution)
}

func (ds FoodContributionService) Update(contribution domains.FoodContribution, userId uint) api.IResponse {
	err := ds.Validator.Validate(contribution)
	if err != nil {
		return api.ErrorValidation(err.Errors)
	}

	oldContribution, err2 := ds.ContributionRepository.FindById(contribution.ID, partyDomains.FullPartyNestedPreload...)
	if err2 != nil {
		return api.ErrorBadRequest(err2.Error())
	}

	if !oldContribution.Party.CanBeAccessedBy(userId) {
		return api.ErrorUnauthorized(domains.NO_ACCESS_TO_PARTY)
	}

	if oldContribution.ContributorId != userId && userId != adminUser.ADMIN_USER_ID {
		return api.ErrorBadRequest("cannot update other people's contribution")
	}

	if oldContribution.FoodReqId != contribution.FoodReqId {
		return api.ErrorBadRequest("cannot change food requirement of the contribution")
	}

	contribution.ContributorId = oldContribution.ContributorId
	contribution.PartyId = oldContribution.PartyId

	if err6 := ds.ContributionRepository.Update(&contribution); err6 != nil {
		return api.ErrorInternalServerError(err6.Error())
	}

	return api.Success(contribution)
}

func (ds FoodContributionService) Delete(contributionId, userId uint) api.IResponse {
	contribution, err := ds.ContributionRepository.FindById(contributionId, "Party.Organizer")
	if err != nil {
		return api.ErrorBadRequest(err.Error())
	}

	if userId != contribution.ContributorId && !contribution.Party.CanBeOrganizedBy(userId) {
		return api.ErrorUnauthorized("cannot delete other people's contribution")
	}

	if err2 := ds.ContributionRepository.Delete(contribution); err2 != nil {
		return api.ErrorInternalServerError(err2.Error())
	}

	return api.Success("contribution deleted successfully")
}

func (ds FoodContributionService) GetByPartyIdAndContributorId(partyId, contributorId, userId uint) api.IResponse {
	party, err := ds.PartyRepository.FindById(partyId, partyDomains.FullPartyPreload...)
	if err != nil {
		return api.ErrorBadRequest(err.Error())
	}

	if !party.CanBeAccessedBy(userId) {
		return api.ErrorUnauthorized(domains.NO_ACCESS_TO_PARTY)
	}

	columnNames := []string{"party_id", "contributor_id"}
	values := []interface{}{partyId, contributorId}

	contributions, err := ds.ContributionRepository.FindAllBy(columnNames, values, "Contributor", "FoodReq")
	if err != nil {
		return api.ErrorInternalServerError(err.Error())
	}

	return api.Success(contributions)
}

func (ds FoodContributionService) GetByRequirementId(requirementId, userId uint) api.IResponse {
	requirement, err := ds.FoodReqRepository.FindById(requirementId, partyDomains.FullPartyNestedPreload...)
	if err != nil {
		return api.ErrorBadRequest(err.Error())
	}

	if !requirement.Party.CanBeAccessedBy(userId) {
		return api.ErrorUnauthorized(domains.NO_ACCESS_TO_PARTY)
	}

	columnNames := []string{"food_req_id"}
	values := []interface{}{requirementId}

	contributions, err := ds.ContributionRepository.FindAllBy(columnNames, values, "Contributor", "FoodReq")
	if err != nil {
		return api.ErrorInternalServerError(err.Error())
	}

	return api.Success(contributions)
}

func (ds FoodContributionService) GetByPartyId(partyId, userId uint) api.IResponse {
	party, err := ds.PartyRepository.FindById(partyId, partyDomains.FullPartyPreload...)
	if err != nil {
		return api.ErrorBadRequest(err.Error())
	}

	if !party.CanBeAccessedBy(userId) {
		return api.ErrorUnauthorized(domains.NO_ACCESS_TO_PARTY)
	}

	columnNames := []string{"party_id"}
	values := []interface{}{partyId}

	contributions, err := ds.ContributionRepository.FindAllBy(columnNames, values, "Contributor", "FoodReq")
	if err != nil {
		return api.ErrorInternalServerError(err.Error())
	}

	return api.Success(contributions)
}
