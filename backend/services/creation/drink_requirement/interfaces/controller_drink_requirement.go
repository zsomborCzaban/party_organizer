package interfaces

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/zsomborCzaban/party_organizer/common/api"
	"github.com/zsomborCzaban/party_organizer/common/jwt"
	"github.com/zsomborCzaban/party_organizer/services/creation/drink_requirement/domains"
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

func (dc DrinkRequirementController) CreateController(w http.ResponseWriter, r *http.Request) {
	var createDrinkRequirementReq domains.DrinkRequirementDTO
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&createDrinkRequirementReq)
	if err != nil {
		br := api.ErrorBadRequest(err.Error())

		//todo: implement response helper that has logger as param
		br.Send(w)
		return
	}

	userId, err2 := jwt.GetIdFromJWT(r.Header.Get("Authorization"))
	if err2 != nil {
		br := api.ErrorBadRequest(domains.BadRequest)

		br.Send(w)
		return
	}

	resp := dc.DrinkRequirementService.CreateDrinkRequirement(createDrinkRequirementReq, userId)
	couldSend := resp.Send(w)
	if !couldSend {
		//todo: handle logging
		return
	}
}

func (dc DrinkRequirementController) GetController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	drinkReqId, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		br := api.ErrorBadRequest(domains.BadRequest)

		br.Send(w)
		return
	}

	userId, err2 := jwt.GetIdFromJWT(r.Header.Get("Authorization"))
	if err2 != nil {
		br := api.ErrorBadRequest(domains.BadRequest)

		br.Send(w)
		return
	}

	resp := dc.DrinkRequirementService.GetDrinkRequirement(uint(drinkReqId), userId)
	couldSend := resp.Send(w)
	if !couldSend {
		//todo: handle logging
		return
	}
}

func (dc DrinkRequirementController) UpdateController(w http.ResponseWriter, r *http.Request) {
	var updateDrinkRequirementReq domains.DrinkRequirementDTO
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&updateDrinkRequirementReq)
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

	vars := mux.Vars(r)
	id, err3 := strconv.ParseUint(vars["id"], 10, 32)
	if err3 != nil {
		br := api.ErrorBadRequest(err3.Error())

		br.Send(w)
		return
	}
	updateDrinkRequirementReq.ID = uint(id)

	resp := dc.DrinkRequirementService.UpdateDrinkRequirement(updateDrinkRequirementReq, userId)
	couldSend := resp.Send(w)
	if !couldSend {
		//todo: handle logging
		return
	}
}

func (dc DrinkRequirementController) DeleteController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	drinkReqId, err := strconv.ParseUint(vars["id"], 10, 32)
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

	resp := dc.DrinkRequirementService.DeleteDrinkRequirement(uint(drinkReqId), userId)
	couldSend := resp.Send(w)
	if !couldSend {
		//todo: handle logging
		return
	}
}

func (dc DrinkRequirementController) GetByPartyIdController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	partyId, err := strconv.ParseUint(vars["party_id"], 10, 32)
	if err != nil {
		br := api.ErrorBadRequest(domains.BadRequest)

		br.Send(w)
		return
	}

	userId, err2 := jwt.GetIdFromJWT(r.Header.Get("Authorization"))
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
