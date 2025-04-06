package interfaces

import (
	"github.com/gorilla/mux"
	"github.com/zsomborCzaban/party_organizer/services/users/user/domains"
)

func NewUserAuthRouter(router *mux.Router, controller domains.IUserController) {
	router.HandleFunc("/login", controller.LoginController).Methods("POST")
	//router.HandleFunc("/register", controller.RegisterController).Methods("POST")
	router.HandleFunc("/resetPassword/{username}", controller.ForgotPassword).Methods("GET")
}

func NewUserPrivateRouter(router *mux.Router, controller domains.IUserController) {
	router.HandleFunc("/user/getFriends", controller.GetFriendsController).Methods("GET")
	router.HandleFunc("/user/uploadProfilePicture", controller.UploadProfilePicture).Methods("POST")
	router.HandleFunc("/resetPassword", controller.ChangePassword).Methods("POST")

}
