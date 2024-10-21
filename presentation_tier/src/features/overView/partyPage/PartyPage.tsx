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
import {acceptInvite} from "./PartyPageApi";


const PartyPage = () => {
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
        //record is good everything is shonw in it, but record.party is undefined

        console.log(record); // allgood
        console.log(record.Party) //undefined
        // console.log(record.Party.ID)
        // acceptInvite(record.Party.ID)
        //     .then(() => {setReload(prev => !prev)} )
        //     .catch(err => { //todo: handle err on the userinterface too
        //         console.log("error while accepting invite: " + err)
        //     });
    }

    const handleInviteDeclined = (record: PartyInvite) => {
        console.log(record)
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
            render: (text: string, record: Party) => (
                <button onClick={() => handleVisitClicked(record)}>Visit</button>
            ),
        },
    ];

    const inviteColumns = [
        {
            title: 'Invited by',
            dataIndex: ['invitor', 'username'],
            key: 'ID',
        },
        {
            title: 'To party',
            dataIndex: ['party', 'name'],
            key: 'ID',
        },
        {
            title: 'Place',
            dataIndex: ['party', 'place'],
            key: 'ID',
        },
        {
            title: 'Time',
            dataIndex: ['party', 'start_time'],
            key: 'ID',
        },

        {
            //todo: to be done in backend
            title: 'Headcount',
            dataIndex: ['party', 'headcount'],
            key: 'ID',
        },
        {
            title: '',
            key: 'ID',
            render: (text: string, record: PartyInvite) => (
                <button onClick={() => handleInviteAccepted(record)}>Accept</button>
            ),
        },
        {
            title: '',
            key: 'ID',
            render: (text: string, record: PartyInvite) => (
            <button onClick={() => handleInviteDeclined(record)}>Decline</button>
        ),
    },
    ];

    console.log(invites)

    const renderParties = (type: string) => {
        let parties: Party[] = []
        if(type === "organized") parties = organizedParties
        if(type === "attended") parties = attendedParties

        if(parties.length === 0){
            return <div>There's no {type} parties at the moment :( </div>
        }
        return (<Table
            dataSource={parties}
            columns={partyColumns}
            pagination={false} // Disable pagination
            scroll={{y: 200}} // Set vertical scroll height
        />)
    }

    const renderInvites = () => {
        if(invites.length === 0){
            return <div>There's no invites at the moment :( </div>
        }
        return (<Table
            dataSource={invites}
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
    )
}

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


export default PartyPage;