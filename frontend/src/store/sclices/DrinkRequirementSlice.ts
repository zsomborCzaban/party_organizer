import { createAsyncThunk, createSlice } from '@reduxjs/toolkit';

import { getDrinkRequirements } from '../../api/apis/RequirementApi';
import { Requirement } from '../../data/types/Requirement';

export interface DrinkRequirementState {
  requirements: Requirement[];
  loading: boolean;
  error: Error | null;
}

const initialState: DrinkRequirementState = {
  requirements: [],
  loading: true,
  error: null,
};

export const loadDrinkRequirements = createAsyncThunk('data/loadDrinkRequirements', async (partyId: number) => getDrinkRequirements(partyId));

export const drinkRequirementSlice = createSlice({
  name: 'drinkRequirements',
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(loadDrinkRequirements.pending, (state) => {
        state.loading = true;
        state.requirements = [];
        state.error = null;
      })
      .addCase(loadDrinkRequirements.fulfilled, (state, action) => {
        state.loading = false;
        state.requirements = action.payload ? action.payload : [];
        state.error = null;
      })
      .addCase(loadDrinkRequirements.rejected, (state, action) => {
        state.loading = false;
        state.requirements = [];
        state.error = new Error(action.error.message || 'Failed to load drink requirements');
      });
  },
});
