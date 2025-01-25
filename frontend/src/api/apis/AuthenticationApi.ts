import { getApiUrl } from '../../api/ApiHelper';
import axios, { AxiosInstance, AxiosResponse } from 'axios';
import { authService } from '../../auth/AuthService';
import { post } from '../Api';

const handleApiResponse = <T>(response: AxiosResponse<T>): T => {
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

export type LoginPostResponse = {
  jwt: string;
};

export interface RegisterPostRequestProps {
  username: string;
  email: string;
  password: string;
  confirm_password: string;
}

export class AuthApi {
  private axiosInstance: AxiosInstance;

  constructor(axiosInstance: AxiosInstance) {
    this.axiosInstance = axiosInstance;
  }

  async postLogin(username: string, password: string): Promise<LoginPostResponse | undefined> {
    try {
        console.log('sa',getApiUrl());
      const response = await this.axiosInstance.post<LoginPostResponse>(`${getApiUrl()}/user/login`, { username, password });
      return handleApiResponse(response);
    } catch (error) {
      handleApiError(error);
    }
    return undefined;
  }

  async postRegister(props: RegisterPostRequestProps): Promise<void> {
    try {
      await this.axiosInstance.post<void>(`${getApiUrl()}/user/register`, props);
    } catch (error) {
      handleApiError(error);
    }
    return undefined;
  }
}

// Deprecated from here!!!!!!!

const LOGIN_PATH = `${getApiUrl()}/user/login`;
const REGISTER_PATH = `${getApiUrl()}/user/register`;

export interface LoginRequestDataInterface {
  username: string;
  password: string;
}

export interface LoginResponseDataInterface {
  jwt: string;
}

export const login = async (username: string, password: string) => {
  const loginRequest: LoginRequestDataInterface = {
    username,
    password,
  };

  try {
    return await post<LoginResponseDataInterface>(LOGIN_PATH, loginRequest)
      .then((response) => {
        authService.userLoggedIn(response.jwt);
        return '';
      })
      .catch((err) => err.response.data.errors);
  } catch (err) {
    return 'An error occurred while logging in. Please try again.';
  }
};

export interface RegisterRequestBody {
  username: string;
  email: string;
  password: string;
  confirm_password: string;
}

export const register = async (requestBody: RegisterRequestBody): Promise<void> =>
  new Promise<void>((resolve, reject) => {
    post<void>(REGISTER_PATH, requestBody)
      .then(() => resolve())
      .catch((error) => reject(error));
  });
