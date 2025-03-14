package domains

import "net/http"

type IFoodContributionController interface {
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)

	GetByPartyIdAndContributorId(w http.ResponseWriter, r *http.Request)
	GetByRequirementId(w http.ResponseWriter, r *http.Request)
	GetByPartyId(w http.ResponseWriter, r *http.Request)
}
