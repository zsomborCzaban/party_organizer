import {get, post} from "../../api/Api";
import {Party} from "../types/Party";
import {User} from "../types/User";
import {getApiUrl} from "../../api/ApiHelper";

const PARTY_PATH =  getApiUrl() + '/party';


export const createParty = async (requestBody: Party): Promise<Party> => {
    return new Promise<Party>((resolve, reject) => {
        post<Party>(PARTY_PATH, requestBody)
            .then((party) => {return resolve(party)})
            .catch(error => {
                return reject(error)
            })
    })
}

export const getPublicParties = async (): Promise<Party[]> => {
    return new Promise<Party[]>((resolve, reject) => {
        get<Party[]>(PARTY_PATH + "/getPublicParties")
            .then((parties: Party[]) => {
                return resolve(parties);
            })
            .catch(err => {
                return reject(err);
            });
    });
};


export const getAttendedParties = async (): Promise<Party[]> => {
    return new Promise<Party[]>((resolve, reject) => {
        get<Party[]>(PARTY_PATH + "/getPartiesByParticipantId")
            .then((parties: Party[]) => {
                return resolve(parties);
            })
            .catch(err => {
                return reject(err);
            });
    });
};

export const getOrganizedParties = async (): Promise<Party[]> => {
    return new Promise<Party[]>((resolve, reject) => {
        get<Party[]>(PARTY_PATH + "/getPartiesByOrganizerId")
            .then((parties: Party[]) => {
                return resolve(parties);
            })
            .catch(err => {
                return reject(err);
            });
    });
};


export const getPartyParticipants = async (partyId: number): Promise<User[]> => {
    return new Promise<User[]>((resolve, reject) => {
        get<User[]>(PARTY_PATH + "/getParticipants/" + partyId)
            .then((participants: User[]) => {
                return resolve(participants);
            })
            .catch(err => {
                return reject(err);
            });
    });
};
