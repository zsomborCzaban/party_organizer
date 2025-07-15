import {PartyPopulated} from "../types/Party.ts";
import {PartyInviteTableRow, PartyTableRow} from "../constants/TableColumns.ts";
import {PartyInvite} from "../types/PartyInvite.ts";
import {User} from "../types/User.ts";
import {FriendInviteTableRow, FriendTableRow} from "../constants/TableColumns.tsx";
import {FriendInvite} from "../types/FriendInvite.ts";

export const convertPartiesToTableDatasource = (parties: PartyPopulated[]): PartyTableRow[] => {
    return parties.map(party => ({
        id: party.ID, name: party.name, organizerName: party.organizer.username, place: party.place, time: party.start_time.toString(), organizerProfilePicture: party.organizer.profile_picture_url,
    }))
}

export const convertPartyInvitesToTableDatasource = (invites: PartyInvite[]) : PartyInviteTableRow[] => {
    return invites.map(invite => ({
        id: invite.party.ID, invitedBy: invite.invitor.username, partyName: invite.party.name, partyPlace: invite.party.place, partyTime: invite.party.start_time.toString(), invitorProfilePicture: invite.invitor.profile_picture_url,
    }))
}

export const convertFriendsToTableData = (friends: User[]): FriendTableRow[] => {
    return friends.map(friend => ({
        id: friend.id,
        username: friend.username,
        email: friend.email,
        friendProfilePicture: friend.profile_picture_url,
    }));
};

export const convertInvitesToTableData = (invites: FriendInvite[]): FriendInviteTableRow[] => {
    return invites.map(invite => ({
        id: invite.invitor.id,
        invitedBy: invite.invitor.username,
        invitorProfilePicture: invite.invitor.profile_picture_url,
    }));
};