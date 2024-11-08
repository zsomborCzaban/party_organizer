import React, {CSSProperties, useEffect, useRef, useState} from "react";
import VisitPartyNavBar from "../../../components/navbar/VisitPartyNavBar";
import VisitPartyProfile from "../../../components/drawer/VisitPartyProfile";
import {User} from "../../../data/types/User";
import {getUser} from "../../../auth/AuthUserUtil";
import {authService} from "../../../auth/AuthService";
import {useDispatch, useSelector} from "react-redux";
import {AppDispatch, RootState} from "../../../store/store";
import {useNavigate} from "react-router-dom";
import {Button, ConfigProvider, Table, theme} from "antd";
import {loadDrinkRequirements} from "../../../data/sclices/DrinkRequirementSlice";
import {loadFoodRequirements} from "../../../data/sclices/FoodRequirementSlice";
import {
    invitedTableColumns,
    requirementTableColumns,
    userTableColumns
} from "../../../data/constants/TableColumns";
import {Requirement} from "../../../data/types/Requirement";
import {loadPartyParticipants} from "../../../data/sclices/PartyParticipantSlice";
import {loadPartyPendingInvites} from "../../../data/sclices/PendingInvitesForPartySlice";
import {inviteToParty, kickFromParty} from "../../../data/apis/PartyAttendanceManagerApi";
import CreateRequirementModal from "./CreateRequirementModal";
import DeleteRequirementModal from "./DeleteRequirementModal";
import backgroundImage from "../../../data/resources/images/gears.png";

