import { createAsyncThunk, createSlice } from '@reduxjs/toolkit';
import { getFriendInvites } from '../../api/apis/FriendInviteManagerApi';
import { FriendInvite } from '../../data/types/FriendInvite';

export interface FriendInviteState {
  invites: FriendInvite[];
  loading: boolean;
  error: Error | null;
}

const initialState: FriendInviteState = {
  invites: [],
  loading: true,
  error: null,
};

export const loadFriendInvites = createAsyncThunk('data/loadFriendInvites', async () => getFriendInvites());

export const friendInviteSlice = createSlice({
  name: 'friendInvites',
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(loadFriendInvites.pending, (state) => {
        state.loading = true;
        state.invites = [];
        state.error = null;
      })
      .addCase(loadFriendInvites.fulfilled, (state, action) => {
        state.loading = false;
        state.invites = action.payload ? action.payload : [];
        state.error = null;
      })
      .addCase(loadFriendInvites.rejected, (state, action) => {
        state.loading = false;
        state.invites = [];
        state.error = new Error(action.error.message || 'Failed to load friend invites');
      });
  },
});
