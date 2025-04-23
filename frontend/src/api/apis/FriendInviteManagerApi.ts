import { get } from '../../api/Api';
import { getApiUrl } from '../../api/ApiHelper';
import { FriendInvite } from '../../data/types/FriendInvite';
import axios, {AxiosInstance, AxiosResponse} from "axios";
import {toast} from "sonner";

const FRIEND_MANGER_PATH = `${getApiUrl()}/friendManager`;

const handleApiResponse = <T>(response: AxiosResponse<T>): T => {
    return response.data;
};

const handleApiError = (error: unknown) => {
    // TODO: handle errors as needed
    if (axios.isAxiosError(error)) {
        console.error(`Axios error: ${error.message}`);
        if(error.status === 404) {
            toast.error('Not found')
        } else {
            toast.error('Unexpected error')
        }
    } else {
        console.error(`Unexpected error: ${error}`);
        toast.error('Unexpected error')
    }
};

export type FriendInviteResponse = {
    data: FriendInvite
}

export type FriendRemoveResponse = {
    data: string
}

export type GetInvitesResponse = {
    data: FriendInvite[]
}


export class FriendInviteManagerApi {
    private axiosInstance: AxiosInstance;

    constructor(axiosInstance: AxiosInstance) {
        this.axiosInstance = axiosInstance;
    }

    async inviteFriend(username: string): Promise<FriendInviteResponse | 'error'> {
        try {
            const response = await this.axiosInstance.get<FriendInviteResponse>(`${FRIEND_MANGER_PATH}/invite/${username}`)
            toast.success('Invite sent')
            return handleApiResponse(response)
        } catch (error) {
            handleApiError(error)
            return 'error'
        }
    }

    async removeFriend(friendId: number): Promise<FriendRemoveResponse | 'error'> {
        try {
            const response = await this.axiosInstance.get<FriendRemoveResponse>(`${FRIEND_MANGER_PATH}/removeFriend/${friendId.toString()}`)
            toast.success('Friend removed')
            return handleApiResponse(response)
        } catch (error) {
            handleApiError(error)
            return 'error'
        }
    }

    async acceptInvite(friendId: number): Promise<FriendInviteResponse | 'error'> {
        try {
            const response = await this.axiosInstance.get<FriendInviteResponse>(`${FRIEND_MANGER_PATH}/accept/${friendId.toString()}`)
            toast.success('Invite accepted')
            return handleApiResponse(response)
        } catch (error) {
            handleApiError(error)
            return 'error'
        }
    }

    async declineInvite(friendId: number): Promise<FriendInviteResponse | 'error'> {
        try {
            const response = await this.axiosInstance.get<FriendInviteResponse>(`${FRIEND_MANGER_PATH}/decline/${friendId.toString()}`)
            toast.success('Invite declined')
            return handleApiResponse(response)
        } catch (error) {
            handleApiError(error)
            return 'error'
        }
    }

    async getPendingInvites(): Promise<GetInvitesResponse | 'error'> {
        try {
            const response = await this.axiosInstance.get<GetInvitesResponse>(`${FRIEND_MANGER_PATH}/getPendingInvites`)
            return handleApiResponse(response)
        } catch (error) {
            handleApiError(error)
            return 'error'
        }
    }



}

export const inviteFriend = async (username: string): Promise<void> =>
  new Promise<void>((resolve, reject) => {
    get<void>(`${FRIEND_MANGER_PATH}/invite/${username}`)
      .then(() => resolve())
      .catch((err) => reject(err));
  });

export const removeFriend = async (friendId: number): Promise<void> =>
  new Promise<void>((resolve, reject) => {
    get<void>(`${FRIEND_MANGER_PATH}/removeFriend/${friendId.toString()}`)
      .then(() => resolve())
      .catch((err) => reject(err));
  });

export const acceptInvite = async (invitorId: number): Promise<void> =>
  new Promise<void>((resolve, reject) => {
    get<void>(`${FRIEND_MANGER_PATH}/accept/${invitorId.toString()}`)
      .then(() => resolve())
      .catch((err) => reject(err));
  });

export const declineInvite = async (invitorId: number): Promise<void> =>
  new Promise<void>((resolve, reject) => {
    get<void>(`${FRIEND_MANGER_PATH}/decline/${invitorId.toString()}`)
      .then(() => resolve())
      .catch((err) => reject(err));
  });

export const getFriendInvites = async (): Promise<FriendInvite[]> =>
  new Promise<FriendInvite[]>((resolve, reject) => {
    get<FriendInvite[]>(`${FRIEND_MANGER_PATH}/getPendingInvites`)
      .then((invites: FriendInvite[]) => resolve(invites))
      .catch((err) => reject(err));
  });
