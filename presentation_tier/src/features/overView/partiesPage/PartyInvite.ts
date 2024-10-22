import {User} from "../User";
import {Party} from "../Party";

export interface PartyInvite {
    ID: number
    invitor: User;
    party: Party
    party_id: number
    state: string
}