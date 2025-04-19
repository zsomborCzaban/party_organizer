import {useApi} from "../../context/ApiContext.ts";
import {useEffect, useState} from "react";
import {toast} from "sonner";
import {FriendInvite} from "../../data/types/FriendInvite.ts";
import {User} from "../../data/types/User.ts";
import {ActionButton, SortableTable} from "../../components/table/SortableTable.tsx";
import styles from './Friends.module.scss';
import {Button, TextField} from "@mui/material";
import {
    FriendInviteTableColumns,
    FriendInviteTableRow,
    FriendTableColumns,
    FriendTableRow
} from "../../data/constants/TableColumns.tsx";
import classes from "../party/parties/Parties.module.scss";
import {convertFriendsToTableData, convertInvitesToTableData} from "../../data/utils/TableUtils.ts";

export const Friends = () => {
    const api = useApi();
    const [pendingInvites, setPendingInvites] = useState<FriendInvite[]>([]);
    const [friends, setFriends] = useState<User[]>([]);
    const [reloadInvites, setReloadInvites] = useState(0);
    const [reloadFriends, setReloadFriends] = useState(0);
    const [inviteUsername, setInviteUsername] = useState('');

    useEffect(() => {
        api.friendManagerApi.getPendingInvites().then(result => {
            if (result === 'error') {
                toast.error('Error while loading invites');
                setPendingInvites([]);
                return;
            }
            setPendingInvites(result.data);
        }).catch(() => {
            toast.error('Error while loading invites');
            setPendingInvites([]);
        });
    }, [api.friendManagerApi, reloadInvites]);

    useEffect(() => {
        api.userApi.getFriends().then(result => {
            if (result === 'error') {
                toast.error('Error while loading friends');
                setFriends([]);
                return;
            }
            setFriends(result.data);
        }).catch(() => {
            toast.error('Error while loading friends');
            setFriends([]);
        });
    }, [api.userApi, reloadFriends]);

    const handleInviteFriend = () => {
        if (!inviteUsername.trim()) {
            toast.error('Please enter a username');
            return;
        }

        api.friendManagerApi.inviteFriend(inviteUsername).then(result => {
            if (result === 'error') {
                return;
            }
            toast.success('Invite sent');
            setInviteUsername('');
        }).catch(() => {
            toast.error('Unexpected error');
        });
    };

    const inviteActionButtons: ActionButton<FriendInviteTableRow>[] = [
        {
            label: 'Accept',
            color: 'success',
            onClick: (row) => {
                api.friendManagerApi.acceptInvite(row.id).then(result => {
                    if (result === 'error') {
                        return;
                    }
                    setReloadInvites((prev) => (prev + 1)%2);
                    setReloadFriends((prev) => (prev + 1)%2);
                }).catch(() => {
                    toast.error('Unexpected error');
                });
            }
        },
        {
            label: 'Decline',
            color: 'error',
            onClick: (row) => {
                api.friendManagerApi.declineInvite(row.id).then(result => {
                    if (result === 'error') {
                        return;
                    }
                    setReloadInvites((prev) => (prev + 1)%2);
                }).catch(() => {
                    toast.error('Unexpected error');
                });
            }
        }
    ];

    const friendActionButtons: ActionButton<FriendTableRow>[] = [
        {
            label: 'Remove',
            color: 'error',
            onClick: (row) => {
                api.friendManagerApi.removeFriend(row.id).then(result => {
                    if (result === 'error') {
                        toast.error('Unexpected error');
                        return;
                    }
                    setReloadFriends((prev) => prev + 1);
                }).catch(() => {
                    toast.error('Unexpected error');
                });
            }
        }
    ];

    return (
        <div className={styles.container}>
            <div className={classes.header}>
                <h1>My Friends</h1>
                <p className={classes.description}>
                    Connect or disconnect with your friends here
                </p>
            </div>

            <div className={styles.section}>
                <h2>Invite Friends</h2>
                <div className={styles.inviteForm}>
                    <TextField
                        value={inviteUsername}
                        onChange={(e) => setInviteUsername(e.target.value)}
                        placeholder="Enter username"
                        variant="outlined"
                        size="small"
                    />
                    <Button
                        variant="contained"
                        color="primary"
                        onClick={handleInviteFriend}
                    >
                        Send Invite
                    </Button>
                </div>
            </div>

            <div className={styles.section}>
                <h2>Pending Invites</h2>
                <div className={classes.tableWrapper}>
                    <SortableTable
                        columns={FriendInviteTableColumns}
                        data={convertInvitesToTableData(pendingInvites)}
                        actionButtons={inviteActionButtons}
                        rowsPerPageOptions={[5, 10, 15]}
                        defaultRowsPerPage={5}
                    />
                </div>
            </div>

                <div className={styles.section}>
                    <h2>My Friends</h2>
                    <div className={classes.tableWrapper}>
                        <SortableTable
                            columns={FriendTableColumns}
                            data={convertFriendsToTableData(friends)}
                            actionButtons={friendActionButtons}
                            rowsPerPageOptions={[5, 10, 15]}
                            defaultRowsPerPage={5}
                        />
                    </div>
                </div>
            </div>
    );
};