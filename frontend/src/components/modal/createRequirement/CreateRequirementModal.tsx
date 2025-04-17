import React, { useEffect, useState } from 'react';
import { Input, Modal } from 'antd';
import { useDispatch, useSelector } from 'react-redux';
import { AppDispatch, RootState } from '../../../store/store.ts';
import { ApiError } from '../../../data/types/ApiResponseTypes.ts';
import { Requirement } from '../../../data/types/Requirement.ts';
import { setForTime } from '../../../data/utils/timeoutSetterUtils.ts';
import { createDrinkRequirement, createFoodRequirement } from '../../../api/apis/RequirementApi.ts';
import { loadDrinkRequirements } from '../../../store/slices/DrinkRequirementSlice.ts';
import { loadFoodRequirements } from '../../../store/slices/FoodRequirementSlice.ts';
import classes from './CreateRequirementModal.module.scss';
import {ContributeModalProps} from "../createContribution/ContributeModal.tsx";

interface Feedbacks {
  type?: string;
  targetQuantity?: string;
  quantityMark?: string;
  buttonError?: string;
  buttonSuccess?: string;
}


const CreateRequirementModal: React.FC<ContributeModalProps> = ({ mode, visible, onClose }) => {
  const [type, setType] = useState('');
  const [targetQuantity, setTargetQuantity] = useState('');
  const [quantityMark, setQuantityMark] = useState('');
  const [feedbacks, setFeedbacks] = useState<Feedbacks>({});

  const dispatch = useDispatch<AppDispatch>();

  const { selectedParty } = useSelector((state: RootState) => state.selectedPartyStore);

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
    if (!validate()) return;
    if (!selectedParty || !selectedParty.ID) return;
    const requirement: Requirement = {
      type,
      target_quantity: Number(targetQuantity),
      quantity_mark: quantityMark,
      party_id: selectedParty.ID,
    };
    const newFeedbacks: Feedbacks = {};

    if (mode === 'drink') {
      createDrinkRequirement(requirement)
        .then((createdRequirement: Requirement) => {
          console.log(createdRequirement);
          newFeedbacks.buttonSuccess = 'created successfully';
          setForTime(setFeedbacks, newFeedbacks, {}, 3000);

          if (!selectedParty || !selectedParty.ID) return;
          dispatch(loadDrinkRequirements(selectedParty.ID));
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
      createFoodRequirement(requirement)
        .then((createdRequirement: Requirement) => {
          console.log(createdRequirement);
          newFeedbacks.buttonSuccess = 'created successfully';
          setForTime(setFeedbacks, newFeedbacks, {}, 3000);

          if (!selectedParty || !selectedParty.ID) return;
          dispatch(loadFoodRequirements(selectedParty.ID));
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
      title={`Create ${mode} requirement`}
      open={visible}
      onCancel={onClose}
      footer={null}
    >
      <div className={classes.modalContent}>
        <label className={classes.label}>Type:</label>
        <Input
          placeholder='(eg: "whisky", "beer", "chips"))'
          value={type}
          onChange={(e) => setType(e.target.value)}
          className={classes.inputField}
        />
        {feedbacks.type && <p className={classes.error}>{feedbacks.type}</p>}

        <div className={classes.quantityContainer}>
          <label className={classes.label}>Target Quantity:</label>
          <Input
            placeholder='Enter desired target'
            value={targetQuantity}
            onChange={(e) => setTargetQuantity(e.target.value)}
            className={classes.inputField}
          />
          {feedbacks.targetQuantity && <p className={classes.error}>{feedbacks.targetQuantity}</p>}

          <label className={classes.label}>Quantity Mark:</label>
          <Input
            placeholder='Enter quantity mark (eg: "l", "kg", "lbs")'
            value={quantityMark}
            onChange={(e) => setQuantityMark(e.target.value)}
            className={classes.inputField}
          />
          {feedbacks.quantityMark && <p className={classes.error}>{feedbacks.quantityMark}</p>}
        </div>

        <div className={classes.buttonContainer}>
          <button
            onClick={handleCreateRequirement}
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
  );
};

export default CreateRequirementModal;
