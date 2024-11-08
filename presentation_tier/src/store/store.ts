import { configureStore } from '@reduxjs/toolkit'
import publicPartySlice from "../data/sclices/PublicPartySlice";
import organizedPartySlice from "../data/sclices/OrganizedPartySlice";
import attendedPartySlice from "../data/sclices/AttendedPartySlice";
import partyInviteSlice from "../data/sclices/PartyInviteSlice";
import friendInviteSlice from "../data/sclices/FriendInviteSlice";
import friendSlice from "../data/sclices/FriendSlice";
import persistedSelectedPartyReducer from "../data/sclices/PartySlice";
import { persistStore } from 'redux-persist';
import drinkRequirementSlice from "../data/sclices/DrinkRequirementSlice";
import foodRequirementSlice from "../data/sclices/FoodRequirementSlice";
import drinkContributionSlice from "../data/sclices/DrinkContributionSlice";
import foodContributionSlice from "../data/sclices/FoodContributionSlice";
import partyParticipantSlice from "../data/sclices/PartyParticipantSlice";
import partyPendingInviteSlice from "../data/sclices/PendingInvitesForPartySlice"

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