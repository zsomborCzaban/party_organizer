import React from 'react';

type OverViewNavBarProps = {
    onProfileClick: () => void;
};

const OverViewNavBar: React.FC<OverViewNavBarProps> = ({ onProfileClick }) => {
    return (
        <header style={styles.header}>
            <h1 style={styles.title}>Party Organizer</h1>
            <nav style={styles.nav}>
                <ul style={styles.navList}>
                    <li style={styles.navItem}><a href="/overview/discover" style={styles.link}>Discover</a></li>
                    <li style={styles.navItem}><a href="/overview/parties" style={styles.link}>Parties</a></li>
                    <li style={styles.navItem}><a href="/overview/friends" style={styles.link}>Friends</a></li>
                    <li style={styles.navItem}><a style={styles.link}
                        onClick={e => {
                            e.preventDefault()
                            onProfileClick()
                        }}
                    >Profile</a></li>
                </ul>
            </nav>
        </header>
    );
};

// Inline styles
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
        textDecoration: 'none',
        color: '#007bff',
        fontSize: '18px',
        cursor: 'pointer',
    },
};

export default OverViewNavBar;
