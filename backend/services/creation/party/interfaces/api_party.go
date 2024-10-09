package interfaces

import (
	"github.com/gorilla/mux"
	"github.com/zsomborCzaban/party_organizer/common/jwt"
	"github.com/zsomborCzaban/party_organizer/services/creation/party/domains"
)

func NewPartyRouter(router *mux.Router, controller domains.IPartyController) {
	r := router.PathPrefix("/party").Subrouter()

	r.Use(jwt.ValidateJWTMiddleware)

	r.HandleFunc("/", controller.CreateController).Methods("POST")
	r.HandleFunc("/{id}", controller.GetController).Methods("GET")
	r.HandleFunc("/{id}", controller.UpdateController).Methods("PUT")
	r.HandleFunc("/{id}", controller.DeleteController).Methods("DELETE")

	r.HandleFunc("/GetPartiesByOrganizerId/{id}", controller.GetPartiesByOrganizerId).Methods("GET")
	r.HandleFunc("/GetPartiesByParticipantId/{id}", controller.GetPartiesByParticipantId).Methods("GET")
	r.HandleFunc("/{id}/Join/", controller.AddUserToParty).Methods("GET")

}
