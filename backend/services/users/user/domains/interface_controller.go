package domains

import "net/http"

type IUserController interface {
	Login(w http.ResponseWriter, r *http.Request)
	AddFriend(w http.ResponseWriter, r *http.Request)
	GetFriends(w http.ResponseWriter, r *http.Request)
	UploadProfilePicture(w http.ResponseWriter, r *http.Request)

	ForgotPassword(w http.ResponseWriter, r *http.Request)
	ChangePassword(w http.ResponseWriter, r *http.Request)
}
