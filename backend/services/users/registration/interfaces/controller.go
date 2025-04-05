package interfaces

import (
	"encoding/json"
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

func (uc *RegistrationController) ConfirmEmail(w http.ResponseWriter, r *http.Request) {
	//todo: implement me
	//var registerReq domains.RegistrationRequest
	//decoder := json.NewDecoder(r.Body)
	//if err := decoder.Decode(&registerReq); err != nil {
	//	br := api.ErrorBadRequest(domains.BadRequest)
	//
	//	br.Send(w)
	//	return
	//}
	//
	//resp := uc.UserService.Register(registerReq)
	//couldSend := resp.Send(w)
	//if !couldSend {
	//	return
	//}
}
