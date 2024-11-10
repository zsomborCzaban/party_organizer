import {post} from "../../api/Api";
import {authService} from "../../auth/AuthService";
import {getApiUrl} from "../../api/ApiHelper";

const LOGIN_PATH = getApiUrl() + '/user/login';
const REGISTER_PATH = getApiUrl() + '/user/register';

export interface LoginRequestDataInterface {
    username: string;
    password: string;
}

export interface LoginResponseDataInterface {
    jwt: string;
}

export const login = async (username: string, password: string) => {
    const loginRequest: LoginRequestDataInterface = {
        username: username,
        password: password,
    };

    try {
        return post<LoginResponseDataInterface>(
            LOGIN_PATH,
            loginRequest,
        )
            .then(response => {
                authService.userLoggedIn(response.jwt)
                return ""
            })
            .catch(err => {
                return err.response.data.errors
            })
    } catch (err) {
        return 'An error occurred while logging in. Please try again.'
    }

}

export interface RegisterRequestBody {
    username: string;
    email: string;
    password: string;
    confirm_password: string;
}

export const register = async (requestBody: RegisterRequestBody): Promise<void> => {
    return new Promise<void>((resolve, reject) => {
        post<void>(REGISTER_PATH, requestBody)
            .then(() => {return resolve()})
            .catch(error => {
                return reject(error)
            })
    })
}