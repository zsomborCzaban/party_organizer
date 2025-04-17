import { createAsyncThunk, createSlice } from '@reduxjs/toolkit';
import { getFoodContributions } from '../../api/apis/ContributionApi';
import { Contribution } from '../../data/types/Contribution';

export interface FoodContributionState {
  contributions: Contribution[];
  loading: boolean;
  error: Error | null;
}

const initialState: FoodContributionState = {
  contributions: [],
  loading: true,
  error: null,
};

export const loadFoodContributions = createAsyncThunk('data/loadFoodContributions', async (partyId: number) => getFoodContributions(partyId));

export const foodContributionSlice = createSlice({
  name: 'foodContributions',
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(loadFoodContributions.pending, (state) => {
        state.loading = true;
        state.contributions = [];
        state.error = null;
      })
      .addCase(loadFoodContributions.fulfilled, (state, action) => {
        state.loading = false;
        state.contributions = action.payload ? action.payload : [];
        state.error = null;
      })
      .addCase(loadFoodContributions.rejected, (state, action) => {
        state.loading = false;
        state.contributions = [];
        state.error = new Error(action.error.message || 'Failed to load food contributions');
      });
  },
});
