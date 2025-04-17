import { createAsyncThunk, createSlice } from '@reduxjs/toolkit';
import { getPartyParticipants } from '../../api/apis/PartyApi';
import { User } from '../../data/types/User';

export interface PartParticipantsState {
  participants: User[];
  loading: boolean;
  error: Error | null;
}

const initialState: PartParticipantsState = {
  participants: [],
  loading: true,
  error: null,
};

export const loadPartyParticipants = createAsyncThunk('data/loadPartyParticipants', async (partyId: number) => getPartyParticipants(partyId));

export const partyParticipantsSlice = createSlice({
  name: 'data/partyParticipants',
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(loadPartyParticipants.pending, (state) => {
        state.loading = true;
        state.participants = [];
        state.error = null;
      })
      .addCase(loadPartyParticipants.fulfilled, (state, action) => {
        state.loading = false;
        state.participants = action.payload ? action.payload : [];
        state.error = null;
      })
      .addCase(loadPartyParticipants.rejected, (state, action) => {
        state.loading = false;
        state.participants = [];
        state.error = new Error(action.error.message || 'Failed to load party participants');
      });
  },
});
