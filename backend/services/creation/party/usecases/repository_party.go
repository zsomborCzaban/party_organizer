package usecases

import (
	"errors"
	"fmt"
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

func (pr PartyRepository) GetPartiesByOrganizerId(id uint) (*[]domains.Party, error) {
	queryParams := []db.QueryParameter{
		{Field: "organizer_id", Operator: "=", Value: id},
	}

	fetchedParties, fetchedError := pr.DbAccess.Query(queryParams)
	if fetchedError != nil {
		//we should return errors from the databaselayer
		return nil, errors.New(fmt.Sprintf("Error while fetching parties for organizer id: %d, this should be only temporary. Error: %s", id, fetchedError.Error()))
	}

	parties, err := fetchedParties.(*[]domains.Party)
	if !err {
		return nil, errors.New("error. fetched parties cannot be transormed to *[]Party")
	}

	//not sure if parties can be nil after the db function call
	if parties == nil {
		return nil, errors.New("Error. Parties were nil")
	}

	return parties, nil
}

func (pr PartyRepository) GetPartiesByParticipantId(id uint) (*[]domains.Party, error) {
	queryParams := []db.QueryParameter{
		{Field: "participant_ids", Operator: "IN", Value: id},
	}

	fetchedParties, fetchedError := pr.DbAccess.Query(queryParams)
	if fetchedError != nil {
		//we should return errors from the databaselayer
		return nil, errors.New(fmt.Sprintf("Error while fetching parties for PARICIPANT.id: %d, this should be only temporary. Error: %s", id, fetchedError.Error()))
	}

	parties, err := fetchedParties.(*[]domains.Party)
	if !err {
		return nil, errors.New("error. fetched parties cannot be transormed to *[]Party")
	}

	//not sure if parties can be nil after the db function call
	if parties == nil {
		return nil, errors.New("Error. Parties were nil")
	}

	return parties, nil
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
