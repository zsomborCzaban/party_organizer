import {createAsyncThunk, createSlice} from "@reduxjs/toolkit";
import {getPartyPendingInvites} from "../apis/PartyAttendanceManagerApi";
import {PartyInvite} from "../types/PartyInvite";

export interface PartyPendingInvitesSlice {
    pendingInvites: PartyInvite[];
    loading: boolean;
    error: Error | null;
}

const initialState: PartyPendingInvitesSlice = {
    pendingInvites: [],
    loading: true,
    error: null,
}

export const loadPartyPendingInvites = createAsyncThunk(
    'data/loadPartyPendingInvites',
    async (partyId: number) => {
        return getPartyPendingInvites(partyId)
    },
);

const partyPartyPendingInviteSlice = createSlice({
    name: 'data/partyPendingInvites',
    initialState,
    reducers: {},
    extraReducers: (builder) => {
        builder
            .addCase(loadPartyPendingInvites.pending, (state) => {
                state.loading = true;
                state.pendingInvites = [];
                state.error = null;
            })
            .addCase(loadPartyPendingInvites.fulfilled, (state, action) => {
                state.loading = false;
                state.pendingInvites = action.payload ? action.payload : [];
                state.error = null;
            })
            .addCase(loadPartyPendingInvites.rejected, (state, action) => {
                state.loading = false;
                state.pendingInvites = [];
                state.error = new Error(action.error.message || 'Failed to load party pending invites',);
            });
    },
});

export default partyPartyPendingInviteSlice.reducer