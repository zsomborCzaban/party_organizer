package usecases

import (
	"errors"
	"fmt"
	"github.com/zsomborCzaban/party_organizer/db"
	"github.com/zsomborCzaban/party_organizer/services/creation/party/domains"
	userDomain "github.com/zsomborCzaban/party_organizer/services/user/domains"
)

type Repository struct {
	DbAccess db.IDatabaseAccess //party DbAccess
}

func NewPartyRepository(databaseAccessManager db.IDatabaseAccessManager) domains.IPartyRepository {
	entityProvider := PartyEntityProvider{}
	databaseAccess := databaseAccessManager.RegisterEntity("partyProvider", entityProvider)

	return &Repository{
		DbAccess: databaseAccess,
	}
}

func (pr Repository) AddUserToParty(party *domains.Party, user *userDomain.User) error {
	//if err := pr.DbAccess.AddToAssociation(party, "Participants", user); err != nil {
	//	return err
	//}

	party.Participants = append(party.Participants, *user)
	return pr.DbAccess.Update(party)
}

func (pr Repository) RemoveUserFromParty(party *domains.Party, user *userDomain.User) error {
	//return pr.DbAccess.DeleteFromAssociation(party, "Participants", user)

	participants := []userDomain.User{}
	for _, participant := range party.Participants {
		if participant.ID != user.ID {
			participants = append(participants, participant)
		}
	}
	associationParam := db.AssociationParameter{
		Model:       party,
		Association: "Participants",
		Values:      participants,
	}

	return pr.DbAccess.ReplaceAssociations(associationParam)
}

func (pr Repository) GetPublicParties() (*[]domains.Party, error) {
	queryParams := []db.QueryParameter{
		{Field: "private", Operator: "=", Value: false},
	}

	fetchedParties, fetchedError := pr.DbAccess.Query(queryParams, "Organizer")
	if fetchedError != nil {
		//we should return errors from the databaselayer
		return nil, errors.New(fmt.Sprintf("Error while fetching public parties. this should be only temporary. Error: %s", fetchedError.Error()))
	}

	parties, err := fetchedParties.(*[]domains.Party)
	if !err {
		return nil, errors.New("error. fetched parties cannot be transormed to *[]Party")
	}

	//not sure if parties can be nil after the db function call
	if parties == nil {
		return nil, errors.New("error. Parties were nil")
	}

	return parties, nil
}

func (pr Repository) GetPartiesByOrganizerId(id uint) (*[]domains.Party, error) {
	queryParams := []db.QueryParameter{
		{Field: "organizer_id", Operator: "=", Value: id},
	}

	fetchedParties, fetchedError := pr.DbAccess.Query(queryParams, "Organizer")
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

func (pr Repository) GetPartiesByParticipantId(id uint) (*[]domains.Party, error) {
	queryCond := db.Many2ManyQueryParameter{
		QueriedTable:            "parties",
		Many2ManyTable:          "party_participants",
		M2MQueriedColumnName:    "party_id",
		M2MConditionColumnName:  "user_id",
		M2MConditionColumnValue: id,
	}

	fetchedParties, fetchedError := pr.DbAccess.Many2ManyQueryId(queryCond, "Organizer")
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

//func (pr Repository) FindUserInParty(userId, partyId uint) error {
//	queryCond := db.Many2ManyQueryParameter{
//		QueriedTable:            "parties",
//		Many2ManyTable:          "party_participants",
//		M2MQueriedColumnName:    "party_id",
//		M2MConditionColumnName:  "user_id",
//		M2MConditionColumnValue: id,
//	}
//}

func (pr Repository) Create(party *domains.Party) error {
	err := pr.DbAccess.Create(party)
	if err != nil {
		return err
	}
	return nil

}

func (pr Repository) FindById(id uint, associations ...string) (*domains.Party, error) {
	party, err := pr.DbAccess.FindById(id, associations...) //todo: causes concurent mapwrites sometimes
	if err != nil {
		return nil, err
	}

	party2, err2 := party.(*domains.Party)
	if !err2 {
		return nil, errors.New("failed to convert database entity to party")
	}
	return party2, nil
}

func (pr Repository) Update(party *domains.Party) error {
	err := pr.DbAccess.Update(party)
	if err != nil {
		return err
	}
	return nil
}

func (pr Repository) Delete(party *domains.Party) error {
	err := pr.DbAccess.Delete(party)
	if err != nil {
		return err
	}
	return nil
}

type PartyEntityProvider struct {
}

func (e PartyEntityProvider) Create() interface{} {
	return &domains.Party{}
}

func (e PartyEntityProvider) CreateArray() interface{} {
	return &[]domains.Party{}
}
