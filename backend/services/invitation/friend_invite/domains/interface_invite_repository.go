package domains

type IFriendInviteRepository interface {
	FindByIds(invitorId uint, invitedId uint) (*FriendInvitation, error)
	Update(*FriendInvitation) error
	Create(*FriendInvitation) error
}
