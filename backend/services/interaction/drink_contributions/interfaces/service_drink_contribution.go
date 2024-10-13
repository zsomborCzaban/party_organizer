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

	party, err3 := ds.PartyRepository.GetParty(contribution.PartyId)
	if err3 != nil {
		return api.ErrorBadRequest(err3.Error())
	}

	drinkReq, err4 := ds.DrinkReqRepository.GetDrinkRequirement(contribution.DrinkReqId)
	if err4 != nil {
		return api.ErrorBadRequest(err4.Error())
	}

	//todo: check if the user is in the party
	if drinkReq.PartyID != contribution.PartyId {
		return api.ErrorBadRequest("drink requirement doesnt belong to party")
	}

	contribution.ContributorId = userId
	contribution.Contributor = *contributor
	contribution.Party = *party
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

	contributor, err3 := ds.UserRepository.FindById(userId)
	if err3 != nil {
		return api.ErrorBadRequest(err3.Error())
	}

	party, err4 := ds.PartyRepository.GetParty(contribution.PartyId)
	if err4 != nil {
		return api.ErrorBadRequest(err4.Error())
	}

	drinkReq, err5 := ds.DrinkReqRepository.GetDrinkRequirement(contribution.DrinkReqId)
	if err5 != nil {
		return api.ErrorBadRequest(err5.Error())
	}

	//todo: check if the user is in the party
	if oldContribution.ContributorId != userId && userId != adminUser.ADMIN_USER_ID {
		return api.ErrorBadRequest("cannot update other people's contribution")
	}

	if drinkReq.PartyID != contribution.PartyId {
		return api.ErrorBadRequest("drink requirement doesnt belong to party")
	}

	contribution.ContributorId = userId
	contribution.Contributor = *contributor
	contribution.Party = *party
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

	if contribution.ContributorId != userId && userId != adminUser.ADMIN_USER_ID {
		return api.ErrorUnauthorized("cannot delete other people's contribution")
	}

	if err2 := ds.ContributionRepository.Delete(contribution); err2 != nil {
		return api.ErrorInternalServerError(err2.Error())
	}

	return api.Success("contribution deleted successfully")
}

func (ds DrinkContributionService) GetByPartyIdAndContributorId(partyId, contributorId, userId uint) api.IResponse {
	//todo: check if user in party

	columnNames := []string{"party_id", "contributor_id"}
	values := []interface{}{partyId, contributorId}

	contributions, err := ds.ContributionRepository.FindAllBy(columnNames, values)
	if err != nil {
		return api.ErrorInternalServerError(err.Error())
	}

	return api.Success(contributions)
}

func (ds DrinkContributionService) GetByPartyIdAndRequirementId(partyId, requirementId, userId uint) api.IResponse {
	//todo: check if user in party

	columnNames := []string{"party_id", "requirement_id"}
	values := []interface{}{partyId, requirementId}

	contributions, err := ds.ContributionRepository.FindAllBy(columnNames, values)
	if err != nil {
		return api.ErrorInternalServerError(err.Error())
	}

	return api.Success(contributions)
}
