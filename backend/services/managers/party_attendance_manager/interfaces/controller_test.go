package interfaces

import (
	"errors"
	"github.com/zsomborCzaban/party_organizer/services/managers/party_attendance_manager/usecases"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/zsomborCzaban/party_organizer/services/managers/party_attendance_manager/domains"
	"github.com/zsomborCzaban/party_organizer/utils/api"
	"github.com/zsomborCzaban/party_organizer/utils/jwt"
)

func setupPartyInviteController() (domains.IPartyInviteController, *usecases.MockService, *httptest.ResponseRecorder) {
	service := new(usecases.MockService)
	controller := NewPartyInviteController(service)
	rr := httptest.NewRecorder()
	return controller, service, rr
}

func TestPartyInviteController_Accept_Success(t *testing.T) {
	controller, service, rr := setupPartyInviteController()
	jwt.GetIdFromJWTFunc = func(string) (uint, error) { return 1, nil }
	service.On("Accept", uint(1), uint(2)).Return(api.Success(nil))

	req, _ := http.NewRequest("POST", "/parties/2/accept", nil)
	req.Header.Set("Authorization", "token")
	req = mux.SetURLVars(req, map[string]string{"party_id": "2"})

	controller.Accept(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	service.AssertExpectations(t)
}

func TestPartyInviteController_Accept_InvalidPartyID(t *testing.T) {
	controller, _, rr := setupPartyInviteController()
	req, _ := http.NewRequest("POST", "/parties/invalid/accept", nil)
	req.Header.Set("Authorization", "token")
	req = mux.SetURLVars(req, map[string]string{"party_id": "invalid"})

	controller.Accept(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestPartyInviteController_Accept_InvalidJWT(t *testing.T) {
	controller, _, rr := setupPartyInviteController()
	jwt.GetIdFromJWTFunc = func(string) (uint, error) { return 0, errors.New("invalid") }

	req, _ := http.NewRequest("POST", "/parties/2/accept", nil)
	req.Header.Set("Authorization", "token")
	req = mux.SetURLVars(req, map[string]string{"party_id": "2"})

	controller.Accept(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestPartyInviteController_Decline_Success(t *testing.T) {
	controller, service, rr := setupPartyInviteController()
	jwt.GetIdFromJWTFunc = func(string) (uint, error) { return 1, nil }
	service.On("Decline", uint(1), uint(2)).Return(api.Success(nil))

	req, _ := http.NewRequest("POST", "/parties/2/decline", nil)
	req.Header.Set("Authorization", "token")
	req = mux.SetURLVars(req, map[string]string{"party_id": "2"})

	controller.Decline(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	service.AssertExpectations(t)
}

func TestPartyInviteController_Decline_InvalidPartyID(t *testing.T) {
	controller, _, rr := setupPartyInviteController()

	req, _ := http.NewRequest("POST", "/parties/invalid/decline", nil)
	req.Header.Set("Authorization", "token")
	req = mux.SetURLVars(req, map[string]string{"party_id": "invalid"})

	controller.Decline(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestPartyInviteController_Decline_InvalidJWT(t *testing.T) {
	controller, _, rr := setupPartyInviteController()
	jwt.GetIdFromJWTFunc = func(string) (uint, error) { return 0, errors.New("invalid token") }

	req, _ := http.NewRequest("POST", "/parties/2/decline", nil)
	req.Header.Set("Authorization", "token")
	req = mux.SetURLVars(req, map[string]string{"party_id": "2"})

	controller.Decline(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestPartyInviteController_Invite_Success(t *testing.T) {
	controller, service, rr := setupPartyInviteController()
	jwt.GetIdFromJWTFunc = func(string) (uint, error) { return 1, nil }
	service.On("Invite", "testuser", uint(1), uint(2)).Return(api.Success(nil))

	req, _ := http.NewRequest("POST", "/parties/2/invite/testuser", nil)
	req.Header.Set("Authorization", "token")
	req = mux.SetURLVars(req, map[string]string{
		"party_id":         "2",
		"invited_username": "testuser",
	})

	controller.Invite(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	service.AssertExpectations(t)
}

func TestPartyInviteController_Invite_InvalidUsername(t *testing.T) {
	controller, _, rr := setupPartyInviteController()
	req, _ := http.NewRequest("POST", "/parties/2/invite/", nil)
	req.Header.Set("Authorization", "token")
	req = mux.SetURLVars(req, map[string]string{
		"party_id": "2",
	})

	controller.Invite(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestPartyInviteController_Invite_InvalidPartyID(t *testing.T) {
	controller, _, rr := setupPartyInviteController()

	req, _ := http.NewRequest("POST", "/parties/invalid/invite/testuser", nil)
	req.Header.Set("Authorization", "token")
	req = mux.SetURLVars(req, map[string]string{
		"party_id":         "invalid",
		"invited_username": "testuser",
	})

	controller.Invite(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestPartyInviteController_Invite_InvalidJWT(t *testing.T) {
	controller, _, rr := setupPartyInviteController()
	jwt.GetIdFromJWTFunc = func(string) (uint, error) { return 0, errors.New("invalid token") }

	req, _ := http.NewRequest("POST", "/parties/2/invite/testuser", nil)
	req.Header.Set("Authorization", "token")
	req = mux.SetURLVars(req, map[string]string{
		"party_id":         "2",
		"invited_username": "testuser",
	})

	controller.Invite(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestPartyInviteController_GetUserPendingInvites_Success(t *testing.T) {
	controller, service, rr := setupPartyInviteController()
	jwt.GetIdFromJWTFunc = func(string) (uint, error) { return 1, nil }
	service.On("GetUserPendingInvites", uint(1)).Return(api.Success(nil))

	req, _ := http.NewRequest("GET", "/invites/pending", nil)
	req.Header.Set("Authorization", "token")

	controller.GetUserPendingInvites(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	service.AssertExpectations(t)
}

func TestPartyInviteController_GetUserPendingInvites_InvalidJWT(t *testing.T) {
	controller, _, rr := setupPartyInviteController()
	jwt.GetIdFromJWTFunc = func(string) (uint, error) { return 0, errors.New("invalid token") }

	req, _ := http.NewRequest("GET", "/invites/pending", nil)
	req.Header.Set("Authorization", "token")

	controller.GetUserPendingInvites(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestPartyInviteController_GetPartyPendingInvites_Success(t *testing.T) {
	controller, service, rr := setupPartyInviteController()
	jwt.GetIdFromJWTFunc = func(string) (uint, error) { return 1, nil }
	service.On("GetPartyPendingInvites", uint(2), uint(1)).Return(api.Success(nil))

	req, _ := http.NewRequest("GET", "/parties/2/pending-invites", nil)
	req.Header.Set("Authorization", "token")
	req = mux.SetURLVars(req, map[string]string{"party_id": "2"})

	controller.GetPartyPendingInvites(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	service.AssertExpectations(t)
}

func TestPartyInviteController_GetPartyPendingInvites_InvalidPartyID(t *testing.T) {
	controller, _, rr := setupPartyInviteController()

	req, _ := http.NewRequest("GET", "/parties/invalid/pending-invites", nil)
	req.Header.Set("Authorization", "token")
	req = mux.SetURLVars(req, map[string]string{"party_id": "invalid"})

	controller.GetPartyPendingInvites(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestPartyInviteController_GetPartyPendingInvites_InvalidJWT(t *testing.T) {
	controller, _, rr := setupPartyInviteController()
	jwt.GetIdFromJWTFunc = func(string) (uint, error) { return 0, errors.New("invalid token") }

	req, _ := http.NewRequest("GET", "/parties/2/pending-invites", nil)
	req.Header.Set("Authorization", "token")
	req = mux.SetURLVars(req, map[string]string{"party_id": "2"})

	controller.GetPartyPendingInvites(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestPartyInviteController_Kick_Success(t *testing.T) {
	controller, service, rr := setupPartyInviteController()
	jwt.GetIdFromJWTFunc = func(string) (uint, error) { return 1, nil }
	service.On("Kick", uint(3), uint(1), uint(2)).Return(api.Success(nil))

	req, _ := http.NewRequest("POST", "/parties/2/kick/3", nil)
	req.Header.Set("Authorization", "token")
	req = mux.SetURLVars(req, map[string]string{
		"party_id":  "2",
		"kicked_id": "3",
	})

	controller.Kick(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	service.AssertExpectations(t)
}

func TestPartyInviteController_Kick_InvalidPartyID(t *testing.T) {
	controller, _, rr := setupPartyInviteController()

	req, _ := http.NewRequest("POST", "/parties/invalid/kick/3", nil)
	req.Header.Set("Authorization", "token")
	req = mux.SetURLVars(req, map[string]string{
		"party_id":  "invalid",
		"kicked_id": "3",
	})

	controller.Kick(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestPartyInviteController_Kick_InvalidKickedID(t *testing.T) {
	controller, _, rr := setupPartyInviteController()

	req, _ := http.NewRequest("POST", "/parties/2/kick/invalid", nil)
	req.Header.Set("Authorization", "token")
	req = mux.SetURLVars(req, map[string]string{
		"party_id":  "2",
		"kicked_id": "invalid",
	})

	controller.Kick(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestPartyInviteController_Kick_InvalidJWT(t *testing.T) {
	controller, _, rr := setupPartyInviteController()
	jwt.GetIdFromJWTFunc = func(string) (uint, error) { return 0, errors.New("invalid token") }

	req, _ := http.NewRequest("POST", "/parties/2/kick/3", nil)
	req.Header.Set("Authorization", "token")
	req = mux.SetURLVars(req, map[string]string{
		"party_id":  "2",
		"kicked_id": "3",
	})

	controller.Kick(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestPartyInviteController_LeaveParty_Success(t *testing.T) {
	controller, service, rr := setupPartyInviteController()
	jwt.GetIdFromJWTFunc = func(string) (uint, error) { return 1, nil }
	service.On("Kick", uint(1), uint(1), uint(2)).Return(api.Success(nil))

	req, _ := http.NewRequest("POST", "/parties/2/leave", nil)
	req.Header.Set("Authorization", "token")
	req = mux.SetURLVars(req, map[string]string{"party_id": "2"})

	controller.LeaveParty(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	service.AssertExpectations(t)
}

func TestPartyInviteController_LeaveParty_InvalidPartyID(t *testing.T) {
	controller, _, rr := setupPartyInviteController()

	req, _ := http.NewRequest("POST", "/parties/invalid/leave", nil)
	req.Header.Set("Authorization", "token")
	req = mux.SetURLVars(req, map[string]string{"party_id": "invalid"})

	controller.LeaveParty(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestPartyInviteController_LeaveParty_InvalidJWT(t *testing.T) {
	controller, _, rr := setupPartyInviteController()
	jwt.GetIdFromJWTFunc = func(string) (uint, error) { return 0, errors.New("invalid token") }

	req, _ := http.NewRequest("POST", "/parties/2/leave", nil)
	req.Header.Set("Authorization", "token")
	req = mux.SetURLVars(req, map[string]string{"party_id": "2"})

	controller.LeaveParty(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestPartyInviteController_JoinPublicParty_Success(t *testing.T) {
	controller, service, rr := setupPartyInviteController()
	jwt.GetIdFromJWTFunc = func(string) (uint, error) { return 1, nil }
	service.On("JoinPublicParty", uint(2), uint(1)).Return(api.Success(nil))

	req, _ := http.NewRequest("POST", "/parties/2/join", nil)
	req.Header.Set("Authorization", "token")
	req = mux.SetURLVars(req, map[string]string{"party_id": "2"})

	controller.JoinPublicParty(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	service.AssertExpectations(t)
}

func TestPartyInviteController_JoinPublicParty_InvalidPartyID(t *testing.T) {
	controller, _, rr := setupPartyInviteController()

	req, _ := http.NewRequest("POST", "/parties/invalid/join", nil)
	req.Header.Set("Authorization", "token")
	req = mux.SetURLVars(req, map[string]string{"party_id": "invalid"})

	controller.JoinPublicParty(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestPartyInviteController_JoinPublicParty_InvalidJWT(t *testing.T) {
	controller, _, rr := setupPartyInviteController()
	jwt.GetIdFromJWTFunc = func(string) (uint, error) { return 0, errors.New("invalid token") }

	req, _ := http.NewRequest("POST", "/parties/2/join", nil)
	req.Header.Set("Authorization", "token")
	req = mux.SetURLVars(req, map[string]string{"party_id": "2"})

	controller.JoinPublicParty(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestPartyInviteController_JoinPrivateParty_Success(t *testing.T) {
	controller, service, rr := setupPartyInviteController()
	jwt.GetIdFromJWTFunc = func(string) (uint, error) { return 1, nil }
	service.On("JoinPrivateParty", uint(1), "secret123").Return(api.Success(nil))

	req, _ := http.NewRequest("POST", "/parties/private/secret123", nil)
	req.Header.Set("Authorization", "token")
	req = mux.SetURLVars(req, map[string]string{"access_code": "secret123"})

	controller.JoinPrivateParty(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	service.AssertExpectations(t)
}

func TestPartyInviteController_JoinPrivateParty_MissingAccessCode(t *testing.T) {
	controller, _, rr := setupPartyInviteController()
	req, _ := http.NewRequest("POST", "/parties/private/", nil)
	req.Header.Set("Authorization", "token")
	req = mux.SetURLVars(req, map[string]string{"access_code": ""})

	controller.JoinPrivateParty(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestPartyInviteController_JoinPrivateParty_InvalidJWT(t *testing.T) {
	controller, _, rr := setupPartyInviteController()
	jwt.GetIdFromJWTFunc = func(string) (uint, error) { return 0, errors.New("invalid token") }

	req, _ := http.NewRequest("POST", "/parties/private/secret123", nil)
	req.Header.Set("Authorization", "token")
	req = mux.SetURLVars(req, map[string]string{"access_code": "secret123"})

	controller.JoinPrivateParty(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}
