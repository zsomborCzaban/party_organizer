package interfaces

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/zsomborCzaban/party_organizer/services/creation/food_requirement/domains"
	"github.com/zsomborCzaban/party_organizer/utils/api"
	"github.com/zsomborCzaban/party_organizer/utils/jwt"
	"net/http"
	"strconv"
)

type FoodRequirementController struct {
	FoodRequirementService domains.IFoodRequirementService
}

func NewFoodRequirementController(service domains.IFoodRequirementService) domains.IFoodRequirementController {
	return &FoodRequirementController{
		FoodRequirementService: service,
	}
}

func (fc FoodRequirementController) CreateController(w http.ResponseWriter, r *http.Request) {
	var createFoodRequirementReq domains.FoodRequirementDTO
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&createFoodRequirementReq)
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

	resp := fc.FoodRequirementService.Create(createFoodRequirementReq, userId)
	couldSend := resp.Send(w)
	if !couldSend {
		//todo: handle logging
		return
	}
}

func (fc FoodRequirementController) GetController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	foodReqId, err := strconv.ParseUint(vars["id"], 10, 32)
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

	resp := fc.FoodRequirementService.FindById(uint(foodReqId), userId)
	couldSend := resp.Send(w)
	if !couldSend {
		//todo: handle logging
		return
	}
}

func (fc FoodRequirementController) DeleteController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	foodReqId, err := strconv.ParseUint(vars["id"], 10, 32)
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

	resp := fc.FoodRequirementService.Delete(uint(foodReqId), userId)
	couldSend := resp.Send(w)
	if !couldSend {
		//todo: handle logging
		return
	}
}

func (fc FoodRequirementController) GetByPartyIdController(w http.ResponseWriter, r *http.Request) {
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

	resp := fc.FoodRequirementService.GetByPartyId(uint(partyId), userId)
	couldSend := resp.Send(w)
	if !couldSend {
		return
	}
}
