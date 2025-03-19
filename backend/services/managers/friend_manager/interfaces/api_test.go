package interfaces

import (
	"github.com/gorilla/mux"
	"github.com/zsomborCzaban/party_organizer/db"
	"github.com/zsomborCzaban/party_organizer/services/managers/friend_manager/usecases"
	usecases2 "github.com/zsomborCzaban/party_organizer/services/user/usecases"
	repoUtils "github.com/zsomborCzaban/party_organizer/utils/repo"
	"testing"
)

var Routes = map[string][]string{
	"/friendManager/accept/{invitor_id}":  {"GET"},
	"/friendManager/decline/{invitor_id}": {"GET"},
	"/friendManager/invite/{username}":    {"GET"},

	"/friendManager/getPendingInvites":        {"GET"},
	"/friendManager/removeFriend/{friend_id}": {"GET"},
}

func Test_NewRouter(t *testing.T) {
	router := mux.NewRouter()
	dbAccess := db.CreateGormDatabaseAccessManager(":memory:", nil)

	userRepo := usecases2.NewUserRepository(dbAccess)
	friendInviteRepo := usecases.NewFriendInviteRepository(dbAccess)

	repoCollector := repoUtils.RepoCollector{
		UserRepo:         userRepo,
		FriendInviteRepo: friendInviteRepo,
	}

	service := usecases.NewFriendInviteService(&repoCollector)
	controller := NewFriendInviteController(service)
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
