package domains

type IPartyInviteRepository interface {
	FindByIds(invitedId, partyId uint) (*PartyInvite, error)
	Save(*PartyInvite) error
	Update(*PartyInvite) error
	Create(*PartyInvite) error

	FindPendingByInvitedId(uint) (*[]PartyInvite, error)
}
