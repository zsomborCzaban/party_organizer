import { useState } from 'react';
import { useNavigate, useSearchParams } from 'react-router-dom';
import classes from './ConfirmEmail.module.scss';
import { useApi } from '../../../context/ApiContext';
import { toast } from 'sonner';
import {Switch} from "@mui/material";
import partyVideo from "../../../data/resources/videos/Subway Surfers (2024) - Gameplay [4K 9x16] No Copyright.mp4";

export const ConfirmEmail = () => {
    const api = useApi();
    const navigate = useNavigate();
    const [searchParams] = useSearchParams();
    const [isLoading, setIsLoading] = useState(false);
    const [error, setError] = useState<string>('');
    const [success, setSuccess] = useState(false);
    const [switchStates, setSwitchStates] = useState({
        first: false,
        second: false,
        third: false,
        fourth: false,
        fifth: false
    });

    const handleSwitchChange = (switchName: keyof typeof switchStates) => {
        setSwitchStates(prev => {
            const newState = { ...prev };
            newState[switchName] = !prev[switchName];
            
            // Reset only the current switch and subsequent switches when turned off
            if (!newState[switchName]) {
                const switchOrder = ['first', 'second', 'third', 'fourth', 'fifth'];
                const currentIndex = switchOrder.indexOf(switchName);
                
                // Reset only the current switch and all switches after it
                switchOrder.forEach((key, index) => {
                    if (index >= currentIndex) {
                        newState[key as keyof typeof switchStates] = false;
                    }
                });
            }
            
            return newState;
        });
    };

    const confirmEmail = () => {
        const hash = searchParams.get('hash');
        const username = searchParams.get('username');

        if (!hash || !username) {
            setError('Invalid or expired confirmation link');
            return;
        }

        setIsLoading(true);
        api.authApi.confirmEmail(username, hash)
            .then(resp => {
                if(resp === 'error'){
                    toast.error('Unexpected error')
                    setError('Failed to confirm email. Please try again.');
                    return
                }

                if (resp.is_error && resp.errors === 'Email already confirmed. try logging in!') {
                    toast.success('Email already confirmed')
                    navigate('/login')
                    return;
                }

                if(resp.is_error && resp.code !== 500){
                    setError('Invalid or expired confirmation link');
                    return;
                }

                if(resp.is_error && resp.code === 500){
                    toast.error('Unexpected error')
                    setError('Failed to confirm email. Please try again.');
                    return
                }

                setSuccess(true)
                toast.success('Email confirmed')
            })
            .catch(() => {
                toast.error('Unexpected error')
                setError('Failed to confirm email. Please try again.');
            })
            .finally(() => {
                setIsLoading(false);
            });
    };

    return (
        <div>
            <div className={classes.container}>
                <h2 className={classes.title}>Confirm Email</h2>
                <>
                    <p className={classes.description}>Click the button below to confirm your email address.</p>
                    <button
                        onClick={confirmEmail}
                        className={classes.confirmButton}
                        disabled={isLoading}
                    >
                        {'Confirm My Email'}
                    </button>
                    {(error || success) && (
                        <>
                            {error && (<p className={classes.error}>{error}</p>)}
                            {success && (<p className={classes.success}>Email confirmed successfully</p>)}
                            <div className={classes.backToLoginContainer}>
                                <a
                                    href=""
                                    onClick={(e) => {
                                        e.preventDefault();
                                        navigate('/login');
                                    }}
                                    className={classes.link}
                                >
                                    Back to Login
                                </a>
                            </div>
                        </>
                    )}
                </>
            </div>
            <div className={classes.funContainer}>
                <div className={classes.firstAttentionSpanContainer}>
                    <div className={classes.videoLabel}>
                        <p>I don't have enough attention span to wait.</p>
                        <Switch 
                            id="first-switch"
                            checked={switchStates.first}
                            onChange={() => handleSwitchChange('first')}
                        />
                    </div>
                    {switchStates.first && (
                        <div className={classes.videoContainer}>
                            <video
                                src={partyVideo}
                                autoPlay
                                loop
                                muted
                                playsInline
                                className={classes.backgroundVideo}
                            />
                        </div>
                    )}
                </div>
                
                {switchStates.first && (
                    <div className={classes.secondAttentionSpanContainer}>
                        <div className={classes.videoLabel}>
                            <p>I don't have enough attention span to wait.</p>
                            <Switch 
                                id="second-switch"
                                checked={switchStates.second}
                                onChange={() => handleSwitchChange('second')}
                            />
                        </div>
                        {switchStates.second && (
                            <div className={classes.videoContainer}>
                                <video
                                    src={partyVideo}
                                    autoPlay
                                    loop
                                    muted
                                    playsInline
                                    className={classes.backgroundVideo}
                                />
                            </div>
                        )}
                    </div>
                )}

                {switchStates.second && (
                    <div className={classes.thirdAttentionSpanContainer}>
                        <div className={classes.videoLabel}>
                            <p>I don't have enough attention span to wait.</p>
                            <Switch 
                                id="third-switch"
                                checked={switchStates.third}
                                onChange={() => handleSwitchChange('third')}
                            />
                        </div>
                        {switchStates.third && (
                            <div className={classes.videoContainer}>
                                <video
                                    src={partyVideo}
                                    autoPlay
                                    loop
                                    muted
                                playsInline
                                className={classes.backgroundVideo}
                            />
                            </div>
                        )}
                    </div>
                )}

                {switchStates.third && (
                    <div className={classes.fourthAttentionSpanContainer}>
                        <div className={classes.videoLabel}>
                            <p>I don't have enough attention span to wait.</p>
                            <Switch 
                                id="fourth-switch"
                                checked={switchStates.fourth}
                                onChange={() => handleSwitchChange('fourth')}
                            />
                        </div>
                        {switchStates.fourth && (
                            <div className={classes.videoContainer}>
                                <div className="video-responsive">
                                    <iframe
                                    width="853"
                                    height="480"
                                    src="https://www.youtube.com/embed/nxSbhVnwdFw"
                                    frameBorder="0"
                                    allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture"
                                    allowFullScreen
                                    title="Embedded youtube"
                                />
                            </div>
                        </div>
                        )}
                    </div>
                )}
                
                {switchStates.fourth && (
                    <div className={classes.fifthAttentionSpanContainer}>
                        <div className={classes.videoLabel}>
                            <p>I don't have enough attention span to wait.</p>
                            <Switch 
                                id="fifth-switch"
                                checked={switchStates.fifth}
                                onChange={() => handleSwitchChange('fifth')}
                            />
                        </div>
                        {switchStates.fifth && (
                            <a href="https://www.youtube.com/watch?v=dQw4w9WgXcQ" target="_blank">click me, trust me bro</a>
                        )}
                    </div>
                )}
            </div>
        </div>
    );
};