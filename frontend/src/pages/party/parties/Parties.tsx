import {useApi} from "../../../context/ApiContext.ts";
import {useEffect, useState} from "react";
import {PartyInvite} from "../../../data/types/PartyInvite.ts";
import {toast} from "sonner";
import {PartyPopulated} from "../../../data/types/Party.ts";
import {ActionButton, SortableTable} from "../../../components/table/SortableTable.tsx";
import {
    PartyInviteTableRow,
    partyInviteTableColumns,
    partyTableColumns,
    PartyTableRow
} from "../../../data/constants/TableColumns.ts";
import {useNavigate} from "react-router-dom";
import classes from './Parties.module.scss';

export const Parties = () => {

    const api = useApi();
    const navigate = useNavigate()
    const [pendingInvites, setPendingInvites] = useState<PartyInvite[]>([])
    const [organizedParties, setOrganizedParties] = useState<PartyPopulated[]>([])
    const [attendedParties, setAttendedParties] = useState<PartyPopulated[]>([])
    const [reloadInvites, setReloadInvites] = useState(0)
    const [reloadAttendiedParties, setReloadAttendedParties] = useState(0)

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
    }, [api.partyAttendanceApi, reloadInvites]);

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
    }, [api.partyApi]);

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
    }, [api.partyApi, reloadAttendiedParties]);

    const convertInvitesToTableDatasource = (invites: PartyInvite[]) : PartyInviteTableRow[] => {
        return invites.map(invite => ({
            id: invite.party.ID, invitedBy: invite.invitor.username, partyName: invite.party.name, partyPlace: invite.party.place, partyTime: invite.party.start_time
        }))
    }

    const convertPartiesToTableDatasource = (parties: PartyPopulated[]): PartyTableRow[] => {
        return parties.map(party => ({
            id: party.ID, name: party.name, organizerName: party.organizer.username, place: party.place, time: party.start_time
        }))
    }

    const partyInviteActionButtons: ActionButton<PartyInviteTableRow>[] = [
        {
            label: 'Accept',
            color: 'success',
            onClick: (invite: PartyInviteTableRow) => {
                api.partyAttendanceApi.acceptInvite(invite.id)
                    .then((response) => {
                        if(response === 'error'){
                            toast.error('Could not accept invite, try again later')
                            return
                        }
                    setReloadInvites((reloadInvites+1)%2)
                    setReloadAttendedParties((reloadAttendiedParties+1)%2)
                })
                    .catch(() => {
                        toast.error('Could not accept invite, try again later')
                    })
            }
        },
        {
            label: 'Decline',
            color: 'error',
            onClick: (invite: PartyInviteTableRow) => {
                api.partyAttendanceApi.declineInvite(invite.id)
                    .then((response) => {
                        if(response === 'error'){
                            toast.error('Could not decline invite, try again later')
                            return
                        }
                        setReloadInvites((reloadInvites+1)%2)
                    })
                    .catch(() => {
                        toast.error('Could not decline invite, try again later')
                    })
            }
        }
    ];

    const partyActionButtons: ActionButton<PartyTableRow>[] = [
        {
            label: 'Visit',
            color: 'info',
            onClick: (party: PartyTableRow) => navigate(`/partyHome?id=${party.id}`) //todo: navigate to party page
        }
    ];

    return (
        <div className={classes.container}>
            <div className={classes.header}>
                <h1>My Parties</h1>
                <p className={classes.description}>
                    Manage your party invitations and view your attended and organized parties
                </p>
            </div>

            <div className={classes.section}>
                <h2>Party Invites</h2>
                <div className={classes.tableWrapper}>
                    <SortableTable
                        columns={partyInviteTableColumns}
                        data={convertInvitesToTableDatasource(pendingInvites)}
                        rowsPerPageOptions={[3,5,10,15]}
                        defaultRowsPerPage={5}
                        actionButtons={partyInviteActionButtons}
                    />
                </div>
            </div>

            <div className={classes.section}>
                <h2>Attended Parties</h2>
                <div className={classes.tableWrapper}>
                    <SortableTable
                        columns={partyTableColumns}
                        data={convertPartiesToTableDatasource(attendedParties)}
                        rowsPerPageOptions={[3,5,10,15]}
                        defaultRowsPerPage={5}
                        actionButtons={partyActionButtons}
                    />
                </div>
            </div>

            <div className={classes.section}>
                <h2>Organized Parties</h2>
                <div className={classes.tableWrapper}>
                    <SortableTable
                        columns={partyTableColumns}
                        data={convertPartiesToTableDatasource(organizedParties)}
                        rowsPerPageOptions={[3,5,10,15]}
                        defaultRowsPerPage={5}
                        actionButtons={partyActionButtons}
                    />
                </div>
            </div>
        </div>
    );
};
