package usecases

import (
	"errors"
	"github.com/zsomborCzaban/party_organizer/db"
	"github.com/zsomborCzaban/party_organizer/services/invitation/party_invite/domains"
)

type PartyInviteRepository struct {
	DbAccess db.IDatabaseAccess
}

func NewPartyInviteRepository(databaseAccessManager db.IDatabaseAccessManager) domains.IPartyInviteRepository {
	entityProvider := EntityProvider{}
	databaseAccess := databaseAccessManager.RegisterEntity("partyInviteProvider", entityProvider)

	return &PartyInviteRepository{
		DbAccess: databaseAccess,
	}
}

func (fr PartyInviteRepository) Create(invitation *domains.PartyInvite) error {
	if err3 := fr.DbAccess.Create(invitation); err3 != nil {
		return err3
	}
	return nil
}

func (fr PartyInviteRepository) Update(invitation *domains.PartyInvite) error {
	if err := fr.DbAccess.Update(invitation); err != nil {
		return err
	}
	return nil
}

func (fr PartyInviteRepository) FindByIds(invitedId, partyId uint) (*domains.PartyInvite, error) {
	queryParams := []db.QueryParameter{
		{Field: "invited_id", Operator: "=", Value: invitedId},
		{Field: "party_id", Operator: "=", Value: partyId},
	}

	fetchedInvites, fetchedError := fr.DbAccess.Query(queryParams)
	if fetchedError != nil {
		return nil, errors.New("error, unexpected error while querying PartyInvite table")
	}

	invites, err := fetchedInvites.(*[]domains.PartyInvite)
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

func (fr PartyInviteRepository) FindPendingByInvitedId(invitedId uint) (*[]domains.PartyInvite, error) {
	queryParams := []db.QueryParameter{
		{Field: "state", Operator: "=", Value: domains.PENDING},
		{Field: "invited_id", Operator: "=", Value: invitedId},
	}

	fetchedInvites, fetchedError := fr.DbAccess.Query(queryParams)
	if fetchedError != nil {
		return nil, fetchedError
	}

	invites, err := fetchedInvites.(*[]domains.PartyInvite)
	if !err {
		return nil, errors.New(domains.FAILED_PARSE_TO_ARRAY)
	}

	return invites, nil
}

type EntityProvider struct {
}

func (e EntityProvider) Create() interface{} {
	return &domains.PartyInvite{}
}

func (e EntityProvider) CreateArray() interface{} {
	return &[]domains.PartyInvite{}
}
