import React, { useState } from 'react';
import {Input, Button, DatePicker, Checkbox} from 'antd';
import dayjs from "dayjs";
import 'antd/dist/reset.css'
import backgroundImage from '../../midjourney_images/blackhole.png'
import {Party} from "../overView/Party";
import {useNavigate} from "react-router-dom";
import {createParty} from "./CreatePartyApi";
import {useDispatch} from "react-redux";
import {AppDispatch} from "../../store/store";
import {setSelectedParty} from '../overView/PartySlice'
import {ApiError} from "../../api/ApiResponse";

interface Feedbacks{
    partyName?: string;
    displayedPlace?: string;
    location?: string;
    startTime?: string;
    facebookLink?: string;
    whatsAppLink?: string;
    isPrivate?: string;
    isAccessCodeEnabled?: string;
    accessCode?: string;
    button?: string;
}

const CreateParty: React.FC = () => {
    const [partyName, setPartyName] = useState('');
    const [displayedPlace, setDisplayedPlace] = useState('');
    const [location, setLocation] = useState('');
    const [startTime, setStartTime] = useState<dayjs.Dayjs>();
    const [facebookLink, setFacebookLink] = useState('');
    const [whatsAppLink, setWhatsAppLink] = useState('');
    const [isPrivate, setIsPrivate] = useState(false);
    const [isAccessCodeEnabled, setIsAccessCodeEnabled] = useState(false);
    const [accessCode, setAccessCode] = useState('');
    const [feedbacks, setFeedbacks] = useState<Feedbacks>({});

    const dispatch = useDispatch<AppDispatch>()

    const navigate = useNavigate()


    const validate = (): boolean => {
        let valid = true;
        const newFeedbacks: Feedbacks = {};

        if (!partyName) {
            newFeedbacks.partyName = 'party name is required.';
            valid = false;
        }
        if (!displayedPlace) {
            newFeedbacks.displayedPlace = 'display name is required.';
            valid = false;
        }
        if (!location) {
            newFeedbacks.location = 'location is required.';
            valid = false;
        }
        if (!startTime) {
            newFeedbacks.startTime = 'party time is required.';
            valid = false;
        }
        if(!startTime?.toDate()){
            newFeedbacks.startTime = 'invalid time format';
            valid = false;
        }
        if (!isAccessCodeEnabled  && accessCode) {
            newFeedbacks.isAccessCodeEnabled = 'to use access code, you should enable it';
            valid = false;
        }
        if (isAccessCodeEnabled  && !accessCode) {
            newFeedbacks.accessCode = 'access code is required if you enable it';
            valid = false;
        }
        if(accessCode && accessCode.length < 6){
            newFeedbacks.accessCode = 'access code must be at least 6 characters long';
            valid = false;
        }

        setFeedbacks(newFeedbacks)
        return valid
    }

    const handleErrors = (errs: ApiError[]) => {
        //todo: implement me!!!
    }

    const handleCreate = () => {
       if(!validate()) return

        const party: Party = {
            name: partyName,
            place: displayedPlace,
            google_maps_link: location,
            facebook_link: facebookLink,
            whatsapp_link: whatsAppLink,
            start_time: startTime?.toDate()!,
            is_private: isPrivate,
            access_code_enabled: isAccessCodeEnabled,
            access_code: accessCode,
        }

        createParty(party)
            .then((returnedParty) => {
                console.log(returnedParty)
                dispatch(setSelectedParty(returnedParty))
                navigate("/setupParty")
            })
            .catch(err => {
                if(err.response){
                    let errors = err.response.data.errors
                    handleErrors(errors)
                } else {
                    const newFeedbacks: Feedbacks = {};
                    newFeedbacks.button = "Something unexpected happened. Try again later!"
                    setFeedbacks(newFeedbacks)
                }
            })
    };

    const handleCancel = () => {
        navigate("/overview/discover");
    }

    return (
        <div style={styles.outerOuterContainer}>
            <div style={styles.outerContainer}>
                {/*<LocationPicker/> todo*/}


                <div style={styles.formContainer}>
                    <h2 style={styles.formTitle}>Create Party</h2>
                    {/* Party Name */}
                    <label style={styles.label}>Party Name</label>
                    <Input
                        placeholder="Enter Party Name"
                        value={partyName}
                        onChange={(e) => setPartyName(e.target.value)}
                        style={styles.input}
                    />
                    {feedbacks.partyName && <p style={styles.error}>{feedbacks.partyName}</p>}

                    {/* Displayed Place */}
                    <label style={styles.label}>Displayed Place</label>
                    <Input
                        placeholder="Enter Displayed Place"
                        value={displayedPlace}
                        onChange={(e) => setDisplayedPlace(e.target.value)}
                        style={styles.input}
                    />
                    {feedbacks.displayedPlace && <p style={styles.error}>{feedbacks.displayedPlace}</p>}

                    {/* Actual Location */}
                    <label style={styles.label}>Actual Location</label>
                    <Input
                        placeholder="Enter googlemaps plus code"
                        value={location}
                        onChange={(e) => setLocation(e.target.value)}
                        style={styles.input}
                    />
                    {feedbacks.location && <p style={styles.error}>{feedbacks.location}</p>}

                    {/* Time Picker */}
                    <label style={styles.label}>Time</label>
                    <DatePicker
                        showTime
                        style={styles.input}
                        onChange={(date) => setStartTime(date)}
                    />
                    {feedbacks.startTime && <p style={styles.error}>{feedbacks.startTime}</p>}

                    {/* Facebook Link */}
                    <label style={styles.label}>Facebook Link</label>
                    <Input
                        placeholder="Enter Facebook Link"
                        value={facebookLink}
                        onChange={(e) => setFacebookLink(e.target.value)}
                        style={styles.input}
                    />
                    {feedbacks.facebookLink && <p style={styles.error}>{feedbacks.facebookLink}</p>}

                    {/* WhatsApp Link */}
                    <label style={styles.label}>WhatsApp Link</label>
                    <Input
                        placeholder="Enter WhatsApp Link"
                        value={whatsAppLink}
                        onChange={(e) => setWhatsAppLink(e.target.value)}
                        style={styles.input}
                    />
                    {feedbacks.whatsAppLink && <p style={styles.error}>{feedbacks.whatsAppLink}</p>}

                    {/* Private Slider */}
                    <div style={styles.checkboxContainer}>
                        <div style={styles.checkbox}>
                            <label style={styles.label}>Private</label>
                            <Checkbox
                                checked={isPrivate}
                                onChange={(e) => setIsPrivate(e.target.checked)}
                                style={styles.slider}
                            />
                            {feedbacks.isPrivate && <p style={styles.error}>{feedbacks.isPrivate}</p>}
                        </div>


                        <div style={styles.checkbox}>
                            {/* Access Code Enable Slider */}
                            <label style={styles.label}>Access Code Enabled</label>
                            <Checkbox
                                checked={isAccessCodeEnabled}
                                onChange={(e) => setIsAccessCodeEnabled(e.target.checked)}
                                style={styles.slider}
                            />
                            {feedbacks.isAccessCodeEnabled && <p style={styles.error}>{feedbacks.isAccessCodeEnabled}</p>}
                        </div>
                    </div>


                {/* Access Code */}
                {isAccessCodeEnabled && (
                        <>
                            <label style={styles.label}>Access Code</label>
                            <Input
                                placeholder="Enter Access Code"
                                value={accessCode}
                                onChange={(e) => setAccessCode(e.target.value)}
                                style={styles.input}
                            />
                            {feedbacks.accessCode && <p style={styles.error}>{feedbacks.accessCode}</p>}
                        </>
                    )}

                    {/* Create Button */}
                    <div style={styles.buttonsContainer}>
                        <Button type="primary" style={styles.createButton} onClick={handleCreate}>
                            Continue
                        </Button>
                        <Button type="primary" style={styles.cancelButton} onClick={handleCancel}>
                            Cancel
                        </Button>
                    </div>
                    {feedbacks.button && <p style={styles.error}>{feedbacks.button}</p>}
                </div>
            </div>
        </div>
            );
            };

