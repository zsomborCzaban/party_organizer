package usecases

import (
	"errors"
	"github.com/zsomborCzaban/party_organizer/db"
	"github.com/zsomborCzaban/party_organizer/services/users/registration/domains"
	"time"
)

type Repository struct {
	DbAccess db.IDatabaseAccess
}

func NewRegistrationRepository(databaseAccessManager db.IDatabaseAccessManager) domains.IRegistrationRepository {
	provider := NewRegistrationEntityProvider()
	access := databaseAccessManager.RegisterEntity("registrationProvider", provider)
	return &Repository{DbAccess: access}
}

func (r Repository) FindById(id uint) (*domains.RegistrationRequest, error) {
	reg, err := r.DbAccess.FindById(id)
	if err != nil {
		return nil, err
	}

	reg2, err2 := reg.(*domains.RegistrationRequest)
	if !err2 {
		return nil, errors.New("failed to convert database entity to party")
	}
	return reg2, nil
}

func (r Repository) FindByUsername(username string) (*domains.RegistrationRequest, error) {
	queryParams := []db.QueryParameter{
		{Field: "username", Operator: "=", Value: username},
	}

	fetchedUser, fetchedError := r.DbAccess.Query(queryParams)
	if fetchedError != nil {
		return nil, errors.New("error while fetching user: " + username + ", error: " + fetchedError.Error())
	}

	if regArray, ok := fetchedUser.(*[]domains.RegistrationRequest); ok {
		if regArray == nil {
			return nil, errors.New(domains.UserNotFound + username)
		}

		if len(*regArray) == 0 {
			return nil, errors.New(domains.UserNotFound + username)
		}

		if len(*regArray) > 1 {
			return nil, errors.New("more than one user found with username: " + username)
		}

		reg := (*regArray)[0]
		return &reg, nil
	} else {
		return nil, errors.New("error while fetching entity: fetched entity cannot be casted to user")
	}
}

func (r Repository) Create(reg *domains.RegistrationRequest) error {
	return r.DbAccess.Create(reg)
}

func (r Repository) Delete(reg *domains.RegistrationRequest) error {
	return r.DbAccess.Delete(reg)
}

func (r Repository) DeleteOlderThan(deleteTime time.Time) error {
	conds := []db.QueryParameter{{
		Field:    "created_at",
		Operator: "<",
		Value:    deleteTime,
	},
	}

	if err := r.DbAccess.BatchDelete(conds); err != nil {
		return err
	}
	return nil
}

type RegistrationEntityProvider struct {
}

func NewRegistrationEntityProvider() *RegistrationEntityProvider {
	return &RegistrationEntityProvider{}
}

func (provider *RegistrationEntityProvider) Create() interface{} {
	return &domains.RegistrationRequest{}
}

func (provider *RegistrationEntityProvider) CreateArray() interface{} {
	return &[]domains.RegistrationRequest{}
}
