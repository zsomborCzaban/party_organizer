import {BACKEND_URL} from "../constants/backend_url";
import {User} from "../types/User";
import {get} from "../../api/Api";

const USER_PATH = BACKEND_URL + "/user"

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
