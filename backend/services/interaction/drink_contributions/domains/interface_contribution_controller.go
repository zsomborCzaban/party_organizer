package domains

import "net/http"

type IDrinkContributionController interface {
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)

	GetByPartyIdAndContributorId(w http.ResponseWriter, r *http.Request)
	GetByPartyIdAndRequirementId(w http.ResponseWriter, r *http.Request)
}
