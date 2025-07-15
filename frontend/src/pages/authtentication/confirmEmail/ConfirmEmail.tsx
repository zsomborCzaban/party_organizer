import { useState, useEffect, useRef } from 'react';
import { useNavigate, useSearchParams } from 'react-router-dom';
import classes from './ConfirmEmail.module.scss';
import { useApi } from '../../../context/ApiContext';
import { toast } from 'sonner';
import {Switch} from "@mui/material";
import partyVideo from "../../../data/resources/videos/Subway Surfers (2024) - Gameplay [4K 9x16] No Copyright.mp4";
import screensaverVideo from "../../../data/resources/videos/screensaver.webm";

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
    });
    const [progress, setProgress] = useState(0);
    const [isProgressActive, setIsProgressActive] = useState(false);
    const progressIntervalRef = useRef<NodeJS.Timeout | null>(null);

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

    const startProgressAndConfirm = () => {
        setError('');
        setSuccess(false);
        setIsProgressActive(true);
        setProgress(0);
    };

    useEffect(() => {
        if (isProgressActive) {
            const totalDuration = 60000;
            const steps = 100;
            const interval = totalDuration / steps;
            progressIntervalRef.current = setInterval(() => {
                setProgress(prev => {
                    if (prev >= 100) {
                        if (progressIntervalRef.current) clearInterval(progressIntervalRef.current);
                        return 100;
                    }
                    return prev + 1;
                });
            }, interval);
        } else {
            if (progressIntervalRef.current) clearInterval(progressIntervalRef.current);
        }
        return () => {
            if (progressIntervalRef.current) clearInterval(progressIntervalRef.current);
        };
    }, [isProgressActive]);

    useEffect(() => {
        if (isProgressActive && progress >= 100) {
            confirmEmail();
        }
    }, [progress, isProgressActive]);

    const confirmEmail = () => {
        const hash = searchParams.get('hash');
        const username = searchParams.get('username');

        if (!hash || !username) {
            setError('Invalid or expired confirmation link');
            setIsProgressActive(false);
            setProgress(0);
            return;
        }

        setIsLoading(true);
        api.authApi.confirmEmail(username, hash)
            .then(resp => {
                if(resp === 'error'){
                    toast.error('Unexpected error')
                    setError('Failed to confirm email. Please try again.');
                    setIsProgressActive(false);
                    setProgress(0);
                    return
                }

                if (resp.is_error && resp.errors === 'Email already confirmed. try logging in!') {
                    toast.success('Email already confirmed')
                    navigate('/login')
                    return;
                }

                if(resp.is_error && resp.code !== 500){
                    setError('Invalid or expired confirmation link');
                    setIsProgressActive(false);
                    setProgress(0);
                    return;
                }

                if(resp.is_error && resp.code === 500){
                    toast.error('Unexpected error')
                    setError('Failed to confirm email. Please try again.');
                    setIsProgressActive(false);
                    setProgress(0);
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
                setIsProgressActive(false);
                setProgress(0);
            });
    };

    useEffect(() => {
        const preventScroll = (e: KeyboardEvent) => {
            if (switchStates.third && (e.code === 'Space' || e.code === 'ArrowUp' || e.code === 'ArrowDown')) {
                e.preventDefault();
            }
        };

        window.addEventListener('keydown', preventScroll);
        return () => {
            window.removeEventListener('keydown', preventScroll);
        };
    }, [switchStates.third]);

    return (
        <div>
            <div className={classes.container}>
                <h2 className={classes.title}>Confirm Email</h2>
                <>
                    <p className={classes.description}>Click the button below to confirm your email address.</p>
                    <button
                        onClick={startProgressAndConfirm}
                        className={classes.confirmButton}
                        disabled={isLoading || isProgressActive}
                    >
                        {'Confirm My Email'}
                    </button>
                    {isProgressActive && (
                        <div className={classes.progressBarWrapper}>
                            <div className={classes.progressBarBg}>
                                <div
                                    className={classes.progressBarFill}
                                    style={{ width: `${progress}%` }}
                                />
                            </div>
                            <span className={classes.progressText}>{Math.floor(progress)}%</span>
                        </div>
                    )}
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
                                src={screensaverVideo}
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
                    <div className={classes.fourthAttentionSpanContainer}>
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
                                <div className={classes.videoResponsive}>
                                    <iframe
                                        src="https://www.youtube.com/embed/nxSbhVnwdFw"
                                        frameBorder="0"
                                        allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture"
                                        allowFullScreen
                                        title="Embedded youtube"
                                        style={{ width: '100%', height: '100%', display: 'block' }}
                                    />
                                </div>
                            </div>
                        )}
                    </div>
                )}
                
                {switchStates.third && (
                    <div className={classes.fifthAttentionSpanContainer}>
                        <div className={classes.videoLabel}>
                            <p>I don't have enough attention span to wait.</p>
                            <Switch 
                                id="fifth-switch"
                                checked={switchStates.fourth}
                                onChange={() => handleSwitchChange('fourth')}
                            />
                        </div>
                        {switchStates.fourth && (
                            <a href="https://www.youtube.com/watch?v=dQw4w9WgXcQ" target="_blank">click me, what could go wrong?</a>
                        )}
                    </div>
                )}
            </div>
        </div>
    );
};