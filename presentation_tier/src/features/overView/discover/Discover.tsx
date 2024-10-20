import {useDispatch, useSelector} from "react-redux";
import {AppDispatch, RootState} from "../../../store/store";
import {CSSProperties, useEffect, useState} from "react";
import {loadPublicParties} from "./PublicPartySlice";
import OverViewNavBar from "../../../components/navbar/OverViewNavBar";
import { Table } from 'antd';
import createImage from '../../../midjourney_images/u5679648646_httpss.mj.runKqg0SHl7m9w_make_a_picture_similar_t_2a92ccce-3fd5-4da4-8398-11898f188cd5_3.png'
import joinImage from '../../../midjourney_images/u5679648646_2_people_dancing_with_a_galactic_lsd_like_trip_li_78e43cba-8023-4e69-9fa1-75bb9790bbe8_0.png'
import {useNavigate} from "react-router-dom";
import {Party} from "../Party";
import AccessCodeModal from "./AccessCodeModal";


const Discover = () => {
    const dispatch = useDispatch<AppDispatch>()

    const {parties, loading, error} = useSelector(
        (state: RootState) => state.publicPartyStore
    )

    useEffect( () => {
        dispatch(loadPublicParties());
    }, []);

    const navigate = useNavigate()

    const [modalVisible, setModalVisible] = useState(false)


    const handleVisitClicked = (record: Party) => {
        console.log(record)
        //todo: set selected party to record and navigate to the parties page
    }
    const handleCreate = () => {
        navigate("/createParty")
    }
    const handleJoin = () =>{
        setModalVisible(true);
    }

    const closeModal = () => {
        setModalVisible(false);
    };

    const columns = [
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
            dataIndex: 'headcount',
            key: 'ID',
            render: (text: string, record: Party) => (
                <button onClick={() => handleVisitClicked(record)}>Visit</button>
            ),
        },
    ];

    const renderPublicParties = () => {
        if(parties.length === 0){
            return <div>There's no public parties at the moment :( </div>
        }
        return (<Table
            dataSource={parties}
            columns={columns}
            pagination={false} // Disable pagination
            scroll={{y: 200}} // Set vertical scroll height
        />)
    }

    if (loading) return <div>Loading...</div>
    if (error) return <div>Error: {error.message}</div>

    return (
        <div style={styles.outerContainer}>
            <OverViewNavBar/>
            <div style={styles.container}>

                <h2 style={styles.label}>Public Parties</h2>

                {/* Scrollable Table using Ant Design Table */}
                <div style={styles.tableContainer}>
                    {renderPublicParties()}
                </div>

                <div style={styles.message}>
                    Didn't find the right party? Don't worry, you can create your own!
                </div>

                {/* Rectangles (Create and Join) */}
                <div style={styles.buttonsContainer}>
                    <div style={styles.createButton} onClick={handleCreate}>Create</div>
                    <div style={styles.joinButton} onClick={handleJoin}>Join</div>
                </div>
                <AccessCodeModal visible={modalVisible} onClose={closeModal} />
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
    createButton: {
        width: '48%',
        backgroundImage: `url(${createImage})`, // Ensure the URL is set properly
        backgroundSize: 'cover',
        backgroundPosition: 'center',
        color: "white",
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
        backgroundImage: `url(${joinImage})`, // Ensure the URL is set properly
        backgroundSize: 'cover',
        backgroundPosition: 'center',
        color: "white",
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
};


export default Discover;