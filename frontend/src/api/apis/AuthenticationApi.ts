import { getApiUrl } from '../ApiHelper.ts';
import axios, { AxiosInstance, AxiosResponse } from 'axios';
import { toast } from 'sonner';
import {ApiResponse} from "../../data/types/ApiResponseTypes.ts";

const handleApiResponse = <T>(response: AxiosResponse<T>): T => {
  if(response.data){
    return response.data
  }
  if(response.response && response.response.data) {
    return response.response.data
  }

  toast.error('Unexpected error')
  return response.data;
};

const handleApiError = (error: unknown) => {
  // TODO: handle errors as needed
  if (axios.isAxiosError(error)) {
    console.error(`Axios error: ${error.message}`);
  } else {
    console.error(`Unexpected error: ${error}`);
  }
};

export type LoginPostProps = {
  username: string;
  password: string;
};

export type LoginPostResponse = ApiResponse<{ jwt: string }>

export type RegisterRequestBody = {
  username: string;
  email: string;
  password: string;
  confirm_password: string;
};

export class AuthApi {
  private axiosInstance: AxiosInstance;

  constructor(axiosInstance: AxiosInstance) {
    this.axiosInstance = axiosInstance;
  }

  async postLogin(username: string, password: string): Promise<LoginPostResponse | 'error'> {
    try {
      const response = await this.axiosInstance.post<LoginPostResponse>(`${getApiUrl()}/login`, { username, password });
      localStorage.removeItem('profile_picture_url') //this is a hacky solution. The problem is we get the profilepictureurl from the jwt payload we get back, but if we upload the jwt stays the same regardless. So whenever a user uploads a profile picture we save it to localStorage. But localstorage can get stuck with old values so whenever we log in we clear it, because at the moment of the login the url in the received jwt is the freshest.
      return handleApiResponse(response);
    } catch (error) {
      handleApiError(error);
      return 'error';
    }
  }

  async postRegister(props: RegisterRequestBody): Promise<ApiResponse<string> | 'error'> {
    try {
      const response = await this.axiosInstance.post<ApiResponse<string>>(`${getApiUrl()}/register`, props);
      return handleApiResponse(response)
    } catch (error) {
      handleApiError(error);
      return 'error';
    }
  }
}
