import { get } from '../Api.ts';

import { getApiUrl } from '../ApiHelper.ts';
import {Party, PartyPopulated} from '../../data/types/Party';
import { PartyInvite } from '../../data/types/PartyInvite';
import axios, {AxiosInstance, AxiosResponse} from "axios";
import {User} from "../../data/types/User.ts";
import {toast} from "sonner";

const PARTY_ATTENDANCE_MANAGER_PATH = `${getApiUrl()}/partyAttendanceManager`;

const handleApiResponse = <T>(response: AxiosResponse<T>): T => {
    return response.data;
};

const handleApiError = (error: unknown) => {
    if (axios.isAxiosError(error)) {
        console.error(`Axios error: ${error.message}`);
    } else {
        console.error(`Unexpected error: ${error}`);
    }
};

export type PartyInviteInvitesResponse = {
    data: {
        ID: number,
        invitor: User,
        invited: User,
        party: PartyPopulated,
        state: string,
    }[]
}

export type JoinPartyResponse = {
    data: PartyPopulated
}

export type LeavePartyResponse = {
    data: string,
}

export class PartyAttendanceManagerApi {
    private axiosInstance: AxiosInstance;

    constructor(axiosInstance: AxiosInstance) {
        this.axiosInstance = axiosInstance;
    }

    async getPendingInvites(): Promise<PartyInviteInvitesResponse | 'error'> {
        try {
            const response = await this.axiosInstance.get<PartyInviteInvitesResponse>(`${PARTY_ATTENDANCE_MANAGER_PATH}/getPendingInvites`)
            // toast.success('Pending invites received')
            return handleApiResponse(response)
        } catch (error) {
            handleApiError(error)
            return 'error'
        }
    }

    async acceptInvite(partyId: number): Promise<PartyInviteInvitesResponse | 'error'> {
        try {
            const response = await this.axiosInstance.get<PartyInviteInvitesResponse>(`${PARTY_ATTENDANCE_MANAGER_PATH}/accept/${partyId.toString()}`)
            toast.success('Invite accepted')
            return handleApiResponse(response)
        } catch (error){
            handleApiError(error)
            return 'error'
        }
    }

    async declineInvite(partyId: number): Promise<PartyInviteInvitesResponse | 'error'> {
        try {
            const response = await this.axiosInstance.get<PartyInviteInvitesResponse>(`${PARTY_ATTENDANCE_MANAGER_PATH}/decline/${partyId.toString()}`)
            toast.success('Invite declined')
            return handleApiResponse(response)
        } catch (error){
            handleApiError(error)
            return 'error'
        }
    }

    async getPartyPendingInvites(partyId: number): Promise<PartyInviteInvitesResponse | 'error'> {
        try {
            const response = await this.axiosInstance.get<PartyInviteInvitesResponse>(`${PARTY_ATTENDANCE_MANAGER_PATH}/getPartyPendingInvites/${partyId}`)
            return handleApiResponse(response)
        } catch (error) {
            handleApiError(error)
            return 'error'
        }
    }

    async joinPublicParty(partyId: number): Promise<JoinPartyResponse | 'error'> {
        try {
            const response = await this.axiosInstance.get<JoinPartyResponse>(`${PARTY_ATTENDANCE_MANAGER_PATH}/joinPublicParty/${partyId.toString()}`)
            return handleApiResponse(response)
        } catch (error) {
            handleApiError(error)
            return 'error'
        }
    }

    async joinPrivateParty(accessCode: string): Promise<JoinPartyResponse | 'error'> {
        try {
            const response = await this.axiosInstance.get<JoinPartyResponse>(`${PARTY_ATTENDANCE_MANAGER_PATH}/joinPrivateParty/${accessCode}`)
            return handleApiResponse(response)
        } catch (error) {
            handleApiError(error)
            return 'error'
        }
    }

    async leaveParty(partyId: number): Promise<LeavePartyResponse | 'error'> {
        try {
            const response = await this.axiosInstance.get<LeavePartyResponse>(`${PARTY_ATTENDANCE_MANAGER_PATH}/leaveParty/${partyId}`)
            console.log(response)
            if(response.response){
                if(response.response.data){
                    return response.response.data
                } else {
                    return response.response.errors
                }
            }
            return handleApiResponse(response)

        } catch (error) {
            handleApiError(error)
            return 'error'
        }
    }

}

export const getPartyPendingInvites = async (partyId: number): Promise<PartyInvite[]> =>
  new Promise<PartyInvite[]>((resolve, reject) => {
    get<PartyInvite[]>(`${PARTY_ATTENDANCE_MANAGER_PATH}/getPartyPendingInvites/${partyId.toString()}`)
      .then((invites: PartyInvite[]) => resolve(invites))
      .catch((err) => reject(err));
  });

export const inviteToParty = async (partyId: number, username: string): Promise<void> =>
  new Promise<void>((resolve, reject) => {
    get<void>(`${PARTY_ATTENDANCE_MANAGER_PATH}/invite/${partyId.toString()}/${username}`)
      .then(() => resolve())
      .catch((err) => reject(err));
  });

export const kickFromParty = async (partyId: number, userId: number): Promise<void> =>
  new Promise<void>((resolve, reject) => {
    get<void>(`${PARTY_ATTENDANCE_MANAGER_PATH}/kick/${partyId.toString()}/${userId.toString()}`)
      .then(() => resolve())
      .catch((err) => reject(err));
  });

export const getPartyInvites = async (): Promise<PartyInvite[]> =>
  new Promise<PartyInvite[]>((resolve, reject) => {
    get<PartyInvite[]>(`${PARTY_ATTENDANCE_MANAGER_PATH}/getPendingInvites`)
      .then((invites: PartyInvite[]) => resolve(invites))
      .catch((err) => reject(err));
  });

export const acceptInvite = async (partyID: number): Promise<void> =>
  new Promise<void>((resolve, reject) => {
    get<void>(`${PARTY_ATTENDANCE_MANAGER_PATH}/accept/${partyID.toString()}`)
      .then(() => resolve())
      .catch((err) => reject(err));
  });

export const declineInvite = async (partyID: number): Promise<void> =>
  new Promise<void>((resolve, reject) => {
    get<void>(`${PARTY_ATTENDANCE_MANAGER_PATH}/decline/${partyID.toString()}`)
      .then(() => resolve())
      .catch((err) => reject(err));
  });

export const joinPublicParty = async (partyId: number): Promise<Party> =>
  new Promise<Party>((resolve, reject) => {
    get<Party>(`${PARTY_ATTENDANCE_MANAGER_PATH}/joinPublicParty/${partyId.toString()}`)
      .then((party: Party) => resolve(party))
      .catch((err) => reject(err));
  });

export const joinPrivateParty = async (accessCode: string): Promise<PartyPopulated> =>
  new Promise<PartyPopulated>((resolve, reject) => {
    get<PartyPopulated>(`${PARTY_ATTENDANCE_MANAGER_PATH}/joinPrivateParty/${accessCode}`)
      .then((party: PartyPopulated) => resolve(party))
      .catch((err) => reject(err));
  });
