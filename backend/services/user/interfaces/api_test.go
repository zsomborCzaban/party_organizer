package interfaces

import (
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/zsomborCzaban/party_organizer/db"
	"github.com/zsomborCzaban/party_organizer/services/user/usecases"
	"github.com/zsomborCzaban/party_organizer/utils/api"
	repoUtils "github.com/zsomborCzaban/party_organizer/utils/repo"
	"testing"
)

var AuthRoutes = map[string][]string{
	"/login":    {"POST"},
	"/register": {"POST"},
}

var UserRoutes = map[string][]string{
	"/user/getFriends":           {"GET"},
	"/user/uploadProfilePicture": {"POST"},
}

func Test_NewUserAuthRouter(t *testing.T) {
	router := mux.NewRouter()

	dbAccess := db.CreateGormDatabaseAccessManager(":memory:", nil)
	vali := api.NewValidator(validator.New())

	userRepo := usecases.NewUserRepository(dbAccess)
	repoCollector := repoUtils.RepoCollector{
		UserRepo: userRepo,
	}

	service := usecases.NewUserService(&repoCollector, vali, nil)
	controller := NewUserController(service)
	NewUserAuthRouter(router, controller)

	router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		path, err := route.GetPathTemplate()
		if err != nil {
			t.FailNow()
		}

		methods, err := route.GetMethods()
		if err != nil {
			t.FailNow()
		}

		val, ok := AuthRoutes[path]
		if !ok {
			t.FailNow()
		}

		found := false
		for _, v := range methods {
			for i, v2 := range val {
				if v2 == v {
					found = true

					AuthRoutes[path] = append(AuthRoutes[path][:i], AuthRoutes[path][i+1:]...)
				}
			}
		}

		if !found {
			t.FailNow()
		}

		return nil
	})
}

func Test_NewUserRouter(t *testing.T) {
	router := mux.NewRouter()

	dbAccess := db.CreateGormDatabaseAccessManager(":memory:", nil)
	vali := api.NewValidator(validator.New())

	userRepo := usecases.NewUserRepository(dbAccess)
	repoCollector := repoUtils.RepoCollector{
		UserRepo: userRepo,
	}

	service := usecases.NewUserService(&repoCollector, vali, nil)
	controller := NewUserController(service)
	NewUserPublicRouter(router, controller)

	router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		path, err := route.GetPathTemplate()
		if err != nil {
			t.FailNow()
		}

		methods, err := route.GetMethods()
		if err != nil {
			t.FailNow()
		}

		val, ok := UserRoutes[path]
		if !ok {
			t.FailNow()
		}

		found := false
		for _, v := range methods {
			for i, v2 := range val {
				if v2 == v {
					found = true

					UserRoutes[path] = append(UserRoutes[path][:i], UserRoutes[path][i+1:]...)
				}
			}
		}

		if !found {
			t.FailNow()
		}

		return nil
	})
}
