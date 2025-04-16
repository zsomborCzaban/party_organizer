import { Button, ConfigProvider, Table, theme } from 'antd';
import { CSSProperties, useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import CreateRequirementModal from './CreateRequirementModal';
import DeleteRequirementModal from './DeleteRequirementModal';
import { User } from '../../../data/types/User';
import { setForTime } from '../../../data/utils/timeoutSetterUtils';
import {Requirement, RequirementPopulated} from '../../../data/types/Requirement';
import { invitedTableColumnsLegacy, requirementTableColumnsLegacy, userTableColumnsLegacy } from '../../../data/constants/TableColumns';
import { inviteToParty, kickFromParty } from '../../../api/apis/PartyAttendanceManagerApi';
import {EMPTY_PARTY_POPULATED, PartyPopulated} from "../../../data/types/Party.ts";
import {PartyInvite} from "../../../data/types/PartyInvite.ts";
import {useApi} from "../../../context/ApiContext.ts";
import {toast} from "sonner";

const styles: { [key: string]: CSSProperties } = {
  background: {
    backgroundImage: `url(${'backgroundImage'})`,
    position: 'fixed',
    backgroundSize: 'cover',
    backgroundPosition: 'center',
    display: 'flex',
  },
  outerContainer: {
    overflowY: 'auto',
    height: '100vh',
    width: '100vw',
    display: 'flex',
    flexDirection: 'column',
    color: '#ffffff',
  },
  container: {
    width: 'min(80%, 1000px)',
    margin: '20px auto',
    padding: '20px',
    display: 'flex',
    flexDirection: 'column',
    // backgroundColor: "#2c2c2c", // Darker gray background for content box
    backgroundColor: 'rgba(33, 33, 33, 0.95)',
    borderRadius: '8px',
    boxShadow: '0 4px 8px rgba(0, 0, 0, 0.4)', // Slightly stronger shadow for depth
    color: '#007bff', // Ensure text is white for readability
  },
  h2: {
    color: '#d3d3d3', // Light gray for headings
    fontSize: '1.8rem',
    fontWeight: 'bold',
    textAlign: 'left',
  },
  inputContainer: {
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'flex-start',
    marginBottom: '20px',
    gap: '10px',
  },
  input: {
    padding: '8px 12px',
    fontSize: '1rem',
    borderRadius: '5px',
    border: '1px solid #444', // Darker border to blend with dark mode
    backgroundColor: '#3a3a3a', // Dark input background
    color: '#ffffff', // Light input text
    width: '60%',
  },
  button: {
    width: 'auto',
    minWidth: '120px',
    padding: '10px 20px',
    borderRadius: '5px',
    marginBottom: '20px',
    fontWeight: 'bold',
    color: '#ffffff',
    backgroundColor: '#007bff', // Accent color to match link color in header
    boxShadow: '0 4px 8px rgba(0, 0, 0, 0.2)',
  },
  errorButton: {
    backgroundColor: '#b30000', // Dark red for delete buttons in dark mode
    textAlign: 'center',
    borderRadius: '10px',
    cursor: 'pointer',
    color: '#ffffff',
    boxShadow: '0 4px 8px rgba(0, 0, 0, 0.3)',
  },
  success: {
    color: '#66ff66', // Light green for success messages
    fontSize: '1rem',
    marginTop: '5px',
  },
  error: {
    color: '#ff6666', // Light red for error messages
    fontSize: '1rem',
    marginTop: '5px',
  },
  requirementContainer: {
    marginTop: '10px',
  },
  requirementTable: {
    border: '1px solid #444',
    borderRadius: '8px',
    padding: '10px',
    backgroundColor: '#3a3a3a',
    marginBottom: '30px',
  },
  loading: {
    textAlign: 'center',
    fontSize: '1rem',
    color: '#d3d3d3',
  },
  errorMessage: {
    textAlign: 'center',
    fontSize: '1rem',
    color: '#ff6666',
  },
};

const ManageParty = () => {
  const navigate = useNavigate();
  const api = useApi()
  const partyId = Number(localStorage.getItem('partyId') || '-1')

  const [usernameInput, setUsernameInput] = useState('');
  const [inviteFeedbackSuccess, setInviteFeedbackSuccess] = useState('');
  const [inviteFeedbackError, setInviteFeedbackError] = useState('');
  const [requirementModalVisible, setRequirementModalVisible] = useState(false);
  const [requirementModalMode, setRequirementModalMode] = useState('');
  const [deleteModalVisible, setDeleteModalVisible] = useState(false);
  const [deleteModalMode, setDeleteModalMode] = useState('');
  const [requirementToDelete, setRequirementToDelete] = useState(-1);

  const [party, setParty] = useState<PartyPopulated>(EMPTY_PARTY_POPULATED)
  const [drinkReqs, setDrinkReqs] = useState<RequirementPopulated[]>([])
  const [foodReqs, setFoodReqs] = useState<RequirementPopulated[]>([])
  const [participants, setParticipants] = useState<User[]>([])
  const [pendingInvites, setPendingInvites] = useState<PartyInvite[]>([])

  // const { selectedParty } = useSelector((state: RootState) => state.selectedPartyStore);
  // const { requirements: dRequirements, loading: dReqLoading, error: dReqError } = useSelector((state: RootState) => state.drinkRequirementStore);
  // const { requirements: fRequirements, loading: fReqLoading, error: fReqError } = useSelector((state: RootState) => state.foodRequirementStore);
  // const { participants, loading: participantLoading, error: participantError } = useSelector((state: RootState) => state.partyParticipantStore);
  // const { pendingInvites, loading: pendingInvitesLoading, error: pendingInvitesError } = useSelector((state: RootState) => state.partyPendingInviteStore);

  useEffect(() => {
    api.partyApi.getParty(partyId)
        .then(result => {
          if(result === 'error'){
            toast.error('Unable to load party')
            return
          }
          if(result === 'private party'){
            toast.error('Navigation error')
            navigate('/partyHome')
            return
          }
          setParty(result.data)
        })
        .catch(() => {
          toast.error('Unexpected error')
        })
  }, [api.partyApi, navigate, partyId]);

  useEffect(() => {
    api.requirementApi.getDrinkRequirementsByPartyId(partyId)
        .then(result => {
          if(result === 'error'){
            toast.error('Unable to load drink requirements')
            return
          }
          setDrinkReqs(result.data)
        })
        .catch(() => {
          toast.error('Unexpected error')
        })
  }, [api.requirementApi, partyId]);

  useEffect(() => {
    api.requirementApi.getFoodRequirementsByPartyId(partyId)
        .then(result => {
          if(result === 'error'){
            toast.error('Unable to load food requirements')
            return
          }
          setFoodReqs(result.data)
        })
        .catch(() => {
          toast.error('Unexpected error')
        })
  }, [api.requirementApi, partyId]);

  useEffect(() => {
    api.partyApi.getPartyParticipants(partyId)
        .then(result => {
          if(result === 'error'){
            toast.error('Unable to load party participants')
            return
          }
          setParticipants(result.data)
        })
        .catch(() => {
          toast.error('Unexpected error')
        })
    
  }, [api.partyApi, partyId]);

  useEffect(() => {
    api.partyAttendanceApi.getPartyPendingInvites(partyId)
        .then(result => {
          if(result === 'error'){
            toast.error('Unable to load party invites')
            return
          }
          setPendingInvites(result.data)
        })
        .catch(() => {
          toast.error('Unexpected error')
        })
  }, [api.partyAttendanceApi, partyId]);


  const handleInviteToParty = (username: string) => {
    inviteToParty(party.ID, username)
      .then(() => {
        //todo: reload pending invites
        setForTime<string>(setInviteFeedbackSuccess, 'Invite sent!', '', 3000);
      })
      .catch((err) => {
        console.log(err);
        setForTime<string>(setInviteFeedbackError, 'Something went wrong!', '', 3000);
      });
  };

  const handleAddRequirement = (mode: string) => {
    if (mode === 'drink') setRequirementModalMode('drink');
    if (mode === 'food') setRequirementModalMode('food');
    setRequirementModalVisible(true);
  };

  const handleDeleteRequirement = (requirement: Requirement, mode: string) => {
    if (mode === 'drink') {
      setDeleteModalMode('drink');
      setRequirementToDelete(requirement.ID || -1);
      setDeleteModalVisible(true);
    }

    if (mode === 'food') {
      setDeleteModalMode('food');
      setRequirementToDelete(requirement.ID || -1);
      setDeleteModalVisible(true);
    }
  };

  const handleKickParticipant = (kickedUser: User) => {
    kickFromParty(party.ID, kickedUser.ID)
      .then(() => {
        //todo: reload party participants
      })
      .catch((err) => {
        console.log(err);
        // todo: make a confirmaction modal and handle feedback
      });
  };

  const participantColumns = [
    ...userTableColumnsLegacy,
    {
      title: '',
      key: 'action 1',
      render: (record: User) => (
        <Button
          style={styles.errorButton}
          onClick={() => handleKickParticipant(record)}
        >
          Kick
        </Button>
      ),
    },
  ];

  const drinkRequirementColumns = [
    ...requirementTableColumnsLegacy,
    {
      title: '',
      key: 'action 1',
      render: (record: Requirement) => (
        <Button
          style={styles.errorButton}
          onClick={() => handleDeleteRequirement(record, 'drink')}
        >
          Delete
        </Button>
      ),
    },
  ];

  const foodRequirementColumns = [
    ...requirementTableColumnsLegacy,
    {
      title: '',
      key: 'action 1',
      render: (record: Requirement) => (
        <Button
          style={styles.errorButton}
          onClick={() => handleDeleteRequirement(record, 'food')}
        >
          Delete
        </Button>
      ),
    },
  ];

  const renderReqs = (requirements: Requirement[], mode: string) => {
    if (!requirements || requirements.length === 0) {
      return <div>There&#39;s no drink requirements yet!</div>;
    }
    return (
      <Table
        dataSource={requirements.map((req) => ({ ...req, key: req.ID }))}
        columns={mode === 'drink' ? drinkRequirementColumns : foodRequirementColumns}
        pagination={false}
        scroll={{ y: 200 }}
      />
    );
  };

  const renderParticipants = () => {
    if (!participants || participants.length === 0) {
      return <div>There&#39;s no drink requirements yet!</div>;
    }
    return (
      <Table
        dataSource={participants.map((person) => ({ ...person, key: person.ID }))}
        columns={participantColumns}
        pagination={false}
        scroll={{ y: 200 }}
      />
    );
  };

  const renderPendingInvites = () => {
    if (!pendingInvites || pendingInvites.length === 0) {
      return <div>There&#39;s no pending invites at the moment</div>;
    }
    return (
      <Table
        dataSource={pendingInvites.map((invite) => ({ ...invite, key: invite.ID }))}
        columns={invitedTableColumnsLegacy}
        pagination={false}
        scroll={{ y: 200 }}
      />
    );
  };

  return (
    <div style={styles.background}>
      <div style={styles.outerContainer}>
        <ConfigProvider theme={{ algorithm: theme.darkAlgorithm }}>
          {/*<VisitPartyNavBar onProfileClick={() => setProfileOpen(true)} />*/}
          {/*<VisitPartyProfile*/}
          {/*  isOpen={profileOpen}*/}
          {/*  onClose={() => setProfileOpen(false)}*/}
          {/*  currentParty={party}*/}
          {/*  user={user}*/}
          {/*  onLeaveParty={() => console.log('leaveparty')}*/}
          {/*/>*/}
          <CreateRequirementModal
            visible={requirementModalVisible}
            onClose={() => setRequirementModalVisible(false)}
            mode={requirementModalMode}
          />
          <DeleteRequirementModal
            visible={deleteModalVisible}
            onClose={() => setDeleteModalVisible(false)}
            mode={deleteModalMode}
            requirementId={requirementToDelete}
          />
          <div style={styles.container}>
            <h2>Invite</h2>
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
                onClick={() => handleInviteToParty(usernameInput)}
              >
                Invite
              </Button>
              {inviteFeedbackSuccess && <p style={styles.success}>{inviteFeedbackSuccess}</p>}
              {inviteFeedbackError && <p style={styles.error}>{inviteFeedbackError}</p>}
            </div>

            <h2>Drink Requirements</h2>
            <div style={styles.requirementContainer}>
              <Button
                type='primary'
                style={styles.button}
                onClick={() => handleAddRequirement('drink')}
              >
                Add
              </Button>
              <div style={styles.requirementTable}>
                {!drinkReqs && <div>Loading...</div>}
                {drinkReqs && renderReqs(drinkReqs, 'drink')}
              </div>
            </div>

            <h2>Food Requirements</h2>
            <div style={styles.requirementContainer}>
              <Button
                type='primary'
                style={styles.button}
                onClick={() => handleAddRequirement('food')}
              >
                Add
              </Button>
              <div style={styles.requirementTable}>
                {!foodReqs && <div>Loading...</div>}
                {foodReqs && renderReqs(foodReqs, 'food')}
              </div>
            </div>

            <h2>Participants</h2>
            <div style={styles.requirementContainer}>
              <div style={styles.requirementTable}>
                {!participants && <div>Loading...</div>}
                {participants && renderParticipants()}
              </div>
            </div>

            <h2>Pending Invites</h2>
            <div style={styles.requirementContainer}>
              <div style={styles.requirementTable}>
                {!pendingInvites && <div>Loading...</div>}
                {pendingInvites && renderPendingInvites()}
              </div>
            </div>
          </div>
        </ConfigProvider>
      </div>
    </div>
  );
};

export default ManageParty;
