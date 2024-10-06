package interfaces

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/zsomborCzaban/party_organizer/common/api"
	"github.com/zsomborCzaban/party_organizer/services/creation/food_requirement/domains"
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

func (dc FoodRequirementController) CreateController(w http.ResponseWriter, r *http.Request) {
	var createFoodRequirementReq domains.FoodRequirementDTO
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&createFoodRequirementReq)
	if err != nil {
		br := api.ErrorBadRequest(err.Error())

		//todo: implement response helper that has logger as param
		br.Send(w)
		return
	}
	fmt.Println("didnt fail there")

	resp := dc.FoodRequirementService.CreateFoodRequirement(createFoodRequirementReq)
	couldSend := resp.Send(w)
	if !couldSend {
		//todo: handle logging
		return
	}
}

func (dc FoodRequirementController) GetController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		br := api.ErrorBadRequest(domains.BadRequest)

		br.Send(w)
		return
	}

	resp := dc.FoodRequirementService.GetFoodRequirement(uint(id))
	couldSend := resp.Send(w)
	if !couldSend {
		//todo: handle logging
		return
	}
}

func (dc FoodRequirementController) UpdateController(w http.ResponseWriter, r *http.Request) {
	var updateFoodRequirementReq domains.FoodRequirementDTO
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&updateFoodRequirementReq)
	if err != nil {
		br := api.ErrorBadRequest(domains.BadRequest)

		//todo: implement response helper that has logger as param
		br.Send(w)
		return
	}

	resp := dc.FoodRequirementService.UpdateFoodRequirement(updateFoodRequirementReq)
	couldSend := resp.Send(w)
	if !couldSend {
		//todo: handle logging
		return
	}
}

func (dc FoodRequirementController) DeleteController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		br := api.ErrorBadRequest(domains.BadRequest)

		br.Send(w)
		return
	}

	resp := dc.FoodRequirementService.DeleteFoodRequirement(uint(id))
	couldSend := resp.Send(w)
	if !couldSend {
		//todo: handle logging
		return
	}
}
