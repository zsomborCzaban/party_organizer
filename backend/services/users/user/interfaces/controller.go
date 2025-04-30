package interfaces

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/zsomborCzaban/party_organizer/services/users/user/domains"
	"github.com/zsomborCzaban/party_organizer/utils/api"
	"github.com/zsomborCzaban/party_organizer/utils/jwt"
	"net/http"
	"strconv"
	"strings"
)

type UserController struct {
	UserService domains.IUserService
}

func NewUserController(userService domains.IUserService) domains.IUserController {
	return &UserController{
		UserService: userService,
	}
}

func (uc *UserController) Login(w http.ResponseWriter, r *http.Request) {
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

func (uc *UserController) AddFriend(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	partyId, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		br := api.ErrorBadRequest(domains.BadRequest)

		br.Send(w)
		return
	}

	userId, err2 := jwt.GetIdFromJWTFunc(r.Header.Get("Authorization"))
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

func (uc *UserController) GetFriends(w http.ResponseWriter, r *http.Request) {
	userId, err := jwt.GetIdFromJWTFunc(r.Header.Get("Authorization"))
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
	userId, err := jwt.GetIdFromJWTFunc(r.Header.Get("Authorization"))
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

	file, mutlipartFileHeader, err3 := r.FormFile("image")
	if err3 != nil {
		br := api.ErrorBadRequest(domains.BadRequest)

		br.Send(w)
		return
	}

	fileHeader := make([]byte, 512)
	if _, err4 := file.Read(fileHeader); err4 != nil {
		br := api.ErrorBadRequest(err4.Error())

		br.Send(w)
		return
	}

	mime := http.DetectContentType(fileHeader)
	if !strings.Contains(mime, "image") {
		br := api.ErrorBadRequest(domains.InvalidFileFormat)

		br.Send(w)
		return

	}

	//sets the file pointer back to the beginning of the file
	if _, err5 := file.Seek(0, 0); err5 != nil {
		br := api.ErrorBadRequest(err5.Error())

		br.Send(w)
		return
	}

	resp := uc.UserService.UploadProfilePicture(userId, file, mutlipartFileHeader)
	couldSend := resp.Send(w)
	if !couldSend {
		//todo: handle logging
		return
	}
}

func (uc *UserController) ForgotPassword(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username, ok := vars["username"]
	if !ok {
		br := api.ErrorBadRequest("Failed to get username")

		br.Send(w)
		return
	}

	resp := uc.UserService.ForgotPassword(username)
	couldSend := resp.Send(w)
	if !couldSend {
		return
	}
}

func (uc *UserController) ChangePassword(w http.ResponseWriter, r *http.Request) {
	var resetReq domains.ChangePasswordRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&resetReq); err != nil {
		br := api.ErrorBadRequest(domains.BadRequest)

		br.Send(w)
		return
	}

	userId, err := jwt.GetIdFromJWTFunc(r.Header.Get("Authorization"))
	if err != nil {
		br := api.ErrorBadRequest(err.Error())

		br.Send(w)
		return
	}

	isAuthorized, err := jwt.GetCanChangePasswordFromJWT(r.Header.Get("Authorization"))
	if err != nil {
		br := api.ErrorBadRequest(err.Error())

		br.Send(w)
		return
	}
	if !isAuthorized {
		br := api.ErrorUnauthorized("Not authorized to change password. Follow the link in you emails")

		br.Send(w)
		return
	}

	resp := uc.UserService.ChangePassword(resetReq, userId)
	couldSend := resp.Send(w)
	if !couldSend {
		return
	}
}
