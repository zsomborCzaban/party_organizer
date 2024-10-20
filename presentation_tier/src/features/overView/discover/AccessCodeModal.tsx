import React, {useEffect, useState} from 'react';
import { Modal } from 'antd';

interface MyModalProps {
    visible: boolean;
    onClose: () => void;
}

const AccessCodeModal: React.FC<MyModalProps> = ({ visible, onClose }) => {
    const [inputValue, setInputValue] = useState<string>('');
    const [feedback, setFeedback] = useState<string>('');

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
        // Add your submit logic here
        //todo: iff correct access code then join the user to the party and navigate him to the parties page
        if (inputValue) {
            setFeedback('Your input has been submitted!');
        } else {
            setFeedback('Please enter a value.');
        }
    };

    return (
        <Modal
            title="Join party"
            open={visible}
            onCancel={onClose}
            footer={null} // Disable default footer
        >
            <div style={styles.modalContent}>
                <label style={styles.label}>Access Code:</label>
                <input
                    type="text"
                    placeholder="Enter the accesscode"
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
        backgroundColor: '#ccc',
        color: '#000',
        border: 'none',
        padding: '10px 20px',
        borderRadius: '4px',
        cursor: 'pointer',
    },
};

export default AccessCodeModal;
