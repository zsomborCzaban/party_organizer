import React from 'react';
import {getUserId} from "../../auth/AuthUserUtil";
import {authService} from "../../auth/AuthService";
import {useSelector} from "react-redux";
import {RootState} from "../../store/store";
import {useNavigate} from "react-router-dom";

type VisitPartyNavBarProps = {
    onProfileClick: () => void;
};

const VisitPartyNavBar: React.FC<VisitPartyNavBarProps> = ({ onProfileClick }) => {
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

    const handlePartyClick = () => {
        navigate("/visitParty/partyHome")
    }
    const handleOverViewClick = () => {
        navigate("/overview/discover")
    }

    return (
        <header style={styles.header}>
            <div style={styles.tittleContainer}>
                <h1 onClick={handlePartyClick} style={styles.title}>{selectedParty.name}</h1>
                <h1 onClick={handleOverViewClick} style={styles.title}>OverView</h1>
            </div>
            <nav style={styles.nav}>
                <ul style={styles.navList}>
                    {userId === selectedParty.organizer.ID &&
                        <li style={styles.navItem}><p onClick={() => navigate("/visitParty/manageParty")}  style={styles.link}>Manage Party</p></li>
                    }
                    {userId === selectedParty.organizer.ID &&
                        <li style={styles.navItem}><p onClick={() => navigate("/visitParty/partySettings")}  style={styles.link}>Party Settings</p></li>
                    }
                    <li style={styles.navItem}><p onClick={() => navigate("/visitParty/partyHome")}  style={styles.link}>Home</p></li>
                    <li style={styles.navItem}><p onClick={() => navigate("/visitParty/contributions")}  style={styles.link}>Contributions</p>
                    </li>
                    <li style={styles.navItem}><p onClick={() => navigate("/visitParty/hallOfFame")}  style={styles.link}>Hall Of Fame</p>
                    </li>
                    <li style={styles.navItem}><p style={styles.link} onClick={(e) => {
                        e.preventDefault();
                        onProfileClick();
                    }}>Profile</p></li>
                </ul>
            </nav>
        </header>
    );
};

const styles: {[key: string]: React.CSSProperties} = {
    header: {
        display: 'flex',
        justifyContent: 'space-between',
        alignItems: 'center',
        padding: '10px 20px',
        backgroundColor: 'rgba(33, 33, 33, 0.8)',
        boxShadow: '0 2px 4px rgba(0, 0, 0, 0.1)',
    },
    tittleContainer: {
        display: 'flex',
        flexDirection: 'row',
        gap: '20%',
        alignItems: 'center',
    },
    title: {
        fontSize: '24px',
        color: '#007bff',
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
        margin: '0',
        cursor: 'pointer',
        textDecoration: 'none',
        color: '#007bff',
        fontSize: '18px',
    },
};

export default VisitPartyNavBar;
