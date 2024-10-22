import { configureStore } from '@reduxjs/toolkit'
import publicPartySlice from "../features/overView/discover/PublicPartySlice";
import organizedPartySlice from "../features/overView/partiesPage/OrganizedPartySlice";
import attendedPartySlice from "../features/overView/partiesPage/AttendedPartySlice";
import partyInviteSlice from "../features/overView/partiesPage/PartyInviteSlice";

export const store = configureStore({
 reducer: {
     publicPartyStore: publicPartySlice,
     organizedPartyStore: organizedPartySlice,
     attendedPartyStore: attendedPartySlice,
     partyInviteStore: partyInviteSlice,
 },
});

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof  store.dispatch;