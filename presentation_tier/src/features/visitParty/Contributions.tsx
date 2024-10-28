import Collapsible from "react-collapsible";
import VisitPartyNavBar from "../../components/navbar/VisitPartyNavBar";
import {useNavigate} from "react-router-dom";
import {useDispatch, useSelector} from "react-redux";
import {AppDispatch, RootState} from "../../store/store";
import {useEffect} from "react";
import {loadDrinkRequirements} from "./data/Slices/DrinkRequirementSlice";
import {loadFoodRequirements} from "./data/Slices/FoodRequirementSlice";
import {loadDrinkContributions} from "./data/Slices/DrinkContributionSlice";
import {loadFoodContributions} from "./data/Slices/FoodContributionSlice";
import {loadPartyParticipants} from "./data/Slices/PartyParticipantSlice";

const Contributions = () => {
    const navigate = useNavigate()
    const dispatch = useDispatch<AppDispatch>()

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

    if(!selectedParty || !selectedParty.ID){
        console.log("error, no selected party or no id of party")
        navigate("/overview/discover")
        return <div>error, selected party was null</div>
    }

    return (
        <div>
            <VisitPartyNavBar/>
            <Collapsible trigger="Start here">
                <p>
                    This is the collapsible content. It can be any element or React
                    component you like.
                </p>
                <Collapsible trigger="Start here">
                    <p>
                        This is the collapsible content. It can be any element or React
                        component you like.
                    </p>
                    <Collapsible trigger="anotherone">
                        <p>
                            This is the collapsible content. It can be any element or React
                            component you like.
                        </p>
                        <p>
                            It can even be another Collapsible component. Check out the next
                            section!
                        </p>
                    </Collapsible>
                </Collapsible>
            </Collapsible>
        </div>
    );
}

export default Contributions