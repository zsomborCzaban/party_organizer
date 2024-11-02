import React, {CSSProperties, useEffect, useState} from 'react';
import { Button, Table } from 'antd';
import OverViewNavBar from "../../../components/navbar/OverViewNavBar";
import {useDispatch, useSelector} from "react-redux";
import {AppDispatch, RootState} from "../../../store/store";
import {loadFriendInvites} from "./FriendInviteSlice";
import {loadFriends} from "./FriendSlice";
import {acceptInvite, declineInvite, inviteFriend, removeFriend} from "./FriendPageApi";
import {FriendInvite} from "./FriendInvite";
import {User} from "../User";

const Friends: React.FC = () => {
    const [reloadInvites, setReloadInvites] = useState(false);
    const [reloadFriends, setReloadFriends] = useState(false);
    const [username, setUsername] = useState('');
    const [inviteFeedbackSuccess, setInviteFeedbackSuccess] = useState('')
    const [inviteFeedbackError, setInviteFeedbackError] = useState('')

    const dispatch = useDispatch<AppDispatch>()

    const {friends, loading: friendLoading, error: friendError} = useSelector(
        (state: RootState) => state.friendStore
    )
    const {invites, loading: inviteLoading, error: inviteError} = useSelector(
        (state: RootState) => state.friendInviteStore
    )

    useEffect( () => {
        dispatch(loadFriendInvites());
    }, [reloadInvites]);

    useEffect( () => {
        dispatch(loadFriends());
    }, [reloadFriends]);

    const handleInviteFriend = (inputUsername: string) => {
        inviteFriend(inputUsername)
            .then(() => {
                setInviteFeedbackSuccess("Invite sent!")
                setUsername('')
                setTimeout(() => {
                    setInviteFeedbackSuccess("")
                }, 4000); // 4000 milliseconds = 4 seconds
            })
            .catch(err => {
                setInviteFeedbackError("something went wrong")
                setUsername('')
                setTimeout(() => {
                    setInviteFeedbackError("")
                }, 4000);
            });
    }

    const handleInviteAccepted = (record: FriendInvite) => {
        acceptInvite(record.invitor.ID)
            .then(() => {
                setReloadInvites(prev => !prev)
                setReloadFriends(prev => !prev)
            })
            .catch(err => { //todo: handle err on the userinterface too
                console.log("error while accepting invite: " + err)
            });
    }

    const handleInviteDeclined = (record: FriendInvite) => {
        declineInvite(record.invitor.ID)
            .then(() => {
                setReloadInvites(prev => !prev)
            })
            .catch(err => { //todo: handle err on the userinterface too
                console.log("error while accepting invite: " + err)
            });
    }

    const handleRemoveFriend = (record: User) => {
        removeFriend(record.ID)
            .then(() => {
                setReloadFriends(prev => !prev)
            })
            .catch(err => { //todo: handle err on the userinterface too
                console.log("error while accepting invite: " + err)
            });
    }


    const inviteColumns = [
        {
            title: 'Invited by',
            dataIndex: ['invitor', 'username'],
            key: 'username',
        },
        {
            title: '',
            key: 'accept',
            render: (record: FriendInvite) => (
                <button onClick={() => handleInviteAccepted(record)}>Accept</button>
            ),
        },
        {
            title: '',
            key: 'decline',
            render: (record: FriendInvite) => (
                <button onClick={() => handleInviteDeclined(record)}>Decline</button>
            ),
        },
    ];

    const friendColumns = [
        {
            title: 'Username',
            dataIndex: 'username',
            key: 'username'
        },
        {
            key: 'remove',
            render: (record: User) => (
                <button onClick={() => handleRemoveFriend(record)}>Remove</button>
            ),
        },
    ];


    const renderInvites = () => {
        if(!invites || invites.length === 0){
            return <div>There's no pending invite at the moment</div>
        }
        return (<Table
            dataSource={invites.map(invite => ({...invite, key: invite.ID}))}
            columns={inviteColumns}
            pagination={false}
            showHeader={false}
            bordered={false}
        />)
    }

    const renderFriends = () => {
        if(!friends || friends.length === 0){
            return <div>You have no friends yet!</div>
        }
        return (<Table
            dataSource={friends.map(friends => ({...friends, key: friends.ID}))}
            columns={friendColumns}
            pagination={false}
            scroll={{y: 200}}
        />)
    }

    return (
        <div style={styles.outerContainer}>
            <OverViewNavBar onProfileClick={() => {console.log('a')}}/>
            <div style={styles.container}>
                {/* Title and input section */}
                <h2 style={styles.label}>Invite Friend</h2>

                <div style={styles.inputContainer}>
                    <input
                        type="text"
                        id="username"
                        value={username}
                        placeholder="Enter username"
                        onChange={(e) => setUsername(e.target.value)}
                        style={styles.input}
                    />
                    <Button type="primary" style={styles.button} onClick={() => handleInviteFriend(username)}>Invite</Button>
                    {inviteFeedbackSuccess && <p style={styles.success}>{inviteFeedbackSuccess}</p>}
                    {inviteFeedbackError && <p style={styles.error}>{inviteFeedbackError}</p>}
                </div>

                {/* Friend Invites Table */}
                <h2 style={styles.label}>Pending Invites</h2>

                <div style={styles.invitesTableContainer}>
                    {inviteLoading && <div>Loading...</div>}
                    {inviteError && <div>Error: Some unexpected error happened</div>}
                    {(!inviteLoading && !inviteError) && renderInvites()}
                </div>

                {/* Friends Table */}
                <h2 style={styles.label}>Friends</h2>

                <div style={styles.friendsTableContainer}>
                    {friendLoading && <div>Loading...</div>}
                    {friendError && <div>Error: Some unexpected error happened</div>}
                    {(!friendLoading && !friendError) && renderFriends()}
                </div>
            </div>
        </div>
    );
};

// Styles object
const styles: { [key: string]: CSSProperties } = {
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
    inputContainer: {
        display: 'flex',
        flexDirection: 'column',
        alignItems: 'flex-start',    // Center items horizontally
        padding: '20px',
        borderRadius: '10px',
        maxWidth: '400px',
    },
    input: {
        width: 'auto', // Let the input size be determined by its content (if required)
        minWidth: '300px', // Ensure a minimum width for good appearance
        marginRight: '10px',
        padding: '10px',
        borderRadius: '5px',
        border: '1px solid #d9d9d9',
        marginBottom: '10px',
    },
    button: {
        width: 'auto', // Button will take as much space as it needs
        minWidth: '120px',
        padding: '10px 20px',
        borderRadius: '5px',
        marginBottom: '20px',
    },
    invitesTableContainer: {
        display: 'flex', // Make this a flex container
        flexShrink: 0, // Prevent the table container from shrinking
        justifyContent: 'flex-start', // Align the table to the left
        marginBottom: '30px',
    },
    friendsTableContainer:{
        display: 'flex', // Make this a flex container
        justifyContent: 'flex-start', // Align the table to the left
        flexShrink: 0, // Prevent the table container from shrinking
        padding: '20px',
        marginBottom: '20px',
        border: '1px solid #ccc',
        borderRadius: '8px',
    },
    error: {
        color: 'red',
        fontSize: '0.875em',
    },
    success: {
        color: 'green',
        fontSize: '0.875em',
    },
};

export default Friends;
