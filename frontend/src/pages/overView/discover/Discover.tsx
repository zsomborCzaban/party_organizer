import { Button, Table } from 'antd';
import { CSSProperties, useEffect, useState } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { useNavigate } from 'react-router-dom';
import AccessCodeModal from './AccessCodeModal';
import { AppDispatch, RootState } from '../../../store/store';
import { User } from '../../../data/types/User';
import { getUser } from '../../../auth/AuthUserUtil';
import { authService } from '../../../auth/AuthService';
import { Party } from '../../../data/types/Party';
import { partyTableColumnsLegacy } from '../../../data/constants/TableColumns';
// import OverViewNavBar from '../../../components/navigation-bar/OverViewNavBar';
import OverViewProfile from '../../../components/drawer/OverViewProfile';
import { joinPublicParty } from '../../../api/apis/PartyAttendanceManagerApi';
import { loadPublicParties } from '../../../store/slices/PublicPartySlice';
import { setSelectedParty } from '../../../store/slices/PartySlice';

const styles: { [key: string]: CSSProperties } = {
  outerContainer: {
    height: '100vh',
    display: 'flex',
    flexDirection: 'column',
  },
  container: {
    width: '80%',
    margin: '0 auto',
    height: '100%',
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
    flexShrink: 0,
    padding: '20px',
    marginBottom: '20px',
    border: '1px solid #ccc',
    borderRadius: '8px',
  },
  message: {
    margin: '10px 0',
    fontSize: '18px',
    textAlign: 'left',
    color: '#555',
  },
  buttonsContainer: {
    display: 'flex',
    justifyContent: 'space-between',
    padding: '0 20px',
    flexGrow: 1,
  },
  createButton: {
    width: '48%',
    backgroundImage: `url(${'createImage'})`,
    backgroundSize: 'cover',
    backgroundPosition: 'center',
    color: 'white',
    fontSize: '48px',
    fontWeight: 'bold',
    textAlign: 'center',
    borderRadius: '10px',
    cursor: 'pointer',
    display: 'flex',
    justifyContent: 'center',
    alignItems: 'center',
    height: '90%',
    border: 'none',
    boxShadow: '0 4px 8px rgba(0, 0, 0, 0.2)',
  },
  joinButton: {
    width: '48%',
    backgroundImage: `url(${'joinImage'})`,
    backgroundSize: 'cover',
    backgroundPosition: 'center',
    color: 'white',
    fontSize: '48px',
    fontWeight: 'bold',
    textAlign: 'center',
    borderRadius: '10px',
    cursor: 'pointer',
    display: 'flex',
    justifyContent: 'center',
    alignItems: 'center',
    height: '90%',
    border: 'none',
    boxShadow: '0 4px 8px rgba(0, 0, 0, 0.2)',
  },
  errorFeedback: {
    margin: '10px 0',
    fontSize: '18px',
    textAlign: 'left',
    color: 'red',
  },
  button: {
    padding: '10px 20px',
    backgroundColor: '#007bff',
    color: 'white',
    border: 'none',
    borderRadius: '5px',
    cursor: 'pointer',
  },
};

const Discover = () => {
  const navigate = useNavigate();
  const dispatch = useDispatch<AppDispatch>();

  const [modalVisible, setModalVisible] = useState(false);
  const [profileOpen, setProfileOpen] = useState(false);
  const [unexpectedError, setUnexpectedError] = useState('');
  const [user, setUser] = useState<User>();

  const { parties, loading, error } = useSelector((state: RootState) => state.publicPartyStore);

  useEffect(() => {
    dispatch(loadPublicParties());
    const currentUser = getUser();

    if (!currentUser) {
      authService.userLoggedOut();
      return;
    }

    setUser(currentUser);
  }, [dispatch]);

  const handleVisitClicked = (record: Party) => {
    joinPublicParty(record.ID || -1)
      .then((joinedParty) => {
        dispatch(setSelectedParty(joinedParty));
        navigate('/visitParty/partyHome');
      })
      .catch(() => {
        console.log('something unexpected happened');
        setUnexpectedError('something unexpected happened while joining public party, try again later ');
      });
  };

  const handleCreate = () => {
    navigate('/createParty');
  };

  const handleJoin = () => {
    setModalVisible(true);
  };

  const closeModal = () => {
    setModalVisible(false);
  };

  const columns = [
    ...partyTableColumnsLegacy,
    {
      title: '',
      key: 'ID',
      // eslint-disable-next-line @typescript-eslint/no-explicit-any
      render: (_: any, record: Party) => (
        <Button
          style={styles.button}
          onClick={() => handleVisitClicked(record)}
        >
          Visit
        </Button>
      ),
    },
  ];

  const renderPublicParties = () => {
    if (!parties || parties.length === 0) {
      return <div>There&#39;s no public parties at the moment :( </div>;
    }
    return (
      <Table
        dataSource={parties.map((party) => ({ ...party, key: party.ID }))}
        columns={columns}
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
      {/*<OverViewNavBar onProfileClick={() => setProfileOpen(true)} />*/}
      <OverViewProfile
        isOpen={profileOpen}
        onClose={() => setProfileOpen(false)}
        user={user}
      />
      <div style={styles.container}>
        <h2 style={styles.label}>Public Parties</h2>

        <div style={styles.tableContainer}>
          {loading && <div>Loading...</div>}
          {error && <div>Error: Some unexpected error happened</div>}
          {!loading && !error && renderPublicParties()}
        </div>

        <div style={styles.message}>Didn&#39;t find the right party? Don&#39;t worry, you can create your own!</div>

        <div style={styles.buttonsContainer}>
          <div
            style={styles.createButton}
            onClick={handleCreate}
          >
            Create
          </div>
          <div
            style={styles.joinButton}
            onClick={handleJoin}
          >
            Join
          </div>
        </div>
        <AccessCodeModal
          visible={modalVisible}
          onClose={closeModal}
        />
        {unexpectedError && <div style={styles.errorFeedback}>{unexpectedError}</div>}
      </div>
    </div>
  );
};

export default Discover;
