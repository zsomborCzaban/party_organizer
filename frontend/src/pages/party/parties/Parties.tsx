import {useApi} from "../../../context/ApiContext.ts";
import {useEffect, useState} from "react";
import {PartyInvite} from "../../../data/types/PartyInvite.ts";
import {toast} from "sonner";
import {PartyPopulated} from "../../../data/types/Party.ts";
import {SortableTable} from "../../../components/table/SortableTable.tsx";
import {
    PartyInviteTableRow,
    partyInviteTableColumns,
    partyTableColumns,
    PartyTableRow
} from "../../../data/constants/TableColumns.ts";

export const Parties = () => {

    const api = useApi();
    const [pendingInvites, setPendingInvites] = useState<PartyInvite[]>([])
    const [organizedParties, setOrganizedParties] = useState<PartyPopulated[]>([])
    const [attendedParties, setAttendedParties] = useState<PartyPopulated[]>([])

    useEffect(() => {
        api.partyAttendanceApi.getPendingInvites().then(result => {
            if(result === 'error'){
                toast.error('Error while loading invites');
                setPendingInvites([]);
                return;
            }
            setPendingInvites(result.data)
        }).catch(() => {
            toast.error('Error while loading invites');
            setPendingInvites([]);
            return;
        })
    }, []);

    useEffect(() => {
        api.partyApi.getOrganizedParties().then(result => {
            if(result === 'error'){
                toast.error('Error while loading organized parties');
                setOrganizedParties([]);
                return;
            }
            setOrganizedParties(result.data)
        }).catch(() => {
            toast.error('Error while loading organized parties');
            setOrganizedParties([]);
            return;
        })
    }, []);

    useEffect(() => {
        api.partyApi.getAttendedParties().then(result => {
            if(result === 'error'){
                toast.error('Error while loading attending parties');
                setAttendedParties([]);
                return;
            }
            setAttendedParties(result.data)
        }).catch(() => {
            toast.error('Error while loading attending parties');
            setAttendedParties([]);
            return;
        })
    }, []);

    const convertInvitesToTableDatasource = (invites: PartyInvite[]) : PartyInviteTableRow[] => {
        return invites.map(invite => ({
            id: invite.ID, invitedBy: invite.invitor.username, partyName: invite.party.name, partyPlace: invite.party.place, partyTime: invite.party.start_time
        }))
    }

    const convertPartiesToTableDatasource = (parties: PartyPopulated[]): PartyTableRow[] => {
        return parties.map(party => ({
            id: party.ID, name: party.name, organizerName: party.organizer.username, place: party.place, time: party.start_time
        }))
    }


    return (<div>
        Parties page
        Show:
        <ul>
            <li>Party Invites</li>
            <li>Attended parties</li>
            <li>Organized parties</li>
        </ul>

        <div>
            <h1>Party Invites </h1>
            <SortableTable
                columns={partyInviteTableColumns}
                data={convertInvitesToTableDatasource(pendingInvites)}
                rowsPerPageOptions={[3,5,10,15]}
                defaultRowsPerPage={5}
                // defaultSortField="name"
            />
        </div>
        <div>
            <h1>Attended parties</h1>
            <SortableTable
                columns={partyTableColumns}
                data={convertPartiesToTableDatasource(attendedParties)}
                rowsPerPageOptions={[3,5,10,15]}
                defaultRowsPerPage={5}
                // defaultSortField="name"
            />
        </div>
        <div>
            <h1>Organized parties</h1>
            <SortableTable
                columns={partyTableColumns}
                data={convertPartiesToTableDatasource(organizedParties)}
                rowsPerPageOptions={[3,5,10,15]}
                defaultRowsPerPage={5}
                // defaultSortField="name"
            />
        </div>

    </div>)
};