package interfaces

import (
	"bytes"
	"encoding/json"
	"errors"
	domains2 "github.com/zsomborCzaban/party_organizer/services/users/user/domains"
	"github.com/zsomborCzaban/party_organizer/services/users/user/usecases"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zsomborCzaban/party_organizer/utils/api"
	"github.com/zsomborCzaban/party_organizer/utils/jwt"
)

func setupUserController() (domains2.IUserController, *usecases.MockService, *httptest.ResponseRecorder) {
	service := new(usecases.MockService)
	controller := NewUserController(service)
	rr := httptest.NewRecorder()
	return controller, service, rr
}

func TestUserController_Login_Success(t *testing.T) {
	controller, service, rr := setupUserController()
	service.On("Login", mock.AnythingOfType("domains.LoginRequest")).Return(api.Success(nil))

	username := "testuser"
	password := "password123"
	loginReq := domains2.LoginRequest{
		Username: &username,
		Password: &password,
	}
	body, _ := json.Marshal(loginReq)
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(body))

	controller.Login(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	service.AssertExpectations(t)
}

func TestUserController_Login_InvalidBody(t *testing.T) {
	controller, _, rr := setupUserController()
	req, _ := http.NewRequest("POST", "/login", bytes.NewBufferString("invalid"))

	controller.Login(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestUserController_AddFriend_Success(t *testing.T) {
	controller, service, rr := setupUserController()
	jwt.GetIdFromJWTFunc = func(string) (uint, error) { return 1, nil }
	service.On("AddFriend", uint(2), uint(1)).Return(api.Success(nil))

	req, _ := http.NewRequest("POST", "/friends/2", nil)
	req.Header.Set("Authorization", "token")
	req = mux.SetURLVars(req, map[string]string{"id": "2"})

	controller.AddFriend(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	service.AssertExpectations(t)
}

func TestUserController_AddFriend_InvalidID(t *testing.T) {
	controller, _, rr := setupUserController()
	req, _ := http.NewRequest("POST", "/friends/invalid", nil)
	req.Header.Set("Authorization", "token")
	req = mux.SetURLVars(req, map[string]string{"id": "invalid"})

	controller.AddFriend(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestUserController_AddFriend_InvalidJWT(t *testing.T) {
	controller, _, rr := setupUserController()
	jwt.GetIdFromJWTFunc = func(string) (uint, error) { return 0, errors.New("invalid") }

	req, _ := http.NewRequest("POST", "/friends/2", nil)
	req.Header.Set("Authorization", "token")
	req = mux.SetURLVars(req, map[string]string{"id": "2"})

	controller.AddFriend(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestUserController_GetFriends_Success(t *testing.T) {
	controller, service, rr := setupUserController()
	jwt.GetIdFromJWTFunc = func(string) (uint, error) { return 1, nil }
	service.On("GetFriends", uint(1)).Return(api.Success(nil))

	req, _ := http.NewRequest("GET", "/friends", nil)
	req.Header.Set("Authorization", "token")

	controller.GetFriends(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	service.AssertExpectations(t)
}

func TestUserController_GetFriends_InvalidJWT(t *testing.T) {
	controller, _, rr := setupUserController()
	jwt.GetIdFromJWTFunc = func(string) (uint, error) { return 0, errors.New("invalid") }

	req, _ := http.NewRequest("GET", "/friends", nil)
	req.Header.Set("Authorization", "token")

	controller.GetFriends(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestUserController_UploadProfilePicture_Success(t *testing.T) {
	controller, service, rr := setupUserController()
	jwt.GetIdFromJWTFunc = func(string) (uint, error) { return 1, nil }

	// Create a buffer to simulate file upload
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("image", "test.jpg")
	part.Write([]byte("test image content"))
	writer.Close()

	service.On("UploadProfilePicture", uint(1), mock.Anything, mock.Anything).Return(api.Success(nil))

	req, _ := http.NewRequest("POST", "/profile/picture", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Authorization", "token")

	controller.UploadProfilePicture(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	service.AssertExpectations(t)
}

func TestUserController_UploadProfilePicture_InvalidJWT(t *testing.T) {
	controller, _, rr := setupUserController()
	jwt.GetIdFromJWTFunc = func(string) (uint, error) { return 0, errors.New("invalid") }

	req, _ := http.NewRequest("POST", "/profile/picture", nil)
	req.Header.Set("Authorization", "token")

	controller.UploadProfilePicture(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestUserController_UploadProfilePicture_NoFile(t *testing.T) {
	controller, _, rr := setupUserController()
	jwt.GetIdFromJWTFunc = func(string) (uint, error) { return 1, nil }

	req, _ := http.NewRequest("POST", "/profile/picture", nil)
	req.Header.Set("Authorization", "token")

	controller.UploadProfilePicture(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestUserController_UploadProfilePicture_InvalidForm(t *testing.T) {
	controller, _, rr := setupUserController()
	jwt.GetIdFromJWTFunc = func(string) (uint, error) { return 1, nil }

	req, _ := http.NewRequest("POST", "/profile/picture", bytes.NewBufferString("invalid"))
	req.Header.Set("Content-Type", "multipart/form-data")
	req.Header.Set("Authorization", "token")

	controller.UploadProfilePicture(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestUserController_UploadProfilePicture_MissingImage(t *testing.T) {
	controller, service, rr := setupUserController()
	jwt.GetIdFromJWTFunc = func(string) (uint, error) { return 1, nil }

	// Create a buffer to simulate file upload
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("notImage", "test.jpg")
	part.Write([]byte("test image content"))
	writer.Close()

	req, _ := http.NewRequest("POST", "/profile/picture", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Authorization", "token")

	controller.UploadProfilePicture(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
	service.AssertExpectations(t)
}
