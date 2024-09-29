package usecases

import "github.com/zsomborCzaban/party_organizer/services/creation/party/domains"

type PartyRepository struct{}

func NewPartiesRepository() domains.IPartyRepository {
	return &PartyRepository{}
}

func (p PartyRepository) CreateParty(partyDTO *domains.PartyDTO) (*domains.Party, error) {
	//TODO implement me
	panic("implement me")
}

func (p PartyRepository) UpdateParty(id uint) (*domains.Party, error) {
	//TODO implement me
	panic("implement me")
}

func (p PartyRepository) GetParty(*domains.PartyDTO) (*domains.Party, error) {
	//TODO implement me
	panic("implement me")
}

func (p PartyRepository) DeleteParty(id uint) (*domains.Party, error) {
	//TODO implement me
	panic("implement me")
}
