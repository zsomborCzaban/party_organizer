package main

import (
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"github.com/zsomborCzaban/party_organizer/services/creation/party/interfaces"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	apiRouter := router.PathPrefix("/api/v0").Subrouter()

	interfaces.NewPartyRouter(apiRouter)

	log.Fatal().Err(http.ListenAndServe(":8080", router))
}
