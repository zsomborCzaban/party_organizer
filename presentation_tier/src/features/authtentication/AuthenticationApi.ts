import {AxiosResponse} from "axios";
import {ApiResponse} from "../../api/ApiResponse";
import {jwtDecode} from "jwt-decode";
import apiClient from "../../api/Api";
import {authService} from "../../auth/AuthService";

const LOGIN_PATH = 'http://localhost:8080/api/v0/user/login';

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

    if(loginResponse.isError || loginResponse.code != 200) {
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
        const response = await apiClient.post<LoginResponseDataInterface>(
            LOGIN_PATH,
            loginRequest,
        );

        const jwt = getJwtFromApiResponse(response);
        authService.userLoggedIn(username, jwt);
    } catch (error) {
        authService.handleUnauthorized();
    }
}

export const logout = () => {
    authService.handleUnauthorized();
}