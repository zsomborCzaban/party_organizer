package usecases

import (
	"github.com/zsomborCzaban/party_organizer/common"
	"github.com/zsomborCzaban/party_organizer/services/users/registration/domains"
	domains2 "github.com/zsomborCzaban/party_organizer/services/users/user/domains"
	"github.com/zsomborCzaban/party_organizer/utils/api"
	"github.com/zsomborCzaban/party_organizer/utils/random"
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
		errorUserAlreadyExists.CollectValidationError("Username", "Username already taken", registrationRequest.Username)
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
		errorUserAlreadyExists.CollectValidationError("Username", "Username already taken", registrationRequest.Username)
		return api.Error(http.StatusBadRequest, errorUserAlreadyExists.Errors)
	}
	if err3.Error() != domains2.UserNotFound+registrationRequest.Username {
		return api.ErrorInternalServerError(err3.Error())
	}

	registrationRequest.ConfirmHash = random.GenerateRandomString(64)
	if err4 := rs.RegistrationRepository.Create(&registrationRequest); err4 != nil {
		return api.ErrorInternalServerError(err4.Error())
	}

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
	m.SetHeader("Subject", "Hello!") //todo: write email body
	m.SetBody("text/plain", "This is the email body")

	d := gomail.NewDialer("smtp.gmail.com", 587, username, password)

	if err := d.DialAndSend(m); err != nil {
		return api.ErrorInternalServerError(err.Error())
	}
	return api.Success("Registration success. To finish, confirm your email")
}

func (rs *RegistrationService) ConfirmEmail(username, confirmHash string) api.IResponse {
	_, err := rs.UserRepository.FindByUsername(username)
	if err == nil {
		return api.ErrorBadRequest("Email already confirmed. try logging in!")
	}
	if err.Error() != domains.UserNotFound+username {
		return api.ErrorInternalServerError(err.Error())
	}

	reg, err2 := rs.RegistrationRepository.FindByUsername(username)
	if err2 != nil {
		return api.ErrorBadRequest(err2.Error())
	}
	if reg.ConfirmHash != confirmHash {
		return api.ErrorBadRequest("Invalid confirm hash")
	}

	user, err3 := reg.TransformToUser()
	if err3 != nil {
		return api.ErrorInternalServerError(err3.Error())
	}

	err4 := rs.UserRepository.CreateUser(user)
	if err4 != nil {
		return api.ErrorInternalServerError(err4.Error())
	}

	return api.Success("Email confirmed")
}
