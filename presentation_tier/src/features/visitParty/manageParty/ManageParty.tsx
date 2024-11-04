import OverViewNavBar from "../../../components/navbar/OverViewNavBar";
import OverViewProfile from "../../../components/drawer/OverViewProfile";
import React, {CSSProperties} from "react";
import VisitPartyNavBar from "../../../components/navbar/VisitPartyNavBar";
import VisitPartyProfile from "../../../components/drawer/VisitPartyProfile";

const ManageParty = () => {
    return <div style={styles.outerContainer}>
        {/*<VisitPartyNavBar onProfileClick={() => setProfileOpen(true)}/>*/}
        {/*<VisitPartyProfile isOpen={profileOpen} onClose={() => setProfileOpen(false)} user={user} onLogout={() => {console.log("logout")}} currentParty={}/>*/}
        {/*<div style={styles.container}>*/}

        {/*    <h2 style={styles.label}>Party Invites</h2>*/}

        {/*    /!* Scrollable Table using Ant Design Table *!/*/}
        {/*    <div style={styles.tableContainer}>*/}
        {/*        {inviteLoading && <div>Loading...</div>}*/}
        {/*        {inviteError && <div>Error: Some unexpected error happened</div>}*/}
        {/*        {(!inviteLoading && !inviteError) && renderInvites()}*/}
        {/*    </div>*/}

        {/*    <h2 style={styles.label}>Attended Parties</h2>*/}

        {/*    /!* Scrollable Table using Ant Design Table *!/*/}
        {/*    <div style={styles.tableContainer}>*/}
        {/*        {attendedLoading && <div>Loading...</div>}*/}
        {/*        {attendedError && <div>Error: Some unexpected error happened</div>}*/}
        {/*        {(!attendedLoading && !attendedError) && renderParties("attended")}*/}
        {/*    </div>*/}

        {/*    <h2 style={styles.label}>Organized Parties</h2>*/}

        {/*    /!* Scrollable Table using Ant Design Table *!/*/}
        {/*    <div style={styles.tableContainer}>*/}
        {/*        {organizedLoading && <div>Loading...</div>}*/}
        {/*        {organizedError && <div>Error: Some unexpected error happened</div>}*/}
        {/*        {(!organizedLoading && !organizedError) && renderParties("organized")}*/}
        {/*    </div>*/}
        {/*</div>*/}
    </div>
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
};


export default ManageParty