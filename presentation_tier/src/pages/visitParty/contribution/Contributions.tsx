import Collapsible from "react-collapsible";
import VisitPartyNavBar from "../../../components/navbar/VisitPartyNavBar";
import {useNavigate} from "react-router-dom";
import {useDispatch, useSelector} from "react-redux";
import {AppDispatch, RootState} from "../../../store/store";
import React, {CSSProperties, useEffect, useRef, useState} from "react";
import {loadDrinkRequirements} from "../../../data/sclices/DrinkRequirementSlice";
import {loadFoodRequirements} from "../../../data/sclices/FoodRequirementSlice";
import {loadDrinkContributions} from "../../../data/sclices/DrinkContributionSlice";
import {loadFoodContributions} from "../../../data/sclices/FoodContributionSlice";
import {loadPartyParticipants} from "../../../data/sclices/PartyParticipantSlice";
import {User} from "../../../data/types/User";
import {Contribution} from "../../../data/types/Contribution";
import {Requirement} from "../../../data/types/Requirement";
import backgroundImage from "../../../data/resources/images/cola-pepsi.png";
import ContributeModal from "./ContributeModal";
import {getUser} from "../../../auth/AuthUserUtil";
import {authService} from "../../../auth/AuthService";
import DeleteContributionModal from "./DeleteContributionModal";
import VisitPartyProfile from "../../../components/drawer/VisitPartyProfile";

