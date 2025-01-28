import React, { CSSProperties, useEffect, useState } from 'react';
import { Button, Table } from 'antd';
import { AppDispatch, RootState } from '../../../store/store';
import { useDispatch, useSelector } from 'react-redux';
import { User } from '../../../data/types/User';
import { getUser } from '../../../auth/AuthUserUtil';
import { authService } from '../../../auth/AuthService';
import { setForTime } from '../../../data/utils/timeoutSetterUtils';
import { FriendInvite } from '../../../data/types/FriendInvite';
import { invitedByTableColumns, userTableColumns } from '../../../data/constants/TableColumns';
import OverViewNavBar from '../../../components/navigation-bar/OverViewNavBar';
import OverViewProfile from '../../../components/drawer/OverViewProfile';
import { acceptInvite, declineInvite, inviteFriend, removeFriend } from '../../../api/apis/FriendInviteManagerApi';
import { loadFriendInvites } from '../../../store/sclices/FriendInviteSlice';
import { loadFriends } from '../../../store/sclices/FriendSlice';

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
    alignItems: 'flex-start', // Center items horizontally
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
  friendsTableContainer: {
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
  tableButton: {
    padding: '7px 15px',
    backgroundColor: '#007bff',
    color: 'white',
    border: 'none',
    borderRadius: '5px',
    cursor: 'pointer',
  },
  acceptButton: {
    padding: '7px 15px',
    backgroundColor: 'green',
    color: 'white',
    border: 'none',
    borderRadius: '5px',
    cursor: 'pointer',
  },
  declineButton: {
    padding: '7px 15px',
    backgroundColor: 'red',
    color: 'white',
    border: 'none',
    borderRadius: '5px',
    cursor: 'pointer',
  },
};

const Friends: React.FC = () => {
  const dispatch = useDispatch<AppDispatch>();

  const [reloadInvites, setReloadInvites] = useState(false);
  const [reloadFriends, setReloadFriends] = useState(false);
  const [usernameInput, setUsernameInput] = useState('');
  const [inviteFeedbackSuccess, setInviteFeedbackSuccess] = useState('');
  const [inviteFeedbackError, setInviteFeedbackError] = useState('');
  const [profileOpen, setProfileOpen] = useState(false);
  const [user, setUser] = useState<User>();

  const { friends, loading: friendLoading, error: friendError } = useSelector((state: RootState) => state.friendStore);
  const { invites, loading: inviteLoading, error: inviteError } = useSelector((state: RootState) => state.friendInviteStore);

  useEffect(() => {
    const currentUser = getUser();

    if (!currentUser) {
      authService.handleUnauthorized();
      return;
    }

    setUser(currentUser);
  }, []);

  useEffect(() => {
    dispatch(loadFriendInvites());
  }, [dispatch, reloadInvites]);

  useEffect(() => {
    dispatch(loadFriends());
  }, [dispatch, reloadFriends]);

  const handleInviteFriend = (inputUsername: string) => {
    inviteFriend(inputUsername)
      .then(() => {
        setForTime<string>(setInviteFeedbackSuccess, 'Invite sent!', '', 4000);
      })
      .catch(() => {
        setForTime<string>(setInviteFeedbackError, 'Something went wrong!', '', 4000);
      });
  };

  const handleInviteAccepted = (record: FriendInvite) => {
    acceptInvite(record.invitor.ID)
      .then(() => {
        setReloadInvites((prev) => !prev);
        setReloadFriends((prev) => !prev);
      })
      .catch((err) => {
        // todo: handle err on the userinterface too
        console.log(`error while accepting invite: ${err}`);
      });
  };

  const handleInviteDeclined = (record: FriendInvite) => {
    declineInvite(record.invitor.ID)
      .then(() => {
        setReloadInvites((prev) => !prev);
      })
      .catch((err) => {
        // todo: handle err on the userinterface too
        console.log(`error while accepting invite: ${err}`);
      });
  };

  const handleRemoveFriend = (record: User) => {
    removeFriend(record.ID)
      .then(() => {
        setReloadFriends((prev) => !prev);
      })
      .catch((err) => {
        // todo: handle err on the userinterface too
        console.log(`error while accepting invite: ${err}`);
      });
  };

  const inviteColumns = [
    ...invitedByTableColumns,
    {
      title: '',
      key: 'accept',
      render: (record: FriendInvite) => (
        <button
          style={styles.acceptButton}
          onClick={() => handleInviteAccepted(record)}
        >
          Accept
        </button>
      ),
    },
    {
      title: '',
      key: 'decline',
      render: (record: FriendInvite) => (
        <button
          style={styles.declineButton}
          onClick={() => handleInviteDeclined(record)}
        >
          Decline
        </button>
      ),
    },
  ];

  const friendColumns = [
    ...userTableColumns,
    {
      title: 'Username',
      dataIndex: 'username',
      key: 'username',
    },
    {
      key: 'remove',
      render: (record: User) => (
        <button
          style={styles.tableButton}
          onClick={() => handleRemoveFriend(record)}
        >
          Remove
        </button>
      ),
    },
  ];

  const renderInvites = () => {
    if (!invites || invites.length === 0) {
      return <div>There&#39;s no pending invite at the moment</div>;
    }
    return (
      <Table
        dataSource={invites.map((invite) => ({ ...invite, key: invite.ID }))}
        columns={inviteColumns}
        pagination={false}
        showHeader={false}
        bordered={false}
      />
    );
  };

  const renderFriends = () => {
    if (!friends || friends.length === 0) {
      return <div>You have no friends yet!</div>;
    }
    return (
      <Table
        dataSource={friends.map((friend) => ({ ...friend, key: friend.ID }))}
        columns={friendColumns}
        pagination={false}
        scroll={{ y: 200 }}
      />
    );
  };

  if (!user) {
    console.log('user was null');
    return <div>Loading...</div>;
  }

  return (
    <div style={styles.outerContainer}>
      <OverViewNavBar onProfileClick={() => setProfileOpen(true)} />
      <OverViewProfile
        isOpen={profileOpen}
        onClose={() => setProfileOpen(false)}
        user={user}
      />
      <div style={styles.container}>
        {/* Title and input section */}
        <h2 style={styles.label}>Invite Friend</h2>

        <div style={styles.inputContainer}>
          <input
            type='text'
            id='username'
            value={usernameInput}
            placeholder='Enter username'
            onChange={(e) => setUsernameInput(e.target.value)}
            style={styles.input}
          />
          <Button
            type='primary'
            style={styles.button}
            onClick={() => handleInviteFriend(usernameInput)}
          >
            Invite
          </Button>
          {inviteFeedbackSuccess && <p style={styles.success}>{inviteFeedbackSuccess}</p>}
          {inviteFeedbackError && <p style={styles.error}>{inviteFeedbackError}</p>}
        </div>

        {/* Friend Invites Table */}
        <h2 style={styles.label}>Pending Invites</h2>

        <div style={styles.invitesTableContainer}>
          {inviteLoading && <div>Loading...</div>}
          {inviteError && <div>Error: Some unexpected error happened</div>}
          {!inviteLoading && !inviteError && renderInvites()}
        </div>

        {/* Friends Table */}
        <h2 style={styles.label}>Friends</h2>

        <div style={styles.friendsTableContainer}>
          {friendLoading && <div>Loading...</div>}
          {friendError && <div>Error: Some unexpected error happened</div>}
          {!friendLoading && !friendError && renderFriends()}
        </div>
      </div>
    </div>
  );
};

export default Friends;