// Inline CSS styles
const styles: {[key: string]: React.CSSProperties} = {
    outerOuterContainer: {
        backgroundImage: `url(${backgroundImage})`,
        width: '100vw',
        height: '100vh',
        backgroundSize: 'cover',
        backgroundPosition: 'center',
    },
    outerContainer: {
        display: 'flex',
        alignItems: 'flex-start',
        flexDirection: 'column',
        paddingLeft: '100px',
        paddingTop: '50px',
        maxWidth: '700px',
    },
    formContainer: {
        flexGrow: '1',
        width: '100%',
        margin: '0 auto',
        padding: '20px',
        backgroundColor: 'rgba(249, 249, 249, 1)',
        borderRadius: '8px',
        boxShadow: '0 2px 10px rgba(0, 0, 0, 0.1)',
    },
    checkboxContainer: {
        display: 'flex',
        alignItems: 'flex-start',
        flexDirection: 'row',
    },
    checkbox: {
        display: 'flex',
        alignItems: 'center',
        flexDirection: 'column',
        marginRight: '20px',
    },
    label: {
        display: 'block',
        marginBottom: '5px',
        fontWeight: 'bold',
        fontSize: '16px',
    },
    input: {
        width: '100%',
        padding: '8px',
        marginBottom: '10px',
        border: '1px solid #d9d9d9',
        borderRadius: '4px',
    },
    slider: {
        width: '100%',
        marginBottom: '10px',
    },
    error: {
        color: 'red',
        fontSize: '0.875em',
    },
    buttonsContainer: {
        display: 'flex',
        justifyContent: 'space-between',
        padding: '0 20px',
        flexGrow: 1, // Allow the buttons container to grow
    },
    createButton: {
        width: '48%',
        textAlign: 'center',
        borderRadius: '10px',
        cursor: 'pointer',
        display: 'flex',
        justifyContent: 'center',
        alignItems: 'center',
        boxShadow: '0 4px 8px rgba(0, 0, 0, 0.2)',
    },
    cancelButton: {
        backgroundColor: 'red',
        width: '48%',
        textAlign: 'center',
        borderRadius: '10px',
        cursor: 'pointer',
        display: 'flex',
        justifyContent: 'center',
        alignItems: 'center',
        boxShadow: '0 4px 8px rgba(0, 0, 0, 0.2)',
    },
    formTitle: {
        textAlign: 'center'
    }
};

export default CreateParty;
