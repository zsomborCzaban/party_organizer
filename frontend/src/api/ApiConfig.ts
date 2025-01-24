import { getApiUrl } from './ApiHelper';

export const getApiConfig = () => ({
  baseUrl: `${getApiUrl()}`,
  headers: {
    'Content-Type': 'application/json',
  },
});

export const getImageUploaderApiConfig = () => ({
  baseUrl: `${getApiUrl()}`,
  headers: {
    'Content-Type': 'multipart/form-data',
  },
});
