import { getApiUrl } from '../../api/ApiHelper';
import axios, { AxiosInstance, AxiosResponse } from 'axios';

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

export class AutchenticationApi {
  private axiosInstance: AxiosInstance;

  constructor(axiosInstance: AxiosInstance) {
    this.axiosInstance = axiosInstance;
  }

  async postLogin(username: string, password: string): Promise<LoginPostResponse | undefined> {
    try {
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
