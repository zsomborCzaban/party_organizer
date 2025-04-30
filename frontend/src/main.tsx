import { StrictMode } from 'react';
import { createRoot } from 'react-dom/client';
import './index.css';
import App from './App.tsx';
import { store } from './store/store.ts';
import { Provider } from 'react-redux';
import { ApiContext } from './context/ApiContext.ts';
import { Api } from './api/Api.ts';
import { Toaster } from '../tailwindcss/components/ui/sonner.tsx';

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <ApiContext.Provider value={new Api()}>
      <Provider store={store}>
          <App />
      </Provider>
    </ApiContext.Provider>
    <Toaster richColors />
  </StrictMode>,
);
