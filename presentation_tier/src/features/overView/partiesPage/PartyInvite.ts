import {User} from "../User";
import {Party} from "../Party";

export interface PartyInvite {
    ID: number,
    invitor: User,
    invited: User,
    party: Party,
    state: string,
}