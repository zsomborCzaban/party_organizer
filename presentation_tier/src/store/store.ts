import { configureStore } from '@reduxjs/toolkit'
import publicPartySlice from "../features/overView/discover/PublicPartySlice";
import organizedPartySlice from "../features/overView/partyPage/OrganizedPartySlice";
import attendedPartySlice from "../features/overView/partyPage/AttendedPartySlice";
import partyInviteSlice from "../features/overView/partyPage/PartyInviteSlice";

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