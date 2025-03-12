package usecases

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/zsomborCzaban/party_organizer/services/user/domains"
	"github.com/zsomborCzaban/party_organizer/utils/api"
	"github.com/zsomborCzaban/party_organizer/utils/repo"
	"mime/multipart"
	"net/http"
	"os"
	"time"
)

type UserService struct {
	Validator      api.IValidator
	UserRepository domains.IUserRepository
	S3Client       *s3.Client
}

func NewUserService(repoCollector *repo.RepoCollector, validator api.IValidator, s3 *s3.Client) *UserService {
	return &UserService{
		UserRepository: *repoCollector.UserRepo,
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

	_, err3 := us.S3Client.PutObject(ctx, &s3.PutObjectInput{
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
