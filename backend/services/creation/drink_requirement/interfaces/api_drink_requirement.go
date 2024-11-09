package interfaces

import (
	"github.com/gorilla/mux"
	"github.com/zsomborCzaban/party_organizer/services/creation/drink_requirement/domains"
)

func NewDrinkRequirementRouter(router *mux.Router, controller domains.IDrinkRequirementController) {
	router.HandleFunc("/drinkRequirement", controller.CreateController).Methods("POST")
	router.HandleFunc("/drinkRequirement/{id}", controller.GetController).Methods("GET")
	router.HandleFunc("/drinkRequirement/{id}", controller.UpdateController).Methods("PUT")
	router.HandleFunc("/drinkRequirement/{id}", controller.DeleteController).Methods("DELETE")

	router.HandleFunc("/drinkRequirement/getByPartyId/{party_id}", controller.GetByPartyIdController).Methods("GET")
}
