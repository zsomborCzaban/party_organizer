package interfaces

import (
	"github.com/gorilla/mux"
	"github.com/zsomborCzaban/party_organizer/services/creation/party/domains"
)

func NewPartyRouter(router *mux.Router, controller domains.IPartyController) {
	router.HandleFunc("/party/getPartiesByOrganizerId", controller.GetPartiesByOrganizerId).Methods("GET")
	router.HandleFunc("/party/getPartiesByParticipantId", controller.GetPartiesByParticipantId).Methods("GET")
	router.HandleFunc("/party/getParticipants/{party_id}", controller.GetParticipants).Methods("GET")

	router.HandleFunc("/party", controller.Create).Methods("POST")
	router.HandleFunc("/party", controller.Update).Methods("PUT")
	router.HandleFunc("/party/{id}", controller.Get).Methods("GET")
	router.HandleFunc("/party/{id}", controller.Delete).Methods("DELETE")
}

func NewPublicPartyRouter(router *mux.Router, controller domains.IPartyController) {
	router.HandleFunc("/publicParties", controller.GetPublicParties).Methods("GET")
	router.HandleFunc("/publicParties/{id}", controller.GetPublicParty).Methods("GET")
}
