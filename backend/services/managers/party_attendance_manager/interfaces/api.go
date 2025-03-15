package interfaces

import (
	"github.com/gorilla/mux"
	"github.com/zsomborCzaban/party_organizer/services/managers/party_attendance_manager/domains"
)

func NewRouter(router *mux.Router, controller domains.IPartyInviteController) {
	router.HandleFunc("/partyAttendanceManager/accept/{party_id}", controller.Accept).Methods("GET")
	router.HandleFunc("/partyAttendanceManager/decline/{party_id}", controller.Decline).Methods("GET")
	router.HandleFunc("/partyAttendanceManager/invite/{party_id}/{invited_username}", controller.Invite).Methods("GET")

	router.HandleFunc("/partyAttendanceManager/getPendingInvites", controller.GetUserPendingInvites).Methods("GET")
	router.HandleFunc("/partyAttendanceManager/getPartyPendingInvites/{party_id}", controller.GetPartyPendingInvites).Methods("GET")

	router.HandleFunc("/partyAttendanceManager/kick/{party_id}/{kicked_id}", controller.Kick).Methods("GET")
	router.HandleFunc("/partyAttendanceManager/leaveParty/{party_id}", controller.LeaveParty).Methods("GET")
	router.HandleFunc("/partyAttendanceManager/joinPublicParty/{party_id}", controller.JoinPublicParty).Methods("GET") // this could be together with the private party join, but its more clearer this way
	router.HandleFunc("/partyAttendanceManager/joinPrivateParty/{access_code}", controller.JoinPrivateParty).Methods("GET")
}
