import { getBackendUrl } from './ApiHelper'

const API_PATH = '/api/v0';

export const getApiConfig = () => {
    return {
        baseUrl : `${getBackendUrl()}${API_PATH}`,
        headers: {
            'Content-Type': 'application/json'
        },
    };
};