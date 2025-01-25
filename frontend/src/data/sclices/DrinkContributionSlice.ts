import {createAsyncThunk, createSlice} from '@reduxjs/toolkit';
import {Contribution} from '../types/Contribution';
import { getDrinkContributions } from '../../api/apis/ContributionApi';


export interface DrinkContributionSlice {
    contributions: Contribution[];
    loading: boolean;
    error: Error | null;
}

const initialState: DrinkContributionSlice = {
    contributions: [],
    loading: true,
    error: null,
};

export const loadDrinkContributions = createAsyncThunk(
    'data/loadDrinkContributions',
    async (partyId: number) => getDrinkContributions(partyId),
);

const drinkContributionSlice = createSlice({
    name: 'drinkContributions',
    initialState,
    reducers: {},
    extraReducers: (builder) => {
        builder
            .addCase(loadDrinkContributions.pending, (state) => {
                state.loading = true;
                state.contributions = [];
                state.error = null;
            })
            .addCase(loadDrinkContributions.fulfilled, (state, action) => {
                state.loading = false;
                state.contributions = action.payload ? action.payload : [];
                state.error = null;
            })
            .addCase(loadDrinkContributions.rejected, (state, action) => {
                state.loading = false;
                state.contributions = [];
                state.error = new Error(action.error.message || 'Failed to load drink contributions');
            });
    },
});

export default drinkContributionSlice.reducer;