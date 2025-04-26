import { configureStore } from '@reduxjs/toolkit';
import { persistStore } from 'redux-persist';
import { userSlice } from './slices/UserSlice';
import { profileDrawersSlice } from "./slices/profileDrawersSlice.ts";

export const store = configureStore({
  reducer: {
    userStore: userSlice.reducer,
    profileDrawers: profileDrawersSlice.reducer,
  },
  middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware({
      serializableCheck: false, // quickfix for development todo: get back here
    }),
});

export const persistor = persistStore(store);

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;
