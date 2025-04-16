package usecases

import (
	"errors"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zsomborCzaban/party_organizer/common"
	domains2 "github.com/zsomborCzaban/party_organizer/services/users/user/domains"
	"github.com/zsomborCzaban/party_organizer/utils/api"
	"github.com/zsomborCzaban/party_organizer/utils/repo"
	mockS3 "github.com/zsomborCzaban/party_organizer/utils/s3"
	"gorm.io/gorm"
	"mime/multipart"
	"os"
	"testing"
)

const JWT_SIGNING_KEY_VALUE = "verysecretkey"

func setupDefaultService() (*UserService, *api.MockValidator, *MockRepository, *mockS3.MockS3ClientWrapper) {
	validator := new(api.MockValidator)
	userRepo := new(MockRepository)
	s3Client := new(mockS3.MockS3ClientWrapper)

	repoCollector := &repo.RepoCollector{
		UserRepo: userRepo,
	}

	service := NewUserService(repoCollector, validator, s3Client)

	return service, validator, userRepo, s3Client
}

func Test_UserService_Login_Success(t *testing.T) {
	service, validator, userRepo, _ := setupDefaultService()
	os.Setenv(common.JWT_SINGING_KEY_ENV_KEY, JWT_SIGNING_KEY_VALUE)
	defer os.Unsetenv(common.JWT_SINGING_KEY_ENV_KEY)

	username := "testuser"
	password := "password123"
	loginRequest := domains2.LoginRequest{
		Username: &username,
		Password: &password,
	}
	user := &domains2.User{
		Username: username,
		Password: "$2a$10$nk2OIw17ipD7mUBoNjzhr.s79s2S2oOIDHUkHOtVUIdNkaGuFniCu", // Mock hashed password
	}

	validator.On("Validate", loginRequest).Return(nil)
	userRepo.On("FindByUsername", username).Return(user, nil)

	response := service.Login(loginRequest)

	assert.False(t, response.GetIsError())
	validator.AssertExpectations(t)
	userRepo.AssertExpectations(t)
}

func Test_UserService_Login_FailValidation(t *testing.T) {
	service, validator, _, _ := setupDefaultService()

	username := ""
	password := ""
	loginRequest := domains2.LoginRequest{
		Username: &username,
		Password: &password,
	}
	validationError := api.NewValidationErrors()

	validator.On("Validate", loginRequest).Return(validationError)

	response := service.Login(loginRequest)

	assert.Equal(t, api.ErrorValidation(validationError.Errors), response)
}

func Test_UserService_Login_FailUserNotFound(t *testing.T) {
	service, validator, userRepo, _ := setupDefaultService()

	username := "nonexistent"
	password := "password123"
	loginRequest := domains2.LoginRequest{
		Username: &username,
		Password: &password,
	}

	validator.On("Validate", loginRequest).Return(nil)
	userRepo.On("FindByUsername", username).Return(&domains2.User{}, errors.New("not found"))

	response := service.Login(loginRequest)

	assert.Equal(t, api.ErrorInvalidCredentials(), response)
}

func Test_UserService_Login_FailInvalidPassword(t *testing.T) {
	service, validator, userRepo, _ := setupDefaultService()

	username := "testuser"
	password := "wrongpassword"
	loginRequest := domains2.LoginRequest{
		Username: &username,
		Password: &password,
	}
	user := &domains2.User{
		Username: username,
		Password: "$2a$10$somehashedpassword", // Mock hashed password
	}

	validator.On("Validate", loginRequest).Return(nil)
	userRepo.On("FindByUsername", username).Return(user, nil)

	response := service.Login(loginRequest)

	assert.Equal(t, api.ErrorInvalidCredentials(), response)
}

func Test_UserService_AddFriend_Success(t *testing.T) {
	service, _, userRepo, _ := setupDefaultService()

	userId := uint(1)
	friendId := uint(2)
	user := &domains2.User{Model: gorm.Model{ID: userId}}
	friend := &domains2.User{Model: gorm.Model{ID: friendId}}

	userRepo.On("FindById", userId, mock.Anything).Return(user, nil)
	userRepo.On("FindById", friendId, mock.Anything).Return(friend, nil)
	userRepo.On("AddFriend", user, friend).Return(nil)

	response := service.AddFriend(friendId, userId)

	assert.False(t, response.GetIsError())
}

