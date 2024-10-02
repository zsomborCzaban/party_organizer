package interfaces

import (
	"github.com/gorilla/mux"
	"github.com/zsomborCzaban/party_organizer/services/creation/drink_requirement/domains"
)

func NewDrinkRequirementRouter(router *mux.Router, controller domains.IDrinkRequirementController) {
	r := router.PathPrefix("/drink_requirement").Subrouter()

	r.HandleFunc("/", controller.CreateController).Methods("POST")
	r.HandleFunc("/{id}", controller.GetController).Methods("GET")
	r.HandleFunc("/", controller.UpdateController).Methods("PUT")
	r.HandleFunc("/{id}", controller.DeleteController).Methods("DELETE")
}
