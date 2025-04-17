import { createAsyncThunk, createSlice, PayloadAction } from '@reduxjs/toolkit';
import { clearJwtAuthToken, getJwtAuthToken, setJwtAuthToken } from '../../auth/AuthStorageUtils';
import { RootState } from '../store';
import { Api } from '../../api/Api';

export interface UserState {
  isLoading: boolean;
  loginError: boolean;
  jwt?: string;
}

const initialState: UserState = {
  isLoading: false,
  loginError: false,
  jwt: getJwtAuthToken() ?? undefined,
};

export const userLogin = createAsyncThunk('user/loginUser', async ({ api, username, password }: { api: Api; username: string; password: string }) => {
  const response = await api.authApi.postLogin(username, password);

  return response;
});

export const userRegister = createAsyncThunk(
  'user/register',
  async ({ api, username, email, password, confirmPassword }: { api: Api; username: string; email: string; password: string; confirmPassword: string }) => {
    const registerResponse = await api.authApi.postRegister({ username, email, password, confirmPassword });

    return registerResponse;
  },
);

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
    deleteLoginError: (state) => {
      state.loginError = false;
    },
  },
  extraReducers: (builder) => {
    builder
      // Login states
      .addCase(userLogin.pending, (state) => {
        state.isLoading = true;
      })
      .addCase(userLogin.fulfilled, (state, action) => {
        state.isLoading = false;
        if (action.payload === 'error') {
          state.loginError = true;
        } else {
          setJwtAuthToken(action.payload.data.jwt);
          state.jwt = action.payload.data.jwt;
          state.loginError = false;
        }
      })
      .addCase(userLogin.rejected, (state) => {
        state.isLoading = false;
      });
    /*  .addCase() */
    //TODO: Register states
  },
});

export const { setUserJwt, deleteUserJwt, deleteLoginError } = userSlice.actions;

export const getUserJwt = (state: RootState) => state.userStore.jwt;
export const isUserLoggedIn = (state: RootState): boolean => !!state.userStore.jwt && state.userStore.isLoading === false;
export const isErrorWhileLoggingIn = (state: RootState): boolean => state.userStore.loginError && state.userStore.isLoading === false;
