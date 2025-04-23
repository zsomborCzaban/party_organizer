import React, { useEffect, useState } from 'react';
import { Button, Modal } from 'antd';
import classes from '../deleteRequirement/DeleteRequirementModal.module.scss'
import {kickFromParty} from "../../../api/apis/PartyAttendanceManagerApi.ts";
import {User} from "../../../data/types/User.ts";

interface DeleteContributeModalProps {
    visible: boolean;
    onClose: () => void;
    user: User
    partyId: number;
}

interface Feedbacks {
    buttonError?: string;
    buttonSuccess?: string;
}

const DeleteContributeModal: React.FC<DeleteContributeModalProps> = ({ partyId, user, visible, onClose }) => {
    const [feedbacks, setFeedbacks] = useState<Feedbacks>({});
    const [countdown, setCountdown] = useState(0);

    useEffect(() => {
        if (visible) {
            setFeedbacks({});
            setCountdown(0);
        }
    }, [visible]);

    const startCloseTimer = () => {
        let count = 5;

        const countdownTimer = () => {
            if (count >= 1) {
                setCountdown(count);
                count -= 1;

                setTimeout(countdownTimer, 1000);
            } else {
                onClose();
            }
        };

        countdownTimer();
    };

    const handleErrors = (errs: string) => {
        const newFeedbacks: Feedbacks = {}
        newFeedbacks.buttonError = errs;
        setFeedbacks(newFeedbacks);
    };

    const handleDelete = () => {
        const newFeedbacks: Feedbacks = {};

        kickFromParty(partyId, user.ID)
            .then(() => {
                newFeedbacks.buttonSuccess = 'deleted successfully';
                setFeedbacks(newFeedbacks);

                startCloseTimer();
            })
            .catch(err => {
                console.log(err)
                if (err.response) {
                    const { errors } = err.response.data;
                    handleErrors(errors);
                } else {
                    newFeedbacks.buttonError = 'Unexpected error. Try again later!';
                    setFeedbacks(newFeedbacks);
                }
            })
    };

    return (
        <Modal
            title={`Kick ${user.username} from party`}
            open={visible}
            onCancel={onClose}
            footer={null}
        >
            <div className={classes.modalContent}>
                <label className={classes.label}>Are you sure you want to kick this user? All of his/her contributions will be deleted!</label>

                <div className={classes.buttonContainer}>
                    <Button
                        onClick={handleDelete}
                        className={classes.submitButton}
                    >
                        Kick
                    </Button>
                    <Button
                        onClick={onClose}
                        className={classes.cancelButton}
                    >
                        Cancel
                    </Button>
                </div>
                {feedbacks.buttonError && <p className={classes.error}>{feedbacks.buttonError}</p>}
                {feedbacks.buttonSuccess && <p className={classes.success}>{feedbacks.buttonSuccess}</p>}
                {countdown !== 0 && <p>The modal will close in {countdown}...</p>}
            </div>
        </Modal>
    );
};

export default DeleteContributeModal;
