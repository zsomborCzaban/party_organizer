const JWT_AUTH_TOKEN_LOCAL_STORAGE_KEY = "JWT_AUTH_TOKEN_LOCAL_STORAGE_KEY";

export const getJwtAuthToken = () => {
    return localStorage.getItem(JWT_AUTH_TOKEN_LOCAL_STORAGE_KEY)
}

export const setJwtAuthToken = (token: string) => {
    return localStorage.setItem(JWT_AUTH_TOKEN_LOCAL_STORAGE_KEY, token)
}

export const clearJwtAuthToken = () => {
    return localStorage.removeItem(JWT_AUTH_TOKEN_LOCAL_STORAGE_KEY)
}