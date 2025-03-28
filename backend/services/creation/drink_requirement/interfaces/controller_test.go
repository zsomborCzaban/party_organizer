package interfaces

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/mock"
	"github.com/zsomborCzaban/party_organizer/services/creation/drink_requirement/domains"
	"github.com/zsomborCzaban/party_organizer/services/creation/drink_requirement/usecases"
	"github.com/zsomborCzaban/party_organizer/utils/api"
	"net/http"
	"net/http/httptest"
	"testing"
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
