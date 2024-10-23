import {get} from "../../../api/Api";
import {FriendInvite} from "./FriendInvite";
import {User} from "../User";

const GET_FRIENDS_PATH = "http://localhost:8080/api/v0/user/getFriends/"
const PENDING_FRIEND_INVITES_PATH = "http://localhost:8080/api/v0/friendManager/getPendingInvites/"
const ACCEPT_INVITE_PATH = "http://localhost:8080/api/v0/friendManager/accept/"
const DECLINE_INVITE_PATH = "http://localhost:8080/api/v0/friendManager/decline/"

export const getFriends = async (): Promise<User[]> => {
    return new Promise<User[]>((resolve, reject) => {
        get<User[]>(GET_FRIENDS_PATH)
            .then((friends: User[]) => {
                return resolve(friends);
            })
            .catch(err => {
                return reject(err);
            });
    });
};

export const getFriendInvites = async (): Promise<FriendInvite[]> => {
    return new Promise<FriendInvite[]>((resolve, reject) => {
        get<FriendInvite[]>(PENDING_FRIEND_INVITES_PATH)
            .then((invites: FriendInvite[]) => {
                return resolve(invites);
            })
            .catch(err => {
                return reject(err);
            });
    });
};

export const acceptInvite = async (invitorId: number): Promise<void> => {
    return new Promise<void>((resolve, reject) => {
        get<void>(ACCEPT_INVITE_PATH + invitorId.toString())
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
        get<void>(DECLINE_INVITE_PATH + invitorId.toString())
            .then(() => {
                return resolve();
            })
            .catch(err => {
                return reject(err);
            });
    });
};