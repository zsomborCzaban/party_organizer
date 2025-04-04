package interfaces

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/mock"
	"github.com/zsomborCzaban/party_organizer/services/creation/food_requirement/usecases"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zsomborCzaban/party_organizer/services/creation/food_requirement/domains"
	"github.com/zsomborCzaban/party_organizer/utils/api"
	"github.com/zsomborCzaban/party_organizer/utils/jwt"
)

func setupDefaultController() (domains.IFoodRequirementController, *usecases.MockService, *httptest.ResponseRecorder) {
	service := new(usecases.MockService)
	controller := NewFoodRequirementController(service)
	responseRecorder := httptest.NewRecorder()

	return controller, service, responseRecorder
}

func TestCreateController_Success(t *testing.T) {
	controller, service, responseRecorder := setupDefaultController()

	jwt.GetIdFromJWTFunc = func(string) (uint, error) { return 1, nil }
	service.On("Create", mock.Anything, uint(1)).Return(api.Success(nil))

	reqBody := domains.FoodRequirementDTO{
		PartyID:        1,
		Type:           "verytest",
		TargetQuantity: 2,
		QuantityMark:   "test",
	}
	body, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(body))
	req.Header.Set("Authorization", "token")

	controller.Create(responseRecorder, req)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	service.AssertExpectations(t)
}

func TestCreateController_InvalidBody(t *testing.T) {
	controller, _, responseRecorder := setupDefaultController()

	req, _ := http.NewRequest("POST", "/", bytes.NewBufferString("invalid"))

	controller.Create(responseRecorder, req)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
}

func TestCreateController_InvalidJWT(t *testing.T) {
	controller, _, responseRecorder := setupDefaultController()

	jwt.GetIdFromJWTFunc = func(string) (uint, error) { return 0, errors.New("err") }

	reqBody := domains.FoodRequirementDTO{
		PartyID:        1,
		Type:           "test",
		TargetQuantity: 2,
		QuantityMark:   "testtest",
	}
	body, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(body))
	req.Header.Set("Authorization", "token")

	controller.Create(responseRecorder, req)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
}

func TestGetController_Success(t *testing.T) {
	controller, service, responseRecorder := setupDefaultController()

	jwt.GetIdFromJWTFunc = func(string) (uint, error) { return 1, nil }
	service.On("FindById", uint(1), uint(1)).Return(api.Success(nil))

	req, _ := http.NewRequest("GET", "/1", nil)
	req.Header.Set("Authorization", "token")
	req = mux.SetURLVars(req, map[string]string{"id": "1"})

	controller.Get(responseRecorder, req)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	service.AssertExpectations(t)
}

func TestGetController_InvalidID(t *testing.T) {
	controller, _, responseRecorder := setupDefaultController()

	req, _ := http.NewRequest("GET", "/invalid", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "invalid"})

	controller.Get(responseRecorder, req)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
}

func TestGetController_InvalidJWT(t *testing.T) {
	controller, _, responseRecorder := setupDefaultController()
	jwt.GetIdFromJWTFunc = func(string) (uint, error) { return 0, errors.New("invalid token") }

	req, _ := http.NewRequest("GET", "/1", nil)
	req.Header.Set("Authorization", "invalid_token")
	req = mux.SetURLVars(req, map[string]string{"id": "1"})

	controller.Get(responseRecorder, req)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
}

func TestDeleteController_Success(t *testing.T) {
	controller, service, responseRecorder := setupDefaultController()

	jwt.GetIdFromJWTFunc = func(string) (uint, error) { return 1, nil }
	service.On("Delete", uint(1), uint(1)).Return(api.Success(nil))

	req, _ := http.NewRequest("DELETE", "/1", nil)
	req.Header.Set("Authorization", "token")
	req = mux.SetURLVars(req, map[string]string{"id": "1"})

	controller.Delete(responseRecorder, req)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	service.AssertExpectations(t)
}

func TestDeleteController_InvalidJWT(t *testing.T) {
	controller, _, responseRecorder := setupDefaultController()

	jwt.GetIdFromJWTFunc = func(string) (uint, error) { return 0, errors.New("invalid token") }

	req, _ := http.NewRequest("DELETE", "/1", nil)
	req.Header.Set("Authorization", "invalid_token")
	req = mux.SetURLVars(req, map[string]string{"id": "1"})

	controller.Delete(responseRecorder, req)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
}

func TestDeleteController_InvalidID(t *testing.T) {
	controller, _, responseRecorder := setupDefaultController()

	req, _ := http.NewRequest("DELETE", "/invalid", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "invalid"})

	controller.Delete(responseRecorder, req)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
}

func TestGetByPartyIdController_Success(t *testing.T) {
	controller, service, responseRecorder := setupDefaultController()

	jwt.GetIdFromJWTFunc = func(string) (uint, error) { return 1, nil }
	service.On("GetByPartyId", uint(1), uint(1)).Return(api.Success(nil))

	req, _ := http.NewRequest("GET", "/party/1", nil)
	req.Header.Set("Authorization", "token")
	req = mux.SetURLVars(req, map[string]string{"party_id": "1"})

	controller.GetByPartyId(responseRecorder, req)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	service.AssertExpectations(t)
}

func TestGetByPartyIdController_InvalidJWT(t *testing.T) {
	controller, _, responseRecorder := setupDefaultController()

	jwt.GetIdFromJWTFunc = func(string) (uint, error) { return 0, errors.New("invalid token") }

	req, _ := http.NewRequest("GET", "/party/1", nil)
	req.Header.Set("Authorization", "invalid_token")
	req = mux.SetURLVars(req, map[string]string{"party_id": "1"})

	controller.GetByPartyId(responseRecorder, req)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
}

func TestGetByPartyIdController_InvalidID(t *testing.T) {
	controller, _, responseRecorder := setupDefaultController()

	req, _ := http.NewRequest("GET", "/party/invalid", nil)
	req = mux.SetURLVars(req, map[string]string{"party_id": "invalid"})

	controller.GetByPartyId(responseRecorder, req)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
}
