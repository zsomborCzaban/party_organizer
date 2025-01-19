import {Party} from '../types/Party';
import {createAsyncThunk, createSlice} from '@reduxjs/toolkit';
import {getPublicParties} from '../apis/PartyApi';

export interface PublicPartySlice {
    parties: Party[];
    loading: boolean;
    error: Error | null;
}

const initialState: PublicPartySlice = {
    parties: [],
    loading: true,
    error: null,
};

export const loadPublicParties = createAsyncThunk(
    'data/loadPublicParties',
    async () => getPublicParties(),
);

const publicPartySlice = createSlice({
    name: 'publicParties',
    initialState,
    reducers: {},
    extraReducers: (builder) => {
        builder
            .addCase(loadPublicParties.pending, (state) => {
                state.loading = true;
                state.parties = [];
                state.error = null;
            })
            .addCase(loadPublicParties.fulfilled, (state, action) => {
                state.loading = false;
                state.error = null;
                state.parties = action.payload ? action.payload : [];
            })
            .addCase(loadPublicParties.rejected, (state, action) => {
                state.loading = false;
                state.parties = [];
                state.error = new Error(action.error.message || 'Failed to load public parties');
            });
    },
});

export default publicPartySlice.reducer;