const ManageParty = () => {
    const navigate = useNavigate()
    const dispatch = useDispatch<AppDispatch>()

    const [user, setUser] = useState<User>()
    const [profileOpen, setProfileOpen] = useState(false)
    const [usernameInput, setUsernameInput] = useState('')
    const [inviteFeedbackSuccess, setInviteFeedbackSuccess] = useState('')
    const [inviteFeedbackError, setInviteFeedbackError] = useState('')
    const [requirementModalVisible, setRequirementModalVisible] = useState(false)
    const [requirementModalMode, setRequirementModalMode] = useState('')
    const [deleteModalVisible, setDeleteModalVisible] = useState(false)
    const [deleteModalMode, setDeleteModalMode] = useState('')
    const [requirementToDelete, setRequirementToDelete] = useState(-1)

    const initialFetchDone = useRef(false);

    const {selectedParty} = useSelector((state: RootState)=> state.selectedPartyStore)
    const {requirements: dRequirements, loading: dReqLoading, error: dReqError} = useSelector((state: RootState) => state.drinkRequirementStore)
    const {requirements: fRequirements, loading: fReqLoading, error: fReqError} = useSelector((state: RootState) => state.foodRequirementStore)
    const {participants, loading: participantLoading, error: participantError} = useSelector((state: RootState) => state.partyParticipantStore)
    const {pendingInvites, loading: pendingInvitesLoading, error: pendingInvitesError} = useSelector((state: RootState) => state.partyPendingInviteStore)

    useEffect(() => {
        if (initialFetchDone.current) return;
        initialFetchDone.current = true;

        const currentUser = getUser()
        if(!currentUser) {
            authService.handleUnauthorized()
            return
        }

        //todo: reload these when needed
        if(!selectedParty || !selectedParty.ID) return
        dispatch(loadDrinkRequirements(selectedParty.ID))
        dispatch(loadFoodRequirements(selectedParty.ID))
        dispatch(loadPartyParticipants(selectedParty.ID))
        dispatch(loadPartyPendingInvites(selectedParty.ID))

        setUser(currentUser)
    }, [])

    if(!selectedParty || !selectedParty.ID){
        console.log("error, selected party was null")
        navigate("/overview/discover")
        return <div>Error selected party was null</div>
    }

    if(!user){
        console.log("user was null")
        return <div>Loading...</div>
    }

    const participantColumns = [...userTableColumns,
        {
            title: '',
            key: 'action 1',
            render: (record: User) => (
                <Button style={styles.errorButton} onClick={() => handleKickParticipant(record)}>Kick</Button>
            ),
        },
    ]

    const drinkRequirementColumns = [...requirementTableColumns,
        {
            title: '',
            key: 'action 1',
            render: (record: Requirement) => (
                <Button style={styles.errorButton} onClick={() => handleDeleteRequirement(record, "drink")}>Delete</Button>
            ),
        },
    ]

    const foodRequirementColumns = [...requirementTableColumns,
        {
            title: '',
            key: 'action 1',
            render: (record: Requirement) => (
                <Button style={styles.errorButton} onClick={() => handleDeleteRequirement(record, "food")}>Delete</Button>
            ),
        },
    ]

    const renderReqs = (requirements: Requirement[], mode: string) => {
        if(!requirements || requirements.length === 0){
            return <div>There's no drink requirements yet!</div>
        }
        return (<Table
            dataSource={requirements.map(req => ({...req, key: req.ID}))}
            columns={mode === 'drink' ? drinkRequirementColumns : foodRequirementColumns}
            pagination={false}
            scroll={{y: 200}}
        />)
    }

    const renderParticipants = () => {
        if(!participants || participants.length === 0){
            return <div>There's no drink requirements yet!</div>
        }
        return (<Table
            dataSource={participants.map(person => ({...person, key: person.ID}))}
            columns={participantColumns}
            pagination={false}
            scroll={{y: 200}}
        />)
    }

    const renderPendingInvites = () => {
        if(!pendingInvites || pendingInvites.length === 0){
            return <div>There's no pending invites at the moment</div>
        }
        return (<Table
            dataSource={pendingInvites.map(invite => ({...invite, key: invite.ID}))}
            columns={invitedTableColumns}
            pagination={false}
            scroll={{y: 200}}
        />)
    }

    const handleInviteToParty = (username: string) => {
        if(!selectedParty || !selectedParty.ID) return
        inviteToParty(selectedParty.ID, username)
            .then(() => {
                if(!selectedParty || !selectedParty.ID) return
                dispatch(loadPartyPendingInvites(selectedParty.ID))

                setInviteFeedbackSuccess("Invite sent!")
                setUsernameInput('')
                setTimeout(() => {
                    setInviteFeedbackSuccess('')
                }, 3000);
            })
            .catch(err => {
                setInviteFeedbackError("something went wrong")
                setUsernameInput('')
                setTimeout(() => {
                    setInviteFeedbackError('')
                }, 3000);
            });
    }

    const handleAddRequirement = (mode: string) => {
        if (mode === "drink") setRequirementModalMode("drink")
        if(mode === "food") setRequirementModalMode("food")
        setRequirementModalVisible(true)
    }

    const handleDeleteRequirement = (requirement: Requirement, mode: string) => {
        if(mode === "drink"){
            setDeleteModalMode("drink")
            setRequirementToDelete(requirement.ID || -1)
            setDeleteModalVisible(true)
        }

        if(mode === "food"){
            setDeleteModalMode("food")
            setRequirementToDelete(requirement.ID || -1)
            setDeleteModalVisible(true)
        }
    }

    const handleKickParticipant = (user: User) => {
        if(!selectedParty || !selectedParty.ID) return
        kickFromParty(selectedParty.ID, user.ID)
            .then(() => {
                if(!selectedParty || !selectedParty.ID) return
                dispatch(loadPartyParticipants(selectedParty.ID))
            })
            .catch(err => {
                //todo: make a confirmaction modal and handle feedback
            });
    }

    return (
        <div style={styles.background}>
            <div style={styles.outerContainer}>
                <ConfigProvider
                    theme={{algorithm: theme.darkAlgorithm,}}
                >
                    <VisitPartyNavBar onProfileClick={() => setProfileOpen(true)}/>
                    <VisitPartyProfile isOpen={profileOpen} onClose={() => setProfileOpen(false)} user={user} onLogout={() => {console.log("logout")}} currentParty={selectedParty} onLeaveParty={() => {}}/>
                    <CreateRequirementModal visible={requirementModalVisible} onClose={() => setRequirementModalVisible(false)} mode={requirementModalMode} />
                    <DeleteRequirementModal visible={deleteModalVisible} onClose={() => setDeleteModalVisible(false)} mode={deleteModalMode} requirementId={requirementToDelete} />
                    <div style={styles.container}>

                        <h2>Invite</h2>
                        <div style={styles.inputContainer}>
                            <input
                                type="text"
                                id="username"
                                value={usernameInput}
                                placeholder="Enter username"
                                onChange={(e) => setUsernameInput(e.target.value)}
                                style={styles.input}
                            />
                            <Button type="primary" style={styles.button}
                                    onClick={() => handleInviteToParty(usernameInput)}>Invite</Button>
                            {inviteFeedbackSuccess && <p style={styles.success}>{inviteFeedbackSuccess}</p>}
                            {inviteFeedbackError && <p style={styles.error}>{inviteFeedbackError}</p>}
                        </div>


                        <h2>Drink Requirements</h2>
                        <div style={styles.requirementContainer}>
                            <Button type="primary" style={styles.button} onClick={() => handleAddRequirement("drink")}>Add</Button>
                            <div style={styles.requirementTable}>
                                {dReqLoading && <div>Loading...</div>}
                                {dReqError && <div>Error: Some unexpected error happened</div>}
                                {(!dReqLoading && !dReqError) && renderReqs(dRequirements, "drink")}
                            </div>
                        </div>

                        <h2>Food Requirements</h2>
                        <div style={styles.requirementContainer}>
                            <Button type="primary" style={styles.button} onClick={() => handleAddRequirement("food")}>Add</Button>
                            <div style={styles.requirementTable}>
                                {fReqLoading && <div>Loading...</div>}
                                {fReqError && <div>Error: Some unexpected error happened</div>}
                                {(!fReqLoading && !fReqError) && renderReqs(fRequirements, "food")}
                            </div>
                        </div>

                        <h2>Participants</h2>
                        <div style={styles.requirementContainer}>
                            <div style={styles.requirementTable}>
                                {participantLoading && <div>Loading...</div>}
                                {participantError && <div>Error: Some unexpected error happened</div>}
                                {(!participantLoading && !participantError) && renderParticipants()}
                            </div>
                        </div>

                        <h2>Pending Invites</h2>
                        <div style={styles.requirementContainer}>
                            <div style={styles.requirementTable}>
                                {pendingInvitesLoading && <div>Loading...</div>}
                                {pendingInvitesError && <div>Error: Some unexpected error happened</div>}
                                {(!pendingInvitesLoading && !pendingInvitesError) && renderPendingInvites()}
                            </div>
                        </div>
                    </div>
                </ConfigProvider>
            </div>
        </div>
    )
}

