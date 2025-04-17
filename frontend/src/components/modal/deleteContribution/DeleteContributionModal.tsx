import { useEffect, useState } from 'react';
import { ConfigProvider, Modal, theme } from 'antd';
import { useDispatch, useSelector } from 'react-redux';
import { AppDispatch, RootState } from '../../../store/store.ts';
import { ApiError } from '../../../data/types/ApiResponseTypes.ts';
import { deleteDrinkContribution, deleteFoodContribution } from '../../../api/apis/ContributionApi.ts';
import { loadDrinkContributions } from '../../../store/sclices/DrinkContributionSlice.ts';
import { loadFoodContributions } from '../../../store/sclices/FoodContributionSlice.ts';
import classes from './DeleteContributionModal.module.scss';

interface DeleteContributeModalProps {
  visible: boolean;
  onClose: () => void;
  mode: string;
  contributionId: number;
}

interface Feedbacks {
  buttonError?: string;
  buttonSuccess?: string;
}

const DeleteContributeModal: React.FC<DeleteContributeModalProps> = ({ mode, contributionId, visible, onClose }) => {
  const [feedbacks, setFeedbacks] = useState<Feedbacks>({});
  const [countdown, setCountdown] = useState(0);

  const dispatch = useDispatch<AppDispatch>();
  const { selectedParty } = useSelector((state: RootState) => state.selectedPartyStore);

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

  const handleErrors = (errs: ApiError[]) => {
    console.log(errs);
    // todo: implement me!!!
  };

  const handleDelete = () => {
    const newFeedbacks: Feedbacks = {};

    if (mode === 'drink') {
      deleteDrinkContribution(contributionId)
        .then(() => {
          newFeedbacks.buttonSuccess = 'deleted successfully';
          setFeedbacks(newFeedbacks);

          if (!selectedParty || !selectedParty.ID) return;
          dispatch(loadDrinkContributions(selectedParty.ID));

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

          if (!selectedParty || !selectedParty.ID) return;
          dispatch(loadFoodContributions(selectedParty.ID));

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
    <ConfigProvider theme={{ algorithm: theme.darkAlgorithm }}>
      <Modal
        title='Delete Contribution'
        open={visible}
        onCancel={onClose}
        footer={null}
        className={classes.modal}
        bodyStyle={{ backgroundColor: 'rgba(33, 33, 33, 0.95)' }}
      >
        <div className={classes.modalContent}>
          <label className={classes.label}>Are you sure you want to delete this contribution?</label>

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
