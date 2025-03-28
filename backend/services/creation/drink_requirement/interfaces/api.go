package interfaces

import (
	"github.com/gorilla/mux"
	"github.com/zsomborCzaban/party_organizer/services/creation/drink_requirement/domains"
)

func NewRouter(router *mux.Router, controller domains.IDrinkRequirementController) {
	router.HandleFunc("/drinkRequirement", controller.Create).Methods("POST")
	router.HandleFunc("/drinkRequirement/{id}", controller.Get).Methods("GET")
	router.HandleFunc("/drinkRequirement/{id}", controller.Delete).Methods("DELETE")

	router.HandleFunc("/drinkRequirement/getByPartyId/{party_id}", controller.GetByPartyId).Methods("GET")
}
