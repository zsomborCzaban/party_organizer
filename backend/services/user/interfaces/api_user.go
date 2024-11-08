package interfaces

import (
	"github.com/gorilla/mux"
	"github.com/zsomborCzaban/party_organizer/services/user/domains"
)

func NewUserRouter(router *mux.Router, controller domains.IUserController) {
	r := router.PathPrefix("/user").Subrouter()

	//todo: refactor this to /login and /register and /getFriends
	r.HandleFunc("/login/", controller.LoginController).Methods("POST")
	r.HandleFunc("/register/", controller.RegisterController).Methods("POST")

	//todo: authenticate these with middleware:
	r.HandleFunc("/addFriend/{id}", controller.AddFriendController).Methods("GET") //wont be used by user
	r.HandleFunc("/getFriends/", controller.GetFriendsController).Methods("GET")

}
