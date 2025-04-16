import { get, postImage } from '../../api/Api';
import { getApiUrl } from '../../api/ApiHelper';
import { User } from '../../data/types/User';
import { LoginResponseDataInterface } from './AuthenticationApi';
import axios, {AxiosInstance, AxiosResponse} from "axios";

const USER_PATH = `${getApiUrl()}/user`;

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

export type UsersResponse = {
    data: User[]
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
