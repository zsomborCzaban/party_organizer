import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';

// https://vite.dev/config/
export default defineConfig({
  plugins: [react()],
  base: process.env.VITE_REACT_APP_URL,
  server: {
    port: 3000,
  },
});
