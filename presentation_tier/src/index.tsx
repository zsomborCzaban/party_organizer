import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import 'bootstrap/dist/css/bootstrap.min.css';
import { Provider } from 'react-redux';
import { persistor, store } from './store/store';
import { PersistGate } from 'redux-persist/integration/react';
import { App } from 'antd';

const root = ReactDOM.createRoot(document.getElementById('root') as HTMLElement);
root.render(
  <React.StrictMode>
    <Provider store={store}>
      {/* todo: make a loading screen */}
      <PersistGate
        loading={<div>Loading...</div>}
        persistor={persistor}
      >
        <App />
      </PersistGate>
    </Provider>
  </React.StrictMode>,
);
