import axios, { InternalAxiosRequestConfig, AxiosResponse } from "axios";

import { getApiConfig } from "./ApiConfig";
import { authService } from "../auth/AuthService";
import  { ApiResponse } from "./ApiResponse";

const apiClient = axios.create(getApiConfig())

const interceptRequest = (request: InternalAxiosRequestConfig) => {
    const accessToken = authService.getJwtToken();
    if(accessToken) {
        request.headers.Authorization = `Bearer ${accessToken}`;
    }
    return request
};

const interceptSuccessResponse = (response: AxiosResponse) => {
    return response
}

const interceptErrorReponse = (error: any) => {
    console.error(`error recieved in API: ${error}`, error);

    if(error.response && error.response.status === 401) {
        authService.handleUnauthorized();
    }

    return Promise.reject(error);
};

apiClient.interceptors.request.use(interceptRequest)
apiClient.interceptors.response.use(
    interceptSuccessResponse,
    interceptErrorReponse,
);

export const parseResponse = <T>(response: AxiosResponse<ApiResponse<T>>) => {
    return new Promise<T>((resolve, reject) => {
        if(response.status !== 200) {
            return reject(`invalid HTTP response status: ${response.status}`);
        }

        const apiResponse: ApiResponse<T> = response.data

        if(apiResponse.isError || apiResponse.code !== 200) {
            return reject(`invalid response status: ${apiResponse.code}; is error: ${apiResponse.errors}`,);
        }
        return resolve(apiResponse.data)
    });
};

export const get = async <T>(url: string) => {
    return new Promise<T>((resolve, reject) => {
        apiClient.get<ApiResponse<T>>(url)
            .then((response) => {
                parseResponse<T>(response)
                    .then((parsedResponse: T) => {
                        return resolve(parsedResponse);
                    })
                    .catch((error) => {
                        return reject(error);
                    });
            })
            .catch((error) => {
                return reject(error);
            });
    });
};

export const post = <T>(url: string, requestBody: object) => {
    return new Promise<T>((resolve, reject) => {
        apiClient.post<ApiResponse<T>>(url, requestBody)
            .then((response) => {
                parseResponse<T>(response)
                    .then((parsedResponse: T) => {
                        return resolve(parsedResponse);
                    })
                    .catch((error) => {
                        return reject(error);
                    });
            })
            .catch((error) => {
                return reject(error);
            });
    });
};

export const DELETE = <T>(url: string) => {
    return new Promise<T>((resolve, reject) => {
        apiClient.delete<ApiResponse<T>>(url)
            .then((response) => {
                parseResponse<T>(response)
                    .then((parsedResponse: T) => {
                        return resolve(parsedResponse);
                    })
                    .catch((error) => {
                        return reject(error);
                    });
            })
            .catch((error) => {
                return reject(error);
            });
    });
};

export default apiClient