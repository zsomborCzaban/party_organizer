import {PartyInvite} from "../types/PartyInvite";
import {get} from "../../api/Api";
import {Party} from "../types/Party";
import {getApiUrl} from "../../api/ApiHelper";

const PARTY_ATTENDANCE_MANAGER_PATH = getApiUrl() + "/partyAttendanceManager"

export const getPartyPendingInvites = async (partyId: number): Promise<PartyInvite[]> => {
    return new Promise<PartyInvite[]>((resolve, reject) => {
        get<PartyInvite[]>(PARTY_ATTENDANCE_MANAGER_PATH + "/getPartyPendingInvites/" + partyId.toString())
            .then((invites: PartyInvite[]) => {
                return resolve(invites);
            })
            .catch(err => {
                return reject(err);
            });
    });
};

export const inviteToParty = async (partyId: number, username: string): Promise<void> => {
    return new Promise<void>((resolve, reject) => {
        get<void>(PARTY_ATTENDANCE_MANAGER_PATH + "/invite/" + partyId.toString() + '/' + username)
            .then(() => {
                return resolve();
            })
            .catch(err => {
                return reject(err);
            });
    });
};

export const kickFromParty = async (partyId: number, userId: number): Promise<void> => {
    return new Promise<void>((resolve, reject) => {
        get<void>(PARTY_ATTENDANCE_MANAGER_PATH + "/kick/" + partyId.toString() + '/' + userId.toString())
            .then(() => {
                return resolve();
            })
            .catch(err => {
                return reject(err);
            });
    });
};

export const getPartyInvites = async (): Promise<PartyInvite[]> => {
    return new Promise<PartyInvite[]>((resolve, reject) => {
        get<PartyInvite[]>(PARTY_ATTENDANCE_MANAGER_PATH + "/getPendingInvites")
            .then((invites: PartyInvite[]) => {
                return resolve(invites);
            })
            .catch(err => {
                return reject(err);
            });
    });
};

export const acceptInvite = async (partyID: number): Promise<void> => {
    return new Promise<void>((resolve, reject) => {
        get<void>(PARTY_ATTENDANCE_MANAGER_PATH + "/accept/" + partyID.toString())
            .then(() => {
                return resolve();
            })
            .catch(err => {
                return reject(err);
            });
    });
};

export const declineInvite = async (partyID: number): Promise<void> => {
    return new Promise<void>((resolve, reject) => {
        get<void>(PARTY_ATTENDANCE_MANAGER_PATH + "/decline/" + partyID.toString())
            .then(() => {
                return resolve();
            })
            .catch(err => {
                return reject(err);
            });
    });
};


export const joinPublicParty = async (partyId: number): Promise<Party> => {
    return new Promise<Party>((resolve, reject) => {
        get<Party>(PARTY_ATTENDANCE_MANAGER_PATH + "/joinPublicParty/" + partyId.toString())
            .then((party: Party) => {
                return resolve(party);
            })
            .catch(err => {
                return reject(err);
            });
    });
};
