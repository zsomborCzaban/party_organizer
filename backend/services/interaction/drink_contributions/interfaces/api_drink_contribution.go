package interfaces

import (
	"github.com/gorilla/mux"
	"github.com/zsomborCzaban/party_organizer/common/jwt"
	"github.com/zsomborCzaban/party_organizer/services/interaction/drink_contributions/domains"
)

func NewDrinkContributionRouter(router *mux.Router, controller domains.IDrinkContributionController) {
	r := router.PathPrefix("/drinkContribution").Subrouter()

	r.Use(jwt.ValidateJWTMiddleware)

	r.HandleFunc("", controller.Create).Methods("POST")
	r.HandleFunc("/{id}", controller.Update).Methods("PUT")
	r.HandleFunc("/{id}", controller.Delete).Methods("DELETE")

	r.HandleFunc("/getByPartyAndContributor/{party_id}/{contributor_id}", controller.GetByPartyIdAndContributorId).Methods("GET")
	r.HandleFunc("/getByPartyAndRequirement/{party_id}/{requirement_id}", controller.GetByPartyIdAndRequirementId).Methods("GET")
}
