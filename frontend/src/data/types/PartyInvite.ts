import {User} from './User';
import {PartyPopulated} from './Party';

export interface PartyInvite {
    ID: number,
    invitor: User,
    invited: User,
    party: PartyPopulated,
    state: string,
}