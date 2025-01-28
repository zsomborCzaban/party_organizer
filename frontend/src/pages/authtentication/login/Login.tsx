import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { login } from '../../../api/apis/AuthenticationApi';
import classes from './Login.module.scss';

type InvalidCred = {
  err: string;
  field: string;
  value: string;
};

export const Login = () => {
  const navigate = useNavigate();

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
        onClick={() => handleLogin()}
        className={classes.loginButton}
      >
        Login
      </button>

      {/* Remove this form of this */}
      {error && <p className={classes.error}>{error}</p>}
      <div className={classes.signUpContainer}>
        <p>Don't have an account yet?</p>
        <a
          href='/register'
          className={classes.link}
        >
          Sign Up
        </a>
      </div>
    </div>
  );
};
