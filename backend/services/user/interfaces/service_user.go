package interfaces

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/zsomborCzaban/party_organizer/common/api"
	"github.com/zsomborCzaban/party_organizer/services/user/domains"
	"mime/multipart"
	"net/http"
	"os"
	"time"
)

type UserService struct {
	UserRepository domains.IUserRepository
	Validator      api.IValidator
	S3Client       *s3.Client
}

func NewUserService(userRepository domains.IUserRepository, validator api.IValidator, s3 *s3.Client) *UserService {
	return &UserService{
		UserRepository: userRepository,
		Validator:      validator,
		S3Client:       s3,
	}
}

func (us *UserService) Login(loginRequest domains.LoginRequest) api.IResponse {
	if err1 := us.Validator.Validate(loginRequest); err1 != nil {
		return api.ErrorValidation(err1.Errors)
	}

	user, err2 := us.UserRepository.FindByUsername(*loginRequest.Username)
	if err2 != nil {
		return api.ErrorInvalidCredentials()
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

func (us *UserService) Register(registerRequest domains.RegisterRequest) api.IResponse {
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
	//todo: parse to dto to avoid password leakingy
	users, err := us.UserRepository.GetFriends(userId)
	if err != nil {
		return api.ErrorInternalServerError(err)
	}

	return api.Success(users)
}

func (us *UserService) UploadProfilePicture(file multipart.File, fileHeader *multipart.FileHeader) api.IResponse {
	defer file.Close()

	buffer := make([]byte, fileHeader.Size)
	_, err := file.Read(buffer)
	if err != nil {
		api.ErrorBadRequest(err.Error())
	}

	key := "user/profile_pictures/" + fileHeader.Filename
	contentType := fileHeader.Header.Get("Content-Type")
	bucketName, exists := os.LookupEnv("AWS_BUCKET_NAME")
	if !exists {
		bucketName = ""
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_, err2 := us.S3Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(bucketName),
		Key:         aws.String(key),
		Body:        bytes.NewReader(buffer),
		ContentType: aws.String(contentType),
	})
	if err2 != nil {
		return api.ErrorInternalServerError(err2.Error())
	}

	return api.Success(fmt.Sprintf("upload success, https://%s.s3.amazonaws.com/%s", bucketName, key))
}
