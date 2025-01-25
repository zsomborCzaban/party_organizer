import React, {useEffect, useState} from 'react';
import { Modal } from 'antd';
import { useNavigate } from 'react-router-dom';
import { setSelectedParty } from '../../../data/sclices/PartySlice';
import { setForTime } from '../../../data/utils/timeoutSetterUtils';
import { Party } from '../../../data/types/Party';
import { joinPrivateParty } from '../../../api/apis/PartyAttendanceManagerApi';

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
        color: 'red',
    },
    buttonContainer: {
        display: 'flex',
        justifyContent: 'space-between',
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
};

interface MyModalProps {
    visible: boolean;
    onClose: () => void;
}

const AccessCodeModal: React.FC<MyModalProps> = ({ visible, onClose }) => {
    const [inputValue, setInputValue] = useState<string>('');
    const [feedback, setFeedback] = useState<string>('');

    const navigate = useNavigate();

    useEffect(() => {
        if (visible) {
            setFeedback(''); // Reset feedback when the modal opens
            setInputValue(''); // Optional: Reset input value when modal opens
        }
    }, [visible]);

    const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        setInputValue(e.target.value);
    };

    const handleSubmit = () => {
        if (inputValue) {
            joinPrivateParty(inputValue)
                .then((party: Party) => {
                    console.log(party);
                    setSelectedParty(party);
                    navigate('/visitParty/partyHome');
                    
                })
                .catch((err) => {
                    if(err.response){
                        const {errors} = err.response.data;
                        setForTime<string>(setFeedback, errors[0].err, '', 4000);
                    } else {
                        setFeedback('Something unexpected happened. Try again later!');
                    }
                    
                });
        } else {
            setForTime<string>(setFeedback, 'Enter an access code', '', 4000);
        }
    };

    return (
      <Modal
        title='Join party'
        open={visible}
        onCancel={onClose}
        footer={null}
      >
        <div style={styles.modalContent}>
          <label style={styles.label}>Access Code:</label>
          <input
            type='text'
            placeholder='Enter the accesscode'
            value={inputValue}
            onChange={handleInputChange}
            style={styles.inputField}
          />
          <div style={styles.feedback}>{feedback}</div>
          <div style={styles.buttonContainer}>
            <button onClick={handleSubmit} style={styles.submitButton}>
              Submit
            </button>
            <button onClick={onClose} style={styles.cancelButton}>
              Cancel
            </button>
          </div>
        </div>
      </Modal>
    );
};

export default AccessCodeModal;
