import { configureStore } from '@reduxjs/toolkit';
import { publicPartySlice } from './sclices/PublicPartySlice';
import { organizedPartySlice } from './sclices/OrganizedPartySlice';
import { attendedPartySlice } from './sclices/AttendedPartySlice';
import { partyInviteSlice } from './sclices/PartyInviteSlice';
import { friendInviteSlice } from './sclices/FriendInviteSlice';
import { friendSlice } from './sclices/FriendSlice';
import { persistedSelectedPartyReducer } from './sclices/PartySlice';
import { persistStore } from 'redux-persist';
import { drinkRequirementSlice } from './sclices/DrinkRequirementSlice';
import { foodRequirementSlice } from './sclices/FoodRequirementSlice';
import { drinkContributionSlice } from './sclices/DrinkContributionSlice';
import { foodContributionSlice } from './sclices/FoodContributionSlice';
import { partyParticipantsSlice } from './sclices/PartyParticipantSlice';
import { partyPendingInviteSlice } from './sclices/PendingInvitesForPartySlice';
import { userSlice } from './sclices/UserSlice';

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
  },
  middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware({
      serializableCheck: false, // quickfix for development todo: get back here
    }),
});

export const persistor = persistStore(store);

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;
