import {getApiUrl} from './ApiHelper'

export const getApiConfig = () => {
    return {
        baseUrl : `${getApiUrl()}`,
        headers: {
            'Content-Type': 'application/json'
        },
    };
};