package interfaces

import (
	"github.com/gorilla/mux"
	"github.com/zsomborCzaban/party_organizer/services/creation/party/domains"
)

func NewPartyRouter(router *mux.Router, controller domains.IPartyController) {
	router.HandleFunc("/party/getPartiesByOrganizerId", controller.GetPartiesByOrganizerId).Methods("GET")
	router.HandleFunc("/party/getPartiesByParticipantId", controller.GetPartiesByParticipantId).Methods("GET")
	router.HandleFunc("/party/getPublicParties", controller.GetPublicParties).Methods("GET")
	router.HandleFunc("/party/getParticipants/{party_id}", controller.GetParticipants).Methods("GET")

	router.HandleFunc("/party", controller.CreateController).Methods("POST")
	router.HandleFunc("/party/{id}", controller.GetController).Methods("GET")
	router.HandleFunc("/party/{id}", controller.UpdateController).Methods("PUT")
	router.HandleFunc("/party/{id}", controller.DeleteController).Methods("DELETE")

}
