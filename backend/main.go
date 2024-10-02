package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"github.com/zsomborCzaban/party_organizer/db"
	drinkRequirementDomains "github.com/zsomborCzaban/party_organizer/services/creation/drink_requirement/domains"
	drinkRequirementInterfaces "github.com/zsomborCzaban/party_organizer/services/creation/drink_requirement/interfaces"
	drinkRequirementUsecases "github.com/zsomborCzaban/party_organizer/services/creation/drink_requirement/usecases"
	partyDomains "github.com/zsomborCzaban/party_organizer/services/creation/party/domains"
	partyInterfaces "github.com/zsomborCzaban/party_organizer/services/creation/party/interfaces"
	partyUsecases "github.com/zsomborCzaban/party_organizer/services/creation/party/usecases"
	"net/http"
)

func main() {
	dbAccess := db.CreateGormDatabaseAccessManager("local.db")

	router := mux.NewRouter()

	apiRouter := router.PathPrefix("/api/v0").Subrouter()

	partyRepository := partyUsecases.NewPartyRepository(dbAccess)
	partyValidator := partyDomains.NewValidator(validator.New())
	partyService := partyInterfaces.NewPartyService(partyRepository, partyValidator)
	partyController := partyInterfaces.NewPartyController(partyService)

	drinkRequirementRepository := drinkRequirementUsecases.NewDrinkRequirementRepository(dbAccess)
	drinkRequirementValidator := drinkRequirementDomains.NewValidator(validator.New())
	drinkRequirementService := drinkRequirementInterfaces.NewDrinkRequirementService(drinkRequirementRepository, drinkRequirementValidator)
	drinkRequirementController := drinkRequirementInterfaces.NewDrinkRequirementController(drinkRequirementService)

	partyInterfaces.NewPartyRouter(apiRouter, partyController)
	drinkRequirementInterfaces.NewDrinkRequirementRouter(apiRouter, drinkRequirementController)

	log.Fatal().Err(http.ListenAndServe(":8080", router))
}
