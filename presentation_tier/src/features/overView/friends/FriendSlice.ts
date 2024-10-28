import {createAsyncThunk, createSlice} from "@reduxjs/toolkit";
import {getFriends} from "./FriendPageApi";
import {User} from "../User";

export interface FriendSlice {
    friends: User[];
    loading: boolean;
    error: Error | null;
}

const initialState: FriendSlice = {
    friends: [],
    loading: true,
    error: null,
}

export const loadFriends = createAsyncThunk(
    'data/loadFriends',
    async () => {
        try {
            return await getFriends();
        } catch (err) {
            console.log("err in loadFriends: " + err)
        }
    },
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
                state.error = new Error(action.error.message || 'Failed to load friends',);
            });
    },
});

export default friendSlice.reducer