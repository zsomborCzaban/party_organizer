import { API_PATH } from '../config/backend_url';

export const getBackendUrl = (): string => import.meta.env.VITE_REACT_APP_BACKEND_URL ?? '';

export const getApiUrl = () => getBackendUrl() + API_PATH;
