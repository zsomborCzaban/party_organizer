import axios, {InternalAxiosRequestConfig, AxiosResponse, AxiosInstance, AxiosError} from 'axios';
import { getApiConfig, getImageUploaderApiConfig } from './ApiConfig';
import { authService } from '../auth/AuthService';
import { ApiResponse } from '../data/types/ApiResponseTypes';
import { AuthApi } from './apis/AuthenticationApi';
import {PartyApi} from "./apis/PartyApi.ts";
import {PartyAttendanceManagerApi} from "./apis/PartyAttendanceManagerApi.ts";
import {FriendInviteManagerApi} from "./apis/FriendInviteManagerApi.ts";
import {UserApi} from "./apis/UserApi.ts";
import {RequirementApi} from "./apis/RequirementApi.ts";
import {ContributionApi} from "./apis/ContributionApi.ts";

export const apiClient = axios.create(getApiConfig());
const imageUploaderApiClient = axios.create(getImageUploaderApiConfig());

const addAuthHeaderToRequest = (request: InternalAxiosRequestConfig) => {
  const accessToken = authService.getJwtToken();
  if (accessToken) {
    request.headers.Authorization = `Bearer ${accessToken}`;
  }
  return request;
};

const checkIsJWTValid = (response: AxiosError) => {
    if (response.response?.data?.errors === 'Authorization missing from header' ||
        response.response?.data?.errors === 'invalid jwt'
    ) {
        authService.userLoggedOut()
    }
    return response;
};

const interceptErrorReponse = (error: { response: { status: number } }) => {
  console.error(`error recieved in API: ${error}`, error); // TODO: remove console error when error handling is done more sophisiticated

  if (error.response && error.response.status === 401) {
    authService.userLoggedOut();
  }

  return Promise.reject(error);
};

// Intercept both APIs
apiClient.interceptors.request.use(addAuthHeaderToRequest);
apiClient.interceptors.response.use((res) => res, interceptErrorReponse);
imageUploaderApiClient.interceptors.request.use(addAuthHeaderToRequest);
imageUploaderApiClient.interceptors.response.use((res) => res, interceptErrorReponse);

export const parseResponse = <T>(response: AxiosResponse<ApiResponse<T>>) =>
  new Promise<T>((resolve, reject) => {
    if (response.status !== 200) {
       
      return reject(`invalid HTTP response status: ${response.status}`);
    }

    const apiResponse: ApiResponse<T> = response.data;

    if (apiResponse.isError || apiResponse.code !== 200) {
       
      return reject(`invalid response status: ${apiResponse.code}; is error: ${apiResponse.errors}`);
    }
     
    return resolve(apiResponse.data);
  });

export const get = async <T>(url: string) =>
  new Promise<T>((resolve, reject) => {
    apiClient
      .get<ApiResponse<T>>(url)
      .then((response) => {
        parseResponse<T>(response)
          .then((parsedResponse: T) => resolve(parsedResponse))
          .catch((error) => reject(error));
      })
      .catch((error) => reject(error));
  });

export const post = <T>(url: string, requestBody: object) =>
  new Promise<T>((resolve, reject) => {
    apiClient
      .post<ApiResponse<T>>(url, requestBody)
      .then((response) => {
        parseResponse<T>(response)
          .then((parsedResponse: T) => resolve(parsedResponse))
          .catch((error) => reject(error));
      })
      .catch((error) => reject(error));
  });

export const postImage = <T>(url: string, requestBody: FormData) =>
  new Promise<T>((resolve, reject) => {
    imageUploaderApiClient
      .post<ApiResponse<T>>(url, requestBody)
      .then((response) => {
        parseResponse<T>(response)
          .then((parsedResponse: T) => resolve(parsedResponse))
          .catch((error) => reject(error));
      })
      .catch((error) => reject(error));
  });

export const put = <T>(url: string, requestBody: object) =>
  new Promise<T>((resolve, reject) => {
    apiClient
      .put<ApiResponse<T>>(url, requestBody)
      .then((response) => {
        parseResponse<T>(response)
          .then((parsedResponse: T) => resolve(parsedResponse))
          .catch((error) => reject(error));
      })
      .catch((error) => reject(error));
  });

export const DELETE = <T>(url: string) =>
  new Promise<T>((resolve, reject) => {
    apiClient
      .delete<ApiResponse<T>>(url)
      .then((response) => {
        parseResponse<T>(response)
          .then((parsedResponse: T) => resolve(parsedResponse))
          .catch((error) => reject(error));
      })
      .catch((error) => reject(error));
  });

export class Api {
  private axiosInstance: AxiosInstance;
  public authApi: AuthApi;
  public partyApi: PartyApi;
  public partyAttendanceApi: PartyAttendanceManagerApi;
  public friendManagerApi: FriendInviteManagerApi;
  public userApi: UserApi;
  public requirementApi: RequirementApi;
  public contributionApi: ContributionApi;

  constructor() {
    this.axiosInstance = axios.create(getApiConfig());
    this.authApi = new AuthApi(this.axiosInstance);
    this.partyApi = new PartyApi(this.axiosInstance);
    this.partyAttendanceApi = new PartyAttendanceManagerApi(this.axiosInstance);
    this.friendManagerApi = new FriendInviteManagerApi(this.axiosInstance);
    this.userApi = new UserApi(this.axiosInstance);
    this.requirementApi = new RequirementApi(this.axiosInstance);
    this.contributionApi = new ContributionApi(this.axiosInstance);
    this.axiosInstance.interceptors.request.use(addAuthHeaderToRequest);
    this.axiosInstance.interceptors.response.use(null, checkIsJWTValid);
  }
}
