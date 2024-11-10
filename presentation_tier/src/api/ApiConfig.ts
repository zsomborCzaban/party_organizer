import {getApiUrl} from './ApiHelper'

export const getApiConfig = () => {
    return {
        baseUrl : `${getApiUrl()}`,
        headers: {
            'Content-Type': 'application/json'
        },
    };
};

export const getImageUploaderApiConfig = () => {
    return {
        baseUrl : `${getApiUrl()}`,
        headers: {
            'Content-Type': 'multipart/form-data'
        },
    };
};