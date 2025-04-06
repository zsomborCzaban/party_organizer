package domains

import (
	"github.com/zsomborCzaban/party_organizer/utils/api"
	"mime/multipart"
)

type IUserService interface {
	Login(LoginRequest) api.IResponse
	AddFriend(uint, uint) api.IResponse
	GetFriends(uint) api.IResponse
	UploadProfilePicture(uint, multipart.File, *multipart.FileHeader) api.IResponse

	ForgotPassword(username string) api.IResponse
	ChangePassword(ChangePasswordRequest) api.IResponse
}
