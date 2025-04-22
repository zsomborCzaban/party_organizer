import { useEffect, useState } from 'react';
import { ConfigProvider, Modal, theme } from 'antd';
import { deleteDrinkContribution, deleteFoodContribution } from '../../../api/apis/ContributionApi.ts';
import classes from './DeleteContributionModal.module.scss';
import { ContributionPopulated } from '../../../data/types/Contribution.ts';
import { RequirementPopulated } from '../../../data/types/Requirement.ts';

interface DeleteContributeModalProps {
  visible: boolean;
  onClose: () => void;
  mode: string;
  contributionId: number;
  contribution: ContributionPopulated;
  requirement: RequirementPopulated;
  onDrinkSuccess: () => void;
  onFoodSuccess: () => void;
}

interface Feedbacks {
  buttonError?: string;
  buttonSuccess?: string;
}

const DeleteContributeModal: React.FC<DeleteContributeModalProps> = ({ 
  mode, 
  contributionId, 
  visible, 
  onClose,
  contribution,
  requirement,
  onFoodSuccess,
  onDrinkSuccess,
}) => {
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

    if (mode === 'drink') {
      deleteDrinkContribution(contributionId)
        .then(() => {
          newFeedbacks.buttonSuccess = 'deleted successfully';
          setFeedbacks(newFeedbacks);
          onDrinkSuccess();
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
      deleteFoodContribution(contributionId)
        .then(() => {
          newFeedbacks.buttonSuccess = 'deleted successfully';
          setFeedbacks(newFeedbacks);
          onFoodSuccess();
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

  const getConfirmationMessage = () => {
    const baseMessage = `Are you sure you want to delete the following contribution: ${requirement.type} ${contribution.quantity} ${requirement.quantity_mark}`;
    return contribution.description 
      ? `${baseMessage} (${contribution.description})?`
      : `${baseMessage}?`;
  };

  return (
    <ConfigProvider theme={{ algorithm: theme.darkAlgorithm }}>
      <Modal
        title='Delete Contribution'
        open={visible}
        onCancel={onClose}
        footer={null}
        className={classes.modal}
      >
        <div className={classes.modalContent}>
          <label className={classes.label}>{getConfirmationMessage()}</label>

          <div className={classes.buttonContainer}>
            <button
              onClick={handleDelete}
              className={classes.submitButton}
            >
              Delete
            </button>
            <button
              onClick={onClose}
              className={classes.cancelButton}
            >
              Cancel
            </button>
          </div>
          {feedbacks.buttonError && <p className={classes.error}>{feedbacks.buttonError}</p>}
          {feedbacks.buttonSuccess && <p className={classes.success}>{feedbacks.buttonSuccess}</p>}
          {countdown !== 0 && <p className={classes.countdown}>The modal will close in {countdown}...</p>}
        </div>
      </Modal>
    </ConfigProvider>
  );
};

export default DeleteContributeModal;
