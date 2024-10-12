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
	friendInvitationInterfaces "github.com/zsomborCzaban/party_organizer/services/invitation/friend_invite/interfaces"
	friendInvitationUsecases "github.com/zsomborCzaban/party_organizer/services/invitation/friend_invite/usecases"
	partyInvitationInterfaces "github.com/zsomborCzaban/party_organizer/services/invitation/party_invite/interfaces"
	partyInvitationUsecases "github.com/zsomborCzaban/party_organizer/services/invitation/party_invite/usecases"
	userInterfaces "github.com/zsomborCzaban/party_organizer/services/user/interfaces"
	userUsecases "github.com/zsomborCzaban/party_organizer/services/user/usecases"
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

	vali := api.NewValidator(validator.New())

	userRepository := userUsecases.NewUserRepository(dbAccess)
	userService := userInterfaces.NewUserService(userRepository, vali)
	userController := userInterfaces.NewUserController(userService)

	partyRepository := partyUsecases.NewPartyRepository(dbAccess)
	partyService := partyInterfaces.NewPartyService(partyRepository, vali, userRepository)
	partyController := partyInterfaces.NewPartyController(partyService)

	drinkRequirementRepository := drinkRequirementUsecases.NewDrinkRequirementRepository(dbAccess)
	drinkRequirementService := drinkRequirementInterfaces.NewDrinkRequirementService(drinkRequirementRepository, vali, partyRepository)
	drinkRequirementController := drinkRequirementInterfaces.NewDrinkRequirementController(drinkRequirementService)

	foodRequirementRepository := foodRequirementUsecases.NewFoodRequirementRepository(dbAccess)
	foodRequirementService := foodRequirementInterfaces.NewFoodRequirementService(foodRequirementRepository, vali, partyRepository)
	foodRequirementController := foodRequirementInterfaces.NewFoodRequirementController(foodRequirementService)

	friendInviteRepository := friendInvitationUsecases.NewFriendInviteRepository(dbAccess)
	friendInviteService := friendInvitationInterfaces.NewFriendInviteService(friendInviteRepository, userRepository)
	friendInviteController := friendInvitationInterfaces.NewFriendInviteController(friendInviteService)

	partyInviteRepository := partyInvitationUsecases.NewPartyInviteRepository(dbAccess)
	partyInviteService := partyInvitationInterfaces.NewPartyInviteService(partyInviteRepository, userRepository, partyRepository)
	partyInviteController := partyInvitationInterfaces.NewPartyInviteController(partyInviteService)

	userInterfaces.NewUserRouter(apiRouter, userController)
	partyInterfaces.NewPartyRouter(apiRouter, partyController)
	drinkRequirementInterfaces.NewDrinkRequirementRouter(apiRouter, drinkRequirementController)
	foodRequirementInterfaces.NewFoodRequirementRouter(apiRouter, foodRequirementController)
	friendInvitationInterfaces.NewFriendInviteRouter(apiRouter, friendInviteController)
	partyInvitationInterfaces.NewPartyInviteRouter(apiRouter, partyInviteController)

	log.Fatal().Err(http.ListenAndServe(":8080", router))
}
