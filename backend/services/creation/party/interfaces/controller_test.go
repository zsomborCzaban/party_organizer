package interfaces

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/zsomborCzaban/party_organizer/services/creation/party/usecases"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zsomborCzaban/party_organizer/services/creation/party/domains"
	"github.com/zsomborCzaban/party_organizer/utils/api"
	"github.com/zsomborCzaban/party_organizer/utils/jwt"
)

func setupPartyController() (domains.IPartyController, *usecases.MockService, *httptest.ResponseRecorder) {
	service := new(usecases.MockService)
	controller := NewController(service)
	rr := httptest.NewRecorder()
	return controller, service, rr
}

func TestPartyController_CreateController_Success(t *testing.T) {
	controller, service, rr := setupPartyController()
	jwt.GetIdFromJWTFunc = func(string) (uint, error) { return 1, nil }
	service.On("Create", mock.Anything, uint(1)).Return(api.Success(nil))

	body, _ := json.Marshal(domains.PartyDTO{})
	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(body))
	req.Header.Set("Authorization", "token")

	controller.Create(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	service.AssertExpectations(t)
}

func TestPartyController_CreateController_InvalidBody(t *testing.T) {
	controller, _, rr := setupPartyController()
	req, _ := http.NewRequest("POST", "/", bytes.NewBufferString("invalid"))
	controller.Create(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestPartyController_CreateController_InvalidJWT(t *testing.T) {
	controller, _, rr := setupPartyController()
	jwt.GetIdFromJWTFunc = func(string) (uint, error) { return 0, errors.New("invalid") }

	body, _ := json.Marshal(domains.PartyDTO{})
	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(body))
	req.Header.Set("Authorization", "token")

	controller.Create(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestPartyController_GetController_Success(t *testing.T) {
	controller, service, rr := setupPartyController()
	jwt.GetIdFromJWTFunc = func(string) (uint, error) { return 1, nil }
	service.On("Get", uint(1), uint(1)).Return(api.Success(nil))

	req, _ := http.NewRequest("GET", "/1", nil)
	req.Header.Set("Authorization", "token")
	req = mux.SetURLVars(req, map[string]string{"id": "1"})

	controller.Get(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	service.AssertExpectations(t)
}

func TestPartyController_GetController_InvalidID(t *testing.T) {
	controller, _, rr := setupPartyController()
	req, _ := http.NewRequest("GET", "/invalid", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "invalid"})
	controller.Get(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestPartyController_GetController_InvalidJWT(t *testing.T) {
	controller, _, rr := setupPartyController()
	jwt.GetIdFromJWTFunc = func(string) (uint, error) { return 0, errors.New("invalid") }

	req, _ := http.NewRequest("GET", "/1", nil)
	req.Header.Set("Authorization", "token")
	req = mux.SetURLVars(req, map[string]string{"id": "1"})

	controller.Get(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestPartyController_UpdateController_Success(t *testing.T) {
	controller, service, rr := setupPartyController()
	jwt.GetIdFromJWTFunc = func(string) (uint, error) { return 1, nil }
	service.On("Update", mock.Anything, uint(1)).Return(api.Success(nil))

	body, _ := json.Marshal(domains.PartyDTO{})
	req, _ := http.NewRequest("PUT", "/", bytes.NewBuffer(body))
	req.Header.Set("Authorization", "token")

	controller.Update(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	service.AssertExpectations(t)
}

func TestPartyController_UpdateController_InvalidBody(t *testing.T) {
	controller, _, rr := setupPartyController()
	req, _ := http.NewRequest("PUT", "/", bytes.NewBufferString("invalid"))
	controller.Update(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestPartyController_UpdateController_InvalidJWT(t *testing.T) {
	controller, _, rr := setupPartyController()
	jwt.GetIdFromJWTFunc = func(string) (uint, error) { return 0, errors.New("invalid") }

	body, _ := json.Marshal(domains.PartyDTO{})
	req, _ := http.NewRequest("PUT", "/", bytes.NewBuffer(body))
	req.Header.Set("Authorization", "token")

	controller.Update(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestPartyController_DeleteController_Success(t *testing.T) {
	controller, service, rr := setupPartyController()
	jwt.GetIdFromJWTFunc = func(string) (uint, error) { return 1, nil }
	service.On("Delete", uint(1), uint(1)).Return(api.Success(nil))

	req, _ := http.NewRequest("DELETE", "/1", nil)
	req.Header.Set("Authorization", "token")
	req = mux.SetURLVars(req, map[string]string{"id": "1"})

	controller.Delete(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	service.AssertExpectations(t)
}

func TestPartyController_DeleteController_InvalidID(t *testing.T) {
	controller, _, rr := setupPartyController()
	req, _ := http.NewRequest("DELETE", "/invalid", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "invalid"})
	controller.Delete(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestPartyController_DeleteController_InvalidJWT(t *testing.T) {
	controller, _, rr := setupPartyController()
	jwt.GetIdFromJWTFunc = func(string) (uint, error) { return 0, errors.New("invalid") }

	req, _ := http.NewRequest("DELETE", "/1", nil)
	req.Header.Set("Authorization", "token")
	req = mux.SetURLVars(req, map[string]string{"id": "1"})

	controller.Delete(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestPartyController_GetPublicParties_Success(t *testing.T) {
	controller, service, rr := setupPartyController()
	service.On("GetPublicParties").Return(api.Success(nil))

	req, _ := http.NewRequest("GET", "/public", nil)
	controller.GetPublicParties(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	service.AssertExpectations(t)
}

func TestPartyController_GetPublicParty_Success(t *testing.T) {
	controller, service, rr := setupPartyController()
	service.On("GetPublicParty", uint(1)).Return(api.Success(nil))

	req, _ := http.NewRequest("GET", "/public/1", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	controller.GetPublicParty(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	service.AssertExpectations(t)
}

func TestPartyController_GetPublicParty_InvalidID(t *testing.T) {
	controller, _, rr := setupPartyController()
	req, _ := http.NewRequest("GET", "/public/invalid", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "invalid"})
	controller.GetPublicParty(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestPartyController_GetPartiesByOrganizerId_Success(t *testing.T) {
	controller, service, rr := setupPartyController()
	jwt.GetIdFromJWTFunc = func(string) (uint, error) { return 1, nil }
	service.On("GetPartiesByOrganizerId", uint(1)).Return(api.Success(nil))

	req, _ := http.NewRequest("GET", "/organizer", nil)
	req.Header.Set("Authorization", "token")
	controller.GetPartiesByOrganizerId(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	service.AssertExpectations(t)
}

func TestPartyController_GetPartiesByOrganizerId_InvalidJWT(t *testing.T) {
	controller, _, rr := setupPartyController()
	jwt.GetIdFromJWTFunc = func(string) (uint, error) { return 0, errors.New("invalid") }

	req, _ := http.NewRequest("GET", "/organizer", nil)
	req.Header.Set("Authorization", "token")
	controller.GetPartiesByOrganizerId(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestPartyController_GetPartiesByParticipantId_Success(t *testing.T) {
	controller, service, rr := setupPartyController()
	jwt.GetIdFromJWTFunc = func(string) (uint, error) { return 1, nil }
	service.On("GetPartiesByParticipantId", uint(1)).Return(api.Success(nil))

	req, _ := http.NewRequest("GET", "/participant", nil)
	req.Header.Set("Authorization", "token")
	controller.GetPartiesByParticipantId(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	service.AssertExpectations(t)
}

func TestPartyController_GetPartiesByParticipantId_InvalidJWT(t *testing.T) {
	controller, _, rr := setupPartyController()
	jwt.GetIdFromJWTFunc = func(string) (uint, error) { return 0, errors.New("invalid") }

	req, _ := http.NewRequest("GET", "/participant", nil)
	req.Header.Set("Authorization", "token")
	controller.GetPartiesByParticipantId(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestPartyController_GetParticipants_Success(t *testing.T) {
	controller, service, rr := setupPartyController()
	jwt.GetIdFromJWTFunc = func(string) (uint, error) { return 1, nil }
	service.On("GetParticipants", uint(1), uint(1)).Return(api.Success(nil))

	req, _ := http.NewRequest("GET", "/1/participants", nil)
	req.Header.Set("Authorization", "token")
	req = mux.SetURLVars(req, map[string]string{"party_id": "1"})
	controller.GetParticipants(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	service.AssertExpectations(t)
}

func TestPartyController_GetParticipants_InvalidID(t *testing.T) {
	controller, _, rr := setupPartyController()
	req, _ := http.NewRequest("GET", "/invalid/participants", nil)
	req = mux.SetURLVars(req, map[string]string{"party_id": "invalid"})
	controller.GetParticipants(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestPartyController_GetParticipants_InvalidJWT(t *testing.T) {
	controller, _, rr := setupPartyController()
	jwt.GetIdFromJWTFunc = func(string) (uint, error) { return 0, errors.New("invalid") }

	req, _ := http.NewRequest("GET", "/1/participants", nil)
	req.Header.Set("Authorization", "token")
	req = mux.SetURLVars(req, map[string]string{"party_id": "1"})
	controller.GetParticipants(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}
