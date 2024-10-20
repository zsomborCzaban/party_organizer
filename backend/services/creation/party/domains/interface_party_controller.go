package domains

import "net/http"

type IPartyController interface {
	GetPublicParties(http.ResponseWriter, *http.Request)
	GetPartiesByOrganizerId(http.ResponseWriter, *http.Request)
	GetPartiesByParticipantId(http.ResponseWriter, *http.Request)
	AddUserToParty(http.ResponseWriter, *http.Request) //probably wont be there

	CreateController(http.ResponseWriter, *http.Request)
	GetController(http.ResponseWriter, *http.Request)
	UpdateController(http.ResponseWriter, *http.Request)
	DeleteController(http.ResponseWriter, *http.Request)
}
