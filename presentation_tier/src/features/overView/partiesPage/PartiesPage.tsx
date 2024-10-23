import {useDispatch, useSelector} from "react-redux";
import {AppDispatch, RootState} from "../../../store/store";
import {CSSProperties, useEffect, useState} from "react";
import OverViewNavBar from "../../../components/navbar/OverViewNavBar";
import { Table } from 'antd';
import {useNavigate} from "react-router-dom";
import {Party} from "../Party";
import {loadOrganizedParties} from "./OrganizedPartySlice";
import {loadAttendedParties} from "./AttendedPartySlice";
import {loadPartyInvites} from "./PartyInviteSlice";
import {PartyInvite} from "./PartyInvite";
import {acceptInvite, declineInvite} from "./PartiesPageApi";


const PartiesPage = () => {
    const [reload, setReload] = useState(false);

    const dispatch = useDispatch<AppDispatch>()

    // the words after the ':' are not types but new names here
    const {parties: organizedParties, loading: organizedLoading, error: organizedError} = useSelector(
        (state: RootState) => state.organizedPartyStore
    )
    const {parties: attendedParties, loading: attendedLoading, error: attendedError} = useSelector(
        (state: RootState) => state.attendedPartyStore
    )
    const {invites, loading: inviteLoading, error: inviteError} = useSelector(
        (state: RootState) => state.partyInviteStore
    )

    useEffect( () => {
        dispatch(loadOrganizedParties());
    }, []);

    useEffect( () => {
        dispatch(loadAttendedParties());
        dispatch(loadPartyInvites());
    }, [reload]);

    const navigate = useNavigate()

    const handleVisitClicked = (record: Party) => {
        console.log(record)
        //todo: set selected party to record and navigate to the parties page
    }

    const handleInviteAccepted = (record: PartyInvite) => {
        acceptInvite(record.party.ID)
            .then(() => {setReload(prev => !prev)} )
            .catch(err => { //todo: handle err on the userinterface too
                console.log("error while accepting invite: " + err)
            });
    }

    const handleInviteDeclined = (record: PartyInvite) => {
        declineInvite(record.party.ID)
            .then(() => {setReload(prev => !prev)} )
            .catch(err => { //todo: handle err on the userinterface too
                console.log("error while accepting invite: " + err)
            });
    }

    const partyColumns = [
        {
            title: 'Name',
            dataIndex: 'name',
            key: 'name',
        },
        {
            title: 'Place',
            dataIndex: 'place',
            key: 'place',
        },
        {
            title: 'Time',
            dataIndex: 'start_time',
            key: 'time',
        },
        {
            title: 'Organizer',
            // dataIndex: ['organizer', 'username'],
            dataIndex: ['organizer', 'username'],
            key: 'organizer',
        },
        {
            //todo: to be done in backend
            title: 'Headcount',
            dataIndex: 'headcount',
            key: 'headcount',
        },
        {
            title: '',
            key: 'ID',
            render: (record: Party) => (
                <button onClick={() => handleVisitClicked(record)}>Visit</button>
            ),
        },
    ];

    const inviteColumns = [
        {
            title: 'Invited by',
            dataIndex: ['invitor', 'username'],
            key: 'invited by',
        },
        {
            title: 'To party',
            dataIndex: ['party', 'name'],
            key: 'to party',
        },
        {
            title: 'Place',
            dataIndex: ['party', 'place'],
            key: 'place',
        },
        {
            title: 'Time',
            dataIndex: ['party', 'start_time'],
            key: 'time',
        },

        {
            //todo: to be done in backend
            title: 'Headcount',
            dataIndex: ['party', 'headcount'],
            key: 'headcount',
        },
        {
            title: '',
            key: 'action 1',
            render: (record: PartyInvite) => (
                <button onClick={() => handleInviteAccepted(record)}>Accept</button>
            ),
        },
        {
            title: '',
            key: 'action 2',
            render: (record: PartyInvite) => (
            <button onClick={() => handleInviteDeclined(record)}>Decline</button>
        ),
    },
    ];

    const renderParties = (type: string) => {
        let parties: Party[] = []
        if(type === "organized") parties = organizedParties
        if(type === "attended") parties = attendedParties

        if(!parties || parties.length === 0){
            return <div>There's no {type} parties at the moment :( </div>
        }
        return (<Table
            dataSource={parties.map(party => ({...party, key: party.ID}))}
            columns={partyColumns}
            pagination={false} // Disable pagination
            scroll={{y: 200}} // Set vertical scroll height
        />)
    }

    const renderInvites = () => {
        if(!invites || invites.length === 0){
            return <div>There's no invites at the moment :( </div>
        }
        return (<Table
            dataSource={invites.map(invite => ({...invite, key: invite.party.ID}))}
            columns={inviteColumns}
            pagination={false} // Disable pagination
            scroll={{y: 200}} // Set vertical scroll height
        />)
    }

    return (
        <div style={styles.outerContainer}>
            <OverViewNavBar/>
            <div style={styles.container}>

                <h2 style={styles.label}>Party Invites</h2>

                {/* Scrollable Table using Ant Design Table */}
                <div style={styles.tableContainer}>
                    {inviteLoading && <div>Loading...</div>}
                    {inviteError && <div>Error: Some unexpected error happened</div>}
                    {(!inviteLoading && !inviteError) && renderInvites()}
                </div>

                <h2 style={styles.label}>Attended Parties</h2>

                {/* Scrollable Table using Ant Design Table */}
                <div style={styles.tableContainer}>
                    {attendedLoading && <div>Loading...</div>}
                    {attendedError && <div>Error: Some unexpected error happened</div>}
                    {(!attendedLoading && !attendedError) && renderParties("attended")}
                </div>

                <h2 style={styles.label}>Organized Parties</h2>

                {/* Scrollable Table using Ant Design Table */}
                <div style={styles.tableContainer}>
                    {organizedLoading && <div>Loading...</div>}
                    {organizedError && <div>Error: Some unexpected error happened</div>}
                    {(!organizedLoading && !organizedError) && renderParties("organized")}
                </div>
            </div>
        </div>
    );
};

const styles: { [key: string]: CSSProperties } = {
    outerContainer: {
        height: '100vh', // Full viewport height
        display: 'flex',
        flexDirection: 'column',
    },
    container: {
        width: '80%',
        margin: '0 auto',
        height: '100%', // Occupy the remaining height
        display: 'flex',
        flexDirection: 'column',
    },
    label: {
        margin: '10px',
        fontSize: '24px',
        fontWeight: 'bold',
        textAlign: 'left',
    },
    tableContainer: {
        flexShrink: 0, // Prevent the table container from shrinking
        padding: '20px',
        marginBottom: '20px',
        border: '1px solid #ccc',
        borderRadius: '8px',
    },
    message: {
        margin: '10px 0', // Space above and below the message
        fontSize: '18px',
        textAlign: 'left', // Center align the message
        color: '#555', // Optional: a softer color for the message
    },
    buttonsContainer: {
        display: 'flex',
        justifyContent: 'space-between',
        padding: '0 20px',
        flexGrow: 1, // Allow the buttons container to grow
    },
};


export default PartiesPage;