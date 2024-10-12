package domains

import "github.com/zsomborCzaban/party_organizer/common/api"

type IUserService interface {
	Login(LoginRequest) api.IResponse
	Register(RegisterRequest) api.IResponse
	AddFriend(uint, uint) api.IResponse
	GetFriends(uint) api.IResponse
}
