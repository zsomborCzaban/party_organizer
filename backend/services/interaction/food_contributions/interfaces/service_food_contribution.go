package interfaces

import (
	"github.com/zsomborCzaban/party_organizer/common/adminUser"
	"github.com/zsomborCzaban/party_organizer/common/api"
	foodReqDomain "github.com/zsomborCzaban/party_organizer/services/creation/food_requirement/domains"
	partyDomains "github.com/zsomborCzaban/party_organizer/services/creation/party/domains"
	"github.com/zsomborCzaban/party_organizer/services/interaction/food_contributions/domains"
	userDomain "github.com/zsomborCzaban/party_organizer/services/user/domains"
)

type FoodContributionService struct {
	Validator              api.IValidator
	ContributionRepository domains.IFoodContributionRepository
	UserRepository         userDomain.IUserRepository
	PartyRepository        partyDomains.IPartyRepository
	FoodReqRepository      foodReqDomain.IFoodRequirementRepository
}

func NewFoodContributionService(contributionRepo domains.IFoodContributionRepository, vali api.IValidator, userRepo userDomain.IUserRepository, partyRepo partyDomains.IPartyRepository, foodReqRepo foodReqDomain.IFoodRequirementRepository) domains.IFoodContributionService {
	return &FoodContributionService{
		Validator:              vali,
		ContributionRepository: contributionRepo,
		UserRepository:         userRepo,
		PartyRepository:        partyRepo,
		FoodReqRepository:      foodReqRepo,
	}
}

func (ds FoodContributionService) Create(contribution domains.FoodContribution, userId uint) api.IResponse {
	err := ds.Validator.Validate(contribution)
	if err != nil {
		return api.ErrorValidation(err.Errors)
	}

	contributor, err2 := ds.UserRepository.FindById(userId)
	if err2 != nil {
		return api.ErrorBadRequest(err2.Error())
	}

	party, err3 := ds.PartyRepository.GetParty(contribution.PartyId)
	if err3 != nil {
		return api.ErrorBadRequest(err3.Error())
	}

	foodReq, err4 := ds.FoodReqRepository.GetFoodRequirement(contribution.FoodReqId)
	if err4 != nil {
		return api.ErrorBadRequest(err4.Error())
	}

	//todo: check if the user is in the party
	if foodReq.PartyID != contribution.PartyId {
		return api.ErrorBadRequest("food requirement doesnt belong to party")
	}

	contribution.ContributorId = userId
	contribution.Contributor = *contributor
	contribution.Party = *party
	contribution.FoodReq = *foodReq

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

	oldContribution, err2 := ds.ContributionRepository.FindById(contribution.ID)
	if err2 != nil {
		return api.ErrorBadRequest(err2.Error())
	}

	contributor, err3 := ds.UserRepository.FindById(userId)
	if err3 != nil {
		return api.ErrorBadRequest(err3.Error())
	}

	party, err4 := ds.PartyRepository.GetParty(contribution.PartyId)
	if err4 != nil {
		return api.ErrorBadRequest(err4.Error())
	}

	foodReq, err5 := ds.FoodReqRepository.GetFoodRequirement(contribution.FoodReqId)
	if err5 != nil {
		return api.ErrorBadRequest(err5.Error())
	}

	//todo: check if the user is in the party
	if oldContribution.ContributorId != userId && userId != adminUser.ADMIN_USER_ID {
		return api.ErrorBadRequest("cannot update other people's contribution")
	}

	if foodReq.PartyID != contribution.PartyId {
		return api.ErrorBadRequest("food requirement doesnt belong to party")
	}

	contribution.ContributorId = userId
	contribution.Contributor = *contributor
	contribution.Party = *party
	contribution.FoodReq = *foodReq

	if err6 := ds.ContributionRepository.Create(&contribution); err6 != nil {
		return api.ErrorInternalServerError(err6.Error())
	}

	return api.Success(contribution)
}

func (ds FoodContributionService) Delete(contributionId, userId uint) api.IResponse {
	contribution, err := ds.ContributionRepository.FindById(contributionId)
	if err != nil {
		return api.ErrorBadRequest(err.Error())
	}

	if contribution.ContributorId != userId && userId != adminUser.ADMIN_USER_ID {
		return api.ErrorUnauthorized("cannot delete other people's contribution")
	}

	if err2 := ds.ContributionRepository.Delete(contribution); err2 != nil {
		return api.ErrorInternalServerError(err2.Error())
	}

	return api.Success("contribution deleted successfully")
}

func (ds FoodContributionService) GetByPartyIdAndContributorId(partyId, contributorId, userId uint) api.IResponse {
	//todo: check if user in party

	columnNames := []string{"party_id", "contributor_id"}
	values := []interface{}{partyId, contributorId}

	contributions, err := ds.ContributionRepository.FindAllBy(columnNames, values)
	if err != nil {
		return api.ErrorInternalServerError(err.Error())
	}

	return api.Success(contributions)
}

func (ds FoodContributionService) GetByRequirementId(requirementId, userId uint) api.IResponse {
	//todo: check if user in party

	columnNames := []string{"food_req_id"}
	values := []interface{}{requirementId}

	contributions, err := ds.ContributionRepository.FindAllBy(columnNames, values)
	if err != nil {
		return api.ErrorInternalServerError(err.Error())
	}

	return api.Success(contributions)
}

func (ds FoodContributionService) GetByPartyId(partyId, userId uint) api.IResponse {
	//todo: check if user in party

	columnNames := []string{"party_id"}
	values := []interface{}{partyId}

	contributions, err := ds.ContributionRepository.FindAllBy(columnNames, values)
	if err != nil {
		return api.ErrorInternalServerError(err.Error())
	}

	return api.Success(contributions)
}
