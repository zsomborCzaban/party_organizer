import {get} from '../../api/Api';
import {FriendInvite} from '../types/FriendInvite';
import {getApiUrl} from '../../api/ApiHelper';

const FRIEND_MANGER_PATH = `${getApiUrl()  }/friendManager`;

export const inviteFriend = async (username: string): Promise<void> => new Promise<void>((resolve, reject) => {
        get<void>(`${FRIEND_MANGER_PATH  }/invite/${  username}`)
            .then(() => resolve())
            .catch((err) => reject(err));
    });

export const removeFriend = async (friendId: number): Promise<void> => new Promise<void>((resolve, reject) => {
        get<void>(`${FRIEND_MANGER_PATH  }/removeFriend/${  friendId.toString()}`)
            .then(() => resolve())
            .catch((err) => reject(err));
    });

export const acceptInvite = async (invitorId: number): Promise<void> => new Promise<void>((resolve, reject) => {
        get<void>(`${FRIEND_MANGER_PATH  }/accept/${  invitorId.toString()}`)
            .then(() => resolve())
            .catch((err) => reject(err));
    });

export const declineInvite = async (invitorId: number): Promise<void> => new Promise<void>((resolve, reject) => {
        get<void>(`${FRIEND_MANGER_PATH  }/decline/${  invitorId.toString()}`)
            .then(() => resolve())
            .catch((err) => reject(err));
    });

export const getFriendInvites = async (): Promise<FriendInvite[]> => new Promise<FriendInvite[]>((resolve, reject) => {
        get<FriendInvite[]>(`${FRIEND_MANGER_PATH  }/getPendingInvites`)
            .then((invites: FriendInvite[]) => resolve(invites))
            .catch((err) => reject(err));
    });
