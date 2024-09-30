package usecases

import (
	"github.com/zsomborCzaban/party_organizer/services/creation/party/domains"
	"gorm.io/gorm"
)

type PartyRepository struct{}

func NewPartiesRepository() domains.IPartyRepository {
	return &PartyRepository{}
}

func (p PartyRepository) CreateParty(partyDTO *domains.PartyDTO) (*domains.Party, error) {

	return &domains.Party{
			Place: "itt",
		},
		nil
}

func (p PartyRepository) GetParty(id uint) (*domains.Party, error) {
	//TODO implement me
	return &domains.Party{
			Model: gorm.Model{ID: 23},
			Place: "itt",
		},
		nil
}

func (p PartyRepository) UpdateParty(*domains.PartyDTO) (*domains.Party, error) {
	//TODO implement me
	return &domains.Party{
			Model: gorm.Model{ID: 23},
			Place: "itt",
		},
		nil
}

func (p PartyRepository) DeleteParty(id uint) (*domains.Party, error) {
	//TODO implement me
	return &domains.Party{
			Model: gorm.Model{ID: 23},
			Place: "itt",
		},
		nil
}
