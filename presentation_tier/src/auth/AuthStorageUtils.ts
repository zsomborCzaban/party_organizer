const JWT_AUTH_TOKEN_LOCAL_STORAGE_KEY = 'JWT_AUTH_TOKEN_LOCAL_STORAGE_KEY';

export const getJwtAuthToken = () => localStorage.getItem(JWT_AUTH_TOKEN_LOCAL_STORAGE_KEY);

export const setJwtAuthToken = (token: string) => localStorage.setItem(JWT_AUTH_TOKEN_LOCAL_STORAGE_KEY, token);

export const clearJwtAuthToken = () => localStorage.removeItem(JWT_AUTH_TOKEN_LOCAL_STORAGE_KEY);