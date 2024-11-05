import { configureStore } from '@reduxjs/toolkit'
import publicPartySlice from "../features/overView/discover/PublicPartySlice";
import organizedPartySlice from "../features/overView/partiesPage/OrganizedPartySlice";
import attendedPartySlice from "../features/overView/partiesPage/AttendedPartySlice";
import partyInviteSlice from "../features/overView/partiesPage/PartyInviteSlice";
import friendInviteSlice from "../features/overView/friends/FriendInviteSlice";
import friendSlice from "../features/overView/friends/FriendSlice";
import persistedSelectedPartyReducer from "../features/overView/PartySlice";
import { persistStore } from 'redux-persist';
import drinkRequirementSlice from "../features/visitParty/data/slices/DrinkRequirementSlice";
import foodRequirementSlice from "../features/visitParty/data/slices/FoodRequirementSlice";
import drinkContributionSlice from "../features/visitParty/data/slices/DrinkContributionSlice";
import foodContributionSlice from "../features/visitParty/data/slices/FoodContributionSlice";
import partyParticipantSlice from "../features/visitParty/data/slices/PartyParticipantSlice";
import partyPendingInviteSlice from "../features/visitParty/data/slices/PendingInvitesForPartySlice"

export const store = configureStore({
    reducer: {
        publicPartyStore: publicPartySlice,
        organizedPartyStore: organizedPartySlice,
        attendedPartyStore: attendedPartySlice,
        partyInviteStore: partyInviteSlice,
        friendInviteStore: friendInviteSlice,
        friendStore: friendSlice,
        selectedPartyStore: persistedSelectedPartyReducer,
        drinkRequirementStore: drinkRequirementSlice,
        foodRequirementStore: foodRequirementSlice,
        drinkContributionStore: drinkContributionSlice,
        foodContributionStore: foodContributionSlice,
        partyParticipantStore: partyParticipantSlice,
        partyPendingInviteStore: partyPendingInviteSlice,
    },
    middleware: (getDefaultMiddleware) =>
        getDefaultMiddleware({
            serializableCheck: false, //quickfix for development todo: get back here
    }),
});

export const persistor = persistStore(store);

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof  store.dispatch;