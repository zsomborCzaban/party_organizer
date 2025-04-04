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

func (ur UserRepository) FindById(id uint, associations ...string) (*domains.User, error) {
	user, err := ur.DbAccess.FindById(id, associations...)
	if err != nil {
		return nil, err
	}

	user2, err2 := user.(*domains.User)
	if !err2 {
		return nil, errors.New("failed to convert database entity to party")
	}
	return user2, nil
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

func (ur UserRepository) AddFriend(user, friend *domains.User) error {
	//should use transaction here for data integrity
	//append both, so they are both ways in the join table. (simpler queries but double data, but for the current use case it's ok)
	//if err := ur.DbAccess.AddToAssociation(user, "Friends", friend); err != nil {
	//	return err
	//}
	//if err2 := ur.DbAccess.AddToAssociation(friend, "Friends", user); err2 != nil {
	//	return err2
	//}
	//not nice workaround
	user.Friends = append(user.Friends, *friend)
	if err := ur.DbAccess.Update(user); err != nil {
		return err
	}

	friend.Friends = append(friend.Friends, *user)
	if err2 := ur.DbAccess.Update(friend); err2 != nil {
		return err2
	}

	return nil
}

func (ur UserRepository) RemoveFriend(user, friend *domains.User) error {
	//if err := ur.DbAccess.DeleteFromAssociation(user, "Friends", friend); err != nil {
	//	return err
	//}
	//
	//if err2 := ur.DbAccess.DeleteFromAssociation(friend, "Friends", user); err != nil {
	//	return err2
	//}

	tr := ur.DbAccess.TransactionBegin()
	var err error
	defer func() {
		err = errors.Join(err, tr.TransactionRollback())
	}()

	var friends []domains.User
	for _, userFriend := range user.Friends {
		if userFriend.ID != friend.ID {
			friends = append(friends, userFriend)
		}
	}
	associationParam := db.AssociationParameter{
		Model:       user,
		Association: "Friends",
		Values:      friends,
	}

	if err = tr.ReplaceAssociations(associationParam); err != nil {
		return err
	}

	friends = []domains.User{}
	for _, friendFriend := range friend.Friends {
		if friendFriend.ID != user.ID {
			friends = append(friends, friendFriend)
		}
	}
	associationParam = db.AssociationParameter{
		Model:       friend,
		Association: "Friends",
		Values:      friends,
	}

	if err = tr.ReplaceAssociations(associationParam); err != nil {
		return err
	}

	if err = tr.TransactionCommit(); err != nil {
		return err
	}
	return nil
}

func (ur UserRepository) UpdateUser(user *domains.User) error {
	err := ur.DbAccess.Update(user)
	if err != nil {
		return err
	}
	return nil
}

func (ur UserRepository) GetFriends(userId uint) (*[]domains.User, error) {
	cond := db.Many2ManyQueryParameter{
		QueriedTable:            "users",
		Many2ManyTable:          "user_friends",
		M2MQueriedColumnName:    "user_id",
		M2MConditionColumnName:  "friend_id",
		M2MConditionColumnValue: userId,
	}

	fetchedUsers, fetchedError := ur.DbAccess.Many2ManyQueryId(cond)
	if fetchedError != nil {
		return nil, fetchedError
	}

	users, err := fetchedUsers.(*[]domains.User)
	if !err {
		return nil, errors.New("failed to convert fetched friends to *[]User type")
	}

	if users == nil {
		return nil, errors.New("friends was nil")
	}

	return users, nil
}

func NewUserEntityProvider() *UserEntityProvider { return &UserEntityProvider{} }

func (provider *UserEntityProvider) Create() interface{} { return &domains.User{} }

func (provider *UserEntityProvider) CreateArray() interface{} { return &[]domains.User{} }
