import {AxiosResponse} from "axios";
import {ApiResponse} from "../../api/ApiResponse";
import {jwtDecode} from "jwt-decode";
import apiClient, {post} from "../../api/Api";
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

const getJwtFromApiResponse = (apiResponse: AxiosResponse) => {
    if(apiResponse.status !== 200) {
        throw new Error(`invalid HTTP response status: ${apiResponse.status}`);
    }

    const loginResponse: ApiResponse<LoginResponseDataInterface> = apiResponse.data

    if(loginResponse.isError || loginResponse.code !== 200) {
        throw new Error(`invalid response status: ${loginResponse.code}; is error: ${loginResponse.errors}`,);
    }

    //will throw error if jwt is invalid
    jwtDecode(loginResponse.data.jwt);

    return loginResponse.data.jwt;
};

export const login = async (username: string, password: string) => {
    const loginRequest: LoginRequestDataInterface = {
        username: username,
        password: password,
    };

    try {
        const response = await apiClient.post<ApiResponse<LoginResponseDataInterface>>(
            LOGIN_PATH,
            loginRequest,
        );

        const jwt = getJwtFromApiResponse(response);
        authService.userLoggedIn(username, jwt);
    } catch (error) {
        authService.handleUnauthorized();
    }
}


export const login2 = async (username: string, password: string) => {
    const loginRequest: LoginRequestDataInterface = {
        username: username,
        password: password,
    };

    try {
        return apiClient.post<ApiResponse<LoginResponseDataInterface>>(
            LOGIN_PATH,
            loginRequest,
        )
            .then(response => {
                if(response.status === 200) {
                    let jwt = response.data.data.jwt
                    if(jwt){
                        authService.userLoggedIn(username, jwt)
                        return ""
                    }
                }
                if(response.status === 406) {
                    console.log("data with 406: " + response)
                    return response.data.errors
                }
                console.log("err here: " + response)
                return "unexpected error: " + response.status
        })
            .catch(err => {
                if(err.response.status === 406) {
                    return err.response.data.errors
                }
                console.log(err)
                return "unexpected err"
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

export const logout = () => {
    authService.handleUnauthorized();
}