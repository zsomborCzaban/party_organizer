import {User} from "../types/User";
import {get, postImage} from "../../api/Api";
import {getApiUrl} from "../../api/ApiHelper";
import {LoginResponseDataInterface} from "./AuthenticationApi";

const USER_PATH = getApiUrl() + "/user"

export const getFriends = async (): Promise<User[]> => {
    return new Promise<User[]>((resolve, reject) => {
        get<User[]>(USER_PATH + "/getFriends")
            .then((friends: User[]) => {
                return resolve(friends);
            })
            .catch(err => {
                return reject(err);
            });
    });
};

//this is a bit hacky. We access our current logged in user by the jwt token. so when we upload a new picture we get a new token with the updated picture
export const uploadPicture = async (formData: FormData): Promise<LoginResponseDataInterface> => {
    return new Promise<LoginResponseDataInterface>((resolve, reject) => {
        postImage<LoginResponseDataInterface>(USER_PATH+ "/uploadProfilePicture", formData)
            .then((response: LoginResponseDataInterface) => {
                return resolve(response)
            })
            .catch(err => {
                return reject(err)
            });

    });
}
