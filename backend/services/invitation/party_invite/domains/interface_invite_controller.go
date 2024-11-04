package domains

import "net/http"

type IPartyInviteController interface {
	Invite(w http.ResponseWriter, r *http.Request)
	Accept(w http.ResponseWriter, r *http.Request)
	Decline(w http.ResponseWriter, r *http.Request)

	GetPendingInvites(w http.ResponseWriter, r *http.Request)
	GetPendingAndAcceptedInvites(w http.ResponseWriter, r *http.Request)

	Kick(w http.ResponseWriter, r *http.Request)
	JoinPublicParty(w http.ResponseWriter, r *http.Request)
	JoinPrivateParty(w http.ResponseWriter, r *http.Request)
}
