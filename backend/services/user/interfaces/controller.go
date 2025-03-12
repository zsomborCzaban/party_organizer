package interfaces

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/zsomborCzaban/party_organizer/services/user/domains"
	"github.com/zsomborCzaban/party_organizer/utils/api"
	"github.com/zsomborCzaban/party_organizer/utils/jwt"
	"net/http"
	"strconv"
)

type UserController struct {
	UserService domains.IUserService
}

func NewUserController(userService domains.IUserService) domains.IUserController {
	return &UserController{
		UserService: userService,
	}
}

func (uc *UserController) LoginController(w http.ResponseWriter, r *http.Request) {
	var loginReq domains.LoginRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&loginReq); err != nil {
		br := api.ErrorBadRequest(domains.BadRequest)

		br.Send(w)
		return
	}

	resp := uc.UserService.Login(loginReq)
	couldSend := resp.Send(w)
	if !couldSend {
		return
	}
}

func (uc *UserController) RegisterController(w http.ResponseWriter, r *http.Request) {
	var registerReq domains.RegisterRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&registerReq); err != nil {
		br := api.ErrorBadRequest(domains.BadRequest)

		br.Send(w)
		return
	}

	resp := uc.UserService.Register(registerReq)
	couldSend := resp.Send(w)
	if !couldSend {
		return
	}
}

func (uc *UserController) AddFriendController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	partyId, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		br := api.ErrorBadRequest(domains.BadRequest)

		br.Send(w)
		return
	}

	userId, err2 := jwt.GetIdFromJWT(r.Header.Get("Authorization"))
	if err2 != nil {
		br := api.ErrorBadRequest(err2.Error())

		br.Send(w)
		return
	}

	resp := uc.UserService.AddFriend(uint(partyId), userId)
	couldSend := resp.Send(w)
	if !couldSend {
		//todo: handle logging
		return
	}
}

func (uc *UserController) GetFriendsController(w http.ResponseWriter, r *http.Request) {
	userId, err := jwt.GetIdFromJWT(r.Header.Get("Authorization"))
	if err != nil {
		br := api.ErrorBadRequest(err.Error())

		br.Send(w)
		return
	}

	resp := uc.UserService.GetFriends(userId)
	couldSend := resp.Send(w)
	if !couldSend {
		//todo: handle logging
		return
	}
}

func (uc *UserController) UploadProfilePicture(w http.ResponseWriter, r *http.Request) {
	userId, err := jwt.GetIdFromJWT(r.Header.Get("Authorization"))
	if err != nil {
		br := api.ErrorBadRequest(err.Error())

		br.Send(w)
		return
	}

	err2 := r.ParseMultipartForm(10 << 20)
	if err2 != nil {
		br := api.ErrorBadRequest(domains.BadRequest)

		br.Send(w)
		return
	}

	file, fileHeader, err3 := r.FormFile("image")
	if err3 != nil {
		br := api.ErrorBadRequest(domains.BadRequest)

		br.Send(w)
		return
	}

	resp := uc.UserService.UploadProfilePicture(userId, file, fileHeader)
	couldSend := resp.Send(w)
	if !couldSend {
		//todo: handle logging
		return
	}
}
