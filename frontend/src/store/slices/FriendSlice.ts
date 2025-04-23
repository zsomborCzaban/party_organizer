import { createAsyncThunk, createSlice } from '@reduxjs/toolkit';
import { getFriends } from '../../api/apis/UserApi';
import { User } from '../../data/types/User';

export interface FriendState {
  friends: User[];
  loading: boolean;
  error: Error | null;
}

const initialState: FriendState = {
  friends: [],
  loading: true,
  error: null,
};

export const loadFriends = createAsyncThunk('data/loadFriends', async () => getFriends());

export const friendSlice = createSlice({
  name: 'friends',
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(loadFriends.pending, (state) => {
        state.loading = true;
        state.friends = [];
        state.error = null;
      })
      .addCase(loadFriends.fulfilled, (state, action) => {
        state.loading = false;
        state.friends = action.payload ? action.payload : [];
        state.error = null;
      })
      .addCase(loadFriends.rejected, (state, action) => {
        state.loading = false;
        state.friends = [];
        state.error = new Error(action.error.message || 'Failed to load friends');
      });
  },
});