func Test_UserService_AddFriend_FailSelfFriend(t *testing.T) {
	service, _, _, _ := setupDefaultService()

	userId := uint(1)
	response := service.AddFriend(userId, userId)

	assert.Equal(t, api.ErrorBadRequest("You cannot be friends with yourself"), response)
}

func Test_UserService_AddFriend_FailFindUser(t *testing.T) {
	service, _, userRepo, _ := setupDefaultService()

	userId := uint(1)
	friendId := uint(2)
	expectedErr := errors.New("not found")

	userRepo.On("FindById", userId, mock.Anything).Return(&domains2.User{}, expectedErr)

	response := service.AddFriend(friendId, userId)

	assert.Equal(t, api.ErrorBadRequest(expectedErr.Error()), response)
}

func Test_UserService_AddFriend_FailFindFriend(t *testing.T) {
	service, _, userRepo, _ := setupDefaultService()

	userId := uint(1)
	friendId := uint(2)
	user := &domains2.User{Model: gorm.Model{ID: userId}}
	expectedErr := errors.New("not found")

	userRepo.On("FindById", userId, mock.Anything).Return(user, nil)
	userRepo.On("FindById", friendId, mock.Anything).Return(&domains2.User{}, expectedErr)

	response := service.AddFriend(friendId, userId)

	assert.Equal(t, api.ErrorBadRequest(expectedErr.Error()), response)
}

func Test_UserService_AddFriend_FailAdd(t *testing.T) {
	service, _, userRepo, _ := setupDefaultService()

	userId := uint(1)
	friendId := uint(2)
	user := &domains2.User{Model: gorm.Model{ID: userId}}
	friend := &domains2.User{Model: gorm.Model{ID: friendId}}
	expectedErr := errors.New("add failed")

	userRepo.On("FindById", userId, mock.Anything).Return(user, nil)
	userRepo.On("FindById", friendId, mock.Anything).Return(friend, nil)
	userRepo.On("AddFriend", user, friend).Return(expectedErr)

	response := service.AddFriend(friendId, userId)

	assert.Equal(t, api.ErrorInternalServerError(expectedErr), response)
}

func Test_UserService_GetFriends_Success(t *testing.T) {
	service, _, userRepo, _ := setupDefaultService()

	userId := uint(1)
	user := &domains2.User{
		Model: gorm.Model{ID: userId},
		Friends: []domains2.User{
			{Model: gorm.Model{ID: 2}, Username: "friend1"},
			{Model: gorm.Model{ID: 3}, Username: "friend2"},
		},
	}

	userRepo.On("FindById", userId, mock.Anything).Return(user, nil)

	response := service.GetFriends(userId)

	assert.False(t, response.GetIsError())
	assert.Len(t, response.GetData().([]domains2.UserDTO), 2)
}

func Test_UserService_GetFriends_Fail(t *testing.T) {
	service, _, userRepo, _ := setupDefaultService()

	userId := uint(1)
	expectedErr := errors.New("not found")

	userRepo.On("FindById", userId, mock.Anything).Return(&domains2.User{}, expectedErr)

	response := service.GetFriends(userId)

	assert.Equal(t, api.ErrorInternalServerError(expectedErr), response)
}

func Test_UserService_UploadProfilePicture_Success(t *testing.T) {
	service, _, userRepo, s3Client := setupDefaultService()

	userId := uint(1)
	user := &domains2.User{Model: gorm.Model{ID: userId}}

	tempFile, err := os.CreateTemp("", "test-*.jpg")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	fileHeader := &multipart.FileHeader{
		Filename: "test.jpg",
		Header:   map[string][]string{"Content-Type": {"image/jpeg"}},
		Size:     0,
	}

	os.Setenv("AWS_BUCKET_NAME", "test-bucket")
	defer os.Unsetenv("AWS_BUCKET_NAME")

	userRepo.On("FindById", userId, mock.Anything).Return(user, nil)
	s3Client.On("PutObject", mock.Anything, mock.AnythingOfType("*s3.PutObjectInput"), mock.Anything).Return(&s3.PutObjectOutput{}, nil)
	userRepo.On("UpdateUser", user).Return(nil)

	response := service.UploadProfilePicture(userId, tempFile, fileHeader)

	assert.False(t, response.GetIsError())
	assert.Contains(t, user.ProfilePictureUrl, "https://test-bucket.s3.amazonaws.com/users/1/profile_pictures/test.jpg")
}

