package domains

type IUserRepository interface {
	FindByUsername(string) (*User, error)
	CreateUser(*User) error
}
