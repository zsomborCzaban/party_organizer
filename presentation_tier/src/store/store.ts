import { configureStore } from '@reduxjs/toolkit'
import publicPartySlice from "../features/overView/discover/PublicPartySlice";
import organizedPartySlice from "../features/overView/partiesPage/OrganizedPartySlice";
import attendedPartySlice from "../features/overView/partiesPage/AttendedPartySlice";
import partyInviteSlice from "../features/overView/partiesPage/PartyInviteSlice";
import friendInviteSlice from "../features/overView/friends/FriendInviteSlice";
import friendSlice from "../features/overView/friends/FriendSlice";

export const store = configureStore({
 reducer: {
     publicPartyStore: publicPartySlice,
     organizedPartyStore: organizedPartySlice,
     attendedPartyStore: attendedPartySlice,
     partyInviteStore: partyInviteSlice,
     friendInviteStore: friendInviteSlice,
     friendStore: friendSlice,
 },
});

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof  store.dispatch;