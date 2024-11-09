package domains

import (
	"github.com/zsomborCzaban/party_organizer/common/api"
	"mime/multipart"
)

type IUserService interface {
	Login(LoginRequest) api.IResponse
	Register(RegisterRequest) api.IResponse
	AddFriend(uint, uint) api.IResponse
	GetFriends(uint) api.IResponse
	UploadProfilePicture(multipart.File, *multipart.FileHeader) api.IResponse
}
