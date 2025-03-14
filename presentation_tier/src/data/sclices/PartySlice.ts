import {Party} from '../types/Party';
import {createSlice} from '@reduxjs/toolkit';
import storage from 'redux-persist/lib/storage';
import { persistReducer } from 'redux-persist';


export interface SelectedPartySlice {
    selectedParty : Party | null;
}

const initialState: SelectedPartySlice = {
    selectedParty: null,
};

const selectedPartySlice = createSlice({
        name: 'selectedParty',
        initialState,
        reducers: {
            setSelectedParty: (state, action:  { payload: Party, type: string }) => {
                state.selectedParty = action.payload; // Set data in the state
                // state.selectedParty.organizer = action.payload.organizer;
            },
        },
});

const persistantConfig = {
    key: 'selectedParty',
    storage,
};

const persistedSelectedPartyReducer = persistReducer(persistantConfig, selectedPartySlice.reducer);

export default persistedSelectedPartyReducer;
export const {setSelectedParty} = selectedPartySlice.actions;