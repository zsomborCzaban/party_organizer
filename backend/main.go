package main

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/rs/zerolog/log"
	"github.com/zsomborCzaban/party_organizer/common/api"
	"github.com/zsomborCzaban/party_organizer/db"
	drinkRequirementInterfaces "github.com/zsomborCzaban/party_organizer/services/creation/drink_requirement/interfaces"
	drinkRequirementUsecases "github.com/zsomborCzaban/party_organizer/services/creation/drink_requirement/usecases"
	foodRequirementInterfaces "github.com/zsomborCzaban/party_organizer/services/creation/food_requirement/interfaces"
	foodRequirementUsecases "github.com/zsomborCzaban/party_organizer/services/creation/food_requirement/usecases"
	partyInterfaces "github.com/zsomborCzaban/party_organizer/services/creation/party/interfaces"
	partyUsecases "github.com/zsomborCzaban/party_organizer/services/creation/party/usecases"
	drinkContributionInterfaces "github.com/zsomborCzaban/party_organizer/services/interaction/drink_contributions/interfaces"
	drinkContributionUsecases "github.com/zsomborCzaban/party_organizer/services/interaction/drink_contributions/usecases"
	foodContributionInterfaces "github.com/zsomborCzaban/party_organizer/services/interaction/food_contributions/interfaces"
	foodContributionUsecases "github.com/zsomborCzaban/party_organizer/services/interaction/food_contributions/usecases"
	friendInvitationInterfaces "github.com/zsomborCzaban/party_organizer/services/managers/friend_manager/interfaces"
	friendInvitationUsecases "github.com/zsomborCzaban/party_organizer/services/managers/friend_manager/usecases"
	partyInvitationInterfaces "github.com/zsomborCzaban/party_organizer/services/managers/party_attendance_manager/interfaces"
	partyInvitationUsecases "github.com/zsomborCzaban/party_organizer/services/managers/party_attendance_manager/usecases"
	userInterfaces "github.com/zsomborCzaban/party_organizer/services/user/interfaces"
	userUsecases "github.com/zsomborCzaban/party_organizer/services/user/usecases"
	"gorm.io/gorm/logger"
	log2 "log"
	"net/http"
	"os"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion("us-east-1"))
	if err != nil {
		log.Fatal().Msg("unable to load SDK config, " + err.Error())
	}

	s3Client := s3.NewFromConfig(cfg)

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

	dbAccess := db.CreateGormDatabaseAccessManager("application.db", newLogger)

	router := mux.NewRouter()

	apiRouter := router.PathPrefix("/api/v1").Subrouter()

	vali := api.NewValidator(validator.New())

	userRepository := userUsecases.NewUserRepository(dbAccess)
	partyRepository := partyUsecases.NewPartyRepository(dbAccess)
	drinkRequirementRepository := drinkRequirementUsecases.NewDrinkRequirementRepository(dbAccess)
	foodRequirementRepository := foodRequirementUsecases.NewFoodRequirementRepository(dbAccess)
	friendInviteRepository := friendInvitationUsecases.NewFriendInviteRepository(dbAccess)
	partyInviteRepository := partyInvitationUsecases.NewPartyInviteRepository(dbAccess)
	drinkContributionRepository := drinkContributionUsecases.NewDrinkContributionRepository(dbAccess)
	foodContributionRepository := foodContributionUsecases.NewFoodContributionRepository(dbAccess)

	userService := userInterfaces.NewUserService(userRepository, vali, s3Client)
	partyService := partyInterfaces.NewPartyService(partyRepository, vali, userRepository)
	drinkRequirementService := drinkRequirementInterfaces.NewDrinkRequirementService(drinkRequirementRepository, vali, partyRepository, drinkContributionRepository)
	foodRequirementService := foodRequirementInterfaces.NewFoodRequirementService(foodRequirementRepository, vali, partyRepository, foodContributionRepository)
	friendInviteService := friendInvitationInterfaces.NewFriendInviteService(friendInviteRepository, userRepository)
	partyInviteService := partyInvitationInterfaces.NewPartyInviteService(partyInviteRepository, userRepository, partyRepository, foodContributionRepository, drinkContributionRepository)
	drinkContributionService := drinkContributionInterfaces.NewDrinkContributionService(drinkContributionRepository, vali, userRepository, partyRepository, drinkRequirementRepository)
	foodContributionService := foodContributionInterfaces.NewFoodContributionService(foodContributionRepository, vali, userRepository, partyRepository, foodRequirementRepository)

	userController := userInterfaces.NewUserController(userService)
	partyController := partyInterfaces.NewPartyController(partyService)
	drinkRequirementController := drinkRequirementInterfaces.NewDrinkRequirementController(drinkRequirementService)
	foodRequirementController := foodRequirementInterfaces.NewFoodRequirementController(foodRequirementService)
	friendInviteController := friendInvitationInterfaces.NewFriendInviteController(friendInviteService)
	partyInviteController := partyInvitationInterfaces.NewPartyInviteController(partyInviteService)
	drinkContributionController := drinkContributionInterfaces.NewDrinkContributionController(drinkContributionService)
	foodContributionController := foodContributionInterfaces.NewFoodContributionController(foodContributionService)

	userInterfaces.NewUserRouter(apiRouter, userController)
	partyInterfaces.NewPartyRouter(apiRouter, partyController)
	drinkRequirementInterfaces.NewDrinkRequirementRouter(apiRouter, drinkRequirementController)
	foodRequirementInterfaces.NewFoodRequirementRouter(apiRouter, foodRequirementController)
	friendInvitationInterfaces.NewFriendInviteRouter(apiRouter, friendInviteController)
	partyInvitationInterfaces.NewPartyInviteRouter(apiRouter, partyInviteController)
	drinkContributionInterfaces.NewDrinkContributionRouter(apiRouter, drinkContributionController)
	foodContributionInterfaces.NewFoodContributionRouter(apiRouter, foodContributionController)

	corsOptions := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "OPTIONS", "DELETE"},
		AllowedHeaders:   []string{"Authorization", "Accept", "Content-Type"},
		AllowCredentials: true,
	})

	handler := corsOptions.Handler(router)

	log.Fatal().Err(http.ListenAndServe(":8080", handler))
}
