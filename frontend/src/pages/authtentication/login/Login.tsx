import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import classes from './Login.module.scss';
import { useApi } from '../../../context/ApiContext';
import {toast} from "sonner";
import {setUserJwt} from "../../../store/slices/UserSlice.ts";
import {useAppDispatch} from "../../../store/store-helper.ts";

interface Feedbacks {
  Username?: string;
  Password?: string;
  buttonError?: string;

  [key: string]: string | undefined;
}

export const Login = () => {
    const api = useApi();
    const navigate = useNavigate();
    const dispatch = useAppDispatch();
    const [username, setUsername] = useState('')
    const [password, setPassword] = useState('')
    const [feedbacks, setFeedbacks] = useState<Feedbacks>({});
    const [isLoginLoading, setIsLoginLoading] = useState(false)

    const validate = () => {
        const newFeedbacks: Feedbacks = {buttonError: '', Password: '', Username: ''}
        let isValid = true

        if(username.trim() === '') {
            newFeedbacks.Username = 'Username is missing'
            isValid = false
        }
        if(password.trim() === ''){
            newFeedbacks.Password = 'Password is missing'
            isValid = false
        }

        setFeedbacks(newFeedbacks)
        return isValid
    }

    const handleLoginClicked = async () => {
        if(!validate()) return

        setIsLoginLoading(true)
        api.authApi.postLogin(username.trim(), password.trim())
            .then(resp => {
                if(resp === 'error'){
                    toast.error('Unexpected error')
                    return
                }

                if(resp.is_error){
                    setFeedbacks({buttonError: resp.errors.toString()})
                    return;
                }

                //todo: login the user
                dispatch(setUserJwt(resp.data.jwt))
            })
            .catch(() => {
              toast.error('Unexpected error')
            })
            .finally(() => setIsLoginLoading(false))
        };

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
                {feedbacks.Username && (
                    <p className={classes.error}>{feedbacks.Username}</p>
                )}
            </div>

            <div className={classes.inputGroup}>
                <div className={classes.passwordLabelContainer}>
                    <label
                        htmlFor='password'
                        className={classes.inputLabel}
                    >
                        Password
                    </label>
                </div>
                <input
                    type='password'
                    id='password'
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                    required
                    className={classes.input}
                />
                {feedbacks.Password && (
                    <p className={classes.error}>{feedbacks.Password}</p>
                )}
            </div>
            <a
                href=''
                onClick={() => navigate('/forgot-password')}
                className={classes.link}
            >
                Forgot Password?
            </a>
            <button
                onClick={handleLoginClicked}
                className={classes.loginButton}
                disabled={isLoginLoading}
            >
                Sign In
            </button>
            {feedbacks.buttonError && (
                <p className={classes.error}>{feedbacks.buttonError}</p>
            )}

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
