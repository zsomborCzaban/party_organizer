package usecases

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/zsomborCzaban/party_organizer/common"
	registrationDomains "github.com/zsomborCzaban/party_organizer/services/users/registration/domains"
	"github.com/zsomborCzaban/party_organizer/services/users/user/domains"
	"github.com/zsomborCzaban/party_organizer/utils/api"
	"github.com/zsomborCzaban/party_organizer/utils/email"
	"github.com/zsomborCzaban/party_organizer/utils/repo"
	s3Wrapper "github.com/zsomborCzaban/party_organizer/utils/s3"
	"gopkg.in/gomail.v2"
	"mime/multipart"
	"os"
	"time"
)

type UserService struct {
	Validator              api.IValidator
	UserRepository         domains.IUserRepository
	RegistrationRepository registrationDomains.IRegistrationRepository
	S3ClientWrapper        s3Wrapper.IS3ClientWrapper
}

func NewUserService(repoCollector *repo.RepoCollector, validator api.IValidator, s3Wrapper s3Wrapper.IS3ClientWrapper) *UserService {
	return &UserService{
		UserRepository:         repoCollector.UserRepo,
		RegistrationRepository: repoCollector.RegistrationRepo,
		Validator:              validator,
		S3ClientWrapper:        s3Wrapper,
	}
}

func (us *UserService) Login(loginRequest domains.LoginRequest) api.IResponse {
	if err1 := us.Validator.Validate(loginRequest); err1 != nil {
		return api.ErrorValidation(err1.Errors)
	}

	_, err := us.RegistrationRepository.FindByUsername(*loginRequest.Username)
	if err == nil {
		return api.ErrorUnauthorized(domains.EmailNotConfirmed)
	}

	user, err3 := us.UserRepository.FindByUsername(*loginRequest.Username)
	if err3 != nil {
		return api.ErrorInvalidCredentials()
	}

	if errPassword := loginRequest.CheckPassword(user.Password); errPassword != nil {
		return api.ErrorInvalidCredentials()
	}

	jwt, err4 := user.GenerateJWT()
	if err4 != nil {
		return api.ErrorBadRequest("error while generating jwt")
	}

	return api.Success(
		domains.JWTData{Jwt: *jwt},
	)
}

func (us *UserService) AddFriend(friendId, userId uint) api.IResponse {
	//wont be used by user
	if friendId == userId {
		return api.ErrorBadRequest("You cannot be friends with yourself")
	} //unnecessary but good for safety

	user, err := us.UserRepository.FindById(userId)
	if err != nil {
		return api.ErrorBadRequest(err.Error())
	}

	friend, err := us.UserRepository.FindById(friendId)
	if err != nil {
		return api.ErrorBadRequest(err.Error())
	}

	if err3 := us.UserRepository.AddFriend(user, friend); err3 != nil {
		return api.ErrorInternalServerError(err3)
	}

	return api.Success("friend_added")
}

func (us *UserService) GetFriends(userId uint) api.IResponse {
	user, err := us.UserRepository.FindById(userId, "Friends")
	if err != nil {
		return api.ErrorInternalServerError(err)
	}

	friendDTOs := []domains.UserDTO{}
	for _, friend := range user.Friends {
		friendDTOs = append(friendDTOs, *friend.TransformToUserDTO())
	}

	return api.Success(friendDTOs)
}

func (us *UserService) UploadProfilePicture(userId uint, file multipart.File, fileHeader *multipart.FileHeader) api.IResponse {
	user, err := us.UserRepository.FindById(userId)
	if err != nil {
		return api.ErrorInternalServerError(err.Error())
	}

	defer file.Close()

	buffer := make([]byte, fileHeader.Size)
	_, err2 := file.Read(buffer)
	if err2 != nil {
		api.ErrorBadRequest(err2.Error())
	}

	key := fmt.Sprintf("users/%d/profile_pictures/%s", userId, fileHeader.Filename)
	contentType := fileHeader.Header.Get("Content-Type")
	bucketName, exists := os.LookupEnv("AWS_BUCKET_NAME")
	if !exists {
		bucketName = ""
		return api.ErrorInternalServerError("AWS_BUCKET_NAME environment variable is not set.")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_, err3 := us.S3ClientWrapper.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(bucketName),
		Key:         aws.String(key),
		Body:        bytes.NewReader(buffer),
		ContentType: aws.String(contentType),
	})
	if err3 != nil {
		return api.ErrorInternalServerError(err3.Error())
	}

	user.ProfilePictureUrl = fmt.Sprintf("https://%s.s3.amazonaws.com/%s", bucketName, key)
	err4 := us.UserRepository.UpdateUser(user)
	if err4 != nil {
		return api.ErrorInternalServerError(err4.Error())
	}

	return api.Success(user.TransformToUserDTO())
}

func (us *UserService) ForgotPassword(username string) api.IResponse {
	user, err := us.UserRepository.FindByUsername(username)
	if err != nil {
		return api.ErrorBadRequest(err.Error())
	}

	jwt, err2 := user.GenerateJWTForPasswordChange()
	if err2 != nil {
		return api.ErrorBadRequest("error while generating jwt")
	}

	emailUsername := os.Getenv(common.EMAIL_USERNAME_ENV_KEY)
	emailPassword := os.Getenv(common.EMAIL_PASSWORD_ENV_KEY)
	emailFull := os.Getenv(common.EMAIL_FULL_ENV_KEY)
	frontendUrl := os.Getenv(common.FRONTEND_URL_ENV_KEY)
	frontendLink := frontendUrl + "/changePassword?xzs=" + *jwt

	m := gomail.NewMessage()
	m.SetHeader("From", emailFull)
	m.SetHeader("To", user.Email)
	m.SetHeader("Subject", "Forgot Password!")
	m.SetBody("text/html", email.ParseForgotPasswordEmailBody(frontendLink))

	d := gomail.NewDialer("smtp.gmail.com", 587, emailUsername, emailPassword)

	if err3 := d.DialAndSend(m); err3 != nil {
		return api.ErrorInternalServerError(err3.Error())
	}
	return api.Success("Check your emails to change your password")
}

func (us *UserService) ChangePassword(req domains.ChangePasswordRequest, userId uint) api.IResponse {
	err := us.Validator.Validate(req)
	if err != nil {
		return api.ErrorValidation(err.Errors)
	}

	user, err2 := us.UserRepository.FindById(userId)
	if err2 != nil {
		return api.ErrorBadRequest(err2.Error())
	}

	user.Password = req.Password
	err3 := us.UserRepository.UpdateUser(user)
	if err3 != nil {
		return api.ErrorInternalServerError(err3.Error())
	}

	return api.Success("Password changed")

}
