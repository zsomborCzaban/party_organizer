package configuration

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/rs/zerolog/log"
	"github.com/zsomborCzaban/party_organizer/common"
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
	"github.com/zsomborCzaban/party_organizer/services/users/user/interfaces"
	"github.com/zsomborCzaban/party_organizer/services/users/user/usecases"
	"github.com/zsomborCzaban/party_organizer/utils/api"
	"github.com/zsomborCzaban/party_organizer/utils/jwt"
	"github.com/zsomborCzaban/party_organizer/utils/repo"
	"net/http"
	"os"
	"time"
)

func SetupRoutes(router *mux.Router, dbAccessManager db.IDatabaseAccessManager) *mux.Router {
	apiRouter := router.PathPrefix("/api/v1").Subrouter()
	apiRouter.Use(jwt.ValidateJWTMiddleware)

	vali := api.NewValidator(validator.New())

	partyRepo := partyUsecases.NewPartyRepository(dbAccessManager)
	userRepo := usecases.NewUserRepository(dbAccessManager)
	drinkRequirementRepo := drinkRequirementUsecases.NewDrinkRequirementRepository(dbAccessManager)
	drinkContributionRepo := drinkContributionUsecases.NewDrinkContributionRepository(dbAccessManager)
	foodRequirementRepo := foodRequirementUsecases.NewFoodRequirementRepository(dbAccessManager)
	foodContributionRepo := foodContributionUsecases.NewFoodContributionRepository(dbAccessManager)
	partyInviteRepo := partyInvitationUsecases.NewPartyInviteRepository(dbAccessManager)
	friendInviteRepo := friendInvitationUsecases.NewFriendInviteRepository(dbAccessManager)

	repoCollector := repo.NewRepoCollector(
		&partyRepo,
		&userRepo,
		&drinkRequirementRepo,
		&drinkContributionRepo,
		&foodRequirementRepo,
		&foodContributionRepo,
		&partyInviteRepo,
		&friendInviteRepo,
	)

	partyService := partyUsecases.NewPartyService(repoCollector, vali)
	userService := usecases.NewUserService(repoCollector, vali, GetAwsS3Client())
	drinkRequirementService := drinkRequirementUsecases.NewDrinkRequirementService(repoCollector, vali)
	foodRequirementService := foodRequirementUsecases.NewFoodRequirementService(repoCollector, vali)
	friendInviteService := friendInvitationUsecases.NewFriendInviteService(repoCollector)
	partyInviteService := partyInvitationUsecases.NewPartyInviteService(repoCollector)
	drinkContributionService := drinkContributionUsecases.NewDrinkContributionService(repoCollector, vali)
	foodContributionService := foodContributionUsecases.NewFoodContributionService(repoCollector, vali)

	partyController := partyInterfaces.NewController(partyService)
	userController := interfaces.NewUserController(userService)
	drinkRequirementController := drinkRequirementInterfaces.NewDrinkRequirementController(drinkRequirementService)
	foodRequirementController := foodRequirementInterfaces.NewFoodRequirementController(foodRequirementService)
	friendInviteController := friendInvitationInterfaces.NewFriendInviteController(friendInviteService)
	partyInviteController := partyInvitationInterfaces.NewPartyInviteController(partyInviteService)
	drinkContributionController := drinkContributionInterfaces.NewDrinkContributionController(drinkContributionService)
	foodContributionController := foodContributionInterfaces.NewFoodContributionController(foodContributionService)

	partyInterfaces.NewPublicPartyRouter(router.PathPrefix("/api/v1").Subrouter(), partyController)
	partyInterfaces.NewPartyRouter(apiRouter, partyController)
	interfaces.NewUserAuthRouter(router.PathPrefix("/api/v1").Subrouter(), userController)
	interfaces.NewUserPrivateRouter(apiRouter, userController)
	drinkRequirementInterfaces.NewRouter(apiRouter, drinkRequirementController)
	foodRequirementInterfaces.NewRouter(apiRouter, foodRequirementController)
	friendInvitationInterfaces.NewRouter(apiRouter, friendInviteController)
	partyInvitationInterfaces.NewRouter(apiRouter, partyInviteController)
	drinkContributionInterfaces.NewRouter(apiRouter, drinkContributionController)
	foodContributionInterfaces.NewRouter(apiRouter, foodContributionController)

	return router
}

func AddCorsSettings(router *mux.Router) http.Handler {
	corsOptions := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "OPTIONS", "DELETE"},
		AllowedHeaders:   []string{"Authorization", "Accept", "Content-Type"},
		AllowCredentials: true,
	})

	return corsOptions.Handler(router)
}

func GetAwsS3Client() *s3.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(os.Getenv(common.AWS_REGION_ENV_KEY)))
	if err != nil {
		log.Fatal().Msg("unable to load SDK config, " + err.Error())
	}

	return s3.NewFromConfig(cfg)
}

func DefaultApplicationSetup() http.Handler {
	LoadEnvVariables()

	return AddCorsSettings(
		SetupRoutes(mux.NewRouter(), CreateDbAccessManager(db.CreateGormDatabaseAccessManager)))
}

//func InitRepoCollector(dbAccess db.IDatabaseAccessManager) *repo.RepoCollector {
//	partyRepo := partyUsecases.NewPartyRepository(dbAccess)
//	userRepo := userUsecases.NewUserRepository(dbAccess)
//	drinkRequirementRepo := drinkRequirementUsecases.NewDrinkRequirementRepository(dbAccess)
//	drinkContributionRepo := drinkContributionUsecases.NewDrinkContributionRepository(dbAccess)
//	foodRequirementRepo := foodRequirementUsecases.NewFoodRequirementRepository(dbAccess)
//	foodContributionRepo := foodContributionUsecases.NewFoodContributionRepository(dbAccess)
//	partyInviteRepo := partyInvitationUsecases.NewPartyInviteRepository(dbAccess)
//	friendInviteRepo := friendInvitationUsecases.NewFriendInviteRepository(dbAccess)
//
//	return repo.NewRepoCollector(
//		&partyRepo,
//		&userRepo,
//		&drinkRequirementRepo,
//		&drinkContributionRepo,
//		&foodRequirementRepo,
//		&foodContributionRepo,
//		&partyInviteRepo,
//		&friendInviteRepo,
//	)
//}
