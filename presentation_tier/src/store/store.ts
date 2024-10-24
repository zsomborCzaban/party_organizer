import { configureStore } from '@reduxjs/toolkit'
import publicPartySlice from "../features/overView/discover/PublicPartySlice";
import organizedPartySlice from "../features/overView/partiesPage/OrganizedPartySlice";
import attendedPartySlice from "../features/overView/partiesPage/AttendedPartySlice";
import partyInviteSlice from "../features/overView/partiesPage/PartyInviteSlice";
import friendInviteSlice from "../features/overView/friends/FriendInviteSlice";
import friendSlice from "../features/overView/friends/FriendSlice";
import persistedSelectedPartyReducer from "../features/overView/PartySlice";
import { persistStore } from 'redux-persist';

export const store = configureStore({
    reducer: {
        publicPartyStore: publicPartySlice,
        organizedPartyStore: organizedPartySlice,
        attendedPartyStore: attendedPartySlice,
        partyInviteStore: partyInviteSlice,
        friendInviteStore: friendInviteSlice,
        friendStore: friendSlice,
        selectedPartyStore: persistedSelectedPartyReducer,
    },
    middleware: (getDefaultMiddleware) =>
        getDefaultMiddleware({
            serializableCheck: false, //quickfix for development todo: get back here
    }),
});

export const persistor = persistStore(store);

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof  store.dispatch;