import { useDispatch, useSelector } from 'react-redux';
import { CSSProperties, useEffect, useState } from 'react';
import { Table } from 'antd';
import { useNavigate } from 'react-router-dom';
import { AppDispatch, RootState } from '../../../store/store';
import { User } from '../../../data/types/User';
import { getUser } from '../../../auth/AuthUserUtil';
import { authService } from '../../../auth/AuthService';
import { Party } from '../../../data/types/Party';
import { PartyInvite } from '../../../data/types/PartyInvite';
import { partyInviteTableColumns, partyTableColumns } from '../../../data/constants/TableColumns';
import OverViewNavBar from '../../../components/navigation-bar/OverViewNavBar';
import OverViewProfile from '../../../components/drawer/OverViewProfile';
import { acceptInvite, declineInvite } from '../../../api/apis/PartyAttendanceManagerApi';
import { loadOrganizedParties } from '../../../store/sclices/OrganizedPartySlice';
import { loadAttendedParties } from '../../../store/sclices/AttendedPartySlice';
import { loadPartyInvites } from '../../../store/sclices/PartyInviteSlice';
import { setSelectedParty } from '../../../store/sclices/PartySlice';

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

const PartiesPage = () => {
  const dispatch = useDispatch<AppDispatch>();
  const navigate = useNavigate();

  const [reload, setReload] = useState(false);
  const [profileOpen, setProfileOpen] = useState(false);
  const [user, setUser] = useState<User>();

  // the words after the ':' are not types but new names here
  const { parties: organizedParties, loading: organizedLoading, error: organizedError } = useSelector((state: RootState) => state.organizedPartyStore);
  const { parties: attendedParties, loading: attendedLoading, error: attendedError } = useSelector((state: RootState) => state.attendedPartyStore);
  const { invites, loading: inviteLoading, error: inviteError } = useSelector((state: RootState) => state.partyInviteStore);

  useEffect(() => {
    dispatch(loadOrganizedParties());
    const currentUser = getUser();

    if (!currentUser) {
      authService.handleUnauthorized();
      return;
    }

    setUser(currentUser);
  }, [dispatch]);

  useEffect(() => {
    dispatch(loadAttendedParties());
    dispatch(loadPartyInvites());
  }, [dispatch, reload]);

  const handleVisitClicked = (record: Party) => {
    dispatch(setSelectedParty(record));
    navigate('/visitParty/partyHome');
  };

  const handleInviteAccepted = (record: PartyInvite) => {
    if (!record.party.ID) return; // todo: handle error message
    acceptInvite(record.party.ID)
      .then(() => {
        setReload((prev) => !prev);
      })
      .catch((err) => {
        // todo: handle err on the userinterface too
        console.log(`error while accepting invite: ${err}`);
      });
  };

  const handleInviteDeclined = (record: PartyInvite) => {
    if (!record.party.ID) return; // todo: handle error message
    declineInvite(record.party.ID)
      .then(() => {
        setReload((prev) => !prev);
      })
      .catch((err) => {
        // todo: handle err on the userinterface too
        console.log(`error while accepting invite: ${err}`);
      });
  };

  const partyColumns = [
    ...partyTableColumns,
    {
      title: '',
      key: 'ID',
      render: (record: Party) => (
        <button
          style={styles.tableButton}
          onClick={() => handleVisitClicked(record)}
        >
          Visit
        </button>
      ),
    },
  ];

  const inviteColumns = [
    ...partyInviteTableColumns,
    {
      title: '',
      key: 'action 1',
      render: (record: PartyInvite) => (
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
      key: 'action 2',
      render: (record: PartyInvite) => (
        <button
          style={styles.declineButton}
          onClick={() => handleInviteDeclined(record)}
        >
          Decline
        </button>
      ),
    },
  ];

  const renderParties = (type: string) => {
    let parties: Party[] = [];
    if (type === 'organized') parties = organizedParties;
    if (type === 'attended') parties = attendedParties;

    if (!parties || parties.length === 0) {
      return <div>There&#39;s no {type} parties at the moment :( </div>;
    }
    return (
      <Table
        dataSource={parties.map((party) => ({ ...party, key: party.ID }))}
        columns={partyColumns}
        pagination={false} // Disable pagination
        scroll={{ y: 200 }}
      />
    );
  };

  const renderInvites = () => {
    if (!invites || invites.length === 0) {
      return <div>There&#39;s no invites at the moment :( </div>;
    }
    return (
      <Table
        dataSource={invites.map((invite) => ({ ...invite, key: invite.party.ID }))}
        columns={inviteColumns}
        pagination={false} // Disable pagination
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
        <h2 style={styles.label}>Party Invites</h2>

        {/* Scrollable Table using Ant Design Table */}
        <div style={styles.tableContainer}>
          {inviteLoading && <div>Loading...</div>}
          {inviteError && <div>Error: Some unexpected error happened</div>}
          {!inviteLoading && !inviteError && renderInvites()}
        </div>

        <h2 style={styles.label}>Attended Parties</h2>

        {/* Scrollable Table using Ant Design Table */}
        <div style={styles.tableContainer}>
          {attendedLoading && <div>Loading...</div>}
          {attendedError && <div>Error: Some unexpected error happened</div>}
          {!attendedLoading && !attendedError && renderParties('attended')}
        </div>

        <h2 style={styles.label}>Organized Parties</h2>

        {/* Scrollable Table using Ant Design Table */}
        <div style={styles.tableContainer}>
          {organizedLoading && <div>Loading...</div>}
          {organizedError && <div>Error: Some unexpected error happened</div>}
          {!organizedLoading && !organizedError && renderParties('organized')}
        </div>
      </div>
    </div>
  );
};

export default PartiesPage;
