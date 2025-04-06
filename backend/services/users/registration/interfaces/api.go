package interfaces

import (
	"github.com/gorilla/mux"
	"github.com/zsomborCzaban/party_organizer/services/users/registration/domains"
)

func NewRegistrationRouter(router *mux.Router, controller domains.IRegistrationController) {
	router.HandleFunc("/register", controller.Register).Methods("POST")
	router.HandleFunc("/confirmEmail/{username}/{confirm_hash}", controller.ConfirmEmail).Methods("GET") //todo: rewrite to post instead of get
}
