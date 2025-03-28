package interfaces

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/mock"
	"github.com/zsomborCzaban/party_organizer/services/creation/drink_requirement/usecases"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/zsomborCzaban/party_organizer/services/creation/drink_requirement/domains"
	"github.com/zsomborCzaban/party_organizer/utils/api"
	"github.com/zsomborCzaban/party_organizer/utils/jwt"
)

func setupDefaultController() (domains.IDrinkRequirementController, *usecases.MockService, *api.MockResponseWriter) {
	service := new(usecases.MockService)
	controller := NewDrinkRequirementController(service)
	writer := new(api.MockResponseWriter)

	return controller, service, writer
}

func Test_ControllerCreate_Success(t *testing.T) {
	controller, service, writer := setupDefaultController()

	req := domains.DrinkRequirement{}
	requestData, _ := json.Marshal(req)
	expectedResponse := api.Success("")
	respJson, _ := json.Marshal(expectedResponse)

	writer.On("Header").Return(make(http.Header))
	writer.On("WriteHeader", mock.Anything).Return()
	writer.On("Write", mock.Anything).Return(0, nil)
	service.On("CreateDrinkRequirement", mock.Anything).Return(expectedResponse, nil)
	request := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestData))

	controller.CreateController(writer, request)

	writer.AssertCalled(t, "Write", respJson)
}

func TestCreateController_Success(t *testing.T) {
	mockService := &MockDrinkRequirementService{}
	mockService.On("Create", mock.Anything, uint(1)).Return(api.SuccessResponse(http.StatusCreated, nil))

	jwt.GetIdFromJWT = func(string) (uint, error) { return 1, nil }

	controller := NewDrinkRequirementController(mockService)
	reqBody := domains.DrinkRequirementDTO{Name: "Beer", Amount: 10, PartyID: 1}
	body, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(body))
	req.Header.Set("Authorization", "token")
	rr := httptest.NewRecorder()

	controller.CreateController(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
	mockService.AssertExpectations(t)
}

func TestCreateController_InvalidBody(t *testing.T) {
	controller := NewDrinkRequirementController(nil)
	req, _ := http.NewRequest("POST", "/", bytes.NewBufferString("invalid"))
	rr := httptest.NewRecorder()

	controller.CreateController(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestCreateController_InvalidJWT(t *testing.T) {
	jwt.GetIdFromJWT = func(string) (uint, error) { return 0, errors.New("invalid") }

	controller := NewDrinkRequirementController(nil)
	reqBody := domains.DrinkRequirementDTO{Name: "Beer"}
	body, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(body))
	req.Header.Set("Authorization", "token")
	rr := httptest.NewRecorder()

	controller.CreateController(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestGetController_Success(t *testing.T) {
	mockService := &MockDrinkRequirementService{}
	mockService.On("FindById", uint(1), uint(1)).Return(api.SuccessResponse(http.StatusOK, nil))

	jwt.GetIdFromJWT = func(string) (uint, error) { return 1, nil }

	controller := NewDrinkRequirementController(mockService)
	req, _ := http.NewRequest("GET", "/1", nil)
	req.Header.Set("Authorization", "token")
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	rr := httptest.NewRecorder()

	controller.GetController(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	mockService.AssertExpectations(t)
}

func TestGetController_InvalidID(t *testing.T) {
	controller := NewDrinkRequirementController(nil)
	req, _ := http.NewRequest("GET", "/invalid", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "invalid"})
	rr := httptest.NewRecorder()

	controller.GetController(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestDeleteController_Success(t *testing.T) {
	mockService := &MockDrinkRequirementService{}
	mockService.On("Delete", uint(1), uint(1)).Return(api.SuccessResponse(http.StatusOK, nil))

	jwt.GetIdFromJWT = func(string) (uint, error) { return 1, nil }

	controller := NewDrinkRequirementController(mockService)
	req, _ := http.NewRequest("DELETE", "/1", nil)
	req.Header.Set("Authorization", "token")
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	rr := httptest.NewRecorder()

	controller.DeleteController(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	mockService.AssertExpectations(t)
}

func TestGetByPartyIdController_Success(t *testing.T) {
	mockService := &MockDrinkRequirementService{}
	mockService.On("GetByPartyId", uint(1), uint(1)).Return(api.SuccessResponse(http.StatusOK, nil))

	jwt.GetIdFromJWT = func(string) (uint, error) { return 1, nil }

	controller := NewDrinkRequirementController(mockService)
	req, _ := http.NewRequest("GET", "/party/1", nil)
	req.Header.Set("Authorization", "token")
	req = mux.SetURLVars(req, map[string]string{"party_id": "1"})
	rr := httptest.NewRecorder()

	controller.GetByPartyIdController(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	mockService.AssertExpectations(t)
}
