import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import classes from './Login.module.scss';
import { useApi } from '../../../context/ApiContext';
import { useAppDispatch } from '../../../store/store-helper';
import { userLogin } from '../../../store/sclices/UserSlice';

type InvalidCred = {
  err: string;
  field: string;
  value: string;
};

export const Login = () => {
  const navigate = useNavigate();
  const dispatch = useAppDispatch();

  // TODO: Remove these
  const [username, setUsername] = useState<string>('');
  const [password, setPassword] = useState<string>('');
  const [error, setError] = useState<string | null>(null);
  const handleError = (data: InvalidCred[]) => {
    setError(data[0].err);
  };

  const api = useApi();


  return (
    <div className={classes.container}>
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
          value={username}
          onChange={(e) => setUsername(e.target.value)}
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
            href='/forgot-password'
            className={classes.link}
          >
            Forgot Password?
          </a>
        </div>
        <input
          type='password'
          id='password'
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          required
          className={classes.input}
        />
      </div>

      <button
        onClick={() => dispatch(userLogin({api, password, username}))}
        className={classes.loginButton}
      >
        Sign In
      </button>

      {error && <p className={classes.error}>{error}</p>}
      
      <div className={classes.signUpContainer}>
        <p>New to the platform?</p>
        <a
          href='/register'
          className={classes.link}
        >
          Create an account
        </a>
      </div>
    </div>
  );
};
