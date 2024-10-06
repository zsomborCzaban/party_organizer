package interfaces

import (
	"encoding/json"
	"github.com/zsomborCzaban/party_organizer/common/api"
	"github.com/zsomborCzaban/party_organizer/services/user/domains"
	"net/http"
)

type UserController struct {
	UserService domains.IUserService
}

func NewUserController(userService domains.IUserService) domains.IUserController {
	return &UserController{
		UserService: userService,
	}
}

func (uc UserController) LoginController(w http.ResponseWriter, r *http.Request) {
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

func (uc UserController) RegisterController(w http.ResponseWriter, r *http.Request) {
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
