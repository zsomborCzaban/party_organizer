import {User} from "../types/User";
import {get} from "../../api/Api";
import {getApiUrl} from "../../api/ApiHelper";

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
