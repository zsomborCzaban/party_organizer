package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"github.com/zsomborCzaban/party_organizer/common/api"
	"github.com/zsomborCzaban/party_organizer/db"
	drinkRequirementInterfaces "github.com/zsomborCzaban/party_organizer/services/creation/drink_requirement/interfaces"
	drinkRequirementUsecases "github.com/zsomborCzaban/party_organizer/services/creation/drink_requirement/usecases"
	foodRequirementInterfaces "github.com/zsomborCzaban/party_organizer/services/creation/food_requirement/interfaces"
	foodRequirementUsecases "github.com/zsomborCzaban/party_organizer/services/creation/food_requirement/usecases"
	partyInterfaces "github.com/zsomborCzaban/party_organizer/services/creation/party/interfaces"
	partyUsecases "github.com/zsomborCzaban/party_organizer/services/creation/party/usecases"
	interfaces2 "github.com/zsomborCzaban/party_organizer/services/user/user/interfaces"
	"github.com/zsomborCzaban/party_organizer/services/user/user/usecases"
	"gorm.io/gorm/logger"
	log2 "log"
	"net/http"
	"os"
	"time"
)

func main() {
	newLogger := logger.New(
		log2.New(os.Stdout, "\r\n", log2.LstdFlags), //io writer
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: false,
			ParameterizedQueries:      true,
			Colorful:                  true,
		},
	)

	dbAccess := db.CreateGormDatabaseAccessManager("local.db", newLogger)

	router := mux.NewRouter()

	apiRouter := router.PathPrefix("/api/v0").Subrouter()

	userRepository := usecases.NewUserRepository(dbAccess)
	userValidator := api.NewValidator(validator.New())
	userService := interfaces2.NewUserService(userRepository, userValidator)
	userController := interfaces2.NewUserController(userService)

	partyRepository := partyUsecases.NewPartyRepository(dbAccess, userRepository)
	partyValidator := api.NewValidator(validator.New())
	partyService := partyInterfaces.NewPartyService(partyRepository, partyValidator)
	partyController := partyInterfaces.NewPartyController(partyService)

	drinkRequirementRepository := drinkRequirementUsecases.NewDrinkRequirementRepository(dbAccess)
	drinkRequirementValidator := api.NewValidator(validator.New())
	drinkRequirementService := drinkRequirementInterfaces.NewDrinkRequirementService(drinkRequirementRepository, drinkRequirementValidator, partyRepository)
	drinkRequirementController := drinkRequirementInterfaces.NewDrinkRequirementController(drinkRequirementService)

	foodRequirementRepository := foodRequirementUsecases.NewFoodRequirementRepository(dbAccess)
	foodRequirementValidator := api.NewValidator(validator.New())
	foodRequirementService := foodRequirementInterfaces.NewFoodRequirementService(foodRequirementRepository, foodRequirementValidator, partyRepository)
	foodRequirementController := foodRequirementInterfaces.NewFoodRequirementController(foodRequirementService)

	partyInterfaces.NewPartyRouter(apiRouter, partyController)
	drinkRequirementInterfaces.NewDrinkRequirementRouter(apiRouter, drinkRequirementController)
	foodRequirementInterfaces.NewFoodRequirementRouter(apiRouter, foodRequirementController)
	interfaces2.NewUserRouter(apiRouter, userController)

	log.Fatal().Err(http.ListenAndServe(":8080", router))
}
