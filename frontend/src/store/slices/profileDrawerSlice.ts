import { createSlice } from '@reduxjs/toolkit';

interface DrawerState {
  isOpen: boolean;
}

const initialState: DrawerState = {
  isOpen: false,
};

export const profileDrawerSlice = createSlice({
  name: 'drawer',
  initialState,
  reducers: {
    toggleDrawer: (state) => {
      state.isOpen = !state.isOpen
      console.log('toggleing')
    },
    openDrawer: (state) => {
      state.isOpen = true;
    },
    closeDrawer: (state) => {
      state.isOpen = false;
    },
  },
});

export const { openDrawer, closeDrawer, toggleDrawer } = profileDrawerSlice.actions;