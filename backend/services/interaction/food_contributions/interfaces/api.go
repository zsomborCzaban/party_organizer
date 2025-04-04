package interfaces

import (
	"github.com/gorilla/mux"
	"github.com/zsomborCzaban/party_organizer/services/interaction/food_contributions/domains"
)

func NewRouter(router *mux.Router, controller domains.IFoodContributionController) {
	router.HandleFunc("/foodContribution", controller.Create).Methods("POST")
	router.HandleFunc("/foodContribution/{id}", controller.Update).Methods("PUT")
	router.HandleFunc("/foodContribution/{id}", controller.Delete).Methods("DELETE")

	router.HandleFunc("/foodContribution/getByPartyAndContributor/{party_id}/{contributor_id}", controller.GetByPartyIdAndContributorId).Methods("GET")
	router.HandleFunc("/foodContribution/getByRequirement/{requirement_id}", controller.GetByRequirementId).Methods("GET")
	router.HandleFunc("/foodContribution/getByParty/{party_id}", controller.GetByPartyId).Methods("GET")
}
