package interfaces

import "github.com/gorilla/mux"

func NewDrinkRequirementRouter(router *mux.Router) {
	r := router.PathPrefix("/drink_requirement").Subrouter()

	drinkRequirementController := NewDrinkRequirementController()

	r.HandleFunc("/", drinkRequirementController.CreateController).Methods("POST")
	r.HandleFunc("/{id}", drinkRequirementController.GetController).Methods("GET")
	r.HandleFunc("/", drinkRequirementController.UpdateController).Methods("PUT")
	r.HandleFunc("/{id}", drinkRequirementController.DeleteController).Methods("DELETE")
}
