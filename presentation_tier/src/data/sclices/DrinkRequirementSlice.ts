import {createAsyncThunk, createSlice} from "@reduxjs/toolkit";
import {Requirement} from "../types/Requirement";
import {getDrinkRequirements} from "../apis/RequirementApi";

export interface DrinkRequirementSlice {
    requirements: Requirement[];
    loading: boolean;
    error: Error | null;
}

const initialState: DrinkRequirementSlice = {
    requirements: [],
    loading: true,
    error: null,
}

export const loadDrinkRequirements = createAsyncThunk(
    'data/loadDrinkRequirements',
    async (partyId: number) => {
        return getDrinkRequirements(partyId)
    },
);

const drinkRequirementSlice = createSlice({
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
                state.error = new Error(action.error.message || 'Failed to load drink requirements',);
            });
    },
});

export default drinkRequirementSlice.reducer