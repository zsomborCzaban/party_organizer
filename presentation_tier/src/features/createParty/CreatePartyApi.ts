import {post} from "../../api/Api";
import {Party} from "../overView/Party";

const PARTY_PATH = 'http://localhost:8080/api/v0/party/';


export const createParty = async (requestBody: Party): Promise<Party> => {
    return new Promise<Party>((resolve, reject) => {
        post<Party>(PARTY_PATH, requestBody)
            .then((party) => {return resolve(party)})
            .catch(error => {
                return reject(error)
            })
    })
}