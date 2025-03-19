package interfaces

import (
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/zsomborCzaban/party_organizer/db"
	usecases3 "github.com/zsomborCzaban/party_organizer/services/creation/drink_requirement/usecases"
	usecases5 "github.com/zsomborCzaban/party_organizer/services/creation/food_requirement/usecases"
	"github.com/zsomborCzaban/party_organizer/services/creation/party/usecases"
	usecases4 "github.com/zsomborCzaban/party_organizer/services/interaction/drink_contributions/usecases"
	usecases6 "github.com/zsomborCzaban/party_organizer/services/interaction/food_contributions/usecases"
	usecases7 "github.com/zsomborCzaban/party_organizer/services/managers/party_attendance_manager/usecases"
	usecases2 "github.com/zsomborCzaban/party_organizer/services/user/usecases"
	"github.com/zsomborCzaban/party_organizer/utils/api"
	repoUtils "github.com/zsomborCzaban/party_organizer/utils/repo"
	"testing"
)

var Routes = map[string][]string{
	"/party":                            {"POST", "PUT"},
	"/party/{id}":                       {"GET", "DELETE"},
	"/party/getPartiesByOrganizerId":    {"GET"},
	"/party/getPartiesByParticipantId":  {"GET"},
	"/party/getParticipants/{party_id}": {"GET"},
}

var PublicRoutes = map[string][]string{
	"/publicParties":      {"GET"},
	"/publicParties/{id}": {"GET"},
}

func Test_NewPartyRouter(t *testing.T) {
	router := mux.NewRouter()

	dbAccess := db.CreateGormDatabaseAccessManager(":memory:", nil)
	vali := api.NewValidator(validator.New())

	partyRepo := usecases.NewPartyRepository(dbAccess)
	userRepo := usecases2.NewUserRepository(dbAccess)
	drinkReqRepo := usecases3.NewDrinkRequirementRepository(dbAccess)
	drinkContribRepo := usecases4.NewDrinkContributionRepository(dbAccess)
	foodReqRepo := usecases5.NewFoodRequirementRepository(dbAccess)
	foodContribRepo := usecases6.NewFoodContributionRepository(dbAccess)
	partyInviteRepo := usecases7.NewPartyInviteRepository(dbAccess)
	repoCollector := repoUtils.RepoCollector{
		PartyRepo:        partyRepo,
		UserRepo:         userRepo,
		DrinkReqRepo:     drinkReqRepo,
		DrinkContribRepo: drinkContribRepo,
		FoodReqRepo:      foodReqRepo,
		FoodContribRepo:  foodContribRepo,
		PartyInviteRepo:  partyInviteRepo,
	}

	service := usecases.NewPartyService(&repoCollector, vali)
	controller := NewPartyController(service)
	NewPartyRouter(router, controller)

	router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		path, err := route.GetPathTemplate()
		if err != nil {
			t.FailNow()
		}

		methods, err := route.GetMethods()
		if err != nil {
			t.FailNow()
		}

		val, ok := Routes[path]
		if !ok {
			t.FailNow()
		}

		found := false
		for _, v := range methods {
			for i, v2 := range val {
				if v2 == v {
					found = true

					Routes[path] = append(Routes[path][:i], Routes[path][i+1:]...)
				}
			}
		}

		if !found {
			t.FailNow()
		}

		return nil
	})
}

func Test_NewPublicPartyRouter(t *testing.T) {
	router := mux.NewRouter()

	dbAccess := db.CreateGormDatabaseAccessManager(":memory:", nil)
	vali := api.NewValidator(validator.New())

	partyRepo := usecases.NewPartyRepository(dbAccess)
	userRepo := usecases2.NewUserRepository(dbAccess)
	drinkReqRepo := usecases3.NewDrinkRequirementRepository(dbAccess)
	drinkContribRepo := usecases4.NewDrinkContributionRepository(dbAccess)
	foodReqRepo := usecases5.NewFoodRequirementRepository(dbAccess)
	foodContribRepo := usecases6.NewFoodContributionRepository(dbAccess)
	partyInviteRepo := usecases7.NewPartyInviteRepository(dbAccess)
	repoCollector := repoUtils.RepoCollector{
		PartyRepo:        partyRepo,
		UserRepo:         userRepo,
		DrinkReqRepo:     drinkReqRepo,
		DrinkContribRepo: drinkContribRepo,
		FoodReqRepo:      foodReqRepo,
		FoodContribRepo:  foodContribRepo,
		PartyInviteRepo:  partyInviteRepo,
	}

	service := usecases.NewPartyService(&repoCollector, vali)
	controller := NewPartyController(service)
	NewPublicPartyRouter(router, controller)

	router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		path, err := route.GetPathTemplate()
		if err != nil {
			return nil
		}

		methods, err := route.GetMethods()
		if err != nil {
			return nil
		}

		val, ok := PublicRoutes[path]
		if !ok {
			return nil
		}

		found := false
		for _, v := range methods {
			for i, v2 := range val {
				if v2 == v {
					found = true

					PublicRoutes[path] = append(PublicRoutes[path][:i], PublicRoutes[path][i+1:]...)
				}
			}
		}

		if !found {
			t.FailNow()
		}

		return nil
	})
}
