package main

import (
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	drinkRequirementInterfaces "github.com/zsomborCzaban/party_organizer/services/creation/drink_requirement/interfaces"
	partyInterfaces "github.com/zsomborCzaban/party_organizer/services/creation/party/interfaces"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	apiRouter := router.PathPrefix("/api/v0").Subrouter()

	partyInterfaces.NewPartyRouter(apiRouter)
	drinkRequirementInterfaces.NewDrinkRequirementRouter(apiRouter)

	log.Fatal().Err(http.ListenAndServe(":8080", router))
}
