import {createAsyncThunk, createSlice} from '@reduxjs/toolkit';
import {getFriends} from '../apis/UserApi';
import {User} from '../types/User';

export interface FriendSlice {
    friends: User[];
    loading: boolean;
    error: Error | null;
}

const initialState: FriendSlice = {
    friends: [],
    loading: true,
    error: null,
};

export const loadFriends = createAsyncThunk(
    'data/loadFriends',
    async () => getFriends(),
);

const friendSlice = createSlice({
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

export default friendSlice.reducer;