import {API_PATH} from '../config/backend_url';

export const getBackendUrl = () => process.env.REACT_APP_BACKEND_URL
        ? process.env.REACT_APP_BACKEND_URL
        : '';

export const getApiUrl = () => getBackendUrl() + API_PATH;