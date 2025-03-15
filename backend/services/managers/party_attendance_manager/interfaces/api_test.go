package interfaces

import (
	"github.com/gorilla/mux"
	"github.com/zsomborCzaban/party_organizer/db"
	usecases2 "github.com/zsomborCzaban/party_organizer/services/creation/party/usecases"
	usecases4 "github.com/zsomborCzaban/party_organizer/services/interaction/drink_contributions/usecases"
	usecases5 "github.com/zsomborCzaban/party_organizer/services/interaction/food_contributions/usecases"
	"github.com/zsomborCzaban/party_organizer/services/managers/party_attendance_manager/usecases"
	usecases3 "github.com/zsomborCzaban/party_organizer/services/user/usecases"
	repoUtils "github.com/zsomborCzaban/party_organizer/utils/repo"
	"testing"
)

var Routes = map[string][]string{
	"/partyAttendanceManager/accept/{party_id}":                    {"GET"},
	"/partyAttendanceManager/decline/{party_id}":                   {"GET"},
	"/partyAttendanceManager/invite/{party_id}/{invited_username}": {"GET"},

	"/partyAttendanceManager/getPendingInvites":                 {"GET"},
	"/partyAttendanceManager/getPartyPendingInvites/{party_id}": {"GET"},

	"/partyAttendanceManager/kick/{party_id}/{kicked_id}":    {"GET"},
	"/partyAttendanceManager/leaveParty/{party_id}":          {"GET"},
	"/partyAttendanceManager/joinPublicParty/{party_id}":     {"GET"},
	"/partyAttendanceManager/joinPrivateParty/{access_code}": {"GET"},
}

func Test_NewRouter(t *testing.T) {
	router := mux.NewRouter()
	dbAccess := db.CreateGormDatabaseAccessManager(":memory:", nil)

	partyIniviteRepo := usecases.NewPartyInviteRepository(dbAccess)
	partyRepo := usecases2.NewPartyRepository(dbAccess)
	userRepo := usecases3.NewUserRepository(dbAccess)
	drinkContribRepo := usecases4.NewDrinkContributionRepository(dbAccess)
	foodContribRepo := usecases5.NewFoodContributionRepository(dbAccess)

	repoCollector := repoUtils.RepoCollector{
		PartyRepo:        &partyRepo,
		UserRepo:         &userRepo,
		DrinkContribRepo: &drinkContribRepo,
		FoodContribRepo:  &foodContribRepo,
		PartyInviteRepo:  &partyIniviteRepo,
	}

	service := usecases.NewPartyInviteService(&repoCollector)
	controller := NewPartyInviteController(service)
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
