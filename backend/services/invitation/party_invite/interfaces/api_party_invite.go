package interfaces

import (
	"github.com/gorilla/mux"
	"github.com/zsomborCzaban/party_organizer/common/jwt"
	"github.com/zsomborCzaban/party_organizer/services/invitation/party_invite/domains"
)

func NewPartyInviteRouter(router *mux.Router, controller domains.IPartyInviteController) {
	r := router.PathPrefix("/partyAttendanceManager").Subrouter()

	r.Use(jwt.ValidateJWTMiddleware)

	r.HandleFunc("/accept/{party_id}", controller.Accept).Methods("GET")
	r.HandleFunc("/decline/{party_id}", controller.Decline).Methods("GET")
	r.HandleFunc("/invite/{party_id}/{invited_username}", controller.Invite).Methods("GET")

	r.HandleFunc("/getPendingInvites/", controller.GetUserPendingInvites).Methods("GET")
	r.HandleFunc("/getPartyPendingInvites/{party_id}", controller.GetPartyPendingInvites).Methods("GET")

	r.HandleFunc("/kick/{party_id}/{kicked_id}", controller.Kick).Methods("GET")
	r.HandleFunc("/joinPublicParty/{party_id}", controller.JoinPublicParty).Methods("GET") // this could be together with the private party join, but its more clearer this way
	r.HandleFunc("/joinPrivateParty/{party_id}/{access_code}", controller.JoinPrivateParty).Methods("GEt")
}
