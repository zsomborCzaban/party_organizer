import {getJwtAuthToken} from "./AuthStorageUtils";
import { jwtDecode } from 'jwt-decode';

export const getUserId = () => {
    const authToken = getJwtAuthToken()
    if (!authToken) {
        return null
    }

    try {
        const decoded = jwtDecode(authToken)
        return decoded.sub;
    } catch (e) {
        return null
    }
}