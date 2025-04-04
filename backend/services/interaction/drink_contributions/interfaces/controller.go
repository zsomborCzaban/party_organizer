package interfaces

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/zsomborCzaban/party_organizer/services/interaction/drink_contributions/domains"
	"github.com/zsomborCzaban/party_organizer/utils/api"
	"github.com/zsomborCzaban/party_organizer/utils/jwt"
	"net/http"
	"strconv"
)

type DrinkContributionController struct {
	ContributionService domains.IDrinkContributionService
}

func NewDrinkContributionController(service domains.IDrinkContributionService) domains.IDrinkContributionController {
	return &DrinkContributionController{ContributionService: service}
}

func (dc DrinkContributionController) Create(w http.ResponseWriter, r *http.Request) {
	var createContributionReq domains.DrinkContribution
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&createContributionReq)
	if err != nil {
		br := api.ErrorBadRequest(err.Error())

		//todo: implement response helper that has logger as param
		br.Send(w)
		return
	}

	userId, err2 := jwt.GetIdFromJWTFunc(r.Header.Get("Authorization"))
	if err2 != nil {
		br := api.ErrorBadRequest(err2.Error())

		br.Send(w)
		return
	}

	resp := dc.ContributionService.Create(createContributionReq, userId)
	couldSend := resp.Send(w)
	if !couldSend {
		//todo: handle logging
		return
	}
}

func (dc DrinkContributionController) Update(w http.ResponseWriter, r *http.Request) {
	var updateContributionReq domains.DrinkContribution
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&updateContributionReq)
	if err != nil {
		br := api.ErrorBadRequest(err.Error())

		//todo: implement response helper that has logger as param
		br.Send(w)
		return
	}

	userId, err2 := jwt.GetIdFromJWTFunc(r.Header.Get("Authorization"))
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
	updateContributionReq.ID = uint(id)

	resp := dc.ContributionService.Update(updateContributionReq, userId)
	couldSend := resp.Send(w)
	if !couldSend {
		//todo: handle logging
		return
	}
}

func (dc DrinkContributionController) Delete(w http.ResponseWriter, r *http.Request) {
	userId, err := jwt.GetIdFromJWTFunc(r.Header.Get("Authorization"))
	if err != nil {
		br := api.ErrorBadRequest(err.Error())

		br.Send(w)
		return
	}

	vars := mux.Vars(r)
	contributionId, err2 := strconv.ParseUint(vars["id"], 10, 32)
	if err2 != nil {
		br := api.ErrorBadRequest(err2.Error())

		br.Send(w)
		return
	}

	resp := dc.ContributionService.Delete(uint(contributionId), userId)
	couldSend := resp.Send(w)
	if !couldSend {
		//todo: handle logging
		return
	}
}

func (dc DrinkContributionController) GetByPartyIdAndContributorId(w http.ResponseWriter, r *http.Request) {
	userId, err := jwt.GetIdFromJWTFunc(r.Header.Get("Authorization"))
	if err != nil {
		br := api.ErrorBadRequest(err.Error())

		br.Send(w)
		return
	}

	vars := mux.Vars(r)
	partyId, err2 := strconv.ParseUint(vars["party_id"], 10, 32)
	if err2 != nil {
		br := api.ErrorBadRequest(err2.Error())

		br.Send(w)
		return
	}

	contributorId, err3 := strconv.ParseUint(vars["contributor_id"], 10, 32)
	if err3 != nil {
		br := api.ErrorBadRequest(err3.Error())

		br.Send(w)
		return
	}

	resp := dc.ContributionService.GetByPartyIdAndContributorId(uint(partyId), uint(contributorId), userId)
	couldSend := resp.Send(w)
	if !couldSend {
		//todo: handle logging
		return
	}
}

func (dc DrinkContributionController) GetByRequirementId(w http.ResponseWriter, r *http.Request) {
	userId, err := jwt.GetIdFromJWTFunc(r.Header.Get("Authorization"))
	if err != nil {
		br := api.ErrorBadRequest(err.Error())

		br.Send(w)
		return
	}

	vars := mux.Vars(r)

	requirementId, err3 := strconv.ParseUint(vars["requirement_id"], 10, 32)
	if err3 != nil {
		br := api.ErrorBadRequest(err3.Error())

		br.Send(w)
		return
	}

	resp := dc.ContributionService.GetByRequirementId(uint(requirementId), userId)
	couldSend := resp.Send(w)
	if !couldSend {
		//todo: handle logging
		return
	}
}

func (dc DrinkContributionController) GetByPartyId(w http.ResponseWriter, r *http.Request) {
	userId, err := jwt.GetIdFromJWTFunc(r.Header.Get("Authorization"))
	if err != nil {
		br := api.ErrorBadRequest(err.Error())

		br.Send(w)
		return
	}

	vars := mux.Vars(r)
	partyId, err2 := strconv.ParseUint(vars["party_id"], 10, 32)
	if err2 != nil {
		br := api.ErrorBadRequest(err2.Error())

		br.Send(w)
		return
	}

	resp := dc.ContributionService.GetByPartyId(uint(partyId), userId)
	couldSend := resp.Send(w)
	if !couldSend {
		//todo: handle logging
		return
	}
}
