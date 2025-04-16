import {useApi} from "../../../context/ApiContext.ts";
import {useNavigate} from "react-router-dom";
import {useEffect, useState} from "react";
import {EMPTY_PARTY_POPULATED, PartyPopulated} from "../../../data/types/Party.ts";
import {toast} from "sonner";
import {useAppSelector} from "../../../store/store-helper.ts";
import {isUserLoggedIn} from "../../../store/sclices/UserSlice.ts";
import partyVideo from '../../../data/resources/videos/party_video.mp4';
import classes from "./PartyHome.module.scss"
import Map from '@mui/icons-material/Map';
import FacebookIcon from '@mui/icons-material/Facebook';
import WhatsAppIcon from '@mui/icons-material/WhatsApp';

export const PartyHome = ()=>{

    const api = useApi()
    const navigate = useNavigate()
    const userLoggedIn = useAppSelector(isUserLoggedIn);
    // const [partyId, setPartyId] = useState(0)
    const [party, setParty] = useState<PartyPopulated>(EMPTY_PARTY_POPULATED)

    useEffect(() => {
        const getPartyId = localStorage.getItem('partyId')
        if(!getPartyId || !Number(getPartyId)){
            toast.error('Navigation error')
            navigate('/')
            return
        }

        const receivedPartyId = Number(getPartyId)
        // setPartyId(receivedPartyId)
        
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


    const handleContributeClick = () => {
        navigate('/contributions');
    };

    const formatDate = (date: Date) => {
        const d = new Date(date);

        return `${d.toLocaleString('default', { month: 'long' })} ${d.getDate()}, ${d.toLocaleTimeString('default', { hour: '2-digit', minute: '2-digit' })}`;
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
                            <div className={classes.progressFill} style={{ width: '60%' }} />
                            <p className={classes.progressText}>60% Complete</p>
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