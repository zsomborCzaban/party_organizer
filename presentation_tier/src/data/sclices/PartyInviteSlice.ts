import {createAsyncThunk, createSlice} from "@reduxjs/toolkit";
import {getPartyInvites} from "../apis/PartyAttendanceManagerApi";
import {PartyInvite} from "../types/PartyInvite";

export interface PartyInviteSlice {
    invites: PartyInvite[];
    loading: boolean;
    error: Error | null;
}

const initialState: PartyInviteSlice = {
    invites: [],
    loading: true,
    error: null,
}

export const loadPartyInvites = createAsyncThunk(
    'data/loadPartyInvites',
    async () => {
        try {
            return await getPartyInvites();
        } catch (err) {
            console.log("err in loadPartyInvites: " + err)
        }
    },
);

const partyInviteSlice = createSlice({
    name: 'partyInvites',
    initialState,
    reducers: {},
    extraReducers: (builder) => {
        builder
            .addCase(loadPartyInvites.pending, (state) => {
                state.loading = true;
                state.invites = [];
                state.error = null;
            })
            .addCase(loadPartyInvites.fulfilled, (state, action) => {
                state.loading = false;
                state.invites = action.payload!;
                state.error = null;
            })
            .addCase(loadPartyInvites.rejected, (state, action) => {
                state.loading = false;
                state.invites = [];
                state.error = new Error(action.error.message || 'Failed to load party invites',);
            });
    },
});

export default partyInviteSlice.reducer