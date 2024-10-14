package interfaces

import (
	"github.com/gorilla/mux"
	"github.com/zsomborCzaban/party_organizer/common/jwt"
	"github.com/zsomborCzaban/party_organizer/services/invitation/party_invite/domains"
)

func NewPartyInviteRouter(router *mux.Router, controller domains.IPartyInviteController) {
	r := router.PathPrefix("/partyAttendance").Subrouter()

	r.Use(jwt.ValidateJWTMiddleware)

	r.HandleFunc("/accept/{party_id}", controller.Accept).Methods("GET")
	r.HandleFunc("/decline/{party_id}", controller.Decline).Methods("GET")
	r.HandleFunc("/invite/{party_id}/{invitedUser_id}", controller.Invite).Methods("GET")

	r.HandleFunc("/getPendingInvites", controller.GetPendingInvites).Methods("GET")

	r.HandleFunc("/joinPublicParty/{party_id}", controller.JoinPublicParty).Methods("GET") // this could be together with the private party join, but its more clearer this way
	r.HandleFunc("/joinPrivateParty/{party_id}/{access_code}", controller.JoinPrivateParty).Methods("GEt")
}
