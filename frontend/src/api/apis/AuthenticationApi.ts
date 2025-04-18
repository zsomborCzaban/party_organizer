import { getApiUrl } from '../../api/ApiHelper';
import axios, { AxiosInstance, AxiosResponse } from 'axios';
import { toast } from 'sonner';

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
  data: {
    jwt: string;
  };
};

export type RegisterPostRequestProps = {
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
      localStorage.removeItem('profile_picture_url')
      return handleApiResponse(response);
    } catch (error) {
      handleApiError(error);
      return 'error';
    }
  }

  async postRegister(props: RegisterPostRequestProps): Promise<void | 'error'> {
    try {
      await this.axiosInstance.post<void>(`${getApiUrl()}/register`, props);
      toast.success('Successfully registered');
    } catch (error) {
      handleApiError(error);
      return 'error';
    }
  }
}
