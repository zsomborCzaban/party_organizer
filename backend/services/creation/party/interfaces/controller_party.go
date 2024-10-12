package interfaces

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/zsomborCzaban/party_organizer/common/api"
	"github.com/zsomborCzaban/party_organizer/common/jwt"
	"github.com/zsomborCzaban/party_organizer/services/creation/party/domains"
	"net/http"
	"strconv"
)

type PartyController struct {
	PartyService domains.IPartyService
}

func NewPartyController(service domains.IPartyService) domains.IPartyController {
	return &PartyController{
		PartyService: service,
	}
}

func (pc PartyController) CreateController(w http.ResponseWriter, r *http.Request) {
	var createPartyReq domains.PartyDTO
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&createPartyReq)
	if err != nil {
		br := api.ErrorBadRequest(err.Error())

		//todo: implement response helper that has logger as param
		br.Send(w)
		return
	}

	userId, err2 := jwt.GetIdFromJWT(r.Header.Get("Authorization"))
	if err2 != nil {
		br := api.ErrorBadRequest(err2.Error())

		br.Send(w)
		return
	}

	resp := pc.PartyService.CreateParty(createPartyReq, userId)
	couldSend := resp.Send(w)
	if !couldSend {
		//todo: handle logging
		return
	}
}

func (pc PartyController) GetController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		br := api.ErrorBadRequest(domains.BadRequest)

		br.Send(w)
		return
	}

	resp := pc.PartyService.GetParty(uint(id))
	couldSend := resp.Send(w)
	if !couldSend {
		//todo: handle logging
		return
	}
}

func (pc PartyController) UpdateController(w http.ResponseWriter, r *http.Request) {
	var updatePartyReq domains.PartyDTO
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&updatePartyReq)
	if err != nil {
		br := api.ErrorBadRequest(err.Error())

		//todo: implement response helper that has logger as param
		br.Send(w)
		return
	}

	userId, err2 := jwt.GetIdFromJWT(r.Header.Get("Authorization"))
	if err2 != nil {
		br := api.ErrorBadRequest(err2.Error())

		br.Send(w)
		return
	}

	vars := mux.Vars(r)
	id, err3 := strconv.ParseUint(vars["id"], 10, 32)
	if err3 != nil {
		br := api.ErrorBadRequest(err3.Error())

		br.Send(w)
		return
	}
	updatePartyReq.ID = uint(id)

	resp := pc.PartyService.UpdateParty(updatePartyReq, userId)
	couldSend := resp.Send(w)
	if !couldSend {
		//todo: handle logging
		return
	}
}

func (pc PartyController) DeleteController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		br := api.ErrorBadRequest(domains.BadRequest)

		br.Send(w)
		return
	}

	resp := pc.PartyService.DeleteParty(uint(id))
	couldSend := resp.Send(w)
	if !couldSend {
		//todo: handle logging
		return
	}
}

func (pc PartyController) GetPartiesByOrganizerId(w http.ResponseWriter, r *http.Request) {
	userId, err := jwt.GetIdFromJWT(r.Header.Get("Authorization"))
	if err != nil {
		br := api.ErrorBadRequest(err.Error())

		br.Send(w)
		return
	}

	resp := pc.PartyService.GetPartiesByOrganizerId(userId)
	couldSend := resp.Send(w)
	if !couldSend {
		//todo: handle logging
		return
	}
}

func (pc PartyController) GetPartiesByParticipantId(w http.ResponseWriter, r *http.Request) {
	userId, err := jwt.GetIdFromJWT(r.Header.Get("Authorization"))
	if err != nil {
		br := api.ErrorBadRequest(err.Error())

		br.Send(w)
		return
	}

	resp := pc.PartyService.GetPartiesByParticipantId(userId)
	couldSend := resp.Send(w)
	if !couldSend {
		//todo: handle logging
		return
	}
}

func (pc PartyController) AddUserToParty(w http.ResponseWriter, r *http.Request) {
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

	resp := pc.PartyService.AddUserToParty(uint(partyId), userId)
	couldSend := resp.Send(w)
	if !couldSend {
		//todo: handle logging
		return
	}
}
