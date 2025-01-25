import { getAttendedParties } from '../../api/apis/PartyApi';
import { Party } from '../types/Party';
import { createAsyncThunk, createSlice } from '@reduxjs/toolkit';

export interface AttendedPartySlice {
  parties: Party[];
  loading: boolean;
  error: Error | null;
}

const initialState: AttendedPartySlice = {
  parties: [],
  loading: true,
  error: null,
};

export const loadAttendedParties = createAsyncThunk(
  'data/loadAttendedParties',
  // eslint-disable-next-line consistent-return
  async () => {
    try {
      return await getAttendedParties();
    } catch (err) {
      console.log(`err in loadAttendedParties: ${err}`);
    }
  },
);

const attendedPartySlice = createSlice({
  name: 'attendedParties',
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(loadAttendedParties.pending, (state) => {
        state.loading = true;
        state.parties = [];
        state.error = null;
      })
      .addCase(loadAttendedParties.fulfilled, (state, action) => {
        state.loading = false;
        state.parties = action.payload!;
        state.error = null;
      })
      .addCase(loadAttendedParties.rejected, (state, action) => {
        state.loading = false;
        state.parties = [];
        state.error = new Error(action.error.message || 'Failed to load attended parties');
      });
  },
});

export default attendedPartySlice.reducer;
