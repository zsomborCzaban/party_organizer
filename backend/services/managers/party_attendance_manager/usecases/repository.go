package usecases

import (
	"errors"
	"github.com/zsomborCzaban/party_organizer/db"
	"github.com/zsomborCzaban/party_organizer/services/managers/party_attendance_manager/domains"
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

func (pr PartyInviteRepository) Create(invitation *domains.PartyInvite) error {
	if err3 := pr.DbAccess.Create(invitation); err3 != nil {
		return err3
	}
	return nil
}

func (pr PartyInviteRepository) Update(invitation *domains.PartyInvite) error {
	if err := pr.DbAccess.Update(invitation); err != nil {
		return err
	}
	return nil
}

func (pr PartyInviteRepository) DeleteByPartyId(partyId uint) error {
	queryParams := []db.QueryParameter{
		{Field: "party_id", Operator: "=", Value: partyId},
	}

	err := pr.DbAccess.BatchDelete(queryParams)
	return err
}

func (pr PartyInviteRepository) FindByIds(invitedId, partyId uint) (*domains.PartyInvite, error) {
	queryParams := []db.QueryParameter{
		{Field: "invited_id", Operator: "=", Value: invitedId},
		{Field: "party_id", Operator: "=", Value: partyId},
	}

	fetchedInvites, fetchedError := pr.DbAccess.Query(queryParams)
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

func (pr PartyInviteRepository) FindPendingByInvitedId(invitedId uint) (*[]domains.PartyInvite, error) {
	queryParams := []db.QueryParameter{
		{Field: "state", Operator: "=", Value: domains.PENDING},
		{Field: "invited_id", Operator: "=", Value: invitedId},
	}

	fetchedInvites, fetchedError := pr.DbAccess.Query(queryParams, "Party", "Invitor")
	if fetchedError != nil {
		return nil, fetchedError
	}

	invites, err := fetchedInvites.(*[]domains.PartyInvite)
	if !err {
		return nil, errors.New(domains.FAILED_PARSE_TO_ARRAY)
	}

	return invites, nil
}

func (pr PartyInviteRepository) FindPendingByPartyId(partyId uint) (*[]domains.PartyInvite, error) {
	queryParams := []db.QueryParameter{
		{Field: "state", Operator: "=", Value: domains.PENDING},
		{Field: "party_id", Operator: "=", Value: partyId},
	}

	fetchedInvites, fetchedError := pr.DbAccess.Query(queryParams, "Party", "Invitor", "Invited") //could cause concurrent map writes
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
