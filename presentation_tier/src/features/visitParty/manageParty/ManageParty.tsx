import OverViewNavBar from "../../../components/navbar/OverViewNavBar";
import OverViewProfile from "../../../components/drawer/OverViewProfile";
import React, {CSSProperties, useEffect, useRef, useState} from "react";
import VisitPartyNavBar from "../../../components/navbar/VisitPartyNavBar";
import VisitPartyProfile from "../../../components/drawer/VisitPartyProfile";
import {User} from "../../overView/User";
import {getUser} from "../../../auth/AuthUserUtil";
import {authService} from "../../../auth/AuthService";
import {useDispatch, useSelector} from "react-redux";
import {AppDispatch, RootState} from "../../../store/store";
import {useNavigate} from "react-router-dom";
import {Button, Table} from "antd";
import {inviteFriend} from "../../overView/friends/FriendPageApi";
import {loadDrinkRequirements} from "../data/slices/DrinkRequirementSlice";
import {loadFoodRequirements} from "../data/slices/FoodRequirementSlice";
import {requirementTableColumns, userTableColumns} from "../../../constants/tableColumns/TableColumns";
import {Requirement} from "../data/Requirement";
import {PartyInvite} from "../../overView/partiesPage/PartyInvite";
import {loadPartyParticipants} from "../data/slices/PartyParticipantSlice";

const ManageParty = () => {
    const navigate = useNavigate()
    const dispatch = useDispatch<AppDispatch>()

    const [user, setUser] = useState<User>()
    const [profileOpen, setProfileOpen] = useState(false)
    const [usernameInput, setUsernameInput] = useState("")
    const [inviteFeedbackSuccess, setInviteFeedbackSuccess] = useState("")
    const [inviteFeedbackError, setInviteFeedbackError] = useState("")
    const [addDrinkReqFeedbackSuccess, setAddDrinkReqFeedbackSuccess] = useState("")
    const [addDrinkReqFeedbackError, setAddDrinkReqFeedbackError] = useState("")
    const [addFoodReqFeedbackSuccess, setAddFoodReqFeedbackSuccess] = useState("")
    const [addFoodReqFeedbackError, setAddFoodReqFeedbackError] = useState("")

    const initialFetchDone = useRef(false);

    const {selectedParty} = useSelector((state: RootState)=> state.selectedPartyStore)
    const {requirements: dRequirements, loading: dReqLoading, error: dReqError} = useSelector((state: RootState) => state.drinkRequirementStore)
    const {requirements: fRequirements, loading: fReqLoading, error: fReqError} = useSelector((state: RootState) => state.foodRequirementStore)
    const {participants, loading: participantLoading, error: participantError} = useSelector((state: RootState) => state.partyParticipantStore)

    useEffect(() => {
        if (initialFetchDone.current) return;
        initialFetchDone.current = true;

        const currentUser = getUser()
        if(!currentUser) {
            authService.handleUnauthorized()
            return
        }

        if(!selectedParty || !selectedParty.ID) return
        dispatch(loadDrinkRequirements(selectedParty.ID))
        dispatch(loadFoodRequirements(selectedParty.ID))
        dispatch(loadPartyParticipants(selectedParty.ID))

        setUser(currentUser)
    }, [])

    if(!selectedParty){
        console.log("error, selected party was null")
        navigate("/overview/discover")
        return <div>Error selected party was null</div>
    }

    if(!user){
        console.log("user was null")
        return <div>Loading...</div>
    }

    const requirementColumns = requirementTableColumns

    const participantColumns = [...userTableColumns,
        {
            title: '',
            key: 'action 1',
            render: (record: User) => (
                <Button onClick={() => handleKickParticipant(record)}>Kick</Button>
            ),
        },
    ]

    const renderReqs = (requirements: Requirement[]) => {
        if(!requirements || requirements.length === 0){
            return <div>There's no drink requirements yet!</div>
        }
        return (<Table
            dataSource={requirements.map(req => ({...req, key: req.ID}))}
            columns={requirementColumns}
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

    const handleInviteFriend = (inputUsername: string) => {
        inviteFriend(inputUsername)
            .then(() => {
                setInviteFeedbackSuccess("Invite sent!")
                setUsernameInput('')
                setTimeout(() => {
                    setInviteFeedbackSuccess("")
                }, 3000);
            })
            .catch(err => {
                setInviteFeedbackError("something went wrong")
                setUsernameInput('')
                setTimeout(() => {
                    setInviteFeedbackError("")
                }, 3000);
            });
    }

    const handleAddRequirement = () => {

    }

    const handleKickParticipant = (user: User) => {

    }

    return <div style={styles.outerContainer}>
        <VisitPartyNavBar onProfileClick={() => setProfileOpen(true)}/>
        <VisitPartyProfile isOpen={profileOpen} onClose={() => setProfileOpen(false)} user={user} onLogout={() => {console.log("logout")}} currentParty={selectedParty} onLeaveParty={() => {}}/>
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
                        onClick={() => handleInviteFriend(usernameInput)}>Invite</Button>
                {inviteFeedbackSuccess && <p style={styles.success}>{inviteFeedbackSuccess}</p>}
                {inviteFeedbackError && <p style={styles.error}>{inviteFeedbackError}</p>}
            </div>


            <h2>Drink Requirements</h2>
            <div style={styles.requirementContainer}>
                <Button type="primary" style={styles.button} onClick={() => handleAddRequirement()}>Add</Button>
                {addDrinkReqFeedbackSuccess && <p style={styles.success}>{addDrinkReqFeedbackSuccess}</p>}
                {addDrinkReqFeedbackError && <p style={styles.error}>{addDrinkReqFeedbackError}</p>}
                <div style={styles.requirementTable}>
                    {dReqLoading && <div>Loading...</div>}
                    {dReqError && <div>Error: Some unexpected error happened</div>}
                    {(!dReqLoading && !dReqError) && renderReqs(dRequirements)}
                </div>
            </div>

            <h2>Food Requirements</h2>
            <div style={styles.requirementContainer}>
                <Button type="primary" style={styles.button} onClick={() => handleAddRequirement()}>Add</Button>
                {addFoodReqFeedbackSuccess && <p style={styles.success}>{addFoodReqFeedbackSuccess}</p>}
                {addFoodReqFeedbackError && <p style={styles.error}>{addFoodReqFeedbackError}</p>}
                <div style={styles.requirementTable}>
                    {fReqLoading && <div>Loading...</div>}
                    {fReqError && <div>Error: Some unexpected error happened</div>}
                    {(!fReqLoading && !fReqError) && renderReqs(fRequirements)}
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

            {/*<h2 style={styles.label}>Attended Parties</h2>*/}

            {/*/!* Scrollable Table using Ant Design Table *!/*/}
            {/*<div style={styles.tableContainer}>*/}
            {/*    {attendedLoading && <div>Loading...</div>}*/}
            {/*    {attendedError && <div>Error: Some unexpected error happened</div>}*/}
            {/*    {(!attendedLoading && !attendedError) && renderParties("attended")}*/}
            {/*</div>*/}

            {/*<h2 style={styles.label}>Organized Parties</h2>*/}

            {/*/!* Scrollable Table using Ant Design Table *!/*/}
            {/*<div style={styles.tableContainer}>*/}
            {/*    {organizedLoading && <div>Loading...</div>}*/}
            {/*    {organizedError && <div>Error: Some unexpected error happened</div>}*/}
            {/*    {(!organizedLoading && !organizedError) && renderParties("organized")}*/}
            {/*</div>*/}
        </div>
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