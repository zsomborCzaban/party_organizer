package interfaces

import (
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/zsomborCzaban/party_organizer/db"
	"github.com/zsomborCzaban/party_organizer/services/creation/drink_requirement/usecases"
	usecases2 "github.com/zsomborCzaban/party_organizer/services/creation/party/usecases"
	usecases3 "github.com/zsomborCzaban/party_organizer/services/interaction/drink_contributions/usecases"
	"github.com/zsomborCzaban/party_organizer/utils/api"
	repoUtils "github.com/zsomborCzaban/party_organizer/utils/repo"
	"testing"
)

var Routes = map[string][]string{
	"/drinkRequirement":                         {"POST"},
	"/drinkRequirement/{id}":                    {"GET", "DELETE"},
	"/drinkRequirement/getByPartyId/{party_id}": {"GET"},
}

func Test_NewRouter(t *testing.T) {
	router := mux.NewRouter()

	dbAccess := db.CreateGormDatabaseAccessManager(":memory:", nil)
	vali := api.NewValidator(validator.New())
	repo := repoUtils.RepoCollector{}
	drinkReqRepo := usecases.NewDrinkRequirementRepository(dbAccess)
	partyRepo := usecases2.NewPartyRepository(dbAccess)
	drinkContribRepo := usecases3.NewDrinkContributionRepository(dbAccess)
	repo.DrinkReqRepo = drinkReqRepo
	repo.PartyRepo = partyRepo
	repo.DrinkContribRepo = drinkContribRepo
	service := usecases.NewDrinkRequirementService(&repo, vali)
	controller := NewDrinkRequirementController(service)
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
