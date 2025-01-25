import { createAsyncThunk, createSlice } from '@reduxjs/toolkit';
import { getPartyPendingInvites } from '../../api/apis/PartyAttendanceManagerApi';
import { PartyInvite } from '../../data/types/PartyInvite';

export interface PartyPendingInvitesState {
  pendingInvites: PartyInvite[];
  loading: boolean;
  error: Error | null;
}

const initialState: PartyPendingInvitesState = {
  pendingInvites: [],
  loading: true,
  error: null,
};

export const loadPartyPendingInvites = createAsyncThunk('data/loadPartyPendingInvites', async (partyId: number) => getPartyPendingInvites(partyId));

export const partyPendingInviteSlice = createSlice({
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
        state.error = new Error(action.error.message || 'Failed to load party pending invites');
      });
  },
});
