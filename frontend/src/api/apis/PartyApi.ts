import { get, post, put } from '../../api/Api';
import { getApiUrl } from '../../api/ApiHelper';
import {Party, PartyPopulated} from '../../data/types/Party';
import { User } from '../../data/types/User';
import axios, {AxiosInstance, AxiosResponse} from "axios";

const PARTY_PATH = `${getApiUrl()}/party`;

const handleApiResponse = <T>(response: AxiosResponse<T>): T => {
    return response.data;
};

const handleApiError = (error: unknown) => {
    // TODO: handle errors as needed
    if (axios.isAxiosError(error)) {
        console.error(`Axios error: ${error.message}`);
    } else {
        console.error(`Unexpected error: ${error}`);
    }
};

export type PublicPartiesResponse = {
    data: PartyPopulated[]
}

export class PartyApi {
    private axiosInstance: AxiosInstance;

    constructor(axiosInstance: AxiosInstance) {
        this.axiosInstance = axiosInstance;
    }

    async getPublicParties(): Promise< PublicPartiesResponse | 'error'> {
        try {
            const response = await this.axiosInstance.get<PublicPartiesResponse>(`${getApiUrl()}/publicParties`)
            // toast.success('Public parties received')
            return handleApiResponse(response)
        } catch (error) {
            handleApiError(error)
            return 'error'
        }
    }

    async getAttendedParties(): Promise< PublicPartiesResponse | 'error'> {
        try {
            const response = await this.axiosInstance.get<PublicPartiesResponse>(`${PARTY_PATH}/getPartiesByParticipantId`)
            // toast.success('Public parties received')
            return handleApiResponse(response)
        } catch (error) {
            handleApiError(error)
            return 'error'
        }
    }

    async getOrganizedParties(): Promise< PublicPartiesResponse | 'error'> {
        try {
            const response = await this.axiosInstance.get<PublicPartiesResponse>(`${PARTY_PATH}/getPartiesByOrganizerId`)
            // toast.success('Public parties received')
            return handleApiResponse(response)
        } catch (error) {
            handleApiError(error)
            return 'error'
        }
    }

}

export const createParty = async (requestBody: Party): Promise<Party> =>
  new Promise<Party>((resolve, reject) => {
    post<Party>(PARTY_PATH, requestBody)
      .then((party) => resolve(party))
      .catch((error) => reject(error));
  });

export const updateParty = async (requestBody: Party): Promise<Party> =>
  new Promise<Party>((resolve, reject) => {
    put<Party>(PARTY_PATH, requestBody)
      .then((party) => resolve(party))
      .catch((error) => reject(error));
  });

export const getPublicParties = async (): Promise<Party[]> =>
  new Promise<Party[]>((resolve, reject) => {
    get<Party[]>(`${PARTY_PATH}/getPublicParties`)
      .then((parties: Party[]) => resolve(parties))
      .catch((err) => reject(err));
  });

export const getAttendedParties = async (): Promise<Party[]> =>
  new Promise<Party[]>((resolve, reject) => {
    get<Party[]>(`${PARTY_PATH}/getPartiesByParticipantId`)
      .then((parties: Party[]) => resolve(parties))
      .catch((err) => reject(err));
  });

export const getOrganizedParties = async (): Promise<Party[]> =>
  new Promise<Party[]>((resolve, reject) => {
    get<Party[]>(`${PARTY_PATH}/getPartiesByOrganizerId`)
      .then((parties: Party[]) => resolve(parties))
      .catch((err) => reject(err));
  });

export const getPartyParticipants = async (partyId: number): Promise<User[]> =>
  new Promise<User[]>((resolve, reject) => {
    get<User[]>(`${PARTY_PATH}/getParticipants/${partyId}`)
      .then((participants: User[]) => resolve(participants))
      .catch((err) => reject(err));
  });
