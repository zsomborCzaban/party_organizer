package interfaces

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/zsomborCzaban/party_organizer/services/creation/drink_requirement/domains"
	"github.com/zsomborCzaban/party_organizer/utils/api"
	"github.com/zsomborCzaban/party_organizer/utils/jwt"
	"net/http"
	"strconv"
)

type DrinkRequirementController struct {
	DrinkRequirementService domains.IDrinkRequirementService
}

func NewDrinkRequirementController(service domains.IDrinkRequirementService) domains.IDrinkRequirementController {
	return &DrinkRequirementController{
		DrinkRequirementService: service,
	}
}

func (dc DrinkRequirementController) Create(w http.ResponseWriter, r *http.Request) {
	var createDrinkRequirementReq domains.DrinkRequirementDTO
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&createDrinkRequirementReq)
	if err != nil {
		br := api.ErrorBadRequest(err.Error())

		//todo: implement response helper that has logger as param
		br.Send(w)
		return
	}

	userId, err2 := jwt.GetIdFromJWTFunc(r.Header.Get("Authorization"))
	if err2 != nil {
		br := api.ErrorBadRequest(domains.BadRequest)

		br.Send(w)
		return
	}

	resp := dc.DrinkRequirementService.Create(createDrinkRequirementReq, userId)
	couldSend := resp.Send(w)
	if !couldSend {
		//todo: handle logging
		return
	}
}

func (dc DrinkRequirementController) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	drinkReqId, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		br := api.ErrorBadRequest(domains.BadRequest)

		br.Send(w)
		return
	}

	userId, err2 := jwt.GetIdFromJWTFunc(r.Header.Get("Authorization"))
	if err2 != nil {
		br := api.ErrorBadRequest(domains.BadRequest)

		br.Send(w)
		return
	}

	resp := dc.DrinkRequirementService.FindById(uint(drinkReqId), userId)
	couldSend := resp.Send(w)
	if !couldSend {
		//todo: handle logging
		return
	}
}

func (dc DrinkRequirementController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	drinkReqId, err := strconv.ParseUint(vars["id"], 10, 32)
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

	resp := dc.DrinkRequirementService.Delete(uint(drinkReqId), userId)
	couldSend := resp.Send(w)
	if !couldSend {
		//todo: handle logging
		return
	}
}

func (dc DrinkRequirementController) GetByPartyId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	partyId, err := strconv.ParseUint(vars["party_id"], 10, 32)
	if err != nil {
		br := api.ErrorBadRequest(domains.BadRequest)

		br.Send(w)
		return
	}

	userId, err2 := jwt.GetIdFromJWTFunc(r.Header.Get("Authorization"))
	if err2 != nil {
		br := api.ErrorBadRequest(domains.BadRequest)

		br.Send(w)
		return
	}

	resp := dc.DrinkRequirementService.GetByPartyId(uint(partyId), userId)
	couldSend := resp.Send(w)
	if !couldSend {
		//todo: handle logging
		return
	}
}
