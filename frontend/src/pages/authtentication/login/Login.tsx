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
  ButtonError?: string;

  [key: string]: string | undefined;
}

export const Login = () => {
    const api = useApi();
    const navigate = useNavigate();
    const dispatch = useAppDispatch();
    const [username, setUsername] = useState('')
    const [password, setPassword] = useState('')
    const [feedbacks, setFeedbacks] = useState<Feedbacks>({});
    const [isLoginRequestLoading, setIsLoginRequestLoading] = useState(false)

    const validate = () => {
        const newFeedbacks: Feedbacks = {ButtonError: '', Password: '', Username: ''}
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

        setIsLoginRequestLoading(true)
        api.authApi.postLogin(username.trim(), password.trim())
            .then(resp => {
                if(resp === 'error'){
                    toast.error('Unexpected error')
                    return
                }

                if(resp.is_error){
                    setFeedbacks({ButtonError: resp.errors.toString()})
                    return;
                }

                dispatch(setUserJwt(resp.data.jwt))
                navigate('/')
            })
            .catch(() => {
              toast.error('Unexpected error')
            })
            .finally(() => setIsLoginRequestLoading(false))
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
                    onChange={(e) => {
                        setUsername(e.target.value)
                        feedbacks.Username = ''
                        feedbacks.ButtonError = ''
                        setFeedbacks(feedbacks)
                    }}
                    className={classes.input}
                />
                {feedbacks.Username && (
                    <p className={classes.error}>{feedbacks.Username}</p>
                )}
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
                    onChange={(e) => {
                        setPassword(e.target.value)
                        feedbacks.Password = ''
                        feedbacks.ButtonError = ''
                        setFeedbacks(feedbacks)
                    }}
                    required
                    className={classes.input}
                />
                {feedbacks.Password && (
                    <p className={classes.error}>{feedbacks.Password}</p>
                )}
            </div>
            <a
                href=''
                onClick={() => navigate('/forgotPassword')}
                className={classes.link}
            >
                Forgot Password?
            </a>
            <button
                onClick={handleLoginClicked}
                className={classes.loginButton}
                disabled={isLoginRequestLoading}
            >
                Sign In
            </button>
            {feedbacks.ButtonError && (
                <p className={classes.error}>{feedbacks.ButtonError}</p>
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
