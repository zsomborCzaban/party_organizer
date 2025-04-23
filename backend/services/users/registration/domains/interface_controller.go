package domains

import "net/http"

type IRegistrationController interface {
	Register(w http.ResponseWriter, r *http.Request)
	ConfirmEmail(w http.ResponseWriter, r *http.Request)
	ResendConfirmEmail(w http.ResponseWriter, r *http.Request)
}
