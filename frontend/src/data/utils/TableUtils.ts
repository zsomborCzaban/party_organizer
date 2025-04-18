import {PartyPopulated} from "../types/Party.ts";
import {PartyInviteTableRow, PartyTableRow} from "../constants/TableColumns.ts";
import {PartyInvite} from "../types/PartyInvite.ts";

export const convertPartiesToTableDatasource = (parties: PartyPopulated[]): PartyTableRow[] => {
    return parties.map(party => ({
        id: party.ID, name: party.name, organizerName: party.organizer.username, place: party.place, time: party.start_time.toString()
    }))
}

export const convertPartyInvitesToTableDatasource = (invites: PartyInvite[]) : PartyInviteTableRow[] => {
    return invites.map(invite => ({
        id: invite.party.ID, invitedBy: invite.invitor.username, partyName: invite.party.name, partyPlace: invite.party.place, partyTime: invite.party.start_time.toString()
    }))
}