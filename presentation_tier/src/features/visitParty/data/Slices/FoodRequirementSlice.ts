import {createAsyncThunk, createSlice} from "@reduxjs/toolkit";
import {Requirement} from "../Requirement";
import {getFoodRequirements} from "../VisitPartyApi";

export interface FoodRequirementSlice {
    requirements: Requirement[];
    loading: boolean;
    error: Error | null;
}

const initialState: FoodRequirementSlice = {
    requirements: [],
    loading: true,
    error: null,
}

export const loadFoodRequirements = createAsyncThunk(
    'data/loadFoodRequirements',
    async (partyId: number) => {
        try {
            return await getFoodRequirements(partyId);
        } catch (err) {
            console.log("err in loadFoodRequirements: " + err)
        }
    },
);

const foodRequirementSlice = createSlice({
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
                state.error = new Error(action.error.message || 'Failed to load food requirements',);
            });
    },
});

export default foodRequirementSlice.reducer