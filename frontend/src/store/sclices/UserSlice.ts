import { createAsyncThunk, createSlice, PayloadAction } from '@reduxjs/toolkit';
import { clearJwtAuthToken, getJwtAuthToken, setJwtAuthToken } from '../../auth/AuthStorageUtils';
import { RootState } from '../store';
import { Api } from '../../api/Api';

export interface UserState {
  isLoading: boolean;
  jwt?: string;
}

const initialState: UserState = {
  isLoading: false,
  jwt: getJwtAuthToken() ?? undefined,
};

export const userLogin = createAsyncThunk('user/loginUser', async ({ api, username, password }: { api: Api; username: string; password: string }) => {
  const response = await api.authApi.postLogin(username, password);

  return response;
});

export const userRegister = createAsyncThunk('user/register', async ({ api }: { api: Api }) => {
  const registerResponse = await api.authApi.postRegister({});
});

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
  extraReducers: (builder) => {
    builder
      .addCase(userLogin.pending, (state) => {
        state.isLoading = true;
      })
      .addCase(userLogin.fulfilled, (state, action) => {
        state.isLoading = false;
        console.log('jwt,', action.payload);
        if (action.payload?.data.jwt) {
          setJwtAuthToken(action.payload.data.jwt);
          state.jwt = action.payload.data.jwt;
        }
      })
      .addCase(userLogin.rejected, (state) => {
        state.isLoading = false;
      });
  },
});

export const { setUserJwt, deleteUserJwt } = userSlice.actions;

export const getUserJwt = (state: RootState) => state.userStore.jwt;
export const isUserLoggedIn = (state: RootState): boolean => !!state.userStore.jwt && state.userStore.isLoading === false;
