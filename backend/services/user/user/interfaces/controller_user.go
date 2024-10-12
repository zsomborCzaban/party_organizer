package interfaces

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/zsomborCzaban/party_organizer/common/api"
	domains2 "github.com/zsomborCzaban/party_organizer/services/user/user/domains"
	"net/http"
	"strconv"
)

type UserController struct {
	UserService domains2.IUserService
}

func NewUserController(userService domains2.IUserService) domains2.IUserController {
	return &UserController{
		UserService: userService,
	}
}

func (uc UserController) LoginController(w http.ResponseWriter, r *http.Request) {
	var loginReq domains2.LoginRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&loginReq); err != nil {
		br := api.ErrorBadRequest(domains2.BadRequest)

		br.Send(w)
		return
	}

	resp := uc.UserService.Login(loginReq)
	couldSend := resp.Send(w)
	if !couldSend {
		return
	}
}

func (uc UserController) RegisterController(w http.ResponseWriter, r *http.Request) {
	var registerReq domains2.RegisterRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&registerReq); err != nil {
		br := api.ErrorBadRequest(domains2.BadRequest)

		br.Send(w)
		return
	}

	resp := uc.UserService.Register(registerReq)
	couldSend := resp.Send(w)
	if !couldSend {
		return
	}
}

func (uc UserController) AddFriendController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	partyId, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		br := api.ErrorBadRequest(domains2.BadRequest)

		br.Send(w)
		return
	}

	resp := uc.UserService.AddFriend(uint(partyId), 3)
	couldSend := resp.Send(w)
	if !couldSend {
		//todo: handle logging
		return
	}
}

func (uc UserController) GetFriendsController(w http.ResponseWriter, r *http.Request) {
	resp := uc.UserService.GetFriends(3)
	couldSend := resp.Send(w)
	if !couldSend {
		//todo: handle logging
		return
	}
}
