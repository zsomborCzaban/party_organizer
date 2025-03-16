package domains

type IUserRepository interface {
	FindByUsername(string) (*User, error)
	FindById(id uint, associations ...string) (*User, error)
	CreateUser(*User) error
	UpdateUser(*User) error
	AddFriend(*User, *User) error
	RemoveFriend(*User, *User) error
	GetFriends(uint) (*[]User, error)
}
