import React, { useState } from 'react';
import {useNavigate} from 'react-router-dom';
import { login } from '../../api/apis/AuthenticationApi';

const styles: { [key: string]: React.CSSProperties } = {
    container: {
        display: 'flex',
        justifyContent: 'center',
        alignItems: 'center',
        height: '100vh',
        backgroundColor: '#f0f0f0',
    },
    form: {
        backgroundColor: '#fff',
        padding: '20px',
        borderRadius: '8px',
        boxShadow: '0 2px 10px rgba(0,0,0,0.1)',
        width: '300px',
    },
    inputGroup: {
        marginBottom: '15px',
    },
    input: {
        width: '100%',
        padding: '8px',
        marginTop: '5px',
        borderRadius: '4px',
        border: '1px solid #ccc',
    },
    button: {
        width: '100%',
        padding: '10px',
        backgroundColor: '#007BFF',
        color: '#fff',
        border: 'none',
        borderRadius: '4px',
        cursor: 'pointer',
        marginTop: '10px',
    },
    textGroup: {
        textAlign: 'center',
        marginTop: '10px',
    },
    link: {
        color: '#007BFF',
        textDecoration: 'none',
    },
};


type InvalidCred = {
    err: string;
    field: string;
    value: string;
};


const Login: React.FC = () => {
    const [username, setUsername] = useState<string>('');
    const [password, setPassword] = useState<string>('');
    const [error, setError] = useState<string | null>(null);

    const navigate = useNavigate();


    const handleError = (data: InvalidCred[]) => {
        setError(data[0].err);
    };

    const handleLogin = (e: React.FormEvent) => {
        e.preventDefault();
        login(username, password)
            .then((responseData: InvalidCred[]) => {
                if(!responseData){
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
      <div style={styles.container}>
        <form onSubmit={handleLogin} style={styles.form}>
          <h2>Login</h2>
          <div style={styles.inputGroup}>
            <label htmlFor='username'>Username</label>
            <input
              type='text'
              id='username'
              value={username}
              onChange={(e) => setUsername(e.target.value)}
              required
              style={styles.input}
            />
          </div>

          <div style={styles.inputGroup}>
            <label htmlFor='password'>Password</label>
            <input
              type='password'
              id='password'
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              required
              style={styles.input}
            />
          </div>

          <button type='submit' style={styles.button}>Login</button>

          {error && <p style={styles.error}>{error}</p>}

          <div style={styles.textGroup}>
            <p><a href='/register' style={styles.link}>Sign Up!</a></p>
            <p><a href='/forgot-password' style={styles.link}>Forgot Password?</a></p>
          </div>
        </form>
      </div>
    );
};

export default Login;
