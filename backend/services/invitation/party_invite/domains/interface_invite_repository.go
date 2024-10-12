package domains

type IPartyInviteRepository interface {
	FindByIds(invitedId, partyId uint) (*PartyInvite, error)
	Update(invite *PartyInvite) error
	Create(*PartyInvite) error

	FindPendingByInvitedId(uint) (*[]PartyInvite, error)
}
