package domains

type IFriendInviteRepository interface {
	FindByIds(invitorId uint, invitedId uint) (*FriendInvite, error)
	Update(*FriendInvite) error
	Create(*FriendInvite) error
}
