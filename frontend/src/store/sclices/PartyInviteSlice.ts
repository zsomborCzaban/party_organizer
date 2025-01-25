import { createAsyncThunk, createSlice } from '@reduxjs/toolkit';
import { getPartyInvites } from '../../api/apis/PartyAttendanceManagerApi';
import { PartyInvite } from '../../data/types/PartyInvite';

export interface PartyInviteState {
  invites: PartyInvite[];
  loading: boolean;
  error: Error | null;
}

const initialState: PartyInviteState = {
  invites: [],
  loading: true,
  error: null,
};

export const loadPartyInvites = createAsyncThunk('data/loadPartyInvites', async () => getPartyInvites());

export const partyInviteSlice = createSlice({
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
        state.invites = action.payload ? action.payload : [];
        state.error = null;
      })
      .addCase(loadPartyInvites.rejected, (state, action) => {
        state.loading = false;
        state.invites = [];
        state.error = new Error(action.error.message || 'Failed to load party invites');
      });
  },
});
