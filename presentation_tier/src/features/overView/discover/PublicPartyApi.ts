import {Party} from "../Party";
import {get} from "../../../api/Api";

const PUBLIC_PARTY_PATH = "http://localhost:8080/api/v0/party/getPublicParties/"
const JOIN_PUBLIC_PARTY_PATH = "http://localhost:8080/api/v0/partyAttendanceManager/joinPublicParty/"

export const getPublicParties = async (): Promise<Party[]> => {
    return new Promise<Party[]>((resolve, reject) => {
        get<Party[]>(PUBLIC_PARTY_PATH)
            .then((parties: Party[]) => {
                return resolve(parties);
            })
            .catch(err => {
                return reject(err);
            });
    });
};

export const joinPublicParty = async (partyId: number): Promise<Party> => {
    return new Promise<Party>((resolve, reject) => {
        get<Party>(JOIN_PUBLIC_PARTY_PATH + partyId.toString())
            .then((party: Party) => {
                return resolve(party);
            })
            .catch(err => {
                return reject(err);
            });
    });
};