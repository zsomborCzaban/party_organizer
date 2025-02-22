import React, { useState, useEffect } from 'react';
import { AxiosError } from 'axios';
import { ApiError, ApiResponse } from '../../../data/types/ApiResponseTypes';
import { register, RegisterRequestBody } from '../../../api/apis/AuthenticationApi';
import classes from './Register.module.scss';

interface Feedbacks {
  username?: string;
  password?: string;
  email?: string;
  confirmPassword?: string;
  buttonError?: string;
  buttonSuccess?: string;
}

const Register: React.FC = () => {
  const [username, setUsername] = useState<string>('');
  const [password, setPassword] = useState<string>('');
  const [email, setEmail] = useState('');
  const [confirmPassword, setConfirmPassword] = useState<string>('');
  const [feedbacks, setFeedbacks] = useState<Feedbacks>({});
  const [passwordsMatch, setPasswordsMatch] = useState<boolean | null>(null);

  useEffect(() => {
    if (password && confirmPassword) {
      setPasswordsMatch(password === confirmPassword);
    } else {
      setPasswordsMatch(null);
    }
  }, [password, confirmPassword]);

  // Validation logic
  const validate = (): boolean => {
    let valid = true;
    const newFeedbacks: Feedbacks = {};

    if (!username) {
      newFeedbacks.username = 'Username is required.';
      valid = false;
    }

    if (!password) {
      newFeedbacks.password = 'Password is required.';
      valid = false;
    }

    if (!email) {
      newFeedbacks.email = 'Email is required.';
      valid = false;
    }

    if (!confirmPassword) {
      newFeedbacks.confirmPassword = "'Passwords do not match.";
      valid = false;
    } else if (password !== confirmPassword) {
      newFeedbacks.confirmPassword = 'Passwords do not match.';
      valid = false;
    }

    setFeedbacks(newFeedbacks);
    return valid;
  };

  const handleErrors = (errs: ApiError[]) => {
    let newFeedbacks: Feedbacks = {};
    const keysOfFeedbacks = ['username', 'password', 'confirmpassword', 'email'];

    errs.forEach((err: ApiError) => {
      const key: keyof Feedbacks = err.field.toLowerCase();
      if (keysOfFeedbacks.includes(key)) {
        newFeedbacks[key] = err.err;
      } else {
        const unexpectedErr: Feedbacks = {};
        unexpectedErr.buttonError = 'unexpected error while registering';
        newFeedbacks = unexpectedErr;
      }
    });
    setFeedbacks(newFeedbacks);
  };

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();

    if (validate()) {
      const registerBody: RegisterRequestBody = {
        username,
        email,
        password,
        confirm_password: confirmPassword,
      };
      register(registerBody)
        .then(() => {
          setFeedbacks({ buttonSuccess: 'registered successfully' });
        })
        .catch((err: AxiosError<ApiResponse<void>>) => {
          if (err.response) {
            handleErrors(err.response.data.errors);
          } else {
            setFeedbacks({ buttonError: 'unexpected error while registering' });
          }
        });
    }
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
            onChange={(e) => setUsername(e.target.value)}
          />
          {feedbacks.username && <p className={classes.error}>{feedbacks.username}</p>}
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
            onChange={(e) => setEmail(e.target.value)}
          />
          {feedbacks.email && <p className={classes.error}>{feedbacks.email}</p>}
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
            onChange={(e) => setPassword(e.target.value)}
          />
          {feedbacks.password && <p className={classes.error}>{feedbacks.password}</p>}
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
            onChange={(e) => setConfirmPassword(e.target.value)}
          />
          {passwordsMatch !== null && (
            <div className={`${classes.passwordMatch} ${passwordsMatch ? classes.match : classes.mismatch}`}>
              {passwordsMatch ? 'Passwords match ✓' : 'Passwords do not match ✗'}
            </div>
          )}
          {feedbacks.confirmPassword && <p className={classes.error}>{feedbacks.confirmPassword}</p>}
        </div>

        <button
          type='submit'
          className={classes.registerButton}
          disabled={passwordsMatch === false}
        >
          Create Account
        </button>

        {feedbacks.buttonError && <p className={classes.error}>{feedbacks.buttonError}</p>}
        {feedbacks.buttonSuccess && <p className={classes.success}>{feedbacks.buttonSuccess}</p>}

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
