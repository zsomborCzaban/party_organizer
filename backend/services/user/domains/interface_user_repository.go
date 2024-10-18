package domains

type IUserRepository interface {
	FindByUsername(string) (*User, error)
	FindById(uint) (*User, error)
	CreateUser(*User) error
	AddFriend(*User, *User) error
	RemoveFriend(*User, *User) error
	GetFriends(uint) (*[]User, error)
}
