import React, { useEffect, useState } from 'react';
import { RegisterPostRequestProps } from '../../../api/apis/AuthenticationApi';
import classes from './Register.module.scss';
import { useApi } from '../../../context/ApiContext';
import { toast } from 'sonner';
import { useNavigate } from 'react-router-dom';

const Register = () => {
  const [formInput, setFormInput] = useState<RegisterPostRequestProps>({ username: '', email: '', password: '', confirm_password: '' });
  const api = useApi();
  const navigate = useNavigate();
  const [feedbacks, setFeedbacks] = useState<Partial<RegisterPostRequestProps>>({});
  const [passwordsMatch, setPasswordsMatch] = useState<boolean | null>(null);

  useEffect(() => {
    if (formInput.password && formInput.confirm_password) {
      setPasswordsMatch(formInput.password === formInput.confirm_password);
    } else {
      setPasswordsMatch(null);
    }
  }, [formInput]);

  // Validation logic
  const validate = (): boolean => {
    let valid = true;
    const newFeedbacks: Partial<RegisterPostRequestProps> = {};

    if (!formInput.username) {
      newFeedbacks.username = 'Username is required.';
      valid = false;
    }

    if (!formInput.password) {
      newFeedbacks.password = 'Password is required.';
      valid = false;
    }

    if (!formInput.email) {
      newFeedbacks.email = 'Email is required.';
      valid = false;
    }

    if (!formInput.confirm_password) {
      newFeedbacks.confirm_password = 'Passwords do not match.';
      valid = false;
    } else if (formInput.password !== formInput.confirm_password) {
      newFeedbacks.confirm_password = 'Passwords do not match.';
      valid = false;
    }

    setFeedbacks(newFeedbacks);
    return valid;
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    const validationResult = validate();
    if (!validationResult) {
      return;
    }

    const registerPostResult = await api.authApi.postRegister(formInput);

    if (registerPostResult === 'error') {
      toast.error('Error while creating an account');
      setFeedbacks({});
      return;
    }

    navigate('/login');
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
            value={formInput.username}
            onChange={(e) => setFormInput((prev) => ({ ...prev, username: e.target.value }))}
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
            value={formInput.email}
            onChange={(e) => setFormInput((prev) => ({ ...prev, email: e.target.value }))}
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
            value={formInput.password}
            onChange={(e) => setFormInput((prev) => ({ ...prev, password: e.target.value }))}
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
            value={formInput.confirm_password}
            onChange={(e) => setFormInput((prev) => ({ ...prev, confirm_password: e.target.value }))}
          />
          {passwordsMatch !== null && (
            <div className={`${classes.passwordMatch} ${passwordsMatch ? classes.match : classes.mismatch}`}>
              {passwordsMatch ? 'Passwords match ✓' : 'Passwords do not match ✗'}
            </div>
          )}
          {feedbacks.confirm_password && <p className={classes.error}>{feedbacks.confirm_password}</p>}
        </div>

        <button
          type='submit'
          className={classes.registerButton}
          disabled={passwordsMatch === false}
        >
          Create Account
        </button>

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
