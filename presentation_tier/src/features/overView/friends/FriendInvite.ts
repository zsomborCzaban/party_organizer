import {User} from "../User";

export interface FriendInvite {
    ID: number,
    invitor: User,
    invited: User,
    state: string,
}