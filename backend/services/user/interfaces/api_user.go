package interfaces

import (
	"github.com/gorilla/mux"
	"github.com/zsomborCzaban/party_organizer/services/user/domains"
)

func NewUserRouter(router *mux.Router, controller domains.IUserController) {
	r := router.PathPrefix("/user").Subrouter()

	r.HandleFunc("/login", controller.LoginController).Methods("POST")
	r.HandleFunc("/register", controller.RegisterController).Methods("POST")

	//todo: authenticate these:
	r.HandleFunc("/addFriend/{id}", controller.AddFriendController).Methods("GET")
	r.HandleFunc("/getFriends", controller.GetFriendsController).Methods("GET")

}
