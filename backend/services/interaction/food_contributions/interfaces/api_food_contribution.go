package interfaces

import (
	"github.com/gorilla/mux"
	"github.com/zsomborCzaban/party_organizer/common/jwt"
	"github.com/zsomborCzaban/party_organizer/services/interaction/food_contributions/domains"
)

func NewFoodContributionRouter(router *mux.Router, controller domains.IFoodContributionController) {
	r := router.PathPrefix("/foodContribution").Subrouter()

	r.Use(jwt.ValidateJWTMiddleware)

	r.HandleFunc("", controller.Create).Methods("POST")
	r.HandleFunc("/{id}", controller.Update).Methods("PUT")
	r.HandleFunc("/{id}", controller.Delete).Methods("DELETE")

	r.HandleFunc("/getByPartyAndContributor/{party_id}/{contributor_id}", controller.GetByPartyIdAndContributorId).Methods("GET")
	r.HandleFunc("/getByRequirement/{requirement_id}", controller.GetByRequirementId).Methods("GET")
	r.HandleFunc("/getByParty/{party_id}", controller.GetByPartyId).Methods("GET")
}
