import { createAsyncThunk, createSlice } from '@reduxjs/toolkit';
import { getFoodRequirements } from '../../api/apis/RequirementApi';
import { Requirement } from '../../data/types/Requirement';

export interface FoodRequirementState {
  requirements: Requirement[];
  loading: boolean;
  error: Error | null;
}

const initialState: FoodRequirementState = {
  requirements: [],
  loading: true,
  error: null,
};

export const loadFoodRequirements = createAsyncThunk('data/loadFoodRequirements', async (partyId: number) => getFoodRequirements(partyId));

export const foodRequirementSlice = createSlice({
  name: 'foodRequirements',
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(loadFoodRequirements.pending, (state) => {
        state.loading = true;
        state.requirements = [];
        state.error = null;
      })
      .addCase(loadFoodRequirements.fulfilled, (state, action) => {
        state.loading = false;
        state.requirements = action.payload ? action.payload : [];
        state.error = null;
      })
      .addCase(loadFoodRequirements.rejected, (state, action) => {
        state.loading = false;
        state.requirements = [];
        state.error = new Error(action.error.message || 'Failed to load food requirements');
      });
  },
});
