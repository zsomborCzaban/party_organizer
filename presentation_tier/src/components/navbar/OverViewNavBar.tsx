import React from 'react';
import {useNavigate} from 'react-router-dom';

const styles: {[key: string]: React.CSSProperties} = {
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
        margin: '0',
        textDecoration: 'none',
        color: '#007bff',
        fontSize: '18px',
        cursor: 'pointer',
    },
};

type OverViewNavBarProps = {
    onProfileClick: () => void;
};

const OverViewNavBar: React.FC<OverViewNavBarProps> = ({ onProfileClick }) => {
    const navigate = useNavigate();

    return (
      <header style={styles.header}>
        <h1 style={styles.title}>Party Organizer</h1>
        <nav style={styles.nav}>
          <ul style={styles.navList}>
            <li style={styles.navItem}><p onClick={() => navigate('/overview/discover')} style={styles.link}>Discover</p></li>
            <li style={styles.navItem}><p onClick={() => navigate('/overview/parties')} style={styles.link}>Parties</p></li>
            <li style={styles.navItem}><p onClick={() => navigate('/overview/friends')} style={styles.link}>Friends</p></li>
            <li style={styles.navItem}><p
              style={styles.link}
              onClick={e => {
                            e.preventDefault();
                            onProfileClick();
                        }}
                                       >Profile</p></li>
          </ul>
        </nav>
      </header>
    );
};


export default OverViewNavBar;
