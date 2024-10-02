package interfaces

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
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
		br := domains.ErrorBadRequest(err.Error())

		//todo: implement response helper that has logger as param
		br.Send(w)
		return
	}
	fmt.Println("didnt fail there")

	resp := pc.PartyService.CreateParty(createPartyReq)
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
		br := domains.ErrorBadRequest(domains.BadRequest)

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
		br := domains.ErrorBadRequest(domains.BadRequest)

		//todo: implement response helper that has logger as param
		br.Send(w)
		return
	}

	resp := pc.PartyService.UpdateParty(updatePartyReq)
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
		br := domains.ErrorBadRequest(domains.BadRequest)

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
