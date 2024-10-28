import Collapsible from "react-collapsible";
import VisitPartyNavBar from "../../../components/navbar/VisitPartyNavBar";
import {useNavigate} from "react-router-dom";
import {useDispatch, useSelector} from "react-redux";
import {AppDispatch, RootState} from "../../../store/store";
import React, {CSSProperties, useEffect, useState} from "react";
import {loadDrinkRequirements} from "../data/slices/DrinkRequirementSlice";
import {loadFoodRequirements} from "../data/slices/FoodRequirementSlice";
import {loadDrinkContributions} from "../data/slices/DrinkContributionSlice";
import {loadFoodContributions} from "../data/slices/FoodContributionSlice";
import {loadPartyParticipants} from "../data/slices/PartyParticipantSlice";
import {User} from "../../overView/User";
import {Contribution} from "../data/Contribution";
import {Requirement} from "../data/Requirement";
import backgroundImage from "../../../midjourney_images/cola-pepsi.png";
import ContributeModal from "./ContributeModal";

const Contributions = () => {
    const navigate = useNavigate()
    const dispatch = useDispatch<AppDispatch>()

    const [participantMap, setparticipantMap] = useState<Record<string, User>>({});
    const [drinkModalVisible, setDrinkModalVisible] = useState(false)
    const [foodModalVisible, setFoodModalVisible] = useState(false)

    const {selectedParty} = useSelector((state: RootState)=> state.selectedPartyStore)

    const {requirements: drinkRequirements, loading: DReqLoading, error: DReqError} = useSelector(
        (state: RootState) => state.drinkRequirementStore
    )
    const {requirements: foodRequirements, loading: FReqLoading, error: FReqError} = useSelector(
        (state: RootState) => state.foodRequirementStore
    )
    const {contributions: drinkContributions, loading: DConLoading, error: DConError} = useSelector(
        (state: RootState) => state.drinkContributionStore
    )
    const {contributions: foodContributions, loading: FConLoading, error: FConError} = useSelector(
        (state: RootState) => state.foodContributionStore
    )
    const {participants, loading: participantsLoading, error: particiapntsError} = useSelector(
        (state: RootState) => state.partyParticipantStore
    )

    useEffect( () => {
        if(!selectedParty || !selectedParty.ID) return
        dispatch(loadDrinkRequirements(selectedParty.ID));
        dispatch(loadFoodRequirements(selectedParty.ID));
        dispatch(loadDrinkContributions(selectedParty.ID));
        dispatch(loadFoodContributions(selectedParty.ID));
        dispatch(loadPartyParticipants(selectedParty.ID))
    }, []);

    useEffect(() => {
        if(!participants) return

        const newMap = participants.reduce((pMap: Record<number, User>, participant) => {
            pMap[participant.ID] = participant;
            return pMap
        }, {})

        setparticipantMap(newMap)
    }, [participants]);

    if(!selectedParty || !selectedParty.ID){
        console.log("error, no selected party or no id of party")
        navigate("/overview/discover")
        return <div>error, selected party was null</div>
    }

    const createContributionDiv = (req: Requirement, contribution: Contribution) => {
        let contributorName//not easily readable :( = contribution.contributor_id ? participantMap[contribution.contributor_id] ? participantMap[contribution.contributor_id].username : "" : ""
        if(contribution.contributor_id && participantMap[contribution.contributor_id]){
            contributorName = participantMap[contribution.contributor_id].username
        } else {
            contributorName = ""
        }

        return <div key={contribution.ID} style={styles.contribution}> {contributorName}: {contribution.quantity} {req.quantity_mark}, {contribution.description}</div>
    }

    const handleContribute = (mode: string) => {
        if(mode === "drink") setDrinkModalVisible(true)
        if(mode === "food") setFoodModalVisible(true)
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
            <VisitPartyNavBar/>
            <ContributeModal mode="drink" options={makeOptions(drinkRequirements)} visible={drinkModalVisible} onClose={() => {setDrinkModalVisible(false)}} />
            <ContributeModal mode="food" options={makeOptions(foodRequirements)} visible={foodModalVisible} onClose={() => {setFoodModalVisible(false)}} />
            <div style={styles.container}>
                <div style={styles.outerCollapsible}>
                    <Collapsible trigger="Drinks">
                        <button style={styles.button} onClick={() => {handleContribute("drink")}}>Contribute</button>
                        {drinkRequirements.map(req => (
                            <div style={styles.collapsible}>

                                <Collapsible trigger={req.type} key={req.ID}>
                                    {drinkContributions
                                        .filter(con => con.requirement_id === req.ID)
                                        .map(contribution => {
                                            return createContributionDiv(req, contribution)
                                        })
                                    }
                                </Collapsible>
                            </div>

                        ))}
                    </Collapsible>
                </div>
                <div style={styles.outerCollapsible}>
                    <Collapsible trigger="Foods">
                        <button style={styles.button} onClick={() => { handleContribute("food") } }> Contribute </button>
                        {foodRequirements.map(req => (
                            <div style={styles.collapsible}>
                                <Collapsible trigger={req.type} key={req.ID}>
                                    {foodContributions
                                        .filter(con => con.requirement_id === req.ID)
                                        .map(contribution => {
                                            return createContributionDiv(req, contribution)
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
        marginBottom: '10px',
        background: 'linear-gradient(to top right, #8B0000, #00008B)',
        border: '1px solid #ddd',
        borderRadius: '5px',
        padding: '10px',
        color: '#007bff',
        fontWeight: 'bold',
    },
    collapsible: {
        marginBottom: '10px',
        backgroundColor: '#333333',
        border: '1px solid #ddd',
        borderRadius: '5px',
        padding: '10px',
        color: '#1E90FF',
    },
    contribution: {
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
};

export default Contributions