package interfaces

import (
	"github.com/gorilla/mux"
	"github.com/zsomborCzaban/party_organizer/common/jwt"
	"github.com/zsomborCzaban/party_organizer/services/invitation/friend_invite/domains"
)

func NewFriendInviteRouter(router *mux.Router, controller domains.IFriendInviteController) {
	r := router.PathPrefix("/friendInvite").Subrouter()

	r.Use(jwt.ValidateJWTMiddleware)

	r.HandleFunc("/accept/{invitor_id}", controller.Accept).Methods("GET")
	r.HandleFunc("/decline/{invitor_id}", controller.Decline).Methods("GET")
	r.HandleFunc("/invite/{invited_id}", controller.Invite).Methods("GET")

	r.HandleFunc("/getPendingInvites", controller.GetPendingInvites).Methods("GET")
	r.HandleFunc("/removeFriend/{friend_id}", controller.RemoveFriend).Methods("GET")
}
