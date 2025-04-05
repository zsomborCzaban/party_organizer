package usecases

import (
	"github.com/zsomborCzaban/party_organizer/common"
	"github.com/zsomborCzaban/party_organizer/services/users/registration/domains"
	domains2 "github.com/zsomborCzaban/party_organizer/services/users/user/domains"
	"github.com/zsomborCzaban/party_organizer/utils/api"
	"github.com/zsomborCzaban/party_organizer/utils/repo"
	"gopkg.in/gomail.v2"
	"net/http"
	"os"
)

type RegistrationService struct {
	Validator              api.IValidator
	RegistrationRepository domains.IRegistrationRepository
	UserRepository         domains2.IUserRepository
}

func NewRegistrationService(repoCollector *repo.RepoCollector, validator api.IValidator) *RegistrationService {
	return &RegistrationService{
		RegistrationRepository: repoCollector.RegistrationRepo,
		UserRepository:         repoCollector.UserRepo,
		Validator:              validator,
	}
}

func (rs *RegistrationService) Register(registrationRequest domains.RegistrationRequest) api.IResponse {
	if err1 := rs.Validator.Validate(registrationRequest); err1 != nil {
		return api.ErrorValidation(err1.Errors)
	}

	reg, err2 := rs.RegistrationRepository.FindByUsername(registrationRequest.Username)
	if err2 == nil && reg.Email != registrationRequest.Email {
		errorUserAlreadyExists := api.NewValidationErrors()
		errorUserAlreadyExists.CollectValidationError("username", "username already taken", registrationRequest.Username)
		return api.Error(http.StatusBadRequest, errorUserAlreadyExists.Errors)
	}
	if err2 == nil && reg.Email == registrationRequest.Email {
		rs.SendConfirmEmail(registrationRequest)
		return api.Success("You already registered with that username and email, please confirm your email to finish")
	}
	if err2.Error() != domains2.UserNotFound+registrationRequest.Username {
		return api.ErrorInternalServerError(err2.Error())
	}

	_, err3 := rs.UserRepository.FindByUsername(registrationRequest.Username)
	if err3 == nil {
		errorUserAlreadyExists := api.NewValidationErrors()
		errorUserAlreadyExists.CollectValidationError("username", "username already taken", registrationRequest.Username)
		return api.Error(http.StatusBadRequest, errorUserAlreadyExists.Errors)
	}
	if err3.Error() != domains2.UserNotFound+registrationRequest.Username {
		return api.ErrorInternalServerError(err3.Error())
	}

	if err4 := rs.RegistrationRepository.Create(&registrationRequest); err4 != nil {
		return api.ErrorInternalServerError(err4.Error())
	}
	//todo: create confirm hash

	resp := rs.SendConfirmEmail(registrationRequest)
	if resp.GetCode() != http.StatusOK {
		rs.RegistrationRepository.Delete(&registrationRequest) //todo: handle error on delete
	}
	return resp
}

func (rs *RegistrationService) SendConfirmEmail(registerRequest domains.RegistrationRequest) api.IResponse {
	username := os.Getenv(common.EMAIL_USERNAME_ENV_KEY)
	password := os.Getenv(common.EMAIL_PASSWORD_ENV_KEY)
	email := os.Getenv(common.EMAIL_FULL_ENV_KEY)

	m := gomail.NewMessage()
	m.SetHeader("From", email)
	m.SetHeader("To", registerRequest.Email)
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/plain", "This is the email body")

	d := gomail.NewDialer("smtp.gmail.com", 587, username, password)

	if err := d.DialAndSend(m); err != nil {
		return api.ErrorInternalServerError(err.Error())
	}
	return api.Success("Registration success. To finish, confirm your email")
}

func (rs *RegistrationService) ConfirmEmail() api.IResponse {
	return nil
}
