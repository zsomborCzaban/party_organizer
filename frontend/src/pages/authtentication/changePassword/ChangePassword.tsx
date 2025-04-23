import { useState } from 'react';
import { useNavigate, useSearchParams } from 'react-router-dom';
import classes from './ChangePassword.module.scss';
import { useApi } from '../../../context/ApiContext';
import { toast } from 'sonner';
import {ApiError} from "../../../data/types/ApiResponseTypes.ts";

interface Feedbacks {
  Password?: string;
  ButtonError?: string;
  [key: string]: string | undefined;
}

export const ChangePassword = () => {
  const api = useApi();
  const navigate = useNavigate();
  const [searchParams] = useSearchParams();
  const [password, setPassword] = useState('');
  const [confirmPassword, setConfirmPassword] = useState('');
  const [feedbacks, setFeedbacks] = useState<Feedbacks>({});
  const [isLoading, setIsLoading] = useState(false);

  const validate = (): boolean => {
    const newFeedbacks: Feedbacks = {};
    let isValid = true;

    if (!password) {
      newFeedbacks.Password = 'Password is required';
      isValid = false;
    }

    if (password !== confirmPassword) {
      newFeedbacks.ConfirmPassword = 'Passwords do not match';
      isValid = false;
    }

    setFeedbacks(newFeedbacks);
    return isValid;
  };

  const isFeedbacksEmpty = () => {
    if(feedbacks.Password) return false
    return true
  }

  const handleErrors = (errs: ApiError[] | string) => {
    const newFeedbacks: Feedbacks = {
      Password: '',
      ButtonError: '',
    };

    if(typeof errs === 'string'){
      newFeedbacks.ButtonError = 'Invalid or expired change password link'
    } else {
      Array.from(errs).forEach((err) => {
        if (newFeedbacks[err.field] !== undefined) {
          newFeedbacks[err.field] = err.err;
        }
      });
    }

    setFeedbacks(newFeedbacks);
  };

  const handleChangePassword = () => {
    if (!validate()) return;

    const xzs = searchParams.get('xzs');

    if (!xzs) {
      setFeedbacks({ButtonError: 'Invalid or expired change password link'})
      return;
    }

    setIsLoading(true);
    api.userApi.changePassword({password: password, confirm_password: confirmPassword}, xzs)
        .then(resp => {
          if (resp === 'error') {
            toast.error('Unexpected error occurred');
            return;
          }

          if(resp.is_error){
            if(resp.code === 500){
              toast.error('Unexpected error occurred');
              return;
            }

            handleErrors(resp.errors)
            return;
          }

          toast.success('Password changed')
          navigate('/login')
        })
        .catch(() => {
          toast.error('Unexpected error occurred');
          return;
        })
        .finally(() => setIsLoading(false))
  };

  return (
    <div className={classes.container}>
      <h2 className={classes.title}>Change Password</h2>
      <p className={classes.description}>Enter your new password below.</p>
        <div className={classes.inputGroup}>
          <label htmlFor="password" className={classes.inputLabel}>
            New Password
          </label>
          <input
            type="password"
            id="password"
            className={classes.input}
            value={password}
            onChange={(e) => {
              setPassword(e.target.value);
              setFeedbacks({});
            }}
            placeholder="Enter your new password"
          />
          {feedbacks.Password && <p className={classes.error}>{feedbacks.Password}</p>}
        </div>

        <div className={classes.inputGroup}>
          <label htmlFor="confirmPassword" className={classes.inputLabel}>
            Confirm Password
          </label>
          <input
            type="password"
            id="confirmPassword"
            className={classes.input}
            value={confirmPassword}
            onChange={(e) => {
              setConfirmPassword(e.target.value);
              setFeedbacks({});
            }}
            placeholder="Confirm your new password"
          />
          {password !== confirmPassword && <p className={classes.errorLast}>'Passwords do not match ✗'</p>}
        </div>

        <button
          onClick={handleChangePassword}
          className={classes.submitButton}
          disabled={password !== confirmPassword || !isFeedbacksEmpty() || isLoading}
        >
          Change Password
        </button>

        {feedbacks.ButtonError && (
          <p className={classes.error}>{feedbacks.ButtonError}</p>
        )}

        <div className={classes.backToLoginContainer}>
          <p>Remember your password?</p>
          <a
            href=""
            onClick={(e) => {
              e.preventDefault();
              navigate('/login');
            }}
            className={classes.link}
          >
            Back to Login
          </a>
        </div>
    </div>
  );
};