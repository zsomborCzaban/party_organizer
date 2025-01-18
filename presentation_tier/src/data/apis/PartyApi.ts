import {get, post, put} from '../../api/Api';
import {Party} from '../types/Party';
import {User} from '../types/User';
import {getApiUrl} from '../../api/ApiHelper';

const PARTY_PATH =  `${getApiUrl()  }/party`;


export const createParty = async (requestBody: Party): Promise<Party> => new Promise<Party>((resolve, reject) => {
        post<Party>(PARTY_PATH, requestBody)
            .then((party) => resolve(party))
            .catch((error) => reject(error));
    });

export const updateParty = async (requestBody: Party): Promise<Party> => new Promise<Party>((resolve, reject) => {
        put<Party>(PARTY_PATH, requestBody)
            .then((party) => resolve(party))
            .catch((error) => reject(error));
    });

export const getPublicParties = async (): Promise<Party[]> => new Promise<Party[]>((resolve, reject) => {
        get<Party[]>(`${PARTY_PATH  }/getPublicParties`)
            .then((parties: Party[]) => resolve(parties))
            .catch((err) => reject(err));
    });


export const getAttendedParties = async (): Promise<Party[]> => new Promise<Party[]>((resolve, reject) => {
        get<Party[]>(`${PARTY_PATH  }/getPartiesByParticipantId`)
            .then((parties: Party[]) => resolve(parties))
            .catch((err) => reject(err));
    });

export const getOrganizedParties = async (): Promise<Party[]> => new Promise<Party[]>((resolve, reject) => {
        get<Party[]>(`${PARTY_PATH  }/getPartiesByOrganizerId`)
            .then((parties: Party[]) => resolve(parties))
            .catch((err) => reject(err));
    });


export const getPartyParticipants = async (partyId: number): Promise<User[]> => new Promise<User[]>((resolve, reject) => {
        get<User[]>(`${PARTY_PATH  }/getParticipants/${  partyId}`)
            .then((participants: User[]) => resolve(participants))
            .catch((err) => reject(err));
    });
