package usecases

import (
	"errors"
	"github.com/zsomborCzaban/party_organizer/db"
	"github.com/zsomborCzaban/party_organizer/services/invitation/friend_invite/domains"
	userDomain "github.com/zsomborCzaban/party_organizer/services/user/domains"
)

type FriendInviteRepository struct {
	DbAccess       db.IDatabaseAccess
	UserRepository userDomain.IUserRepository
}

func NewFriendInviteRepository(databaseAccessManager db.IDatabaseAccessManager, ur userDomain.IUserRepository) domains.IFriendInviteRepository {
	entityProvider := EntityProvider{}
	databaseAccess := databaseAccessManager.RegisterEntity("partyProvider", entityProvider)

	return &FriendInviteRepository{
		DbAccess:       databaseAccess,
		UserRepository: ur,
	}
}

func (fr FriendInviteRepository) Create(invitation *domains.FriendInvitation) error {
	invitor, err := fr.UserRepository.FindById(invitation.InvitorId)
	if err != nil {
		return err
	}

	invited, err2 := fr.UserRepository.FindById(invitation.InvitedId)
	if err2 != nil {
		return err2
	}

	invitation.Invited = *invited
	invitation.Invitor = *invitor

	if err3 := fr.DbAccess.Create(invitation); err != nil {
		return err3
	}
	return nil
}

func (fr FriendInviteRepository) Update(invitation *domains.FriendInvitation) error {
	if err := fr.DbAccess.Update(invitation); err != nil {
		return err
	}
	return nil
}

func (fr FriendInviteRepository) FindByIds(invitorId, invitedId uint) (*domains.FriendInvitation, error) {
	queryParams := []db.QueryParameter{
		{Field: "invitor_id", Operator: "=", Value: invitorId},
		{Field: "invited_id", Operator: "=", Value: invitedId},
	}

	fetchedInvite, fetchedError := fr.DbAccess.Query(queryParams)
	if fetchedError != nil {
		return nil, errors.New("error, unexpected error while querying FriendInvite table")
	}

	//possible is database state is invalid and more than 1 invite is found
	invite, err := fetchedInvite.(*domains.FriendInvitation)
	if !err {
		return nil, errors.New("error, fetched invites cannot be transformed to *[]FriendInvitation")
	}

	//not sure if parties can be nil after the db function call
	if invite == nil {
		return nil, errors.New("error, invite was nil")
	}

	return invite, nil
}

type EntityProvider struct {
}

func (e EntityProvider) Create() interface{} {
	return &domains.FriendInvitation{}
}

func (e EntityProvider) CreateArray() interface{} {
	return &[]domains.FriendInvitation{}
}
