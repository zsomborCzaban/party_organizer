import {Party} from "../types/Party";
import {createAsyncThunk, createSlice} from "@reduxjs/toolkit";
import {getOrganizedParties} from "../apis/PartyApi";

export interface OrganizedPartySlice {
    parties: Party[];
    loading: boolean;
    error: Error | null;
}

const initialState: OrganizedPartySlice = {
    parties: [],
    loading: true,
    error: null,
}

export const loadOrganizedParties = createAsyncThunk(
    'data/loadOrganizedParties',
    async () => {
        try {
            return await getOrganizedParties();
        } catch (err) {
            console.log("err in loadOrganizedParties: " + err)
        }
    },
);

const organizedPartySlice = createSlice({
    name: 'publicParties',
    initialState,
    reducers: {},
    extraReducers: (builder) => {
        builder
            .addCase(loadOrganizedParties.pending, (state) => {
                state.loading = true;
                state.parties = [];
                state.error = null;
            })
            .addCase(loadOrganizedParties.fulfilled, (state, action) => {
                state.loading = false;
                state.parties = action.payload!;
                state.error = null;
            })
            .addCase(loadOrganizedParties.rejected, (state, action) => {
                state.loading = false;
                state.parties = [];
                state.error = new Error(action.error.message || 'Failed to load organized parties',);
            });
    },
});

export default organizedPartySlice.reducer