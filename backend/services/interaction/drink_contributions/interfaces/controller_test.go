package interfaces

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/zsomborCzaban/party_organizer/services/interaction/drink_contributions/usecases"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zsomborCzaban/party_organizer/services/interaction/drink_contributions/domains"
	"github.com/zsomborCzaban/party_organizer/utils/api"
	"github.com/zsomborCzaban/party_organizer/utils/jwt"
)

func setupDrinkContributionController() (domains.IDrinkContributionController, *usecases.MockService, *httptest.ResponseRecorder) {
	service := new(usecases.MockService)
	controller := NewDrinkContributionController(service)
	rr := httptest.NewRecorder()
	return controller, service, rr
}

func TestDrinkContributionController_Create_Success(t *testing.T) {
	controller, service, rr := setupDrinkContributionController()
	jwt.GetIdFromJWTFunc = func(string) (uint, error) { return 1, nil }
	service.On("Create", mock.Anything, uint(1)).Return(api.Success(nil))

	body, _ := json.Marshal(domains.DrinkContribution{})
	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(body))
	req.Header.Set("Authorization", "token")

	controller.Create(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	service.AssertExpectations(t)
}

func TestDrinkContributionController_Create_InvalidBody(t *testing.T) {
	controller, _, rr := setupDrinkContributionController()
	req, _ := http.NewRequest("POST", "/", bytes.NewBufferString("invalid"))

	controller.Create(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestDrinkContributionController_Create_InvalidJWT(t *testing.T) {
	controller, _, rr := setupDrinkContributionController()
	jwt.GetIdFromJWTFunc = func(string) (uint, error) { return 0, errors.New("invalid") }

	body, _ := json.Marshal(domains.DrinkContribution{})
	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(body))
	req.Header.Set("Authorization", "token")

	controller.Create(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestDrinkContributionController_Update_Success(t *testing.T) {
	controller, service, rr := setupDrinkContributionController()
	jwt.GetIdFromJWTFunc = func(string) (uint, error) { return 1, nil }
	service.On("Update", mock.Anything, uint(1)).Return(api.Success(nil))

	body, _ := json.Marshal(domains.DrinkContribution{})
	req, _ := http.NewRequest("PUT", "/1", bytes.NewBuffer(body))
	req.Header.Set("Authorization", "token")
	req = mux.SetURLVars(req, map[string]string{"id": "1"})

	controller.Update(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	service.AssertExpectations(t)
}

func TestDrinkContributionController_Update_InvalidBody(t *testing.T) {
	controller, _, rr := setupDrinkContributionController()
	req, _ := http.NewRequest("PUT", "/1", bytes.NewBufferString("invalid"))
	req = mux.SetURLVars(req, map[string]string{"id": "1"})

	controller.Update(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestDrinkContributionController_Update_InvalidJWT(t *testing.T) {
	controller, _, rr := setupDrinkContributionController()
	jwt.GetIdFromJWTFunc = func(string) (uint, error) { return 0, errors.New("invalid") }

	body, _ := json.Marshal(domains.DrinkContribution{})
	req, _ := http.NewRequest("PUT", "/1", bytes.NewBuffer(body))
	req.Header.Set("Authorization", "token")
	req = mux.SetURLVars(req, map[string]string{"id": "1"})

	controller.Update(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestDrinkContributionController_Update_InvalidID(t *testing.T) {
	controller, _, rr := setupDrinkContributionController()
	jwt.GetIdFromJWTFunc = func(string) (uint, error) { return 1, nil }
	body, _ := json.Marshal(domains.DrinkContribution{})
	req, _ := http.NewRequest("PUT", "/invalid", bytes.NewBuffer(body))
	req = mux.SetURLVars(req, map[string]string{"id": "invalid"})

	controller.Update(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestDrinkContributionController_Delete_Success(t *testing.T) {
	controller, service, rr := setupDrinkContributionController()
	jwt.GetIdFromJWTFunc = func(string) (uint, error) { return 1, nil }
	service.On("Delete", uint(1), uint(1)).Return(api.Success(nil))

	req, _ := http.NewRequest("DELETE", "/1", nil)
	req.Header.Set("Authorization", "token")
	req = mux.SetURLVars(req, map[string]string{"id": "1"})

	controller.Delete(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	service.AssertExpectations(t)
}

func TestDrinkContributionController_Delete_InvalidJWT(t *testing.T) {
	controller, _, rr := setupDrinkContributionController()
	jwt.GetIdFromJWTFunc = func(string) (uint, error) { return 0, errors.New("invalid") }

	req, _ := http.NewRequest("DELETE", "/1", nil)
	req.Header.Set("Authorization", "token")
	req = mux.SetURLVars(req, map[string]string{"id": "1"})

	controller.Delete(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestDrinkContributionController_Delete_InvalidID(t *testing.T) {
	controller, _, rr := setupDrinkContributionController()
	jwt.GetIdFromJWTFunc = func(string) (uint, error) { return 1, nil }
	req, _ := http.NewRequest("DELETE", "/invalid", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "invalid"})

	controller.Delete(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestDrinkContributionController_GetByPartyIdAndContributorId_Success(t *testing.T) {
	controller, service, rr := setupDrinkContributionController()
	jwt.GetIdFromJWTFunc = func(string) (uint, error) { return 1, nil }
	service.On("GetByPartyIdAndContributorId", uint(1), uint(2), uint(1)).Return(api.Success(nil))

	req, _ := http.NewRequest("GET", "/parties/1/contributors/2", nil)
	req.Header.Set("Authorization", "token")
	req = mux.SetURLVars(req, map[string]string{"party_id": "1", "contributor_id": "2"})

	controller.GetByPartyIdAndContributorId(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	service.AssertExpectations(t)
}

func TestDrinkContributionController_GetByPartyIdAndContributorId_InvalidJWT(t *testing.T) {
	controller, _, rr := setupDrinkContributionController()
	jwt.GetIdFromJWTFunc = func(string) (uint, error) { return 0, errors.New("invalid") }

	req, _ := http.NewRequest("GET", "/parties/1/contributors/2", nil)
	req.Header.Set("Authorization", "token")
	req = mux.SetURLVars(req, map[string]string{"party_id": "1", "contributor_id": "2"})

	controller.GetByPartyIdAndContributorId(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestDrinkContributionController_GetByPartyIdAndContributorId_InvalidPartyID(t *testing.T) {
	controller, _, rr := setupDrinkContributionController()
	jwt.GetIdFromJWTFunc = func(string) (uint, error) { return 1, nil }
	req, _ := http.NewRequest("GET", "/parties/invalid/contributors/2", nil)
	req = mux.SetURLVars(req, map[string]string{"party_id": "invalid", "contributor_id": "2"})

	controller.GetByPartyIdAndContributorId(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestDrinkContributionController_GetByPartyIdAndContributorId_InvalidContributorID(t *testing.T) {
	controller, _, rr := setupDrinkContributionController()
	jwt.GetIdFromJWTFunc = func(string) (uint, error) { return 1, nil }
	req, _ := http.NewRequest("GET", "/parties/1/contributors/invalid", nil)
	req = mux.SetURLVars(req, map[string]string{"party_id": "1", "contributor_id": "invalid"})

	controller.GetByPartyIdAndContributorId(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestDrinkContributionController_GetByRequirementId_Success(t *testing.T) {
	controller, service, rr := setupDrinkContributionController()
	jwt.GetIdFromJWTFunc = func(string) (uint, error) { return 1, nil }
	service.On("GetByRequirementId", uint(1), uint(1)).Return(api.Success(nil))

	req, _ := http.NewRequest("GET", "/requirements/1", nil)
	req.Header.Set("Authorization", "token")
	req = mux.SetURLVars(req, map[string]string{"requirement_id": "1"})

	controller.GetByRequirementId(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	service.AssertExpectations(t)
}

func TestDrinkContributionController_GetByRequirementId_InvalidJWT(t *testing.T) {
	controller, _, rr := setupDrinkContributionController()
	jwt.GetIdFromJWTFunc = func(string) (uint, error) { return 0, errors.New("invalid") }

	req, _ := http.NewRequest("GET", "/requirements/1", nil)
	req.Header.Set("Authorization", "token")
	req = mux.SetURLVars(req, map[string]string{"requirement_id": "1"})

	controller.GetByRequirementId(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestDrinkContributionController_GetByRequirementId_InvalidID(t *testing.T) {
	controller, _, rr := setupDrinkContributionController()
	jwt.GetIdFromJWTFunc = func(string) (uint, error) { return 1, nil }
	req, _ := http.NewRequest("GET", "/requirements/invalid", nil)
	req = mux.SetURLVars(req, map[string]string{"requirement_id": "invalid"})

	controller.GetByRequirementId(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestDrinkContributionController_GetByPartyId_Success(t *testing.T) {
	controller, service, rr := setupDrinkContributionController()
	jwt.GetIdFromJWTFunc = func(string) (uint, error) { return 1, nil }
	service.On("GetByPartyId", uint(1), uint(1)).Return(api.Success(nil))

	req, _ := http.NewRequest("GET", "/parties/1", nil)
	req.Header.Set("Authorization", "token")
	req = mux.SetURLVars(req, map[string]string{"party_id": "1"})

	controller.GetByPartyId(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	service.AssertExpectations(t)
}

func TestDrinkContributionController_GetByPartyId_InvalidJWT(t *testing.T) {
	controller, _, rr := setupDrinkContributionController()
	jwt.GetIdFromJWTFunc = func(string) (uint, error) { return 0, errors.New("invalid") }

	req, _ := http.NewRequest("GET", "/parties/1", nil)
	req.Header.Set("Authorization", "token")
	req = mux.SetURLVars(req, map[string]string{"party_id": "1"})

	controller.GetByPartyId(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestDrinkContributionController_GetByPartyId_InvalidID(t *testing.T) {
	controller, _, rr := setupDrinkContributionController()
	jwt.GetIdFromJWTFunc = func(string) (uint, error) { return 1, nil }
	req, _ := http.NewRequest("GET", "/parties/invalid", nil)
	req = mux.SetURLVars(req, map[string]string{"party_id": "invalid"})

	controller.GetByPartyId(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}
