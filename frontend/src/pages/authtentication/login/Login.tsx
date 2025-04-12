import { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import classes from './Login.module.scss';
import { useApi } from '../../../context/ApiContext';
import { useAppDispatch, useAppSelector } from '../../../store/store-helper';
import { deleteLoginError, userLogin } from '../../../store/sclices/UserSlice';

export const Login = () => {
  const [loginDetails, setLoginDetails] = useState<{ username: string; password: string }>({ username: '', password: '' });
  const [loginDetailsMissing, setLoginDetailsMissing] = useState(false);
  const dispatch = useAppDispatch();
  const api = useApi();
  const { loginError: errorWhileLoginPost, isLoading, jwt } = useAppSelector((state) => state.userStore);
  const navigate = useNavigate();

  const handleLoginClicked = async () => {
    dispatch(deleteLoginError());
    if (loginDetails.password.trim() === '' || loginDetails.username.trim() === '') {
      setLoginDetailsMissing(true);
      return;
    }
    setLoginDetailsMissing(false);
    await dispatch(userLogin({ api, password: loginDetails.password, username: loginDetails.username }));
  };

  // Route to home if user is logged in
  useEffect(() => {
    if (!isLoading && !errorWhileLoginPost && jwt) {
      navigate('/');
    }
  }, [isLoading, errorWhileLoginPost, jwt, navigate]);

  return (
    <div className={classes.container}>
      {isLoading && <div className={classes.overlay}>Loading (TODO: get a loading icon here)</div>}
      <h2 className={classes.title}>Welcome Back</h2>
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
          value={loginDetails.username}
          onChange={(e) => setLoginDetails((prev) => ({ ...prev, username: e.target.value }))}
          className={classes.input}
        />
      </div>

      <div className={classes.inputGroup}>
        <div className={classes.passwordLabelContainer}>
          <label
            htmlFor='password'
            className={classes.inputLabel}
          >
            Password
          </label>
          <a
            href=''
            onClick={() => navigate('/forgot-password')}
            className={classes.link}
          >
            Forgot Password?
          </a>
        </div>
        <input
          type='password'
          id='password'
          value={loginDetails.password}
          onChange={(e) => setLoginDetails((prev) => ({ ...prev, password: e.target.value }))}
          required
          className={classes.input}
        />
      </div>
      {loginDetailsMissing && <p className={classes.error}>Both fields are required to log in</p>}
      {errorWhileLoginPost && <p className={classes.error}>Username or password is incorrect</p>}
      <button
        onClick={handleLoginClicked}
        className={classes.loginButton}
        disabled={isLoading}
      >
        Sign In
      </button>

      <div className={classes.signUpContainer}>
        <p>New to the platform?</p>
        <a
          href=''
          onClick={() => navigate('/register')}
          className={classes.link}
        >
          Create an account
        </a>
      </div>
    </div>
  );
};
