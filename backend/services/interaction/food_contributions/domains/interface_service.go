package domains

import (
	"github.com/zsomborCzaban/party_organizer/utils/api"
)

type IFoodContributionService interface {
	Create(contribution FoodContribution, userId uint) api.IResponse
	Update(contribution FoodContribution, userId uint) api.IResponse
	Delete(contributionId, userId uint) api.IResponse

	GetByPartyIdAndContributorId(partyId, contributorId, userId uint) api.IResponse
	GetByRequirementId(requirementId, userId uint) api.IResponse
	GetByPartyId(partyId, userId uint) api.IResponse
}
