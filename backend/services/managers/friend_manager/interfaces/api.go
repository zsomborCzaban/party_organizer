package interfaces

import (
	"github.com/gorilla/mux"
	"github.com/zsomborCzaban/party_organizer/services/managers/friend_manager/domains"
)

func NewRouter(router *mux.Router, controller domains.IFriendInviteController) {
	router.HandleFunc("/friendManager/accept/{invitor_id}", controller.Accept).Methods("GET")
	router.HandleFunc("/friendManager/decline/{invitor_id}", controller.Decline).Methods("GET")
	router.HandleFunc("/friendManager/invite/{username}", controller.Invite).Methods("GET")

	router.HandleFunc("/friendManager/getPendingInvites", controller.GetPendingInvites).Methods("GET")
	router.HandleFunc("/friendManager/removeFriend/{friend_id}", controller.RemoveFriend).Methods("GET")
}
