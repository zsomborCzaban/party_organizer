import {get} from "../../api/Api";
import {FriendInvite} from "../types/FriendInvite";
import {BACKEND_URL} from "../constants/backend_url";

const FRIEND_MANGER_PATH = BACKEND_URL + "/friendManager"

export const inviteFriend = async (username: string): Promise<void> => {
    return new Promise<void>((resolve, reject) => {
        get<void>(FRIEND_MANGER_PATH + "/invite/" + username)
            .then(() => {
                return resolve();
            })
            .catch(err => {
                return reject(err);
            });
    });
};

export const removeFriend = async (friendId: number): Promise<void> => {
    return new Promise<void>((resolve, reject) => {
        get<void>(FRIEND_MANGER_PATH + "/removeFriend/" + friendId.toString())
            .then(() => {
                return resolve();
            })
            .catch(err => {
                return reject(err);
            });
    });
};

export const acceptInvite = async (invitorId: number): Promise<void> => {
    return new Promise<void>((resolve, reject) => {
        get<void>(FRIEND_MANGER_PATH + "/accept/" + invitorId.toString())
            .then(() => {
                return resolve();
            })
            .catch(err => {
                return reject(err);
            });
    });
};

export const declineInvite = async (invitorId: number): Promise<void> => {
    return new Promise<void>((resolve, reject) => {
        get<void>(FRIEND_MANGER_PATH + "/decline/" + invitorId.toString())
            .then(() => {
                return resolve();
            })
            .catch(err => {
                return reject(err);
            });
    });
};

export const getFriendInvites = async (): Promise<FriendInvite[]> => {
    return new Promise<FriendInvite[]>((resolve, reject) => {
        get<FriendInvite[]>(FRIEND_MANGER_PATH + "/getPendingInvites")
            .then((invites: FriendInvite[]) => {
                if (invites[0]){
                }
                return resolve(invites);
            })
            .catch(err => {
                return reject(err);
            });
    });
};
