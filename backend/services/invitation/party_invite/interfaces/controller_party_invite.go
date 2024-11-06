package interfaces

import (
	"github.com/gorilla/mux"
	"github.com/zsomborCzaban/party_organizer/common/api"
	"github.com/zsomborCzaban/party_organizer/common/jwt"
	"github.com/zsomborCzaban/party_organizer/services/invitation/party_invite/domains"
	"net/http"
	"strconv"
)

type PartyInviteController struct {
	PartyInviteService domains.IPartyInviteService
}

func NewPartyInviteController(service domains.IPartyInviteService) domains.IPartyInviteController {
	return &PartyInviteController{PartyInviteService: service}
}

func (fc PartyInviteController) Accept(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	partyId, err := strconv.ParseUint(vars["party_id"], 10, 32)
	if err != nil {
		br := api.ErrorBadRequest(err.Error())

		br.Send(w)
		return
	}

	userId, err2 := jwt.GetIdFromJWT(r.Header.Get("Authorization"))
	if err2 != nil {
		br := api.ErrorBadRequest(err2.Error())

		br.Send(w)
		return
	}

	resp := fc.PartyInviteService.Accept(userId, uint(partyId))
	couldSend := resp.Send(w)
	if !couldSend {
		//todo: handle logging
		return
	}
}

func (fc PartyInviteController) Decline(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	partyId, err := strconv.ParseUint(vars["party_id"], 10, 32)
	if err != nil {
		br := api.ErrorBadRequest(err.Error())

		br.Send(w)
		return
	}

	userId, err2 := jwt.GetIdFromJWT(r.Header.Get("Authorization"))
	if err2 != nil {
		br := api.ErrorBadRequest(err2.Error())

		br.Send(w)
		return
	}

	resp := fc.PartyInviteService.Decline(userId, uint(partyId))
	couldSend := resp.Send(w)
	if !couldSend {
		//todo: handle logging
		return
	}
}

func (fc PartyInviteController) Invite(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	invitedUsername, err := vars["invited_username"]
	if !err {
		br := api.ErrorBadRequest("wrong parameter as 'invited_username'")

		br.Send(w)
		return
	}

	partyId, err2 := strconv.ParseUint(vars["party_id"], 10, 32)
	if err2 != nil {
		br := api.ErrorBadRequest(err2.Error())

		br.Send(w)
		return
	}

	userId, err3 := jwt.GetIdFromJWT(r.Header.Get("Authorization"))
	if err3 != nil {
		br := api.ErrorBadRequest(err3.Error())

		br.Send(w)
		return
	}

	resp := fc.PartyInviteService.Invite(invitedUsername, userId, uint(partyId))
	couldSend := resp.Send(w)
	if !couldSend {
		//todo: handle logging
		return
	}
}

func (fc PartyInviteController) GetUserPendingInvites(w http.ResponseWriter, r *http.Request) {
	userId, err := jwt.GetIdFromJWT(r.Header.Get("Authorization"))
	if err != nil {
		br := api.ErrorBadRequest(err.Error())

		br.Send(w)
		return
	}

	resp := fc.PartyInviteService.GetUserPendingInvites(userId)
	couldSend := resp.Send(w)
	if !couldSend {
		//todo: log here
		return
	}
}

func (fc PartyInviteController) GetPartyPendingInvites(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, err := jwt.GetIdFromJWT(r.Header.Get("Authorization"))
	if err != nil {
		br := api.ErrorBadRequest(err.Error())

		br.Send(w)
		return
	}

	partyId, err2 := strconv.ParseUint(vars["party_id"], 10, 32)
	if err2 != nil {
		br := api.ErrorBadRequest(err2.Error())

		br.Send(w)
		return
	}

	resp := fc.PartyInviteService.GetPartyPendingInvites(uint(partyId), userId)
	couldSend := resp.Send(w)
	if !couldSend {
		//todo: log here
		return
	}
}

func (fc PartyInviteController) Kick(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	partyId, err := strconv.ParseUint(vars["party_id"], 10, 32)
	if err != nil {
		br := api.ErrorBadRequest(err.Error())

		br.Send(w)
		return
	}

	kickedId, err2 := strconv.ParseUint(vars["kicked_id"], 10, 32)
	if err2 != nil {
		br := api.ErrorBadRequest(err2.Error())

		br.Send(w)
		return
	}

	userId, err3 := jwt.GetIdFromJWT(r.Header.Get("Authorization"))
	if err3 != nil {
		br := api.ErrorBadRequest(err3.Error())

		br.Send(w)
		return
	}

	resp := fc.PartyInviteService.Kick(uint(kickedId), userId, uint(partyId))
	couldSend := resp.Send(w)
	if !couldSend {
		//todo: handle logging
		return
	}
}

func (fc PartyInviteController) JoinPublicParty(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	partyId, err := strconv.ParseUint(vars["party_id"], 10, 32)
	if err != nil {
		br := api.ErrorBadRequest(err.Error())

		br.Send(w)
		return
	}

	userId, err2 := jwt.GetIdFromJWT(r.Header.Get("Authorization"))
	if err2 != nil {
		br := api.ErrorBadRequest(err2.Error())

		br.Send(w)
		return
	}

	resp := fc.PartyInviteService.JoinPublicParty(uint(partyId), userId)
	couldSend := resp.Send(w)
	if !couldSend {
		//todo: handle logging
		return
	}
}

func (fc PartyInviteController) JoinPrivateParty(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	partyId, err := strconv.ParseUint(vars["party_id"], 10, 32)
	if err != nil {
		br := api.ErrorBadRequest(err.Error())

		br.Send(w)
		return
	}

	accessCode := vars["access_code"]
	if accessCode != "" {
		br := api.ErrorBadRequest("missing access_code")

		br.Send(w)
		return
	}

	userId, err2 := jwt.GetIdFromJWT(r.Header.Get("Authorization"))
	if err2 != nil {
		br := api.ErrorBadRequest(err2.Error())

		br.Send(w)
		return
	}

	resp := fc.PartyInviteService.JoinPrivateParty(uint(partyId), userId, accessCode)
	couldSend := resp.Send(w)
	if !couldSend {
		//todo: handle logging
		return
	}
}
