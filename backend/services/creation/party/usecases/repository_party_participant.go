package usecases

import (
	"github.com/zsomborCzaban/party_organizer/db"
	"github.com/zsomborCzaban/party_organizer/services/creation/party/domains"
)

type PartyParticipantsRepository struct {
	DbAccess db.IDatabaseAccess
}

func NewPartyParticipantsRepository(databaseAccessManager db.IDatabaseAccessManager) *PartyParticipantsRepository {
	entityProvider := EntityProvider2{}
	databaseAccess := databaseAccessManager.RegisterEntity("partyParticipantProvider", entityProvider)

	return &PartyParticipantsRepository{
		DbAccess: databaseAccess,
	}
}

type EntityProvider2 struct {
}

func (e EntityProvider2) Create() interface{} {
	return &domains.Party{}
}

func (e EntityProvider2) CreateArray() interface{} {
	return &[]domains.Party{}
}
