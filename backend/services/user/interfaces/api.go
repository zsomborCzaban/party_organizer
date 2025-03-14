package interfaces

import (
	"github.com/gorilla/mux"
	"github.com/zsomborCzaban/party_organizer/services/user/domains"
)

func NewUserAuthRouter(router *mux.Router, controller domains.IUserController) {
	router.HandleFunc("/login", controller.LoginController).Methods("POST")
	router.HandleFunc("/register", controller.RegisterController).Methods("POST")
}

func NewUserRouter(router *mux.Router, controller domains.IUserController) {
	router.HandleFunc("/user/getFriends", controller.GetFriendsController).Methods("GET")
	router.HandleFunc("/user/uploadProfilePicture", controller.UploadProfilePicture).Methods("POST")

}
