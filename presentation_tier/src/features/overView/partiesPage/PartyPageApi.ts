import {Party} from "../Party";
import {get} from "../../../api/Api";
import {PartyInvite} from "./PartyInvite";

const ATTENDED_PARTY_PATH = "http://localhost:8080/api/v0/party/getPartiesByParticipantId/"
const ORGANIZED_PARTY_PATH = "http://localhost:8080/api/v0/party/getPartiesByOrganizerId/"
const PENDING_PARTY_INVITES = "http://localhost:8080/api/v0/partyAttendanceManager/getPendingInvites/"
const ACCEPT_INVITE = "http://localhost:8080/api/v0/partyAttendanceManager/accept/"
const DECLINE_INVITE = "http://localhost:8080/api/v0/partyAttendanceManager/decline/"

export const getAttendedParties = async (): Promise<Party[]> => {
    return new Promise<Party[]>((resolve, reject) => {
        get<Party[]>(ATTENDED_PARTY_PATH)
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
        get<Party[]>(ORGANIZED_PARTY_PATH)
            .then((parties: Party[]) => {
                return resolve(parties);
            })
            .catch(err => {
                return reject(err);
            });
    });
};

export const getPartyInvites = async (): Promise<PartyInvite[]> => {
    return new Promise<PartyInvite[]>((resolve, reject) => {
        get<PartyInvite[]>(PENDING_PARTY_INVITES)
            .then((invites: PartyInvite[]) => {
                return resolve(invites);
            })
            .catch(err => {
                return reject(err);
            });
    });
};

export const acceptInvite = async (invitorId: number): Promise<void> => {
    return new Promise<void>((resolve, reject) => {
        console.log(invitorId)
        get<void>(ACCEPT_INVITE + invitorId.toString())
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
        console.log(invitorId)
        get<void>(DECLINE_INVITE + invitorId.toString())
            .then(() => {
                return resolve();
            })
            .catch(err => {
                return reject(err);
            });
    });
};