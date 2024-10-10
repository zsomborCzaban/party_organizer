package domains

import "net/http"

type IFoodRequirementController interface {
	CreateController(http.ResponseWriter, *http.Request)
	GetController(http.ResponseWriter, *http.Request)
	UpdateController(http.ResponseWriter, *http.Request)
	DeleteController(http.ResponseWriter, *http.Request)
	GetByPartyIdController(http.ResponseWriter, *http.Request)
}
