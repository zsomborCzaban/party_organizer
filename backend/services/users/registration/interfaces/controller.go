package interfaces

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/zsomborCzaban/party_organizer/services/users/registration/domains"
	"github.com/zsomborCzaban/party_organizer/utils/api"
	"net/http"
)

type RegistrationController struct {
	RegistrationService domains.IRegistrationService
}

func NewRegistrationController(registrationService domains.IRegistrationService) domains.IRegistrationController {
	return &RegistrationController{
		RegistrationService: registrationService,
	}
}

func (c *RegistrationController) Register(w http.ResponseWriter, r *http.Request) {
	var registerReq domains.RegistrationRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&registerReq); err != nil {
		br := api.ErrorBadRequest(domains.BadRequest)

		br.Send(w)
		return
	}

	resp := c.RegistrationService.Register(registerReq)
	couldSend := resp.Send(w)
	if !couldSend {
		return
	}
}

func (c *RegistrationController) ConfirmEmail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username, ok := vars["username"]
	if !ok {
		br := api.ErrorBadRequest("Failed to get username")

		br.Send(w)
		return
	}

	confirmHash, ok2 := vars["confirm_hash"]
	if !ok2 {
		br := api.ErrorBadRequest("Failed to get confirm_hash")

		br.Send(w)
		return
	}

	resp := c.RegistrationService.ConfirmEmail(username, confirmHash)
	couldSend := resp.Send(w)
	if !couldSend {
		return
	}
}
