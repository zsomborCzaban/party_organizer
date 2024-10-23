package interfaces

import (
	"github.com/gorilla/mux"
	"github.com/zsomborCzaban/party_organizer/common/api"
	"github.com/zsomborCzaban/party_organizer/common/jwt"
	"github.com/zsomborCzaban/party_organizer/services/invitation/friend_invite/domains"
	"net/http"
	"strconv"
)

type FriendInviteController struct {
	FriendInviteService domains.IFriendInviteService
}

func NewFriendInviteController(service domains.IFriendInviteService) domains.IFriendInviteController {
	return &FriendInviteController{FriendInviteService: service}
}

func (fc FriendInviteController) Accept(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	invitorId, err := strconv.ParseUint(vars["invitor_id"], 10, 32)
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

	resp := fc.FriendInviteService.Accept(uint(invitorId), userId)
	couldSend := resp.Send(w)
	if !couldSend {
		//todo: handle logging
		return
	}
}

func (fc FriendInviteController) Decline(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	invitorId, err := strconv.ParseUint(vars["invitor_id"], 10, 32)
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

	resp := fc.FriendInviteService.Decline(uint(invitorId), userId)
	couldSend := resp.Send(w)
	if !couldSend {
		//todo: handle logging
		return
	}
}

func (fc FriendInviteController) Invite(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	invitedUsername, err := vars["username"]
	if !err {
		br := api.ErrorBadRequest("cannot parse given username")

		br.Send(w)
		return
	}

	userId, err2 := jwt.GetIdFromJWT(r.Header.Get("Authorization"))
	if err2 != nil {
		br := api.ErrorBadRequest(err2.Error())

		br.Send(w)
		return
	}

	resp := fc.FriendInviteService.Invite(invitedUsername, userId)
	couldSend := resp.Send(w)
	if !couldSend {
		//todo: handle logging
		return
	}
}

func (fc FriendInviteController) GetPendingInvites(w http.ResponseWriter, r *http.Request) {
	userId, err := jwt.GetIdFromJWT(r.Header.Get("Authorization"))
	if err != nil {
		br := api.ErrorBadRequest(err.Error())

		br.Send(w)
		return
	}

	resp := fc.FriendInviteService.GetPendingInvites(userId)
	couldSend := resp.Send(w)
	if !couldSend {
		//todo: log here
		return
	}
}

func (fc FriendInviteController) RemoveFriend(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	friendId, err := strconv.ParseUint(vars["friend_id"], 10, 32)
	if err != nil {
		br := api.ErrorBadRequest(err.Error())

		br.Send(w)
		return
	}

	userId, err := jwt.GetIdFromJWT(r.Header.Get("Authorization"))
	if err != nil {
		br := api.ErrorBadRequest(err.Error())

		br.Send(w)
		return
	}

	resp := fc.FriendInviteService.RemoveFriend(userId, uint(friendId))
	couldSend := resp.Send(w)
	if !couldSend {
		//todo: log here
		return
	}
}
