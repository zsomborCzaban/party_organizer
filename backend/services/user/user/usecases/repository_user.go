package usecases

import (
	"errors"
	"github.com/zsomborCzaban/party_organizer/db"
	domains2 "github.com/zsomborCzaban/party_organizer/services/user/user/domains"
)

type UserEntityProvider struct {
}

type UserRepository struct {
	DbAccess db.IDatabaseAccess
}

func NewUserRepository(databaseAccessManager db.IDatabaseAccessManager) domains2.IUserRepository {
	provider := NewUserEntityProvider()
	access := databaseAccessManager.RegisterEntity("userProvider", provider)
	return &UserRepository{DbAccess: access}
}

func (ur UserRepository) FindById(id uint) (*domains2.User, error) {
	user, err := ur.DbAccess.FindById(id)
	if err != nil {
		return nil, err
	}

	user2, err2 := user.(*domains2.User)
	if !err2 {
		return nil, errors.New("failed to convert database entity to party")
	}
	return user2, nil
}

func (ur UserRepository) FindByUsername(username string) (*domains2.User, error) {
	queryParams := []db.QueryParameter{
		{Field: "username", Operator: "=", Value: username},
	}

	fetchedUser, fetchedError := ur.DbAccess.Query(queryParams)
	if fetchedError != nil {
		return nil, errors.New("error while fetching user: " + username + ", error: " + fetchedError.Error())
	}

	if userArray, ok := fetchedUser.(*[]domains2.User); ok {
		if userArray == nil {
			return nil, errors.New(domains2.UserNotFound + username)
		}

		if len(*userArray) == 0 {
			return nil, errors.New(domains2.UserNotFound + username)
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

func (ur UserRepository) CreateUser(user *domains2.User) error {
	err := ur.DbAccess.Create(user)
	if err != nil {
		return err
	}
	return nil
}

func (ur UserRepository) AddFriend(userId, friendId uint) error {
	user, err := ur.FindById(userId)
	if err != nil {
		return err
	}

	friend, err2 := ur.FindById(friendId)
	if err2 != nil {
		return err2
	}

	//should use transaction here for data integrity
	user.Friends = append(user.Friends, *friend)
	if err3 := ur.UpdateUser(user); err3 != nil {
		return err3
	}

	friend.Friends = append(friend.Friends, *user)
	if err4 := ur.UpdateUser(friend); err4 != nil {
		return err4
	}

	return nil
}

func (ur UserRepository) UpdateUser(user *domains2.User) error {
	err := ur.DbAccess.Update(user)
	if err != nil {
		return err
	}
	return nil
}

func (ur UserRepository) GetFriends(userId uint) (*[]domains2.User, error) {
	cond := db.Many2ManyQueryParameter{
		QueriedTable:            "users",
		Many2ManyTable:          "user_friends",
		M2MQueriedColumnName:    "user_id",
		M2MConditionColumnName:  "friend_id",
		M2MConditionColumnValue: userId,
		//OrActive:                true,
		//OrConditionColumnName:   "friend_id",
		//OrConditionColumnValue:  userId,
	}

	fetchedUsers, fetchedError := ur.DbAccess.Many2ManyQueryId(cond)
	if fetchedError != nil {
		return nil, fetchedError
	}

	users, err := fetchedUsers.(*[]domains2.User)
	if !err {
		return nil, errors.New("failed to convert fetched friends to *[]User type")
	}

	if users == nil {
		return nil, errors.New("friends was nil")
	}

	return users, nil
}

func NewUserEntityProvider() *UserEntityProvider { return &UserEntityProvider{} }

func (provider *UserEntityProvider) Create() interface{} { return &domains2.User{} }

func (provider *UserEntityProvider) CreateArray() interface{} { return &[]domains2.User{} }
