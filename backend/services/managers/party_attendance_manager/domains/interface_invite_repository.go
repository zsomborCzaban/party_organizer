package domains

type IPartyInviteRepository interface {
	FindByIds(invitedId, partyId uint) (*PartyInvite, error)
	Save(*PartyInvite) error
	Update(*PartyInvite) error
	Create(*PartyInvite) error
	DeleteByPartyId(uint) error

	FindPendingByInvitedId(uint) (*[]PartyInvite, error)
	FindPendingByPartyId(uint) (*[]PartyInvite, error)
}
