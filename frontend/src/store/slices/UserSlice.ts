import { createSlice, PayloadAction } from '@reduxjs/toolkit';
import { clearJwtAuthToken, getJwtAuthToken, setJwtAuthToken } from '../../auth/AuthStorageUtils';
import { RootState } from '../store';

export interface UserState {
  jwt?: string;
}

const initialState: UserState = {
  jwt: getJwtAuthToken() ?? undefined,
};

export const userSlice = createSlice({
  name: 'user',
  initialState,
  reducers: {
    setUserJwt: (state, action: PayloadAction<string>) => {
      setJwtAuthToken(action.payload);
      state.jwt = action.payload;
    },
    deleteUserJwt: (state) => {
      clearJwtAuthToken();
      state.jwt = undefined;
    },
  },
});

export const { setUserJwt, deleteUserJwt } = userSlice.actions;

export const isUserLoggedIn = (state: RootState): boolean => !!state.userStore.jwt;
