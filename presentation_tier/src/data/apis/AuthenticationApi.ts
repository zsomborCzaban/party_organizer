import {post} from '../../api/Api';
import {authService} from '../../auth/AuthService';
import {getApiUrl} from '../../api/ApiHelper';

const LOGIN_PATH = `${getApiUrl()  }/login`;
const REGISTER_PATH = `${getApiUrl()  }/register`;

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
        return await post<LoginResponseDataInterface>(
            LOGIN_PATH,
            loginRequest,
        )
            .then(response => {
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

export const register = async (requestBody: RegisterRequestBody): Promise<void> => new Promise<void>((resolve, reject) => {
        post<void>(REGISTER_PATH, requestBody)
            .then(() => resolve())
            .catch((error) => reject(error));
    });