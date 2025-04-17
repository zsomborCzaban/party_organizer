import { createSlice } from '@reduxjs/toolkit';
import storage from 'redux-persist/lib/storage';
import { persistReducer } from 'redux-persist';
import { Party } from '../../data/types/Party';

export interface SelectedPartyState {
  selectedParty?: Party;
}

const initialState: SelectedPartyState = {
  selectedParty: undefined,
};

export const selectedPartySlice = createSlice({
  name: 'selectedParty',
  initialState,
  reducers: {
    setSelectedParty: (state, action: { payload: Party; type: string }) => {
      state.selectedParty = action.payload; // Set data in the state
    },
  },
});

const persistantConfig = {
  key: 'selectedParty',
  storage,
};

export const persistedSelectedPartyReducer = persistReducer(persistantConfig, selectedPartySlice.reducer);

export const { setSelectedParty } = selectedPartySlice.actions;
