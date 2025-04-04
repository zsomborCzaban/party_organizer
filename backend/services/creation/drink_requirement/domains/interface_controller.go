package domains

import "net/http"

type IDrinkRequirementController interface {
	Create(http.ResponseWriter, *http.Request)
	Get(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
	GetByPartyId(http.ResponseWriter, *http.Request)
}
