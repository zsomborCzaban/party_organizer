package interfaces

import (
	"github.com/gorilla/mux"
	"github.com/zsomborCzaban/party_organizer/common/jwt"
	"github.com/zsomborCzaban/party_organizer/services/creation/food_requirement/domains"
)

func NewFoodRequirementRouter(router *mux.Router, controller domains.IFoodRequirementController) {
	r := router.PathPrefix("/food_requirement").Subrouter()

	r.Use(jwt.ValidateJWTMiddleware)

	r.HandleFunc("/", controller.CreateController).Methods("POST")
	r.HandleFunc("/{id}", controller.GetController).Methods("GET")
	r.HandleFunc("/", controller.UpdateController).Methods("PUT")
	r.HandleFunc("/{id}", controller.DeleteController).Methods("DELETE")

	r.HandleFunc("/getByPartyId/{party_id}", controller.GetByPartyIdController).Methods("GET")
}
