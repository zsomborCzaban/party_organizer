package domains

import "net/http"

type IPartyController interface {
	GetParticipants(w http.ResponseWriter, r *http.Request)
	GetPublicParties(http.ResponseWriter, *http.Request)
	GetPartiesByOrganizerId(http.ResponseWriter, *http.Request)
	GetPartiesByParticipantId(http.ResponseWriter, *http.Request)

	CreateController(http.ResponseWriter, *http.Request)
	GetController(http.ResponseWriter, *http.Request)
	UpdateController(http.ResponseWriter, *http.Request)
	DeleteController(http.ResponseWriter, *http.Request)
}
