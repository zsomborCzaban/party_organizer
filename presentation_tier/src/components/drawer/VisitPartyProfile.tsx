import React, {ChangeEvent, useState} from 'react';
import {User} from "../../data/types/User";
import defaultProfilePicture from "../../data/resources/images/default_profile_picture.png"
import {Button} from "antd";
import {Party} from "../../data/types/Party";
import {handleProfilePictureUpload} from "../../data/utils/imageUtils";

type DrawerProps = {
    isOpen: boolean;
    onClose: () => void;
    currentParty: Party;
    user: User;
    onLogout: () => void;
    onLeaveParty: () => void;
};

const Drawer: React.FC<DrawerProps> = ({ isOpen, onClose, user, currentParty, onLogout, onLeaveParty }) => {
    const [profileErrorMessage, setProfileErrorMessage] = useState("")
    const [profilePictureUrl, setProfilePictureUrl] = useState(user.profile_picture_url ? user.profile_picture_url : defaultProfilePicture)

    const handleFileUpload = (event: ChangeEvent<HTMLInputElement>) => {
        handleProfilePictureUpload(event, setProfilePictureUrl, setProfileErrorMessage)
    }


    return (
        <div>
            {isOpen && (
                <div style={styles.backdrop} onClick={onClose} />
            )}
            <div
                style={{
                    ...styles.drawerContainer,
                    right: isOpen ? 0 : '-100%',
                }}
            >
                <svg
                    style={styles.closeIcon}
                    onClick={onClose}
                    xmlns="http://www.w3.org/2000/svg"
                    width="40"
                    height="40"
                    fill="none"
                    viewBox="0 0 24 24"
                >
                    <path
                        stroke="#fff"
                        strokeWidth="2"
                        d="M6 18L18 6M6 6l12 12"
                    />
                </svg>

                <div style={styles.profileContainer}>
                    <img src={profilePictureUrl} alt="Profile"
                         style={styles.profilePicture}/>

                    <input style={{display: 'none'}} id="file-input" type="file" accept="image/*"
                           onChange={handleFileUpload}/>
                    <label htmlFor="file-input" style={styles.changeProfilePicture}>
                        Upload Picture
                    </label>
                    {profileErrorMessage && <p style={styles.errorMessage}>{profileErrorMessage}</p>}
                </div>


                <div style={styles.infoContainer}>
                    <div style={styles.userInfo}>
                        <label style={styles.label}>Username:</label>
                        <div style={styles.userData}>{user.username}</div>
                    </div>

                    <div style={styles.userInfo}>
                        <label style={styles.label}>Email:</label>
                        <div style={styles.userData}>{user.email}</div>
                    </div>

                    <div style={styles.buttonContainer}>
                        <Button type="primary" onClick={onLogout} style={styles.logoutButton}>
                            Logout
                        </Button>
                        <Button type="primary" onClick={onLogout} style={styles.leavePartyButton}>
                            Leave Party
                        </Button>
                    </div>
                </div>
            </div>
        </div>
    );
};

const styles: { [key: string]: React.CSSProperties } = {
    backdrop: {
        position: 'fixed',
        top: 0,
        left: 0,
        width: '100vw',
        height: '100vh',
        backgroundColor: 'rgba(0, 0, 0, 0.8)',  // Darker backdrop
        zIndex: 999,
        transition: 'opacity 0.3s ease',
    },
    drawerContainer: {
        position: 'fixed',
        top: 0,
        right: 0,
        width: 'min(400px, 80%)',
        height: '100%',
        backgroundColor: '#333',
        color: '#fff',
        boxShadow: '0 0 10px rgba(0, 0, 0, 0.7)',
        transition: 'right 0.3s ease',
        padding: '20px',
        display: 'flex',
        flexDirection: 'column',
        alignItems: 'center',
        zIndex: 1000,
    },
    profileContainer: {
        width: 'min(300px, 90%)',
        display: 'flex',
        flexDirection: 'column',
        alignItems: 'center',
        justifyItems: 'flex-start',
        marginBottom: '20px',
    },
    infoContainer: {
        width: 'min(300px, 90%)',
        height: '100%',
        display: 'flex',
        flexDirection: 'column',
        alignItems: 'flex-start'
    },
    closeIcon: {
        alignSelf: 'flex-start',
        cursor: 'pointer',
        marginBottom: '20px',
    },
    profilePicture: {
        width: '100px',
        height: '100px',
        borderRadius: '50%',
        marginBottom: '20px',
    },
    changeProfilePicture: {
        display: 'inline-block',
        padding: '5px',
        color: '#fff',
        backgroundColor: '#007bff',
        borderRadius: '5px',
        cursor: 'pointer',
        fontWeight: 'bold',
        fontSize: '18px',
        width: '100%',
        textAlign: 'center',
    },
    userInfo: {
        textAlign: 'left',
        marginBottom: '10px',
    },
    label: {
        fontWeight: 'bold',
    },
    userData: {
        marginLeft: '30px',
    },
    buttonContainer: {
        width: '100%',
        marginTop: 'auto',
        display: 'flex',
        flexDirection: 'row',
        justifyContent: 'space-between',
    },
    logoutButton: {
        // padding: '10px 20px',
        // border: 'none',
        // borderRadius: '5px',
        fontWeight: 'bold',
        fontSize: '18px',
        width: '48%',
    },
    leavePartyButton: {
        fontWeight: 'bold',
        fontSize: '18px',
        width: '48%',
        backgroundColor: 'red',
    },
};


export default Drawer;
