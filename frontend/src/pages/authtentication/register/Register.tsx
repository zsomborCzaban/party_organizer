import React, { useState } from 'react';
import { RegisterRequestBody } from '../../../api/apis/AuthenticationApi';
import classes from './Register.module.scss';
import { useApi } from '../../../context/ApiContext';
import { toast } from 'sonner';
import { useNavigate } from 'react-router-dom';
import {ApiError} from "../../../data/types/ApiResponseTypes.ts";

interface Feedbacks {
  Username?: string;
  Email?: string;
  Password?: string;
  ButtonError?: string;

  [key: string]: string | undefined;
}


const Register = () => {
  const api = useApi();
  const navigate = useNavigate();
  const [username, setUsername] = useState('')
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const [confirmPassword, setConfirmPassword] = useState('')
  const [feedbacks, setFeedbacks] = useState<Feedbacks>({});
  const [isRegisterRequestLoading, setIsRegisterRequestLoading] = useState(false)

  const validate = (): boolean => {
    let valid = true;
    const newFeedbacks: Feedbacks = {};

    if (!username) {
      newFeedbacks.Username = 'Username is required.';
      valid = false;
    }

    if (!email) {
      newFeedbacks.Email = 'Email is required.';
      valid = false;
    }

    if (!password) {
      newFeedbacks.Password = 'Password is required.';
      valid = false;
    }

    setFeedbacks(newFeedbacks);
    return valid;
  };

  const isFeedbacksEmpty = () => {
    if(feedbacks.Username) return false
    if(feedbacks.Email) return false
    if(feedbacks.ButtonError) return false
    if(feedbacks.Password) return false
    return true
  }

  const handleErrors = (errs: ApiError[] | string) => {
    const newFeedbacks: Feedbacks = {
      Username: '',
      Email: '',
      Password: '',
      ButtonError: '',
    };

    console.log(errs)

    if(typeof errs === 'string'){
      newFeedbacks.ButtonError = errs
    } else {
      Array.from(errs).forEach((err) => {
        if (newFeedbacks[err.field] !== undefined) {
          newFeedbacks[err.field] = err.err;
        }
      });
    }

    setFeedbacks(newFeedbacks);
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if(!validate()) return

    setIsRegisterRequestLoading(true)
    const registerRequestBody: RegisterRequestBody = {
      username: username.trim(),
      email: email,
      password: password.trim(),
      confirm_password: confirmPassword.trim(),
    }

    api.authApi.postRegister(registerRequestBody)
        .then(resp => {
          if(resp === 'error'){
            toast.error('Unexpected error')
            return
          }

          if(resp.is_error){
            handleErrors(resp.errors)
            return;
          }

          navigate('/login')
          toast.success('Registration done, to login confirm you email first.')
        })
        .catch(() => toast.error('Unexpected error'))
        .finally(() => setIsRegisterRequestLoading(false))
  };

  return (
    <div className={classes.container}>
      <h2 className={classes.title}>Create Account</h2>
      <form onSubmit={handleSubmit}>
        <div className={classes.inputGroup}>
          <label
            htmlFor='username'
            className={classes.inputLabel}
          >
            Username
          </label>
          <input
            type='text'
            id='username'
            className={classes.input}
            value={username}
            onChange={(e) => {
              setUsername(e.target.value)
              feedbacks.Username = ''
              feedbacks.ButtonError = ''
              setFeedbacks(feedbacks)
            }}
          />
          {feedbacks.Username && <p className={classes.error}>{feedbacks.Username}</p>}
        </div>

        <div className={classes.inputGroup}>
          <label
            htmlFor='email'
            className={classes.inputLabel}
          >
            Email
          </label>
          <input
            type='email'
            id='email'
            className={classes.input}
            value={email}
            onChange={(e) => {
              setEmail(e.target.value)
              feedbacks.Email = ''
              feedbacks.ButtonError = ''
              setFeedbacks(feedbacks)
            }}
          />
          {feedbacks.Email && <p className={classes.error}>{feedbacks.Email}</p>}
        </div>

        <div className={classes.inputGroup}>
          <label
            htmlFor='password'
            className={classes.inputLabel}
          >
            Password
          </label>
          <input
            type='password'
            id='password'
            className={classes.input}
            value={password}
            onChange={(e) => {
              setPassword(e.target.value)
              feedbacks.Password = ''
              feedbacks.ButtonError = ''
              setFeedbacks(feedbacks)
            }}
          />
          {feedbacks.Password && <p className={classes.error}>{feedbacks.Password}</p>}
        </div>

        <div className={classes.inputGroup}>
          <label
            htmlFor='confirmPassword'
            className={classes.inputLabel}
          >
            Confirm Password
          </label>
          <input
            type='password'
            id='confirmPassword'
            className={classes.input}
            value={confirmPassword}
            onChange={(e) => {
              setConfirmPassword(e.target.value)
              feedbacks.ButtonError = ''
              setFeedbacks(feedbacks)
            }}
          />
          {password !== confirmPassword && <p className={classes.error}>'Passwords do not match ✗'</p>}
        </div>

        <button
          type='submit'
          className={classes.registerButton}
          disabled={password !== confirmPassword || !isFeedbacksEmpty() || isRegisterRequestLoading}
        >
          Create Account
        </button>
        {feedbacks.ButtonError && <p className={classes.error}>{feedbacks.ButtonError}</p>}

        <div className={classes.signInContainer}>
          <div className={classes.divider}>
            <span>or</span>
          </div>
          <a
            href='/login'
            className={classes.signInButton}
          >
            Sign in to existing account
          </a>
        </div>
      </form>
    </div>
  );
};

export default Register;
