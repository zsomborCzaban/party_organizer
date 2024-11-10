import {
    setJwtAuthToken,
    getJwtAuthToken,
    clearJwtAuthToken
} from "./AuthStorageUtils";

export interface AuthService {
    userLoggedIn: (jwt: string) => void
    handleUnauthorized: () => void
    getJwtToken: () => string | null
    isAuthenticated: () => boolean
}

export const authService: AuthService = {
    userLoggedIn(jwt) {
        setJwtAuthToken(jwt)
    },

    handleUnauthorized() {
        clearJwtAuthToken()
        window.location.href = '/login'
    },

    getJwtToken() {
        return getJwtAuthToken()
    },

    isAuthenticated() {
        const authToken = getJwtAuthToken()
        return ( authToken !== null && authToken !== undefined && authToken.length > 0)
    },
}