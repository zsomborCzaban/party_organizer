import { get, postImage } from '../Api.ts';
import { getApiUrl } from '../ApiHelper.ts';
import { User } from '../../data/types/User';
import { LoginResponseDataInterface } from './AuthenticationApi';
import axios, {AxiosInstance, AxiosResponse} from "axios";
import {ApiResponse} from "../../data/types/ApiResponseTypes.ts";
import {clearJwtAuthToken, setJwtAuthToken} from "../../auth/AuthStorageUtils.ts";

const USER_PATH = `${getApiUrl()}/user`;

const handleApiResponse = <T>(response: AxiosResponse<T>): T => {
    return response.data;
};

const handleApiError = (error: unknown) => {
    if (axios.isAxiosError(error)) {
        console.error(`Axios error: ${error.message}`);
    } else {
        console.error(`Unexpected error: ${error}`);
    }
};

export type UsersResponse = {
    data: User[]
}

export type changePasswordBody = {
    password: string,
    confirm_password: string,
}

export class UserApi {
    private axiosInstance: AxiosInstance;

    constructor(axiosInstance: AxiosInstance) {
        this.axiosInstance = axiosInstance;
    }

    async getFriends(): Promise<UsersResponse | 'error'> {
        try {
            const response = await this.axiosInstance.get<UsersResponse>(`${USER_PATH}/getFriends`)
            return handleApiResponse(response)
        } catch (error) {
            handleApiError(error)
            return 'error'
        }
    }

    async changePassword(body: changePasswordBody, xzs: string): Promise<ApiResponse<string> | 'error'> {
        try {
            setJwtAuthToken(xzs)
            const response = await this.axiosInstance.post<ApiResponse<string>>(
                `${USER_PATH}/resetPassword`, body
            );
            return handleApiResponse(response.response ? response.response: response); //axios post gives back different resp on request success and on request error for some reason
        } catch (error) {
            handleApiError(error);
            return 'error';
        } finally {
            clearJwtAuthToken()
        }
    }
}

export const getFriends = async (): Promise<User[]> =>
  new Promise<User[]>((resolve, reject) => {
    get<User[]>(`${USER_PATH}/getFriends`)
      .then((friends: User[]) => resolve(friends))
      .catch((err) => reject(err));
  });

// this is a bit hacky. We access our current logged in user by the jwt token. so when we upload a new picture we get a new token with the updated picture
export const uploadPicture = async (formData: FormData): Promise<LoginResponseDataInterface> =>
  new Promise<LoginResponseDataInterface>((resolve, reject) => {
    postImage<LoginResponseDataInterface>(`${USER_PATH}/uploadProfilePicture`, formData)
      .then((response: LoginResponseDataInterface) => resolve(response))
      .catch((err) => reject(err));
  });
