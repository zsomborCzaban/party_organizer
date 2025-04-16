import {
    setJwtAuthToken,
    getJwtAuthToken,
    clearJwtAuthToken,
} from './AuthStorageUtils';
import {store} from "../store/store.ts";
import {deleteUserJwt} from "../store/sclices/UserSlice.ts";

export interface AuthService {
    userLoggedIn: (jwt: string) => void
    userLoggedOut: () => void
    getJwtToken: () => string | null
    isAuthenticated: () => boolean
}

export const authService: AuthService = {
    userLoggedIn: (jwt) => {
        setJwtAuthToken(jwt);
    },

    userLoggedOut: () => {
        store.dispatch(deleteUserJwt());
        clearJwtAuthToken();
    },

    getJwtToken: () => getJwtAuthToken(),

    isAuthenticated: () => {
        const authToken = getJwtAuthToken();
        return ( authToken !== null && authToken !== undefined && authToken.length > 0);
    },
};