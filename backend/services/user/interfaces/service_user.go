package interfaces

import (
	"github.com/zsomborCzaban/party_organizer/common/api"
	"github.com/zsomborCzaban/party_organizer/services/user/domains"
	"net/http"
)

type UserService struct {
	UserRepository domains.IUserRepository
	Validator      api.IValidator
}

func NewUserService(userRepository domains.IUserRepository, validator api.IValidator) *UserService {
	return &UserService{
		UserRepository: userRepository,
		Validator:      validator,
	}
}

func (us UserService) Login(loginRequest domains.LoginRequest) api.IResponse {
	if err1 := us.Validator.Validate(loginRequest); err1 != nil {
		return api.ErrorValidation(err1.Errors)
	}

	user, err2 := us.UserRepository.FindByUsername(*loginRequest.Username)
	if err2 != nil {
		return api.ErrorNotFound("username", *loginRequest.Username)
	}

	if errPassword := loginRequest.CheckPassword(user.Password); errPassword != nil {
		return api.ErrorInvalidCredentials()
	}

	userDTO := user.TransformToUserDTO()

	jwt, err3 := userDTO.GenerateJWT()
	if err3 != nil {
		return api.ErrorBadRequest("error while generating jwt")
	}

	return api.Success(
		domains.JWTData{Jwt: *jwt},
	)
}

func (us UserService) Register(registerRequest domains.RegisterRequest) api.IResponse {
	if err1 := us.Validator.Validate(registerRequest); err1 != nil {
		return api.ErrorValidation(err1.Errors)
	}

	duplicateUser, err2 := us.UserRepository.FindByUsername(registerRequest.Username)
	if duplicateUser != nil {
		errorUserAlreadyExists := api.NewValidationErrors()
		errorUserAlreadyExists.CollectValidationError("username", "username already taken", registerRequest.Username)
		return api.Error(http.StatusBadRequest, errorUserAlreadyExists.Errors)
	}
	if err2.Error() != domains.UserNotFound+registerRequest.Username {
		return api.ErrorInternalServerError(err2.Error())
	}

	user, err3 := registerRequest.TransformToUser()
	if err3 != nil {
		return api.ErrorInternalServerError(err3.Error())
	}

	if err4 := us.UserRepository.CreateUser(user); err4 != nil {
		return api.ErrorInternalServerError(err4.Error())
	}

	return api.Success("register_success")
}
