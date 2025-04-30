import { configureStore } from '@reduxjs/toolkit';
import { userSlice } from './slices/UserSlice';
import { profileDrawersSlice } from "./slices/profileDrawersSlice.ts";

export const store = configureStore({
  reducer: {
    userStore: userSlice.reducer,
    profileDrawers: profileDrawersSlice.reducer,
  },
  middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware({
      serializableCheck: false,
    }),
});


export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;
