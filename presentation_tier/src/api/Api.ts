import axios, { InternalAxiosRequestConfig, AxiosResponse } from 'axios';

import {getApiConfig, getImageUploaderApiConfig} from './ApiConfig';
import { authService } from '../auth/AuthService';
import  { ApiResponse } from './ApiResponse';

const apiClient = axios.create(getApiConfig());
const imageUploaderApiClient = axios.create(getImageUploaderApiConfig());

const interceptRequest = (request: InternalAxiosRequestConfig) => {
    const accessToken = authService.getJwtToken();
    if(accessToken) {
        request.headers.Authorization = `Bearer ${accessToken}`;
    }
    return request;
};

const interceptSuccessResponse = (response: AxiosResponse) => response;

// todo: get back here and figure the type out
const interceptErrorReponse = (error: { response: { status: number; }; }) => {
    console.error(`error recieved in API: ${error}`, error);

    if(error.response && error.response.status === 401) {
        authService.handleUnauthorized();
    }

    return Promise.reject(error);
};

apiClient.interceptors.request.use(interceptRequest);
apiClient.interceptors.response.use(
    interceptSuccessResponse,
    interceptErrorReponse,
);
imageUploaderApiClient.interceptors.request.use(interceptRequest);
imageUploaderApiClient.interceptors.response.use(
    interceptSuccessResponse,
    interceptErrorReponse,
);

export const parseResponse = <T>(response: AxiosResponse<ApiResponse<T>>) => new Promise<T>((resolve, reject) => {
        if(response.status !== 200) {
            // eslint-disable-next-line prefer-promise-reject-errors,no-promise-executor-return
            return reject(`invalid HTTP response status: ${response.status}`);
        }

        const apiResponse: ApiResponse<T> = response.data;

        if(apiResponse.isError || apiResponse.code !== 200) {
            // eslint-disable-next-line no-promise-executor-return,prefer-promise-reject-errors
            return reject(`invalid response status: ${apiResponse.code}; is error: ${apiResponse.errors}`);
        }
    // eslint-disable-next-line no-promise-executor-return
        return resolve(apiResponse.data);
    });

export const get = async <T>(url: string) => new Promise<T>((resolve, reject) => {
        apiClient.get<ApiResponse<T>>(url)
            .then((response) => {
                parseResponse<T>(response)
                    .then((parsedResponse: T) => resolve(parsedResponse))
                    .catch((error) => reject(error));
            })
            .catch((error) => reject(error));
    });

export const post = <T>(url: string, requestBody: object) => new Promise<T>((resolve, reject) => {
        apiClient.post<ApiResponse<T>>(url, requestBody)
            .then((response) => {
                parseResponse<T>(response)
                    .then((parsedResponse: T) => resolve(parsedResponse))
                    .catch((error) => reject(error));
            })
            .catch((error) => reject(error));
    });

export const postImage = <T>(url: string, requestBody: FormData) => new Promise<T>((resolve, reject) => {
        imageUploaderApiClient.post<ApiResponse<T>>(url, requestBody)
            .then((response) => {
                parseResponse<T>(response)
                    .then((parsedResponse: T) => resolve(parsedResponse))
                    .catch((error) => reject(error));
            })
            .catch((error) => reject(error));
    });

export const put = <T>(url: string, requestBody: object) => new Promise<T>((resolve, reject) => {
        apiClient.put<ApiResponse<T>>(url, requestBody)
            .then((response) => {
                parseResponse<T>(response)
                    .then((parsedResponse: T) => resolve(parsedResponse))
                    .catch((error) => reject(error));
            })
            .catch((error) => reject(error));
    });

export const DELETE = <T>(url: string) => new Promise<T>((resolve, reject) => {
        apiClient.delete<ApiResponse<T>>(url)
            .then((response) => {
                parseResponse<T>(response)
                    .then((parsedResponse: T) => resolve(parsedResponse))
                    .catch((error) => reject(error));
            })
            .catch((error) => reject(error));
    });

export default apiClient;