package domains

type IUserRepository interface {
	FindByUsername(string) (*User, error)
	FindById(uint) (*User, error)
	CreateUser(*User) error
	AddFriend(uint, uint) error
	GetFriends(uint) (*[]User, error)
}
