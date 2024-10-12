package domains

import "net/http"

type IUserController interface {
	LoginController(w http.ResponseWriter, r *http.Request)
	RegisterController(w http.ResponseWriter, r *http.Request)
	AddFriendController(w http.ResponseWriter, r *http.Request)
	GetFriendsController(w http.ResponseWriter, r *http.Request)
}
