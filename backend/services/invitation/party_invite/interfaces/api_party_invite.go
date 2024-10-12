package interfaces

import (
	"github.com/gorilla/mux"
	"github.com/zsomborCzaban/party_organizer/common/jwt"
	"github.com/zsomborCzaban/party_organizer/services/invitation/party_invite/domains"
)

func NewPartyInviteRouter(router *mux.Router, controller domains.IPartyInviteController) {
	r := router.PathPrefix("/partyInvite").Subrouter()

	r.Use(jwt.ValidateJWTMiddleware)

	r.HandleFunc("/accept/{party_id}", controller.Accept).Methods("GET")
	r.HandleFunc("/decline/{party_id}", controller.Decline).Methods("GET")
	r.HandleFunc("/invite/{invitedUser_id}/{party_id}", controller.Invite).Methods("GET")

	r.HandleFunc("/getPendingInvites", controller.GetPendingInvites).Methods("GET")
}
