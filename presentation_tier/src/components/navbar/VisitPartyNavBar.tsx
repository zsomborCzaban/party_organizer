import React from 'react';
import {getUserId} from "../../auth/AuthUserUtil";
import {authService} from "../../auth/AuthService";
import {useSelector} from "react-redux";
import {RootState} from "../../store/store";
import {useNavigate} from "react-router-dom";

const VisitPartyNavBar: React.FC = () => {
    const navigate = useNavigate()
    const {selectedParty} = useSelector((state: RootState)=> state.selectedPartyStore)
    if(!selectedParty){
        console.log("error, no selected party")
        navigate("/overview/discover")
        return <div>error, selected party was null</div>
    }
    if(!selectedParty.organizer){
        console.log("error, selected party had no organizer")
        navigate("/overview/discover")
        return <div>error, selected party had no organizer</div>
    }

    const userIdString = getUserId()
    if(!userIdString){
        authService.handleUnauthorized()
        return <div>error, couldn't get userId</div>
    }
    const userId = parseInt(userIdString)

    return (
        <header style={styles.header}>
            <h1 style={styles.title}>{selectedParty.name}</h1>
            <nav style={styles.nav}>
                <ul style={styles.navList}>
                    {userId === selectedParty.organizer.ID && <li style={styles.navItem}><a href="/visitParty/manageParty" style={styles.link}>Manage Party</a></li>}
                    <li style={styles.navItem}><a href="/visitParty/partyHome" style={styles.link}>Home</a></li>
                    <li style={styles.navItem}><a href="/visitParty/contributions" style={styles.link}>Contributions</a></li>
                    <li style={styles.navItem}><a href="/visitParty/hallOfFame" style={styles.link}>Hall Of Fame</a></li>
                    <li style={styles.navItem}><a href="/profile" style={styles.link}>Profile</a></li>
                </ul>
            </nav>
        </header>
    );
};

// Inline styles
const styles = {
    header: {
        display: 'flex',
        justifyContent: 'space-between',
        alignItems: 'center',
        padding: '10px 20px',
        backgroundColor: '#f8f9fa',
        boxShadow: '0 2px 4px rgba(0, 0, 0, 0.1)',
    },
    title: {
        fontSize: '24px',
        color: '#333',
    },
    nav: {
        display: 'flex',
    },
    navList: {
        listStyleType: 'none',
        display: 'flex',
        margin: 0,
        padding: 0,
    },
    navItem: {
        margin: '0 15px',
    },
    link: {
        textDecoration: 'none',
        color: '#007bff',
        fontSize: '18px',
    },
};

export default VisitPartyNavBar;
