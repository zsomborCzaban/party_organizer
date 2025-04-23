import { configureStore } from '@reduxjs/toolkit';
import { publicPartySlice } from './slices/PublicPartySlice';
import { organizedPartySlice } from './slices/OrganizedPartySlice';
import { attendedPartySlice } from './slices/AttendedPartySlice';
import { partyInviteSlice } from './slices/PartyInviteSlice';
import { friendInviteSlice } from './slices/FriendInviteSlice';
import { friendSlice } from './slices/FriendSlice';
import { persistedSelectedPartyReducer } from './slices/PartySlice';
import { persistStore } from 'redux-persist';
import { drinkRequirementSlice } from './slices/DrinkRequirementSlice';
import { foodRequirementSlice } from './slices/FoodRequirementSlice';
import { drinkContributionSlice } from './slices/DrinkContributionSlice';
import { foodContributionSlice } from './slices/FoodContributionSlice';
import { partyParticipantsSlice } from './slices/PartyParticipantSlice';
import { partyPendingInviteSlice } from './slices/PendingInvitesForPartySlice';
import { userSlice } from './slices/UserSlice';
import { profileDrawersSlice } from "./slices/profileDrawersSlice.ts";

export const store = configureStore({
  reducer: {
    publicPartyStore: publicPartySlice.reducer,
    organizedPartyStore: organizedPartySlice.reducer,
    attendedPartyStore: attendedPartySlice.reducer,
    partyInviteStore: partyInviteSlice.reducer,
    friendInviteStore: friendInviteSlice.reducer,
    friendStore: friendSlice.reducer,
    selectedPartyStore: persistedSelectedPartyReducer,
    drinkRequirementStore: drinkRequirementSlice.reducer,
    foodRequirementStore: foodRequirementSlice.reducer,
    drinkContributionStore: drinkContributionSlice.reducer,
    foodContributionStore: foodContributionSlice.reducer,
    partyParticipantStore: partyParticipantsSlice.reducer,
    partyPendingInviteStore: partyPendingInviteSlice.reducer,
    userStore: userSlice.reducer,
    profileDrawers: profileDrawersSlice.reducer,
  },
  middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware({
      serializableCheck: false, // quickfix for development todo: get back here
    }),
});

export const persistor = persistStore(store);

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;