const Contributions = () => {
    const navigate = useNavigate()
    const dispatch = useDispatch<AppDispatch>()

    const [participantMap, setParticipantMap] = useState<Record<string, User>>({});
    const [dReqContributionMap, setDReqContributionMap] = useState<Record<number, number>>({})
    const [fReqContributionMap, setFReqContributionMap] = useState<Record<number, number>>({})
    const [fulfilledDReqs, setFulfilledDReqs] = useState(0)
    const [fulfilledFReqs, setFulfilledFReqs] = useState(0)
    const [dModalVisible, setDModalVisible] = useState(false)
    const [fModalVisible, setFModalVisible] = useState(false)
    const [deleteContributionModalVisible, setDeleteContributionModalVisible] = useState(false)
    const [contributionIdToDelete, setContributionIdToDelete] = useState(0)
    const [deleteMode, setDeleteMode] = useState('')
    const [user, setUser] = useState<User>()
    const [isOrganizer, setIsOrganizer] = useState(false)
    const [profileOpen, setProfileOpen] = useState(false)

    const initialFetchDone = useRef(false);

    const {selectedParty} = useSelector((state: RootState)=> state.selectedPartyStore)
    const {requirements: dRequirements, loading: dReqLoading, error: DReqError} = useSelector(
        (state: RootState) => state.drinkRequirementStore
    )
    const {requirements: fRequirements, loading: fReqLoading, error: FReqError} = useSelector(
        (state: RootState) => state.foodRequirementStore
    )
    const {contributions: dContributions, loading: dConLoading, error: DConError} = useSelector(
        (state: RootState) => state.drinkContributionStore
    )
    const {contributions: fContributions, loading: fConLoading, error: FConError} = useSelector(
        (state: RootState) => state.foodContributionStore
    )
    const {participants, loading: participantsLoading, error: particiapntsError} = useSelector(
        (state: RootState) => state.partyParticipantStore
    )

    useEffect(() => {
        if(!participants) return
        const newMap = participants.reduce((pMap: Record<number, User>, participant) => {
            pMap[participant.ID] = participant;
            return pMap
        }, {})

        setParticipantMap(newMap)
    }, [participants]);

    useEffect( () => {
        if (initialFetchDone.current) return;

        initialFetchDone.current = true;
        const currentUser = getUser()

        if(!currentUser) {
            authService.handleUnauthorized()
            return
        }

        if(!selectedParty || !selectedParty.ID) return
        dispatch(loadDrinkRequirements(selectedParty.ID));
        dispatch(loadFoodRequirements(selectedParty.ID));
        dispatch(loadDrinkContributions(selectedParty.ID));
        dispatch(loadFoodContributions(selectedParty.ID));
        dispatch(loadPartyParticipants(selectedParty.ID))

        setUser(currentUser)
        setIsOrganizer(selectedParty.organizer ? selectedParty.organizer.ID === currentUser.ID : false)

    }, []);

    useEffect(() => {
        if(dContributions.length === 0) return
        const reqContributionMap: Record<number, number> = {}
        let fulfilled = 0

        dRequirements.forEach(req => {
            dContributions.forEach(contriburion => {
                if(contriburion.requirement_id !== req.ID) return

                const oldCount = reqContributionMap[req.ID] || 0
                const newCount = oldCount + contriburion.quantity

                reqContributionMap[req.ID] = newCount
                if(newCount >= req.target_quantity && oldCount < req.target_quantity){
                    fulfilled += 1
                }
            })
        })

        setDReqContributionMap(reqContributionMap)
        setFulfilledDReqs(fulfilled)

    }, [dContributions]);

    useEffect(() => {
        if(fContributions.length === 0) return
        const reqContributionMap: Record<number, number> = {}
        let fulfilled = 0

        fRequirements.forEach(req => {
            fContributions.forEach(contriburion => {
                if(contriburion.requirement_id !== req.ID) return

                const oldCount = reqContributionMap[req.ID] || 0
                const newCount = oldCount + contriburion.quantity

                reqContributionMap[req.ID] = newCount
                if(newCount >= req.target_quantity && oldCount < req.target_quantity){
                    fulfilled += 1
                }
            })
        })

        setFReqContributionMap(reqContributionMap)
        setFulfilledFReqs(fulfilled)

    }, [fContributions]);

    if(!selectedParty || !selectedParty.ID){
        console.log("error, no selected party or no id of party")
        navigate("/overview/discover")
        return <div>error, selected party was null</div>
    }

    if(!user){
        console.log("user was null")
        return <div>Loading...</div>
    }

    const createContributionDiv = (req: Requirement, contribution: Contribution, mode: string) => {
        let contributorName//not easily readable :( = contribution.contributor_id ? participantMap[contribution.contributor_id] ? participantMap[contribution.contributor_id].username : "" : ""
        let contributorId

        if(contribution.contributor_id && participantMap[contribution.contributor_id]){
            contributorName = participantMap[contribution.contributor_id].username
            contributorId = participantMap[contribution.contributor_id].ID
        } else {
            contributorName = ""
            contributorId = 0
        }

        return <div key={contribution.ID} style={styles.contribution}>
                    <div>{contributorName}: {contribution.quantity} {req.quantity_mark}, {contribution.description}</div>
                    { (contributorId == user.ID || isOrganizer) &&
                            <div>
                                <button style={styles.deleteButton} onClick={() => {
                                    handleDeleteContribution(contribution, mode)
                                }}> Delete
                                </button>
                            </div>
                    }
                </div>
    }

    const handleContribute = (mode: string) => {
        if (mode === "drink") setDModalVisible(true)
        if(mode === "food") setFModalVisible(true)
    }

    const handleDeleteContribution = (contribution: Contribution, mode: string) => {
        if(mode === "drink"){
            setDeleteMode("drink")
            setContributionIdToDelete(contribution.ID || -1)
            setDeleteContributionModalVisible(true)
        }

        if(mode === "food"){
            setDeleteMode("food")
            setContributionIdToDelete(contribution.ID || -1)
            setDeleteContributionModalVisible(true)
        }
    }

    const makeOptions = (requirements: Requirement[]) => {
        return requirements.reduce<{ value: number, label: string }[]>((opts, requirement) => {
            opts.push({value: requirement.ID ? requirement.ID : 0, label: requirement.type + " (" + requirement.quantity_mark + ")"})
            return opts
        }, [])
    };

    return (
        <div style={styles.background}>
            <div style={styles.outerContainer}>
                <VisitPartyNavBar onProfileClick={()=> setProfileOpen(true)}/>
                <VisitPartyProfile isOpen={profileOpen} onClose={() => setProfileOpen(false)} currentParty={selectedParty} user={user} onLogout={() => console.log("logout")} onLeaveParty={() => console.log("leaveparty")} />
                <ContributeModal mode="drink" options={makeOptions(dRequirements)} visible={dModalVisible} onClose={() => {setDModalVisible(false)}} />
                <ContributeModal mode="food" options={makeOptions(fRequirements)} visible={fModalVisible} onClose={() => {setFModalVisible(false)}} />
                <DeleteContributionModal visible={deleteContributionModalVisible} onClose={() => setDeleteContributionModalVisible(false)} mode={deleteMode} contributionId={contributionIdToDelete}/>
                <div style={styles.container}>
                    <div style={styles.outerCollapsible}>
                        <Collapsible trigger={`Drinks ${fulfilledDReqs}/${dRequirements.length} fulfilled`} key={'drinks'}>
                            <button style={styles.button} onClick={() => {handleContribute("drink")}}>Contribute</button>
                            {dRequirements.map(req => (
                                <div style={styles.collapsible} key={req.ID}>
                                    <Collapsible trigger={`${req.type} ${dReqContributionMap[req.ID || -1] || 0}/${req.target_quantity} ${req.quantity_mark}`} key={req.ID}>
                                        {dContributions
                                            .filter(con => con.requirement_id === req.ID)
                                            .map(contribution => {
                                                return createContributionDiv(req, contribution, "drink")
                                            })
                                        }
                                    </Collapsible>
                                </div>

                            ))}
                        </Collapsible>
                    </div>
                    <div style={styles.outerCollapsible}>
                        <Collapsible trigger={`Foods ${fulfilledFReqs}/${fRequirements.length} fulfilled`}>
                            <button  style={styles.button} onClick={()=> { handleContribute("food") } }> Contribute </button>
                            {fRequirements.map(req => (
                                <div style={styles.collapsible} key={req.ID}>
                                    <Collapsible trigger={`${req.type} ${fReqContributionMap[req.ID || -1] || 0}/${req.target_quantity} ${req.quantity_mark}`} key={req.ID}>
                                        {fContributions
                                            .filter(con => con.requirement_id === req.ID)
                                            .map(contribution => {
                                                return createContributionDiv(req, contribution, "food")
                                            })
                                        }
                                    </Collapsible>
                                </div>
                            ))}
                        </Collapsible>
                    </div>
                </div>
            </div>
        </div>
    );
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
        width: '100vw',
        height: '100vh',
        display: 'flex',
        flexDirection: 'column',
    },
    container: {
        margin: '20px',
        display: 'flex',
        borderRadius: '20px',
        flexDirection: 'column',
        width: 'min(80%, 800px)',
        padding: '20px',
        alignSelf: 'center',
        background: 'linear-gradient(to top right, rgba(139, 0, 0, 0.8), rgba(0, 0, 139, 0.8))',
},
    outerCollapsible: {
        minHeight: '50px',
        marginBottom: '10px',
        background: 'linear-gradient(to top right, #8B0000, #00008B)',
        border: '1px solid #ddd',
        borderRadius: '5px',
        padding: '10px',
        color: '#007bff',
        fontWeight: 'bold',
    },
    collapsible: {
        minHeight: '50px',
        marginBottom: '10px',
        backgroundColor: '#333333',
        border: '1px solid #ddd',
        borderRadius: '5px',
        padding: '10px',
        color: '#1E90FF',
    },
    contribution: {
        display: 'flex',
        minHeight: '50px',
        flexDirection: 'row',
        alignItems: 'center',
        justifyContent: 'space-between',
        padding: '5px 10px',
        margin: '5px 0',
        backgroundColor: '#4A4A4A',
        borderRadius: '4px',
        boxShadow: '0 1px 2px rgba(0,0,0,0.1)',
        color: '#1E90FF',
    },
    button: {
        padding: '10px 20px',
        backgroundColor: '#007bff',
        color: 'white',
        border: 'none',
        borderRadius: '5px',
        cursor: 'pointer',
        margin: '20px 0px 20px 0px',
    },
    deleteButton: {
        padding: '5px 20px',
        backgroundColor: 'red',
        color: 'white',
        border: 'none',
        borderRadius: '5px',
        cursor: 'pointer',
    },
};

export default Contributions