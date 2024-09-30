package interfaces

import "github.com/gorilla/mux"

func NewPartyRouter(router *mux.Router) {
	r := router.PathPrefix("/party").Subrouter()

	partyController := NewPartyController()

	r.HandleFunc("/", partyController.CreateController).Methods("POST")
	r.HandleFunc("/{id}", partyController.GetController).Methods("GET")
	r.HandleFunc("/", partyController.UpdateController).Methods("PUT")
	r.HandleFunc("/{id}", partyController.DeleteController).Methods("DELETE")
}
