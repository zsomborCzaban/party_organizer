package interfaces

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
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
		br := domains.ErrorBadRequest(err.Error())

		//todo: implement response helper that has logger as param
		br.Send(w)
		return
	}
	fmt.Println("didnt fail there")

	resp := dc.DrinkRequirementService.CreateDrinkRequirement(createDrinkRequirementReq)
	couldSend := resp.Send(w)
	if !couldSend {
		//todo: handle logging
		return
	}
}

func (dc DrinkRequirementController) GetController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		br := domains.ErrorBadRequest(domains.BadRequest)

		br.Send(w)
		return
	}

	resp := dc.DrinkRequirementService.GetDrinkRequirement(uint(id))
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
		br := domains.ErrorBadRequest(domains.BadRequest)

		//todo: implement response helper that has logger as param
		br.Send(w)
		return
	}

	resp := dc.DrinkRequirementService.UpdateDrinkRequirement(updateDrinkRequirementReq)
	couldSend := resp.Send(w)
	if !couldSend {
		//todo: handle logging
		return
	}
}

func (dc DrinkRequirementController) DeleteController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		br := domains.ErrorBadRequest(domains.BadRequest)

		br.Send(w)
		return
	}

	resp := dc.DrinkRequirementService.DeleteDrinkRequirement(uint(id))
	couldSend := resp.Send(w)
	if !couldSend {
		//todo: handle logging
		return
	}
}
