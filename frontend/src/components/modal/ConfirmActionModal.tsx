import React, {useEffect, useState} from 'react';
import {Button, Modal} from 'antd';
import {ApiError} from '../../data/types/ApiResponseTypes';

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

interface ConfirmActionModalProps {
    visible: boolean;
    onClose: () => void;
    onContinue: () => Promise<void>;
    warningText: string
    title: string
}

interface Feedbacks{
    buttonError?: string,
    buttonSuccess?:string
}

const ConfirmActionModal: React.FC<ConfirmActionModalProps> = ({ warningText, title, visible, onClose, onContinue }) => {
    const [feedbacks, setFeedbacks] = useState<Feedbacks>({});
    const [countdown, setCountdown] = useState(0);

    useEffect(() => {
        if (visible) {
            setFeedbacks({});
            setCountdown(0);
        }
    }, [visible]);

    const startCloseTimer =  () => {
        let count = 5; // Start from 5 seconds

        const countdownTimer = () => {
            if (count >= 1) {
                setCountdown(count); // Update the countdown state
                count -= 1; // Decrement the countdown

                setTimeout(countdownTimer, 1000); // Call the function again after 1 second
            } else {
                // Close the modal after countdown finishes
                onClose(); // Or set visible to false
            }
        };

        countdownTimer(); // Start the countdown
    };


    const handleErrors = (errs: ApiError[]) => {
        console.log(errs);
        // todo: implement me!!!
    };

    const handleContinue = () => {
        const newFeedbacks: Feedbacks = {};

        onContinue()
            .then(() => {
                newFeedbacks.buttonSuccess = 'deleted successfully';
                setFeedbacks(newFeedbacks);

                startCloseTimer();
                
            })
            .catch(err => {
                if(err.response){
                    const {errors} = err.response.data;
                    handleErrors(errors);

                    newFeedbacks.buttonError = 'unhandled errors';
                    setFeedbacks(newFeedbacks);
                    startCloseTimer();
                } else {
                    newFeedbacks.buttonError = 'Something unexpected happened. Try again later!';
                    setFeedbacks(newFeedbacks);
                    startCloseTimer();
                }
                
            });
    };



    return (
      <Modal
        title={`${title}`}
        open={visible}
        onCancel={onClose}
        footer={null}
      >
        <div style={styles.modalContent}>
          <label style={styles.label}>{`${warningText}`}</label>

          <div style={styles.buttonContainer}>
            <Button onClick={handleContinue} style={styles.submitButton}>
              Continue
            </Button>
            <Button onClick={onClose} style={styles.cancelButton}>
              Cancel
            </Button>
          </div>
          {feedbacks.buttonError && <p style={styles.error}>{feedbacks.buttonError}</p>}
          {feedbacks.buttonSuccess && <p style={styles.success}>{feedbacks.buttonSuccess}</p>}
          {countdown !== 0 && <p>The modal will close in {countdown}...</p>}
        </div>
      </Modal>
    );
};



export default ConfirmActionModal;
