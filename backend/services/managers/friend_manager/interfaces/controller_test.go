package interfaces

import (
	"errors"
	"github.com/zsomborCzaban/party_organizer/services/managers/friend_manager/usecases"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/zsomborCzaban/party_organizer/services/managers/friend_manager/domains"
	"github.com/zsomborCzaban/party_organizer/utils/api"
	"github.com/zsomborCzaban/party_organizer/utils/jwt"
)

func setupFriendInviteController() (domains.IFriendInviteController, *usecases.MockService, *httptest.ResponseRecorder) {
	service := new(usecases.MockService)
	controller := NewFriendInviteController(service)
	rr := httptest.NewRecorder()
	return controller, service, rr
}

func TestFriendInviteController_Invite_Success(t *testing.T) {
	controller, service, rr := setupFriendInviteController()
	jwt.GetIdFromJWTFunc = func(string) (uint, error) { return 1, nil }
	service.On("Invite", "testuser", uint(1)).Return(api.Success(nil))

	req, _ := http.NewRequest("POST", "/users/testuser/invite", nil)
	req.Header.Set("Authorization", "token")
	req = mux.SetURLVars(req, map[string]string{"username": "testuser"})

	controller.Invite(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	service.AssertExpectations(t)
}

func TestFriendInviteController_Invite_InvalidUsername(t *testing.T) {
	controller, _, rr := setupFriendInviteController()
	req, _ := http.NewRequest("POST", "/users//invite", nil)
	req.Header.Set("Authorization", "token")

	controller.Invite(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestFriendInviteController_Invite_InvalidJWT(t *testing.T) {
	controller, _, rr := setupFriendInviteController()
	jwt.GetIdFromJWTFunc = func(string) (uint, error) { return 0, errors.New("invalid") }

	req, _ := http.NewRequest("POST", "/users/testuser/invite", nil)
	req.Header.Set("Authorization", "token")
	req = mux.SetURLVars(req, map[string]string{"username": "testuser"})

	controller.Invite(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestFriendInviteController_Accept_Success(t *testing.T) {
	controller, service, rr := setupFriendInviteController()
	jwt.GetIdFromJWTFunc = func(string) (uint, error) { return 1, nil }
	service.On("Accept", uint(2), uint(1)).Return(api.Success(nil))

	req, _ := http.NewRequest("POST", "/invites/2/accept", nil)
	req.Header.Set("Authorization", "token")
	req = mux.SetURLVars(req, map[string]string{"invitor_id": "2"})

	controller.Accept(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	service.AssertExpectations(t)
}

func TestFriendInviteController_Accept_InvalidInvitorID(t *testing.T) {
	controller, _, rr := setupFriendInviteController()
	req, _ := http.NewRequest("POST", "/invites/invalid/accept", nil)
	req.Header.Set("Authorization", "token")
	req = mux.SetURLVars(req, map[string]string{"invitor_id": "invalid"})

	controller.Accept(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestFriendInviteController_Accept_InvalidJWT(t *testing.T) {
	controller, _, rr := setupFriendInviteController()
	jwt.GetIdFromJWTFunc = func(string) (uint, error) { return 0, errors.New("invalid") }

	req, _ := http.NewRequest("POST", "/invites/2/accept", nil)
	req.Header.Set("Authorization", "token")
	req = mux.SetURLVars(req, map[string]string{"invitor_id": "2"})

	controller.Accept(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestFriendInviteController_Decline_Success(t *testing.T) {
	controller, service, rr := setupFriendInviteController()
	jwt.GetIdFromJWTFunc = func(string) (uint, error) { return 1, nil }
	service.On("Decline", uint(2), uint(1)).Return(api.Success(nil))

	req, _ := http.NewRequest("POST", "/invites/2/decline", nil)
	req.Header.Set("Authorization", "token")
	req = mux.SetURLVars(req, map[string]string{"invitor_id": "2"})

	controller.Decline(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	service.AssertExpectations(t)
}

func TestFriendInviteController_Decline_InvalidInvitorID(t *testing.T) {
	controller, _, rr := setupFriendInviteController()
	req, _ := http.NewRequest("POST", "/invites/invalid/decline", nil)
	req.Header.Set("Authorization", "token")
	req = mux.SetURLVars(req, map[string]string{"invitor_id": "invalid"})

	controller.Decline(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestFriendInviteController_Decline_InvalidJWT(t *testing.T) {
	controller, _, rr := setupFriendInviteController()
	jwt.GetIdFromJWTFunc = func(string) (uint, error) { return 0, errors.New("invalid") }

	req, _ := http.NewRequest("POST", "/invites/2/decline", nil)
	req.Header.Set("Authorization", "token")
	req = mux.SetURLVars(req, map[string]string{"invitor_id": "2"})

	controller.Decline(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestFriendInviteController_GetPendingInvites_Success(t *testing.T) {
	controller, service, rr := setupFriendInviteController()
	jwt.GetIdFromJWTFunc = func(string) (uint, error) { return 1, nil }
	service.On("GetPendingInvites", uint(1)).Return(api.Success(nil))

	req, _ := http.NewRequest("GET", "/invites/pending", nil)
	req.Header.Set("Authorization", "token")

	controller.GetPendingInvites(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	service.AssertExpectations(t)
}

func TestFriendInviteController_GetPendingInvites_InvalidJWT(t *testing.T) {
	controller, _, rr := setupFriendInviteController()
	jwt.GetIdFromJWTFunc = func(string) (uint, error) { return 0, errors.New("invalid") }

	req, _ := http.NewRequest("GET", "/invites/pending", nil)
	req.Header.Set("Authorization", "token")

	controller.GetPendingInvites(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestFriendInviteController_RemoveFriend_Success(t *testing.T) {
	controller, service, rr := setupFriendInviteController()
	jwt.GetIdFromJWTFunc = func(string) (uint, error) { return 1, nil }
	service.On("RemoveFriend", uint(1), uint(2)).Return(api.Success(nil))

	req, _ := http.NewRequest("DELETE", "/friends/2", nil)
	req.Header.Set("Authorization", "token")
	req = mux.SetURLVars(req, map[string]string{"friend_id": "2"})

	controller.RemoveFriend(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	service.AssertExpectations(t)
}

func TestFriendInviteController_RemoveFriend_InvalidFriendID(t *testing.T) {
	controller, _, rr := setupFriendInviteController()
	req, _ := http.NewRequest("DELETE", "/friends/invalid", nil)
	req.Header.Set("Authorization", "token")
	req = mux.SetURLVars(req, map[string]string{"friend_id": "invalid"})

	controller.RemoveFriend(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestFriendInviteController_RemoveFriend_InvalidJWT(t *testing.T) {
	controller, _, rr := setupFriendInviteController()
	jwt.GetIdFromJWTFunc = func(string) (uint, error) { return 0, errors.New("invalid") }

	req, _ := http.NewRequest("DELETE", "/friends/2", nil)
	req.Header.Set("Authorization", "token")
	req = mux.SetURLVars(req, map[string]string{"friend_id": "2"})

	controller.RemoveFriend(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}
