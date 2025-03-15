package interfaces

import (
	"github.com/gorilla/mux"
	"github.com/zsomborCzaban/party_organizer/services/creation/food_requirement/domains"
)

func NewRouter(router *mux.Router, controller domains.IFoodRequirementController) {
	router.HandleFunc("/foodRequirement", controller.CreateController).Methods("POST")
	router.HandleFunc("/foodRequirement/{id}", controller.GetController).Methods("GET")
	router.HandleFunc("/foodRequirement/{id}", controller.DeleteController).Methods("DELETE")

	router.HandleFunc("/foodRequirement/getByPartyId/{party_id}", controller.GetByPartyIdController).Methods("GET")
}