func Test_UserService_UploadProfilePicture_FailFindUser(t *testing.T) {
	service, _, userRepo, _ := setupDefaultService()

	userId := uint(1)
	expectedErr := errors.New("not found")

	file, err := os.CreateTemp("", "test-*.jpg")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(file.Name())

	fileHeader := &multipart.FileHeader{}

	userRepo.On("FindById", userId, mock.Anything).Return(&domains2.User{}, expectedErr)

	response := service.UploadProfilePicture(userId, file, fileHeader)

	assert.Equal(t, api.ErrorInternalServerError(expectedErr.Error()), response)
}

func Test_UserService_UploadProfilePicture_FailBucketNotSet(t *testing.T) {
	service, _, userRepo, _ := setupDefaultService()

	userId := uint(1)
	user := &domains2.User{Model: gorm.Model{ID: userId}}

	file, err := os.CreateTemp("", "test-*.jpg")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(file.Name())

	fileHeader := &multipart.FileHeader{
		Size: int64(len("test")),
	}

	userRepo.On("FindById", userId, mock.Anything).Return(user, nil)

	response := service.UploadProfilePicture(userId, file, fileHeader)

	assert.Equal(t, api.ErrorInternalServerError("AWS_BUCKET_NAME environment variable is not set."), response)
}

func Test_UserService_UploadProfilePicture_FailS3Upload(t *testing.T) {
	service, _, userRepo, s3Client := setupDefaultService()

	userId := uint(1)
	user := &domains2.User{Model: gorm.Model{ID: userId}}
	file, err := os.CreateTemp("", "test-*.jpg")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(file.Name())
	fileHeader := &multipart.FileHeader{
		Filename: "test.jpg",
		Header:   map[string][]string{"Content-Type": {"image/jpeg"}},
		Size:     int64(len("test content")),
	}
	expectedErr := errors.New("upload failed")

	os.Setenv("AWS_BUCKET_NAME", "test-bucket")
	defer os.Unsetenv("AWS_BUCKET_NAME")

	userRepo.On("FindById", userId, mock.Anything).Return(user, nil)
	s3Client.On("PutObject", mock.Anything, mock.AnythingOfType("*s3.PutObjectInput"), mock.Anything).Return(&s3.PutObjectOutput{}, expectedErr)

	response := service.UploadProfilePicture(userId, file, fileHeader)

	assert.Equal(t, api.ErrorInternalServerError(expectedErr.Error()), response)
}

func Test_UserService_UploadProfilePicture_FailUpdateUser(t *testing.T) {
	service, _, userRepo, s3Client := setupDefaultService()

	userId := uint(1)
	user := &domains2.User{Model: gorm.Model{ID: userId}}
	file, err := os.CreateTemp("", "test-*.jpg")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(file.Name())
	fileHeader := &multipart.FileHeader{
		Filename: "test.jpg",
		Header:   map[string][]string{"Content-Type": {"image/jpeg"}},
		Size:     0,
	}
	expectedErr := errors.New("update failed")

	os.Setenv("AWS_BUCKET_NAME", "test-bucket")
	defer os.Unsetenv("AWS_BUCKET_NAME")

	userRepo.On("FindById", userId, mock.Anything).Return(user, nil)
	s3Client.On("PutObject", mock.Anything, mock.AnythingOfType("*s3.PutObjectInput"), mock.Anything).Return(&s3.PutObjectOutput{}, nil)
	userRepo.On("UpdateUser", user).Return(expectedErr)

	response := service.UploadProfilePicture(userId, file, fileHeader)

	assert.Equal(t, api.ErrorInternalServerError(expectedErr.Error()), response)
}
