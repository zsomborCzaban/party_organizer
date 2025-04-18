import { createSlice } from '@reduxjs/toolkit';

interface DrawerState {
  isPartyProfileOpen: boolean;
  isDefaultProfileOpen: boolean;
}

const initialState: DrawerState = {
  isDefaultProfileOpen: false,
  isPartyProfileOpen: false,
};

export const profileDrawersSlice = createSlice({
  name: 'drawer',
  initialState,
  reducers: {
    togglePartyProfileDrawer: (state) => {
      state.isPartyProfileOpen = !state.isPartyProfileOpen
    },
    openPartyProfileDrawer: (state) => {
      state.isPartyProfileOpen = true;
    },
    closePartyProfileDrawer: (state) => {
      state.isPartyProfileOpen = false;
    },
    toggleDefaultProfileDrawer: (state) => {
      state.isDefaultProfileOpen = !state.isDefaultProfileOpen
    },
    openDefaultProfileDrawer: (state) => {
      state.isDefaultProfileOpen = true;
    },
    closeDefaultProfileDrawer: (state) => {
      state.isDefaultProfileOpen = false;
    },
  },
});

export const { openPartyProfileDrawer, closePartyProfileDrawer, togglePartyProfileDrawer, openDefaultProfileDrawer, closeDefaultProfileDrawer, toggleDefaultProfileDrawer } = profileDrawersSlice.actions;