import React, {CSSProperties, useEffect, useState} from "react";
import VisitPartyNavBar from "../../../components/navbar/VisitPartyNavBar";
import "bootstrap/dist/css/bootstrap.min.css";
import {useNavigate} from "react-router-dom";
import {useSelector} from "react-redux";
import {RootState} from "../../../store/store";
import videoBackground from "../../../data/resources/videos/party_video.mp4"
import VisitPartyProfile from "../../../components/drawer/VisitPartyProfile";
import {User} from "../../../data/types/User";
import {getUser} from "../../../auth/AuthUserUtil";
import {authService} from "../../../auth/AuthService";


const PartyHome: React.FC = () => {
    const navigate = useNavigate()

    const [profileOpen, setProfileOpen] = useState(false)
    const [user, setUser] = useState<User>()

    const {selectedParty} = useSelector((state: RootState)=> state.selectedPartyStore)

    useEffect(() => {
        const currentUser = getUser()

        if(!currentUser) {
            authService.handleUnauthorized()
            return
        }

        setUser(currentUser)
    }, []);

    if(!selectedParty){
        console.log("error, no selected party")
        navigate("/overview/discover")
        return <div>error, selected party was null</div>
    }

    if(!user){
        console.log("user was null")
        return <div>Loading...</div>
    }

    const handleContributeClick = () => {
        navigate("/visitParty/Contributions")
    }

    return (
        <div style={styles.outerContainer}>

            {/*<div style={{ position: 'relative', width: '100%', height: '100vh', overflow: 'hidden' }}>*/}
                <video
                    src={videoBackground}
                    autoPlay
                    loop
                    muted
                    playsInline
                    style={styles.video}
                />
            {/*</div>*/}

            <VisitPartyNavBar onProfileClick={() => setProfileOpen(true)}/>
            <VisitPartyProfile isOpen={profileOpen} onClose={() => setProfileOpen(false)} currentParty={selectedParty} user={user} onLeaveParty={() => console.log("leaveparty")} />
            <div style={styles.container}>
                {/* Middle Section */}
                <div style={styles.middleSection}>
                    {/* Title */}
                    <h1 style={{fontSize: '48px', margin: '20px 0'}}>{selectedParty.name}</h1>

                    {/* Subtitle */}
                    <p style={{fontSize: '18px', margin: '10px 0'}}>Contributions to the party.</p>

                    {/* Progress Bar */}
                    <div style={styles.progressBarContainer}>
                        <p style={styles.progressText}> 60% Complete</p>
                        <div style={styles.progressBar}/>
                    </div>

                    {/* Button */}
                    <button style={styles.button} onClick={handleContributeClick}>Contribute</button>
                </div>

                {/* Bottom Row */}
                <div style={styles.bottomRow}>
                    {/* Date */}
                    <div>
                        <h5>12<sup>th</sup> June, 18:00</h5>
                    </div>

                    {/* Place */}
                    <div>
                        <h5>Kecskemét, Szabadság tér 13</h5>
                    </div>

                    {/* Link */}
                    <div>
                        <h5>
                            <a href="" style={styles.link}>Join Group Chat</a>
                        </h5>
                    </div>
                </div>
            </div>
        </div>
);
};

const styles: { [key: string]: CSSProperties } = {
    outerContainer: {
        height: '100vh',
        display: 'flex',
        flexDirection: 'column',
        backgroundColor: 'rgba(0, 0, 0, 0.3)',
    },
    container: {
        display: 'flex',
        flexDirection: 'column',
        justifyContent: 'space-between',
        flexGrow: '1',
        padding: '20px',
    },
    middleSection: {
        display: 'flex',
        flexDirection: 'column',
        justifyContent: 'center',
        alignItems: 'center',
        textAlign: 'center',
        flexGrow: 1,
        color: 'white'
    },
    progressBarContainer: {
        width: '75%',
        backgroundColor: '#e0e0e0',
        borderRadius: '10px',
        height: '20px',
        position: 'relative',
        margin: '10px 0 0 0',
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'center',
        color: 'white',
    },
    progressBar: {
        width: '60%',
        backgroundColor: '#4caf50',
        borderRadius: '10px',
        height: '100%',
        position: 'absolute',
        left: 0,
        top: 0,
    },
    progressText: {
      zIndex: 1,
        margin: 0,
    },
    bottomRow: {
        display: 'flex',
        justifyContent: 'space-around',
        alignItems: 'center',
        paddingBottom: '20px',
        color: '#007bff',
    },
    button: {
        padding: '10px 20px',
        backgroundColor: '#007bff',
        color: 'white',
        border: 'none',
        borderRadius: '5px',
        cursor: 'pointer',
        marginTop: '20px',
    },
    link: {
        color: '#007bff',
        textDecoration: 'none',
    },
    video: {
        position: 'absolute',
        top: 0,
        left: 0,
        width: '100%',
        height: '100%',
        objectFit: 'cover',
        zIndex: -1,
    }
};

export default PartyHome;
