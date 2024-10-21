import {User} from "../User";
import {Party} from "../Party";

export interface PartyInvite {
    Invitor: User;
    Party: Party
    State: string
}