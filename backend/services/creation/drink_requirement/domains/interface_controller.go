package domains

import "net/http"

type IDrinkRequirementController interface {
	CreateController(http.ResponseWriter, *http.Request)
	GetController(http.ResponseWriter, *http.Request)
	DeleteController(http.ResponseWriter, *http.Request)
	GetByPartyIdController(http.ResponseWriter, *http.Request)
}
