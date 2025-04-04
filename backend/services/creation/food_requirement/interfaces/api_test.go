package interfaces

import (
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/zsomborCzaban/party_organizer/db"
	"github.com/zsomborCzaban/party_organizer/services/creation/food_requirement/usecases"
	usecases2 "github.com/zsomborCzaban/party_organizer/services/creation/party/usecases"
	usecases3 "github.com/zsomborCzaban/party_organizer/services/interaction/food_contributions/usecases"
	"github.com/zsomborCzaban/party_organizer/utils/api"
	repoUtils "github.com/zsomborCzaban/party_organizer/utils/repo"
	"testing"
)

var Routes = map[string][]string{
	"/foodRequirement":                         {"POST"},
	"/foodRequirement/{id}":                    {"GET", "DELETE"},
	"/foodRequirement/getByPartyId/{party_id}": {"GET"},
}

func Test_NewRouter(t *testing.T) {
	router := mux.NewRouter()

	dbAccess := db.CreateGormDatabaseAccessManager(":memory:", nil)
	vali := api.NewValidator(validator.New())
	repo := repoUtils.RepoCollector{}
	foodReqRepo := usecases.NewFoodRequirementRepository(dbAccess)
	partyRepo := usecases2.NewPartyRepository(dbAccess)
	foodContribRepo := usecases3.NewFoodContributionRepository(dbAccess)
	repo.FoodReqRepo = foodReqRepo
	repo.PartyRepo = partyRepo
	repo.FoodContribRepo = foodContribRepo
	service := usecases.NewFoodRequirementService(&repo, vali)
	controller := NewFoodRequirementController(service)
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
