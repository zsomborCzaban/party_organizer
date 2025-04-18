import React, { useEffect, useState } from 'react';
import { Button, Modal } from 'antd';
import { ApiError } from '../../../data/types/ApiResponseTypes.ts';
import { deleteDrinkRequirement, deleteFoodRequirement } from '../../../api/apis/RequirementApi.ts';
import classes from './DeleteRequirementModal.module.scss'

// todo: instead of creating a delete modal for each delete craete a confirmAction modal, that gets the function to call if confirmed


interface DeleteContributeModalProps {
  visible: boolean;
  onClose: () => void;
  mode: string;
  requirementId: number;
}

interface Feedbacks {
  buttonError?: string;
  buttonSuccess?: string;
}

const DeleteContributeModal: React.FC<DeleteContributeModalProps> = ({ mode, requirementId, visible, onClose }) => {
  const [feedbacks, setFeedbacks] = useState<Feedbacks>({});
  const [countdown, setCountdown] = useState(0);

  useEffect(() => {
    if (visible) {
      setFeedbacks({});
      setCountdown(0);
    }
  }, [visible]);

  const startCloseTimer = () => {
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

  const handleDelete = () => {
    const newFeedbacks: Feedbacks = {};

    if (mode === 'drink') {
      deleteDrinkRequirement(requirementId)
        .then(() => {
          newFeedbacks.buttonSuccess = 'deleted successfully';
          setFeedbacks(newFeedbacks);

          startCloseTimer();
        })
        .catch((err) => {
          if (err.response) {
            const { errors } = err.response.data;
            handleErrors(errors);
          } else {
            newFeedbacks.buttonError = 'Something unexpected happened. Try again later!';
            setFeedbacks(newFeedbacks);
          }
        });
      return;
    }

    if (mode === 'food') {
      deleteFoodRequirement(requirementId)
        .then(() => {
          newFeedbacks.buttonSuccess = 'deleted successfully';
          setFeedbacks(newFeedbacks);

          startCloseTimer();
        })
        .catch((err) => {
          if (err.response) {
            const { errors } = err.response.data;
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
      title={`Delete ${mode} requirement`}
      open={visible}
      onCancel={onClose}
      footer={null}
    >
      <div className={classes.modalContent}>
        <label className={classes.label}>Are you sure you want to delete this requirement? All of its contributions will be deleted too!</label>

        <div className={classes.buttonContainer}>
          <Button
            onClick={handleDelete}
            className={classes.submitButton}
          >
            Submit
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
