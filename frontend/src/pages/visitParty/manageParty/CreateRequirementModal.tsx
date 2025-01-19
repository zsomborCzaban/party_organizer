import React, {useEffect, useState} from 'react';
import {Input, Modal} from 'antd';
import { useDispatch, useSelector } from 'react-redux';
import { AppDispatch, RootState } from '../../../store/store';
import { ApiError } from '../../../api/ApiResponse';
import { Requirement } from '../../../data/types/Requirement';
import { createDrinkRequirement, createFoodRequirement } from '../../../data/apis/RequirementApi';
import { setForTime } from '../../../data/utils/timeoutSetterUtils';
import { loadDrinkRequirements } from '../../../data/sclices/DrinkRequirementSlice';
import { loadFoodRequirements } from '../../../data/sclices/FoodRequirementSlice';



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

interface ContributeModalProps {
    visible: boolean;
    onClose: () => void;
    mode: string;
}

interface Feedbacks{
    type?: string,
    targetQuantity?: string,
    quantityMark?: string,
    buttonError?: string,
    buttonSuccess?:string
}


const CreateRequirementModal: React.FC<ContributeModalProps> = ({ mode, visible, onClose }) => {
    const [type, setType] = useState('');
    const [targetQuantity, setTargetQuantity] = useState('');
    const [quantityMark, setQuantityMark] = useState('');
    const [feedbacks, setFeedbacks] = useState<Feedbacks>({});

    const dispatch = useDispatch<AppDispatch>();

    const {selectedParty} = useSelector((state: RootState)=> state.selectedPartyStore);

    useEffect(() => {
        if (visible) {
            setType('');
            setTargetQuantity('');
            setQuantityMark('');
            setFeedbacks({});
        }
    }, [visible]);

    const validate = (): boolean => {
        let valid = true;
        const newFeedbacks: Feedbacks = {};

        if (!type) {
            newFeedbacks.type = 'type is required';
            valid = false;
        }
        if (!targetQuantity) {
            newFeedbacks.targetQuantity = 'targetQuantity is required';
            valid = false;
        }
        if (!Number(targetQuantity)) {
            newFeedbacks.targetQuantity = 'targetQuantity must be a number';
            valid = false;
        }
        if (!quantityMark) {
            newFeedbacks.quantityMark = 'quantityMark is required';
            valid = false;
        }

        setFeedbacks(newFeedbacks);
        return valid;
    };

    const handleErrors = (errs: ApiError[]) => {
        console.log(errs);
        // todo: implement me!!!
    };

    const handleCreateRequirement = () => {
        if(!validate()) return;
        if(!selectedParty || !selectedParty.ID) return;
        const requirement: Requirement = {
            type,
            target_quantity: Number(targetQuantity),
            quantity_mark: quantityMark,
            party_id: selectedParty.ID,
        };
        const newFeedbacks: Feedbacks = {};


        if(mode === 'drink'){
            createDrinkRequirement(requirement)
                .then((createdRequirement: Requirement) => {
                    console.log(createdRequirement);
                    newFeedbacks.buttonSuccess = 'created successfully';
                    setForTime(setFeedbacks, newFeedbacks, {}, 3000);

                    if(!selectedParty || !selectedParty.ID) return;
                    dispatch(loadDrinkRequirements(selectedParty.ID));
                    
                })
                .catch(err => {
                    if(err.response){
                        const {errors} = err.response.data;
                        handleErrors(errors);
                    } else {
                        newFeedbacks.buttonError = 'Something unexpected happened. Try again later!';
                        setFeedbacks(newFeedbacks);
                    }
                    
                });
            return;
        }

        if(mode === 'food'){
            createFoodRequirement(requirement)
                .then((createdRequirement: Requirement) => {
                    console.log(createdRequirement);
                    newFeedbacks.buttonSuccess = 'created successfully';
                    setForTime(setFeedbacks, newFeedbacks, {}, 3000);

                    if(!selectedParty || !selectedParty.ID) return;
                    dispatch(loadFoodRequirements(selectedParty.ID));
                    
                })
                .catch(err => {
                    if(err.response){
                        const {errors} = err.response.data;
                        handleErrors(errors);
                    } else {
                        newFeedbacks.buttonError = 'Something unexpected happened. Try again later!';
                        setFeedbacks(newFeedbacks);
                    }
                    
                });
            return;
        }

        console.log('unexpected modalMode');
        newFeedbacks.buttonError = 'Unexpected modal mode try again later';
        setFeedbacks(newFeedbacks);
    };

    return (
      <Modal
        title={`Create ${mode} requirement`}
        open={visible}
        onCancel={onClose}
        footer={null}
      >
        <div style={styles.modalContent}>
          <label style={styles.label}>Type:</label>
          <Input
            placeholder='(eg: "whisky", "beer", "chips"))'
            value={type}
            onChange={(e) => setType(e.target.value)}
            style={styles.inputField}
          />
          {feedbacks.type && <p style={styles.error}>{feedbacks.type}</p>}

          <div style={styles.quantityContainer}>
            <label style={styles.label}>Target Quantity:</label>
            <Input
              placeholder='Enter desired target'
              value={targetQuantity}
              onChange={(e) => setTargetQuantity(e.target.value)}
              style={styles.inputField}
            />
            {feedbacks.targetQuantity && <p style={styles.error}>{feedbacks.targetQuantity}</p>}

            <label style={styles.label}>Quantity Mark:</label>
            <Input
              placeholder='Enter quantity mark (eg: "l", "kg", "lbs")'
              value={quantityMark}
              onChange={(e) => setQuantityMark(e.target.value)}
              style={styles.inputField}
            />
            {feedbacks.quantityMark && <p style={styles.error}>{feedbacks.quantityMark}</p>}
          </div>

          <div style={styles.buttonContainer}>
            <button onClick={handleCreateRequirement} style={styles.submitButton}>
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

export default CreateRequirementModal;
