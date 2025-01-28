import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { AuthApi, login } from '../../../api/apis/AuthenticationApi';
import classes from './Login.module.scss';
import axios from 'axios';

type InvalidCred = {
  err: string;
  field: string;
  value: string;
};

export const Login = () => {
  const navigate = useNavigate();
  const authApi = new AuthApi(axios);

  // TODO: Remove these
  const [username, setUsername] = useState<string>('');
  const [password, setPassword] = useState<string>('');
  const [error, setError] = useState<string | null>(null);
  const handleError = (data: InvalidCred[]) => {
    setError(data[0].err);
  };

  const handleLogin = () => {
    login(username, password)
      .then((responseData: InvalidCred[]) => {
        if (!responseData) {
          navigate('/overview/discover');
        }
        handleError(responseData);
      })
      .catch(() => {
        setError('Unexpected error, please try again');
      });
    console.log('Logging in with', { username, password });
  };

  return (
    <div className={classes.container}>
      <h2 className={classes.title}>Login</h2>
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
        <label
          htmlFor='password'
          className={classes.inputLabel}
        >
          Password
        </label>
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
        onClick={() => authApi.postLogin('sandor', 'janos')}
        className={classes.button}
      >
        Login
      </button>

      {error && <p className={classes.error}>{error}</p>}

      <div className={classes.textGroup}>
        <p>
          <a
            href='/register'
            className={classes.link}
          >
            Sign Up!
          </a>
        </p>
        <p>
          <a
            href='/forgot-password'
            className={classes.link}
          >
            Forgot Password?
          </a>
        </p>
      </div>
    </div>
  );
};
