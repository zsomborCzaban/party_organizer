import {Party} from "./Party";
import {createSlice} from "@reduxjs/toolkit";

export interface SelectedPartySlice {
    selectedParty : Party | null;
}

const initialState: SelectedPartySlice = {
    selectedParty: null,
}

const selectedPartySlice = createSlice({
        name: 'selectedParty',
        initialState,
        reducers: {
            setSelectedParty: (state, action:  { payload: Party, type: string }) => {
                state.selectedParty = action.payload; // Set data in the state
            },
        },
});

export default selectedPartySlice.reducer;
export const setSelectedParty = selectedPartySlice.actions.setSelectedParty