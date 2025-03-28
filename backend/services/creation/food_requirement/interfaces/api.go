package interfaces

import (
	"github.com/gorilla/mux"
	"github.com/zsomborCzaban/party_organizer/services/creation/food_requirement/domains"
)

func NewRouter(router *mux.Router, controller domains.IFoodRequirementController) {
	router.HandleFunc("/foodRequirement", controller.Create).Methods("POST")
	router.HandleFunc("/foodRequirement/{id}", controller.Get).Methods("GET")
	router.HandleFunc("/foodRequirement/{id}", controller.Delete).Methods("DELETE")

	router.HandleFunc("/foodRequirement/getByPartyId/{party_id}", controller.GetByPartyId).Methods("GET")
}
