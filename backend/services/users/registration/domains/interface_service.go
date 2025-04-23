package domains

import "github.com/zsomborCzaban/party_organizer/utils/api"

type IRegistrationService interface {
	Register(registrationRequest RegistrationRequest) api.IResponse
	ConfirmEmail(username, confirmHash string) api.IResponse
	ResendConfirmEmail(username string) api.IResponse
}
