import { createSlice } from '@reduxjs/toolkit';

interface DrawerState {
  isOpen: boolean;
}

const initialState: DrawerState = {
  isOpen: false,
};

export const profileDrawersSlice = createSlice({
  name: 'drawer',
  initialState,
  reducers: {
    togglePartyProfileDrawer: (state) => {
      state.isOpen = !state.isOpen
    },
    openPartyProfileDrawer: (state) => {
      state.isOpen = true;
    },
    closePartyProfileDrawer: (state) => {
      state.isOpen = false;
    },
  },
});

export const { openPartyProfileDrawer, closePartyProfileDrawer, togglePartyProfileDrawer } = profileDrawersSlice.actions;