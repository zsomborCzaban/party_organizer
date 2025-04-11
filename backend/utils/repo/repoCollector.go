package repo

import (
	domains3 "github.com/zsomborCzaban/party_organizer/services/creation/drink_requirement/domains"
	domains5 "github.com/zsomborCzaban/party_organizer/services/creation/food_requirement/domains"
	"github.com/zsomborCzaban/party_organizer/services/creation/party/domains"
	domains4 "github.com/zsomborCzaban/party_organizer/services/interaction/drink_contributions/domains"
	domains6 "github.com/zsomborCzaban/party_organizer/services/interaction/food_contributions/domains"
	domains8 "github.com/zsomborCzaban/party_organizer/services/managers/friend_manager/domains"
	domains7 "github.com/zsomborCzaban/party_organizer/services/managers/party_attendance_manager/domains"
	domains9 "github.com/zsomborCzaban/party_organizer/services/users/registration/domains"
	domains2 "github.com/zsomborCzaban/party_organizer/services/users/user/domains"
)

// we could use a map for the repositories so it would be dynamic
type RepoCollector struct {
	PartyRepo        domains.IPartyRepository
	UserRepo         domains2.IUserRepository
	DrinkReqRepo     domains3.IDrinkRequirementRepository
	DrinkContribRepo domains4.IDrinkContributionRepository
	FoodReqRepo      domains5.IFoodRequirementRepository
	FoodContribRepo  domains6.IFoodContributionRepository
	PartyInviteRepo  domains7.IPartyInviteRepository
	FriendInviteRepo domains8.IFriendInviteRepository
	RegistrationRepo domains9.IRegistrationRepository
}

func NewRepoCollector(
	partyRepo *domains.IPartyRepository,
	userRepo *domains2.IUserRepository,
	drinkReqRepo *domains3.IDrinkRequirementRepository,
	drinkContribRepo *domains4.IDrinkContributionRepository,
	foodReqRepo *domains5.IFoodRequirementRepository,
	foodContribRepo *domains6.IFoodContributionRepository,
	partyInviteRepo *domains7.IPartyInviteRepository,
	friendInviteRepo *domains8.IFriendInviteRepository,
	registrationRepo *domains9.IRegistrationRepository,
) *RepoCollector {
	return &RepoCollector{
		PartyRepo:        *partyRepo,
		UserRepo:         *userRepo,
		DrinkReqRepo:     *drinkReqRepo,
		DrinkContribRepo: *drinkContribRepo,
		FoodReqRepo:      *foodReqRepo,
		FoodContribRepo:  *foodContribRepo,
		PartyInviteRepo:  *partyInviteRepo,
		FriendInviteRepo: *friendInviteRepo,
		RegistrationRepo: *registrationRepo,
	}
}
