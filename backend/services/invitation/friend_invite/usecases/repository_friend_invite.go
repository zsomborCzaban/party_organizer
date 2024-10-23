package usecases

import (
	"errors"
	"github.com/zsomborCzaban/party_organizer/db"
	"github.com/zsomborCzaban/party_organizer/services/invitation/friend_invite/domains"
)

type FriendInviteRepository struct {
	DbAccess db.IDatabaseAccess
}

func NewFriendInviteRepository(databaseAccessManager db.IDatabaseAccessManager) domains.IFriendInviteRepository {
	entityProvider := EntityProvider{}
	databaseAccess := databaseAccessManager.RegisterEntity("friendInviteProvider", entityProvider)

	return &FriendInviteRepository{
		DbAccess: databaseAccess,
	}
}

func (fr FriendInviteRepository) Create(invitation *domains.FriendInvite) error {
	if err3 := fr.DbAccess.Create(invitation); err3 != nil {
		return err3
	}
	return nil
}

func (fr FriendInviteRepository) Update(invitation *domains.FriendInvite) error {
	if err := fr.DbAccess.Update(invitation); err != nil {
		return err
	}
	return nil
}

func (fr FriendInviteRepository) FindByIds(invitorId, invitedId uint) (*domains.FriendInvite, error) {
	queryParams := []db.QueryParameter{
		{Field: "invitor_id", Operator: "=", Value: invitorId},
		{Field: "invited_id", Operator: "=", Value: invitedId},
	}

	fetchedInvites, fetchedError := fr.DbAccess.Query(queryParams)
	if fetchedError != nil {
		return nil, errors.New("error, unexpected error while querying FriendInvite table")
	}

	//possible is database state is invalid and more than 1 invite is found
	invites, err := fetchedInvites.(*[]domains.FriendInvite)
	if !err {
		return nil, errors.New(domains.FAILED_PARSE_TO_ARRAY)
	}

	if len(*invites) > 1 {
		return nil, errors.New("invalid datbase state, more than 1 invites found")
	}

	if len(*invites) == 0 {
		return nil, errors.New(domains.NOT_FOUND)
	}

	invite := &(*invites)[0]

	return invite, nil
}

func (fr FriendInviteRepository) FindPendingByInvitedId(invitedId uint) (*[]domains.FriendInvite, error) {
	queryParams := []db.QueryParameter{
		{Field: "state", Operator: "=", Value: domains.PENDING},
		{Field: "invited_id", Operator: "=", Value: invitedId},
	}

	fetchedInvites, fetchedError := fr.DbAccess.Query(queryParams, "Invitor", "Invited")
	if fetchedError != nil {
		return nil, fetchedError
	}

	invites, err := fetchedInvites.(*[]domains.FriendInvite)
	if !err {
		return nil, errors.New(domains.FAILED_PARSE_TO_ARRAY)
	}

	return invites, nil
}

type EntityProvider struct {
}

func (e EntityProvider) Create() interface{} {
	return &domains.FriendInvite{}
}

func (e EntityProvider) CreateArray() interface{} {
	return &[]domains.FriendInvite{}
}
