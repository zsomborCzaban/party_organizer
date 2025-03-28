package domains

import "net/http"

type IPartyController interface {
	//public endpoints
	GetPublicParties(http.ResponseWriter, *http.Request)
	GetPublicParty(http.ResponseWriter, *http.Request)

	//private endpoints
	GetParticipants(w http.ResponseWriter, r *http.Request)
	GetPartiesByOrganizerId(http.ResponseWriter, *http.Request)
	GetPartiesByParticipantId(http.ResponseWriter, *http.Request)

	Create(http.ResponseWriter, *http.Request)
	Get(http.ResponseWriter, *http.Request)
	Update(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
}
