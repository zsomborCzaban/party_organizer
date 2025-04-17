import { getOrganizedParties } from '../../api/apis/PartyApi';
import { createAsyncThunk, createSlice } from '@reduxjs/toolkit';
import { Party } from '../../data/types/Party';

export interface OrganizedPartyState {
  parties: Party[];
  loading: boolean;
  error: Error | null;
}

const initialState: OrganizedPartyState = {
  parties: [],
  loading: true,
  error: null,
};

export const loadOrganizedParties = createAsyncThunk('data/loadOrganizedParties', async () => getOrganizedParties());

export const organizedPartySlice = createSlice({
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
        state.error = new Error(action.error.message || 'Failed to load organized parties');
      });
  },
});
