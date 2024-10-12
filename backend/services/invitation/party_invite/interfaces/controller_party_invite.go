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
	invitedUserId, err := strconv.ParseUint(vars["invitedUser_id"], 10, 32)
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

	userId, err3 := jwt.GetIdFromJWT(r.Header.Get("Authorization"))
	if err3 != nil {
		br := api.ErrorBadRequest(err3.Error())

		br.Send(w)
		return
	}

	resp := fc.PartyInviteService.Invite(uint(invitedUserId), userId, uint(partyId))
	couldSend := resp.Send(w)
	if !couldSend {
		//todo: handle logging
		return
	}
}

func (fc PartyInviteController) GetPendingInvites(w http.ResponseWriter, r *http.Request) {
	userId, err := jwt.GetIdFromJWT(r.Header.Get("Authorization"))
	if err != nil {
		br := api.ErrorBadRequest(err.Error())

		br.Send(w)
		return
	}

	resp := fc.PartyInviteService.GetPendingInvites(userId)
	couldSend := resp.Send(w)
	if !couldSend {
		//todo: log here
		return
	}
}
