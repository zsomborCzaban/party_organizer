package usecases

import (
	"errors"
	"github.com/zsomborCzaban/party_organizer/db"
	"github.com/zsomborCzaban/party_organizer/services/creation/party/domains"
)

type PartyRepository struct {
	DbAccess db.IDatabaseAccess
}

func NewPartyRepository(databaseAccessManager db.IDatabaseAccessManager) domains.IPartyRepository {
	entityProvider := EntityProvider{}
	databaseAccess := databaseAccessManager.RegisterEntity("partyProvider", entityProvider)

	return &PartyRepository{
		DbAccess: databaseAccess,
	}
}

func (pr PartyRepository) CreateParty(party *domains.Party) error {
	err := pr.DbAccess.Create(party)
	if err != nil {
		return err
	}
	return nil

}

func (pr PartyRepository) GetParty(id uint) (*domains.Party, error) {
	party, err := pr.DbAccess.FindById(id)
	if err != nil {
		return nil, err
	}

	party2, err2 := party.(*domains.Party)
	if !err2 {
		return nil, errors.New("failed to convert database entity to party")
	}
	return party2, nil
}

func (pr PartyRepository) UpdateParty(party *domains.Party) error {
	err := pr.DbAccess.Update(party)
	if err != nil {
		return err
	}
	return nil
}

func (pr PartyRepository) DeleteParty(party *domains.Party) error {
	err := pr.DbAccess.Delete(party)
	if err != nil {
		return err
	}
	return nil
}

type EntityProvider struct {
}

func (e EntityProvider) Create() interface{} {
	return &domains.Party{}
}

func (e EntityProvider) CreateArray() interface{} {
	return &[]domains.Party{}
}
