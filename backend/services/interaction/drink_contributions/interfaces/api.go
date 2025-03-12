package interfaces

import (
	"github.com/gorilla/mux"
	"github.com/zsomborCzaban/party_organizer/services/interaction/drink_contributions/domains"
)

func NewDrinkContributionRouter(router *mux.Router, controller domains.IDrinkContributionController) {
	router.HandleFunc("/drinkContribution", controller.Create).Methods("POST")
	router.HandleFunc("/drinkContribution/{id}", controller.Update).Methods("PUT")
	router.HandleFunc("/drinkContribution/{id}", controller.Delete).Methods("DELETE")

	router.HandleFunc("/drinkContribution/getByPartyAndContributor/{party_id}/{contributor_id}", controller.GetByPartyIdAndContributorId).Methods("GET")
	router.HandleFunc("/drinkContribution/getByRequirement/{requirement_id}", controller.GetByRequirementId).Methods("GET")
	router.HandleFunc("/drinkContribution/getByParty/{party_id}", controller.GetByPartyId).Methods("GET")
}
