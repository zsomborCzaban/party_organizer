package usecases

import (
	"errors"
	"github.com/zsomborCzaban/party_organizer/db"
	"github.com/zsomborCzaban/party_organizer/services/user/domains"
)

type UserEntityProvider struct {
}

type UserRepository struct {
	DbAccess db.IDatabaseAccess
}

func NewUserRepository(databaseAccessManager db.IDatabaseAccessManager) domains.IUserRepository {
	provider := NewUserEntityProvider()
	access := databaseAccessManager.RegisterEntity("userProvider", provider)
	return &UserRepository{DbAccess: access}
}

func (ur UserRepository) FindByUsername(username string) (*domains.User, error) {
	queryParams := []db.QueryParameter{
		{Field: "username", Operator: "=", Value: username},
	}

	fetchedUser, fetchedError := ur.DbAccess.Query(queryParams)
	if fetchedError != nil {
		return nil, errors.New("error while fetching user: " + username + ", error: " + fetchedError.Error())
	}

	if userArray, ok := fetchedUser.(*[]domains.User); ok {
		if userArray == nil {
			return nil, errors.New(domains.UserNotFound + username)
		}

		if len(*userArray) == 0 {
			return nil, errors.New(domains.UserNotFound + username)
		}

		if len(*userArray) > 1 {
			return nil, errors.New("more than one user found with username: " + username)
		}

		user := (*userArray)[0]
		return &user, nil
	} else {
		return nil, errors.New("error while fetching entity: fetched entity cannot be casted to user")
	}
}

func (ur UserRepository) CreateUser(user *domains.User) error {
	err := ur.DbAccess.Create(user)
	if err != nil {
		return err
	}
	return nil
}

func NewUserEntityProvider() *UserEntityProvider { return &UserEntityProvider{} }

func (provider *UserEntityProvider) Create() interface{} { return &domains.User{} }

func (provider *UserEntityProvider) CreateArray() interface{} { return &[]domains.User{} }
