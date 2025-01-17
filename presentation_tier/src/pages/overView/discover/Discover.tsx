import { Button, Table } from 'antd';
import { CSSProperties, useEffect, useState } from 'react';
import {useDispatch, useSelector} from 'react-redux';
import { useNavigate } from 'react-router-dom';
import { authService } from 'src/auth/AuthService';
import { getUser } from 'src/auth/AuthUserUtil';
import OverViewProfile from 'src/components/drawer/OverViewProfile';
import OverViewNavBar from 'src/components/navbar/OverViewNavBar';
import { joinPublicParty } from 'src/data/apis/PartyAttendanceManagerApi';
import { partyTableColumns } from 'src/data/constants/TableColumns';
import { setSelectedParty } from 'src/data/sclices/PartySlice';
import { loadPublicParties } from 'src/data/sclices/PublicPartySlice';
import { Party } from 'src/data/types/Party';
import { User } from 'src/data/types/User';
import { AppDispatch, RootState } from 'src/store/store';
import AccessCodeModal from './AccessCodeModal';


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

    const {parties, loading, error} = useSelector(
        (state: RootState) => state.publicPartyStore,
    );

    useEffect( () => {
        dispatch(loadPublicParties());
        const currentUser = getUser();

        if(!currentUser) {
            authService.handleUnauthorized();
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

    const handleJoin = () =>{
        setModalVisible(true);
    };

    const closeModal = () => {
        setModalVisible(false);
    };

    const columns = [...partyTableColumns,
        {
            title: '',
            key: 'ID',
            render: (text: string, record: Party) => (
              <Button style={styles.button} onClick={() => handleVisitClicked(record)}>Visit</Button>
            ),
        },
    ];

    const renderPublicParties = () => {
        if(!parties || parties.length === 0){
            return <div>There&#39;s no public parties at the moment :( </div>;
        }
        return (<Table
          dataSource={parties.map(party => ({...party, key: party.ID}))}
          columns={columns}
          pagination={false} // Disable pagination
          scroll={{y: 200}}
                />);
    };

    if(!user){
        console.log('user was null');
        return <div>Loading...</div>;
    }

    return (
      <div style={styles.outerContainer}>
        <OverViewNavBar onProfileClick={() => setProfileOpen(true)} />
        <OverViewProfile isOpen={profileOpen} onClose={() => setProfileOpen(false)} user={user} />
        <div style={styles.container}>

          <h2 style={styles.label}>Public Parties</h2>

          <div style={styles.tableContainer}>
            {loading && <div>Loading...</div>}
            {error && <div>Error: Some unexpected error happened</div>}
            {(!loading && !error) && renderPublicParties()}
          </div>

          <div style={styles.message}>
            Didn&#39;t find the right party? Don&#39;t worry, you can create your own!
          </div>

          <div style={styles.buttonsContainer}>
            <div style={styles.createButton} onClick={handleCreate}>Create</div>
            <div style={styles.joinButton} onClick={handleJoin}>Join</div>
          </div>
          <AccessCodeModal visible={modalVisible} onClose={closeModal} />
          {unexpectedError && <div style={styles.errorFeedback}>{unexpectedError}</div>}
        </div>
      </div>
    );
};



export default Discover;