import {API_PATH} from "../config/backend_url";

export const getBackendUrl = () => {
    return process.env.REACT_APP_BACKEND_URL
        ? process.env.REACT_APP_BACKEND_URL
        : '';
}

export const getApiUrl = () => {
    return getBackendUrl() + API_PATH
}