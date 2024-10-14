package interfaces

import (
	"fmt"
	"github.com/zsomborCzaban/party_organizer/common/api"
	partyDomains "github.com/zsomborCzaban/party_organizer/services/creation/party/domains"
	"github.com/zsomborCzaban/party_organizer/services/invitation/party_invite/domains"
	userDomain "github.com/zsomborCzaban/party_organizer/services/user/domains"
)

type PartyInviteService struct {
	PartyInviteRepository domains.IPartyInviteRepository
	UserRepository        userDomain.IUserRepository
	PartyRepository       partyDomains.IPartyRepository
}

func NewPartyInviteService(repo domains.IPartyInviteRepository, userRepo userDomain.IUserRepository, partyRepo partyDomains.IPartyRepository) domains.IPartyInviteService {
	return &PartyInviteService{
		PartyInviteRepository: repo,
		UserRepository:        userRepo,
		PartyRepository:       partyRepo,
	}
}

func (ps PartyInviteService) Accept(invitedId, partyId uint) api.IResponse {
	invite, err := ps.PartyInviteRepository.FindByIds(invitedId, partyId)
	if err != nil {
		return api.ErrorInternalServerError(err.Error())
	}

	if invite.State == domains.DECLINED {
		return api.ErrorBadRequest("Cannot accept already declined parties. Try inviting them")
	}

	if invite.State == domains.ACCEPTED {
		return api.Success(invite)
	}

	//todo: put this in a transaction

	invite.State = domains.ACCEPTED
	if err2 := ps.PartyInviteRepository.Update(invite); err2 != nil {
		return api.ErrorInternalServerError(err2.Error())
	}

	invitedUser, err3 := ps.UserRepository.FindById(invitedId)
	if err3 != nil {
		return api.ErrorInternalServerError(err3.Error())
	}

	party, err4 := ps.PartyRepository.FindById(partyId)
	if err4 != nil {
		return api.ErrorBadRequest(err4.Error())
	}

	if err5 := ps.PartyRepository.AddUserToParty(party, invitedUser); err4 != nil {
		return api.ErrorInternalServerError(err5.Error())
	}

	return api.Success(invite)
}

func (ps PartyInviteService) Decline(invitedId, partyId uint) api.IResponse {
	invite, err := ps.PartyInviteRepository.FindByIds(invitedId, partyId)
	if err != nil {
		return api.ErrorInternalServerError(err.Error())
	}

	if invite.State == domains.ACCEPTED {
		return api.ErrorBadRequest("Cannot decline already accepted parties. Try deleting them")
	}

	if invite.State == domains.DECLINED {
		return api.Success(invite)
	}

	invite.State = domains.DECLINED
	if err2 := ps.PartyInviteRepository.Update(invite); err2 != nil {
		return api.ErrorInternalServerError(err2.Error())
	}

	return api.Success(invite)
}

func (ps PartyInviteService) Invite(invitedId, invitorId, partyId uint) api.IResponse {
	if invitorId == invitedId {
		return api.ErrorBadRequest("cannot party invite yourself")
	}

	invite, err := ps.PartyInviteRepository.FindByIds(invitedId, partyId)
	if err != nil && err.Error() == domains.NOT_FOUND {
		return ps.CreateInvitation(invitedId, invitorId, partyId)
	}
	if err != nil {
		return api.ErrorInternalServerError(err.Error())
	}

	if invite.State == domains.ACCEPTED {
		return api.ErrorBadRequest("User already accepted the invite")
	}

	if invite.State == domains.PENDING {
		return api.Success(invite)
	}

	invite.State = domains.PENDING
	if err2 := ps.PartyInviteRepository.Update(invite); err2 != nil {
		return api.ErrorInternalServerError(err2.Error())
	}

	return api.Success(invite)
}

func (ps PartyInviteService) CreateInvitation(invitedId, invitorId, partyId uint) api.IResponse {
	invitor, err := ps.UserRepository.FindById(invitorId)
	if err != nil {
		return api.ErrorBadRequest(fmt.Sprintf("cannot find user with id: %d", invitorId))
	}

	party, err2 := ps.PartyRepository.FindById(partyId)
	if err2 != nil {
		return api.ErrorBadRequest(fmt.Sprintf("cannot find party with id: %d", partyId))
	}

	invited, err3 := ps.UserRepository.FindById(invitedId)
	if err3 != nil {
		return api.ErrorBadRequest(fmt.Sprintf("cannot find user with id: %d", invitedId))
	}

	//would be faster is this would be on top. but the code is more clear this way
	if invitorId != party.OrganizerID {
		return api.ErrorUnauthorized("cannot invite users for other people's party")
	}

	invitation := &domains.PartyInvite{
		InvitorId: invitorId,
		Invited:   *invited,
		PartyId:   partyId,
		Party:     *party,
		InvitedId: invitedId,
		Invitor:   *invitor,
		State:     domains.PENDING,
	}

	if errCreation := ps.PartyInviteRepository.Create(invitation); errCreation != nil {
		return api.ErrorInternalServerError(errCreation.Error())
	}
	return api.Success(invitation)
}

func (ps PartyInviteService) GetPendingInvites(userId uint) api.IResponse {
	invites, err := ps.PartyInviteRepository.FindPendingByInvitedId(userId)
	if err != nil {
		return api.ErrorInternalServerError(err.Error())
	}

	return api.Success(invites)
}
