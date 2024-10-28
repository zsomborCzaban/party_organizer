import React, {useEffect, useState} from 'react';
import {Input, Modal} from 'antd';
import {ApiError} from "../../../api/ApiResponse";
import {Contribution} from "../data/Contribution";
import {createDrinkContribution, createFoodContribution} from "../data/VisitPartyApi";

interface ContributeModalProps {
    visible: boolean;
    onClose: () => void;
    options: { value: number; label: string }[];
    mode: string;
}

interface Feedbacks{
    quantity?: string,
    description?: string,
    requirementId?: string,
    buttonError?: string,
    buttonSuccess?:string
}


const ContributeModal: React.FC<ContributeModalProps> = ({ mode, options, visible, onClose }) => {
    const [quantity, setQuantity] = useState('');
    const [description, setDescription] = useState('');
    const [requirementId, setRequirementId] = useState(0);
    const [feedbacks, setFeedbacks] = useState<Feedbacks>({});

    useEffect(() => {
        if (visible) {
            setQuantity('')
            setDescription('')
            setRequirementId(0)
            setFeedbacks({})
        }
    }, [visible]);

    const validate = (): boolean => {
        let valid = true;
        const newFeedbacks: Feedbacks = {};

        if (!quantity) {
            newFeedbacks.quantity = 'quantity is required';
            valid = false;
        }
        if (!Number(quantity)) {
            newFeedbacks.quantity = 'quantity must be a number';
            valid = false;
        }
        if (!requirementId) {
            newFeedbacks.requirementId = 'requirement is required';
            valid = false;
        }
        if (!options.some(option => option.value === requirementId)) {
            newFeedbacks.requirementId = 'choose requirement from the available options';
            valid = false;
        }

        setFeedbacks(newFeedbacks)
        return valid
    }

    const handleErrors = (errs: ApiError[]) => {
        //todo: implement me!!!
    }

    const handleContribute = () => {
        if(!validate()) return
        const contribution: Contribution = {
            quantity: Number(quantity),
            description: description,
            requirement_id: requirementId,
        }
        const newFeedbacks: Feedbacks = {};


        if(mode === "drink"){
            createDrinkContribution(contribution)
                .then(createdContribution => {
                    newFeedbacks.buttonSuccess = "created successfully"
                    setFeedbacks(newFeedbacks)
                    //todo: reload contributions
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
            createFoodContribution(contribution)
                .then(createdContribution => {
                    newFeedbacks.buttonSuccess = "created successfully"
                    setFeedbacks(newFeedbacks)
                    //todo: reload contributions
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

    return (
        <Modal
            title="Contribute"
            open={visible}
            onCancel={onClose}
            footer={null} // Disable default footer
        >
            <div style={styles.modalContent}>
                <label style={styles.label}>Select Option:</label>
                <select value={requirementId} onChange={(e) => setRequirementId(Number(e.target.value) ? Number(e.target.value) : 0)} style={styles.selectField}>
                    <option value='0'>-- Please select --</option>
                    {options.map(option => (
                        <option key={option.value} value={option.value}>
                            {option.label}
                        </option>
                    ))}
                </select>
                {feedbacks.requirementId && <p style={styles.error}>{feedbacks.requirementId}</p>}


                <label style={styles.label}>Quantity:</label>
                <Input
                    placeholder="Enter the contributed quantity"
                    value={quantity} // Ensure controlled input
                    onChange={(e) => setQuantity(e.target.value)}
                    style={styles.inputField}
                />
                {feedbacks.quantity && <p style={styles.error}>{feedbacks.quantity}</p>}


                <label style={styles.label}>Description:</label>
                <Input
                    placeholder="brands or types"
                    value={description}
                    onChange={(e) => setDescription(e.target.value)}
                    style={styles.inputField}
                />
                {feedbacks.description && <p style={styles.error}>{feedbacks.description}</p>}

                <div style={styles.buttonContainer}>
                    <button onClick={handleContribute} style={styles.submitButton}>
                        Submit
                    </button>
                    <button onClick={onClose} style={styles.cancelButton}>
                        Cancel
                    </button>
                </div>
                {feedbacks.buttonError && <p style={styles.error}>{feedbacks.buttonError}</p>}
                {feedbacks.buttonSuccess && <p style={styles.success}>{feedbacks.buttonSuccess}</p>}
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

export default ContributeModal;
