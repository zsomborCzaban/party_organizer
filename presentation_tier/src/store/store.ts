import { configureStore } from '@reduxjs/toolkit'

export const store = configureStore({
 reducer: {
     //stores here:
 },
});

export type RootState = ReturnType<typeof store.getState>;
export type ApiDispatch = typeof  store.dispatch;