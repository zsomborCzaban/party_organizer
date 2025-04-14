import { getApiUrl } from './ApiHelper';

export const getApiConfig = () => ({
  baseURL: `${getApiUrl()}`,
  headers: {
    'Content-Type': 'application/json',
  },
});

export const getImageUploaderApiConfig = () => ({
  baseURL: `${getApiUrl()}`,
  headers: {
    'Content-Type': 'multipart/form-data',
  },
});
