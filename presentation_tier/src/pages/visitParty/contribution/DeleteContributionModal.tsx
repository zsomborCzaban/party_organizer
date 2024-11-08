import React, {useEffect, useState} from 'react';
import {Modal} from 'antd';
import {ApiError} from "../../../api/ApiResponse";
import {
    deleteDrinkContribution,
    deleteFoodContribution
} from "../../../data/apis/ContributionApi";
import {useDispatch, useSelector} from "react-redux";
import {AppDispatch, RootState} from "../../../store/store";
import {loadDrinkContributions} from "../../../data/sclices/DrinkContributionSlice";
import {loadFoodContributions} from "../../../data/sclices/FoodContributionSlice";

interface DeleteContributeModalProps {
    visible: boolean;
    onClose: () => void;
    mode: string;
    contributionId: number
}

interface Feedbacks{
    buttonError?: string,
    buttonSuccess?:string
}


const DeleteContributeModal: React.FC<DeleteContributeModalProps> = ({ mode, contributionId, visible, onClose }) => {
    const [feedbacks, setFeedbacks] = useState<Feedbacks>({});
    const [countdown, setCountdown] = useState(0)

    const dispatch = useDispatch<AppDispatch>()

    const {selectedParty} = useSelector((state: RootState)=> state.selectedPartyStore)

    useEffect(() => {
        if (visible) {
            setFeedbacks({})
            setCountdown(0)
        }
    }, [visible]);


    const handleErrors = (errs: ApiError[]) => {
        //todo: implement me!!!
    }

    const handleDelete = () => {
        const newFeedbacks: Feedbacks = {};

        if(mode === "drink"){
            deleteDrinkContribution(contributionId)
                .then(() => {
                    newFeedbacks.buttonSuccess = "deleted successfully"
                    setFeedbacks(newFeedbacks)

                    if(!selectedParty || !selectedParty.ID) return
                    dispatch(loadDrinkContributions(selectedParty.ID));

                    startCloseTimer()
                    return
                })
                .catch(err => {
                    if(err.response){
                        let errors = err.response.data.errors
                        handleErrors(errors)
                    } else {
                        newFeedbacks.buttonError = "Something unexpected happened. Try again later!"
                        setFeedbacks(newFeedbacks)
                    }
                })
            return;
        }

        if(mode === "food"){
            deleteFoodContribution(contributionId)
                .then(() => {
                    newFeedbacks.buttonSuccess = "deleted successfully"
                    setFeedbacks(newFeedbacks)

                    if(!selectedParty || !selectedParty.ID) return
                    dispatch(loadFoodContributions(selectedParty.ID));

                    startCloseTimer()
                    return
                })
                .catch(err => {
                    if(err.response){
                        let errors = err.response.data.errors
                        handleErrors(errors)
                    } else {
                        newFeedbacks.buttonError = "Something unexpected happened. Try again later!"
                        setFeedbacks(newFeedbacks)
                    }
                })
            return;
        }

        console.log("unexpected modalMode")
        newFeedbacks.buttonError = "Unexpected modal mode try again later"
        setFeedbacks(newFeedbacks)
    };

    const startCloseTimer =  () => {
        let count = 5; // Start from 5 seconds

        const countdownTimer = () => {
            if (count >= 1) {
                setCountdown(count); // Update the countdown state
                count--; // Decrement the countdown

                setTimeout(countdownTimer, 1000); // Call the function again after 1 second
            } else {
                // Close the modal after countdown finishes
                onClose(); // Or set visible to false
            }
        };

        countdownTimer(); // Start the countdown
    }

    return (
        <Modal
            title="Contribute"
            open={visible}
            onCancel={onClose}
            footer={null} // Disable default footer
        >
            <div style={styles.modalContent}>
                <label style={styles.label}>Are you sure you want to delete this contribution?</label>

                <div style={styles.buttonContainer}>
                    <button onClick={handleDelete} style={styles.submitButton}>
                        Submit
                    </button>
                    <button onClick={onClose} style={styles.cancelButton}>
                        Cancel
                    </button>
                </div>
                {feedbacks.buttonError && <p style={styles.error}>{feedbacks.buttonError}</p>}
                {feedbacks.buttonSuccess && <p style={styles.success}>{feedbacks.buttonSuccess}</p>}
                {countdown !== 0 && <p>The modal will close in {countdown}...</p>}
            </div>
        </Modal>
    );
};

// Styles
const styles: { [key: string]: React.CSSProperties } = {
    modalContent: {
        display: 'flex',
        flexDirection: 'column',
        alignItems: 'flex-start',
    },
    label: {
        marginBottom: '8px',
        fontSize: '16px',
        fontWeight: 'bold',
    },
    inputField: {
        marginBottom: '16px',
        padding: '8px',
        width: '100%',
        border: '1px solid #ccc',
        borderRadius: '4px',
    },
    feedback: {
        marginBottom: '16px',
        fontSize: '14px',
        color: '#555',
    },
    buttonContainer: {
        display: 'flex',
        flexDirection: 'row',
        gap: '20px',
        width: '100%',
    },
    submitButton: {
        backgroundColor: '#007bff',
        color: '#fff',
        border: 'none',
        padding: '10px 20px',
        borderRadius: '4px',
        cursor: 'pointer',
    },
    cancelButton: {
        backgroundColor: 'red',
        color: '#fff',
        border: 'none',
        padding: '10px 20px',
        borderRadius: '4px',
        cursor: 'pointer',
    },
    error: {
        color: 'red',
        fontSize: '0.875em',
    },
    success: {
        color: 'green',
        fontSize: '0.875em',
    },
};

export default DeleteContributeModal;
