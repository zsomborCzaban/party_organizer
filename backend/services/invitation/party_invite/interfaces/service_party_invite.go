package interfaces

import (
	"github.com/zsomborCzaban/party_organizer/common/api"
	partyDomains "github.com/zsomborCzaban/party_organizer/services/creation/party/domains"
	drinkContributionDomains "github.com/zsomborCzaban/party_organizer/services/interaction/drink_contributions/domains"
	foodContributionDomains "github.com/zsomborCzaban/party_organizer/services/interaction/food_contributions/domains"
	"github.com/zsomborCzaban/party_organizer/services/invitation/party_invite/domains"
	userDomain "github.com/zsomborCzaban/party_organizer/services/user/domains"
)

type PartyInviteService struct {
	PartyInviteRepository       domains.IPartyInviteRepository
	UserRepository              userDomain.IUserRepository
	PartyRepository             partyDomains.IPartyRepository
	FoodContributionRepository  foodContributionDomains.IFoodContributionRepository
	DrinkContributionRepository drinkContributionDomains.IDrinkContributionRepository
}

func NewPartyInviteService(repo domains.IPartyInviteRepository, userRepo userDomain.IUserRepository, partyRepo partyDomains.IPartyRepository, fContrRepo foodContributionDomains.IFoodContributionRepository, dContrRepo drinkContributionDomains.IDrinkContributionRepository) domains.IPartyInviteService {
	return &PartyInviteService{
		PartyInviteRepository:       repo,
		UserRepository:              userRepo,
		PartyRepository:             partyRepo,
		FoodContributionRepository:  fContrRepo,
		DrinkContributionRepository: dContrRepo,
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

	invitedUser, err3 := ps.UserRepository.FindById(invitedId)
	if err3 != nil {
		return api.ErrorInternalServerError(err3.Error())
	}

	party, err4 := ps.PartyRepository.FindById(partyId)
	if err4 != nil {
		return api.ErrorBadRequest(err4.Error())
	}

	//todo: put this in a transaction
	invite.State = domains.ACCEPTED
	if err5 := ps.PartyInviteRepository.Update(invite); err5 != nil {
		return api.ErrorInternalServerError(err5.Error())
	}

	if err6 := ps.PartyRepository.AddUserToParty(party, invitedUser); err6 != nil {
		return api.ErrorInternalServerError(err6.Error())
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
		return api.ErrorBadRequest(err.Error())
	}

	party, err2 := ps.PartyRepository.FindById(partyId)
	if err2 != nil {
		return api.ErrorBadRequest(err2.Error())
	}

	invited, err3 := ps.UserRepository.FindById(invitedId)
	if err3 != nil {
		return api.ErrorBadRequest(err3.Error())
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
func (ps PartyInviteService) GetPendingAndAcceptedInvites(partyId, userId uint) api.IResponse {
	party, err := ps.PartyRepository.FindById(partyId)
	if err != nil {
		return api.ErrorBadRequest(err.Error())
	}

	if !party.CanBeOrganizedBy(userId) {
		return api.ErrorUnauthorized("cannot organize this party")
	}

	invites, err := ps.PartyInviteRepository.FindPendingAndAcceptedByPartyId(partyId)
	if err != nil {
		return api.ErrorInternalServerError(err.Error())
	}

	return api.Success(invites)
}

func (ps PartyInviteService) JoinPublicParty(partyId, userId uint) api.IResponse {
	party, err := ps.PartyRepository.FindById(partyId)
	if err != nil {
		return api.ErrorBadRequest(err.Error())
	}

	user, err2 := ps.UserRepository.FindById(userId)
	if err2 != nil {
		return api.ErrorBadRequest(err2.Error())
	}

	if party.HasParticipant(userId) {
		return api.Success(party)
	}

	invite, err3 := ps.PartyInviteRepository.FindByIds(userId, partyId)
	if err3 != nil {
		invite = &domains.PartyInvite{
			InvitorId: party.OrganizerID,
			Invitor:   party.Organizer,
			InvitedId: userId,
			Invited:   *user,
			PartyId:   partyId,
			Party:     *party,
		}
	}
	invite.State = domains.ACCEPTED

	//todo: put this in transaction
	err4 := ps.PartyInviteRepository.Save(invite)
	if err4 != nil {
		return api.ErrorInternalServerError(err4.Error())
	}

	err5 := ps.PartyRepository.AddUserToParty(party, user)
	if err5 != nil {
		//todo: rollback
		return api.ErrorInternalServerError(err5.Error())
	}
	return api.Success(party)

}

func (ps PartyInviteService) JoinPrivateParty(partyId, userId uint, accessCode string) api.IResponse {
	party, err := ps.PartyRepository.FindById(partyId)
	if err != nil {
		return api.ErrorBadRequest(err.Error())
	}

	if !party.AccessCodeEnabled {
		return api.ErrorBadRequest("accessCode is not enabled for party")
	}

	if party.AccessCode != accessCode {
		return api.ErrorUnauthorized("invalid accessCode")
	}

	user, err2 := ps.UserRepository.FindById(userId)
	if err2 != nil {
		return api.ErrorBadRequest(err2.Error())
	}

	if party.HasParticipant(userId) {
		return api.Success(party)
	}

	invite, err3 := ps.PartyInviteRepository.FindByIds(userId, partyId)
	if err3 != nil {
		invite = &domains.PartyInvite{
			InvitorId: party.OrganizerID,
			Invitor:   party.Organizer,
			InvitedId: userId,
			Invited:   *user,
			PartyId:   partyId,
			Party:     *party,
		}
	}
	invite.State = domains.ACCEPTED

	//todo: put this in transaction
	err4 := ps.PartyInviteRepository.Save(invite)
	if err4 != nil {
		return api.ErrorInternalServerError(err4.Error())
	}
	if err5 := ps.PartyRepository.AddUserToParty(party, user); err5 != nil {
		//todo: rollback
		return api.ErrorInternalServerError(err5.Error())
	}
	return api.Success(party)
}

func (ps PartyInviteService) Kick(kickedId, userId, partyId uint) api.IResponse {
	kickedUser, err := ps.UserRepository.FindById(kickedId)
	if err != nil {
		return api.ErrorInternalServerError(err.Error())
	}

	party, err3 := ps.PartyRepository.FindById(partyId)
	if err3 != nil {
		return api.ErrorBadRequest(err3.Error())
	}

	if !party.CanBeOrganizedBy(userId) && kickedId != userId {
		return api.ErrorUnauthorized(domains.UNAUTHORIZED)
	}
	if party.OrganizerID == kickedId {
		return api.ErrorUnauthorized("The organizer cannot leave the party.")
	}
	if !party.HasParticipant(kickedId) {
		return api.Success("user kicked successfully")
	}

	invite, err4 := ps.PartyInviteRepository.FindByIds(kickedId, partyId)
	if err4 != nil {
		return api.ErrorInternalServerError(err4.Error())
	}

	//todo: put this in a transaction
	if err5 := ps.FoodContributionRepository.DeleteByContributorId(kickedId); err5 != nil {
		return api.ErrorInternalServerError(err5.Error())
	}
	if err6 := ps.DrinkContributionRepository.DeleteByContributorId(kickedId); err6 != nil {
		return api.ErrorInternalServerError(err6.Error())
	}

	if err7 := ps.PartyRepository.RemoveUserFromParty(party, kickedUser); err7 != nil {
		return api.ErrorInternalServerError(err7.Error())
	}

	invite.State = domains.DECLINED
	if err8 := ps.PartyInviteRepository.Update(invite); err8 != nil {
		return api.ErrorInternalServerError(err8.Error())
	}

	return api.Success("user kicked successfully")
}