const styles: { [key: string]: CSSProperties } = {
    background: {
        backgroundImage: `url(${backgroundImage})`,
        position: 'fixed',
        backgroundSize: 'cover',
        backgroundPosition: 'center',
        display: 'flex',
    },
    outerContainer: {
        overflowY: 'auto',
        height: '100vh',
        display: 'flex',
        flexDirection: 'column',
        color: '#ffffff',
    },
    container: {
        width: "min(80%, 1000px)",
        margin: "20px auto",
        padding: "20px",
        display: "flex",
        flexDirection: "column",
        // backgroundColor: "#2c2c2c", // Darker gray background for content box
        backgroundColor: 'rgba(33, 33, 33, 0.95)',
        borderRadius: "8px",
        boxShadow: "0 4px 8px rgba(0, 0, 0, 0.4)", // Slightly stronger shadow for depth
        color: "#007bff", // Ensure text is white for readability
    },
    h2: {
        color: "#d3d3d3", // Light gray for headings
        fontSize: "1.8rem",
        fontWeight: "bold",
        textAlign: "left",
    },
    inputContainer: {
        display: "flex",
        flexDirection: "column",
        alignItems: "flex-start",
        marginBottom: "20px",
        gap: "10px",
    },
    input: {
        padding: "8px 12px",
        fontSize: "1rem",
        borderRadius: "5px",
        border: "1px solid #444", // Darker border to blend with dark mode
        backgroundColor: "#3a3a3a", // Dark input background
        color: "#ffffff", // Light input text
        width: "60%",
    },
    button: {
        width: 'auto',
        minWidth: '120px',
        padding: '10px 20px',
        borderRadius: '5px',
        marginBottom: '20px',
        fontWeight: 'bold',
        color: '#ffffff',
        backgroundColor: '#007bff', // Accent color to match link color in header
        boxShadow: '0 4px 8px rgba(0, 0, 0, 0.2)',
    },
    errorButton: {
        backgroundColor: '#b30000', // Dark red for delete buttons in dark mode
        textAlign: 'center',
        borderRadius: '10px',
        cursor: 'pointer',
        color: '#ffffff',
        boxShadow: '0 4px 8px rgba(0, 0, 0, 0.3)',
    },
    success: {
        color: "#66ff66", // Light green for success messages
        fontSize: "1rem",
        marginTop: "5px",
    },
    error: {
        color: "#ff6666", // Light red for error messages
        fontSize: "1rem",
        marginTop: "5px",
    },
    requirementContainer: {
        marginTop: "10px",
    },
    requirementTable: {
        border: "1px solid #444",
        borderRadius: "8px",
        padding: "10px",
        backgroundColor: "#3a3a3a",
        marginBottom: "30px",
    },
    loading: {
        textAlign: "center",
        fontSize: "1rem",
        color: "#d3d3d3",
    },
    errorMessage: {
        textAlign: "center",
        fontSize: "1rem",
        color: "#ff6666",
    },
};


export default ManageParty