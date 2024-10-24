import React, {CSSProperties} from "react";
import VisitPartyNavBar from "../../../components/navbar/VisitPartyNavBar";
import "bootstrap/dist/css/bootstrap.min.css";


const PartyHome: React.FC = () => {
    return (
        <div style={styles.outerContainer}>
            <VisitPartyNavBar/>
            <div style={styles.container}>
                {/* Middle Section */}
                <div style={styles.middleSection}>
                    {/* Title */}
                    <h1 style={{fontSize: '48px', margin: '20px 0'}}>Big Title</h1>

                    {/* Subtitle */}
                    <p style={{fontSize: '18px', margin: '10px 0'}}>This is the smaller text below the title.</p>

                    {/* Progress Bar */}
                    <div style={styles.progressBarContainer}>
                        <div style={styles.progressBar}>60% Complete</div>
                    </div>

                    {/* Button */}
                    <button style={styles.button}>Click Me</button>
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
                            <a href="#" style={styles.link}>Join Group Chat</a>
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
    },
    container: {
        display: 'flex',
        flexDirection: 'column',
        justifyContent: 'space-between',
        height: 'auto',
        flexGrow: '1',
        padding: '20px',
    },
    middleSection: {
        display: 'flex',
        flexDirection: 'column',
        alignItems: 'center',
        textAlign: 'center',
        flexGrow: 1,
    },
    progressBarContainer: {
        width: '75%',
        backgroundColor: '#e0e0e0',
        borderRadius: '10px',
        height: '20px',
        position: 'relative' as 'relative',
        margin: '20px 0',
    },
    progressBar: {
        width: '60%',
        backgroundColor: '#4caf50',
        borderRadius: '10px',
        height: '100%',
        textAlign: 'center',
        color: 'white',
        lineHeight: '20px',
    },
    bottomRow: {
        display: 'flex',
        justifyContent: 'space-around',
        alignItems: 'center',
        paddingBottom: '20px',
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
    }
};

export default PartyHome;
