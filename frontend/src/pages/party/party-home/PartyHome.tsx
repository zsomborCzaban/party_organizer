import {useApi} from "../../../context/ApiContext.ts";
import {useNavigate} from "react-router-dom";
import {useEffect, useState} from "react";
import {EMPTY_PARTY_POPULATED, PartyPopulated} from "../../../data/types/Party.ts";
import {toast} from "sonner";
import {useAppSelector} from "../../../store/store-helper.ts";
import {isUserLoggedIn} from "../../../store/slices/UserSlice.ts";
import partyVideo from '../../../data/resources/videos/party_video.mp4';
import classes from "./PartyHome.module.scss"
import Map from '@mui/icons-material/Map';
import FacebookIcon from '@mui/icons-material/Facebook';
import WhatsAppIcon from '@mui/icons-material/WhatsApp';

export const PartyHome = ()=>{

    const api = useApi()
    const navigate = useNavigate()
    const userLoggedIn = useAppSelector(isUserLoggedIn);
    const [party, setParty] = useState<PartyPopulated>(EMPTY_PARTY_POPULATED)
    const [timeRemaining, setTimeRemaining] = useState<number>(0);

    useEffect(() => {
        const getPartyId = localStorage.getItem('partyId')
        if(!getPartyId || !Number(getPartyId)){
            toast.error('Navigation error')
            navigate('/')
            return
        }

        const receivedPartyId = Number(getPartyId)
        if(userLoggedIn){
            api.partyApi.getParty(receivedPartyId)
                .then(result => {
                    if(result === 'error'){
                        toast.error('Unable to load party')
                        return
                    }
                    if(result === 'private party'){
                        toast.error('Cannot access that party')
                        navigate('/')
                        return
                    }
                    setParty(result.data)
                })
                .catch(() => {
                    toast.error('Unexpected error')
                })
        } else {
            api.partyApi.getPartyUnauthenticated(receivedPartyId)
                .then(result => {
                    if(result === 'error'){
                        toast.error('Unable to load party')
                        return
                    }
                    if(result === 'private party'){
                        toast.error('Cannot access that party')
                        navigate('/')
                        return
                    }
                    setParty(result.data)
                })
                .catch(() => {
                    toast.error('Unexpected error')
                })
        }
        
    }, [api.partyApi, navigate, userLoggedIn]);

    useEffect(() => {
        const partyStartTime = new Date(party.start_time).getTime();
        const currentTime = new Date().getTime();
        const timeDiff = partyStartTime - currentTime;
        const hoursRemaining = timeDiff / (1000 * 60 * 60);

        if (hoursRemaining > 336) {
            setTimeRemaining(336);
        } else if (hoursRemaining < 0) {
            setTimeRemaining(0);
        } else {
            setTimeRemaining(hoursRemaining);
        }
    }, [party.start_time]);

    const handleContributeClick = () => {
        navigate('/contributions');
    };

    const formatDate = (date: Date) => {
        const d = new Date(date);

        return `${d.toLocaleString('default', { month: 'long' })} ${d.getDate()}, ${d.toLocaleTimeString('default', { hour: '2-digit', minute: '2-digit' })}`;
    };

    const getProgressBarWidth = () => {
        if (timeRemaining > 336) {
            return '0%';
        }
        return `${((336 - timeRemaining) / 336) * 100}%`;
    };

    const getProgressText = () => {
        if (timeRemaining > 336) {
            return 'More than 336 hours until party';
        }
        if (timeRemaining <= 0) {
            return 'Party has started!';
        }
        return `${Math.ceil(timeRemaining)} hours until party`;
    };

    return (
        <div className={classes.outerContainer}>
            <video
                src={partyVideo}
                autoPlay
                loop
                muted
                playsInline
                className={classes.backgroundVideo}
            />

            <div className={classes.content}>
                <div className={classes.middleSection}>
                    <h1>{party.name}</h1>
                    <p>Contributions to the party</p>

                    <div className={classes.progressBarContainer}>
                        <div className={classes.progressBar}>
                            <div 
                                className={classes.progressFill} 
                                style={{ width: getProgressBarWidth() }} 
                            />
                            <p className={classes.progressText}>{getProgressText()}</p>
                        </div>
                    </div>

                    <button
                        className={classes.contributeButton}
                        onClick={handleContributeClick}
                    >
                        Contribute
                    </button>
                </div>

                <div className={classes.bottomRow}>
                    <div className={classes.infoItem}>
                        <h5>{formatDate(party.start_time)}</h5>
                    </div>

                    <div className={classes.infoItem}>
                        <div className={classes.locationContainer}>
                            <h5>{party.place}</h5>
                            {party.google_maps_link && (
                                <a 
                                    href={party.google_maps_link} 
                                    target="_blank" 
                                    rel="noopener noreferrer"
                                    className={classes.iconLink}
                                >
                                    <Map />
                                </a>
                            )}
                        </div>
                    </div>

                    <div className={classes.infoItem}>
                        <div className={classes.socialLinks}>
                            {party.facebook_link && (
                                <a 
                                    href={party.facebook_link} 
                                    target="_blank" 
                                    rel="noopener noreferrer"
                                    className={classes.iconLink}
                                >
                                    <FacebookIcon />
                                </a>
                            )}
                            {party.whatsapp_link && (
                                <a 
                                    href={party.whatsapp_link} 
                                    target="_blank" 
                                    rel="noopener noreferrer"
                                    className={classes.iconLink}
                                >
                                    <WhatsAppIcon />
                                </a>
                            )}
                        </div>
                    </div>
                </div>
            </div>
        </div>
    );
};