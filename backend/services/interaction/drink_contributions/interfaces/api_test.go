package interfaces

import (
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/zsomborCzaban/party_organizer/db"
	usecases3 "github.com/zsomborCzaban/party_organizer/services/creation/drink_requirement/usecases"
	usecases4 "github.com/zsomborCzaban/party_organizer/services/creation/party/usecases"
	"github.com/zsomborCzaban/party_organizer/services/interaction/drink_contributions/usecases"
	usecases2 "github.com/zsomborCzaban/party_organizer/services/users/user/usecases"
	"github.com/zsomborCzaban/party_organizer/utils/api"
	repoUtils "github.com/zsomborCzaban/party_organizer/utils/repo"
	"testing"
)

var Routes = map[string][]string{
	"/drinkContribution":      {"POST"},
	"/drinkContribution/{id}": {"PUT", "DELETE"},

	"/drinkContribution/getByPartyAndContributor/{party_id}/{contributor_id}": {"GET"},
	"/drinkContribution/getByRequirement/{requirement_id}":                    {"GET"},
	"/drinkContribution/getByParty/{party_id}":                                {"GET"},
}

func Test_NewRouter(t *testing.T) {
	router := mux.NewRouter()

	dbAccess := db.CreateGormDatabaseAccessManager(":memory:", nil)
	vali := api.NewValidator(validator.New())

	partyRepo := usecases4.NewPartyRepository(dbAccess)
	userRepo := usecases2.NewUserRepository(dbAccess)
	drinkReqRepo := usecases3.NewDrinkRequirementRepository(dbAccess)
	drinkContribRepo := usecases.NewDrinkContributionRepository(dbAccess)
	repoCollector := repoUtils.RepoCollector{
		PartyRepo:        partyRepo,
		UserRepo:         userRepo,
		DrinkReqRepo:     drinkReqRepo,
		DrinkContribRepo: drinkContribRepo,
	}

	service := usecases.NewDrinkContributionService(&repoCollector, vali)
	controller := NewDrinkContributionController(service)
	NewRouter(router, controller)

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
