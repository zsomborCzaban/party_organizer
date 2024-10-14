import {
    setJwtAuthToken,
    getJwtAuthToken,
    clearJwtAuthToken
} from "./AuthStorageUtils";

export interface AuthService {
    userLoggedIn: (username: string, jwt: string) => void
    handleUnauthorized: () => void
    getJwtToken: () => string | null
    isAuthenticated: () => boolean
}

export const authService: AuthService = {
    userLoggedIn(username, jwt) {
        console.log(`User logged in: ${username}`)
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