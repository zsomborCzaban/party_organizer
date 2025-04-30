import React, { useEffect, useState } from 'react';
import {ConfigProvider, Input, Modal, theme} from 'antd';
import { Contribution } from '../../../data/types/Contribution.ts';
import { setForTime } from '../../../data/utils/timeoutSetterUtils.ts';
import { createDrinkContribution, createFoodContribution } from '../../../api/apis/ContributionApi.ts';
import classes from './ContributeModal.module.scss';

export interface ContributeModalProps {
  visible: boolean;
  onClose: () => void;
  options?: { value: number; label: string }[];
  mode: string;
  onFoodSuccess?: () => void;
  onDrinkSuccess?: () => void;
}

interface Feedbacks {
  quantity?: string;
  description?: string;
  requirementId?: string;
  buttonError?: string;
  buttonSuccess?: string;
}

export const ContributeModal: React.FC<ContributeModalProps> = ({ mode, options, visible, onClose, onDrinkSuccess, onFoodSuccess }) => {
  const [quantity, setQuantity] = useState('');
  const [description, setDescription] = useState('');
  const [requirementId, setRequirementId] = useState(0);
  const [feedbacks, setFeedbacks] = useState<Feedbacks>({});

  useEffect(() => {
    if (visible) {
      setQuantity('');
      setDescription('');
      setRequirementId(0);
      setFeedbacks({});
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
    if (!options || !options.some((option) => option.value === requirementId)) {
      newFeedbacks.requirementId = 'choose requirement from the available options';
      valid = false;
    }

    setFeedbacks(newFeedbacks);
    return valid;
  };

  const handleErrors = (errs: string) => {
    const newFeedbacks: Feedbacks = {}
    newFeedbacks.buttonError = errs;
    setFeedbacks(newFeedbacks);
  };

  const handleContribute = () => {
    if (!validate()) return;
    const contribution: Contribution = {
      quantity: Number(quantity),
      description,
      requirement_id: requirementId,
    };
    const newFeedbacks: Feedbacks = {};

    if (mode === 'drink') {
      createDrinkContribution(contribution)
        .then(() => {
          newFeedbacks.buttonSuccess = 'created successfully';
          setForTime(setFeedbacks, newFeedbacks, {}, 3000);
          if(onDrinkSuccess){
            onDrinkSuccess();
          }
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
      createFoodContribution(contribution)
        .then(() => {
          newFeedbacks.buttonSuccess = 'created successfully';
          setForTime(setFeedbacks, newFeedbacks, {}, 3000);
          if(onFoodSuccess){
            onFoodSuccess();
          }
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
        title={mode === 'drink' ? 'Contribute drink' : 'Contribute Food'}
        open={visible}
        onCancel={onClose}
        footer={null}
        className={classes.modal}
      >
        <div className={classes.modalContent}>
          <label className={classes.label}>Select Option:</label>
          <select
            value={requirementId}
            onChange={(e) => setRequirementId(Number(e.target.value) ? Number(e.target.value) : 0)}
            className={classes.selectField}
          >
            <option value='0'>-- Please select --</option>
            {options && options.map((option) => (
              <option
                key={option.value}
                value={option.value}
              >
                {option.label}
              </option>
            ))}
          </select>
          {feedbacks.requirementId && <p className={classes.error}>{feedbacks.requirementId}</p>}

          <label className={classes.label}>Quantity:</label>
          <Input
            placeholder='Enter the contributed quantity'
            value={quantity}
            onChange={(e) => setQuantity(e.target.value)}
            className={classes.inputField}
          />
          {feedbacks.quantity && <p className={classes.error}>{feedbacks.quantity}</p>}

          <label className={classes.label}>Description:</label>
          <Input
            placeholder='brands or types'
            value={description}
            onChange={(e) => setDescription(e.target.value)}
            className={classes.inputField}
          />
          {feedbacks.description && <p className={classes.error}>{feedbacks.description}</p>}

          <div className={classes.buttonContainer}>
            <button
              onClick={handleContribute}
              className={classes.submitButton}
            >
              Submit
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
        </div>
      </Modal>
    </ConfigProvider>
  );
};
