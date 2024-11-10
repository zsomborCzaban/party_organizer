import React, {CSSProperties, useEffect, useState} from "react";
import {useDispatch, useSelector,} from "react-redux";
import {AppDispatch, RootState,} from "../../../store/store";
import {useNavigate} from "react-router-dom";
import {Button, Checkbox, ConfigProvider, DatePicker, Input, theme} from "antd";
import backgroundImage from "../../../data/resources/images/gears.png";
import dayjs from "dayjs";
import {ApiError} from "../../../api/ApiResponse";
import {Party} from "../../../data/types/Party";
import {createParty} from "../../../data/apis/PartyApi";
import {setSelectedParty} from "../../../data/sclices/PartySlice";
import VisitPartyNavBar from "../../../components/navbar/VisitPartyNavBar";
import VisitPartyProfile from "../../../components/drawer/VisitPartyProfile";
import {User} from "../../../data/types/User";
import {getUser} from "../../../auth/AuthUserUtil";
import {authService} from "../../../auth/AuthService";

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


const PartyOptions = () => {
    const navigate = useNavigate()
    const dispatch = useDispatch<AppDispatch>()

    const {selectedParty} = useSelector((state: RootState)=> state.selectedPartyStore)

    const [profileOpen, setProfileOpen] = useState(false)
    const [user, setUser] = useState<User>()
    const [partyName, setPartyName] = useState(selectedParty ? selectedParty.name : '');
    const [displayedPlace, setDisplayedPlace] = useState(selectedParty ? selectedParty.place : '');
    const [location, setLocation] = useState(selectedParty ? selectedParty.google_maps_link : '');
    const [startTime, setStartTime] = useState<dayjs.Dayjs>();
    const [facebookLink, setFacebookLink] = useState(selectedParty ? selectedParty.facebook_link : '');
    const [whatsAppLink, setWhatsAppLink] = useState(selectedParty ? selectedParty.whatsapp_link : '');
    const [isPrivate, setIsPrivate] = useState(selectedParty ? selectedParty.is_private : false);
    const [isAccessCodeEnabled, setIsAccessCodeEnabled] = useState(selectedParty ? selectedParty.access_code_enabled : false);
    const [accessCode, setAccessCode] = useState(selectedParty ? selectedParty.access_code : '');
    const [feedbacks, setFeedbacks] = useState<Feedbacks>({});


    useEffect(() => {
        const currentUser = getUser()

        if(!currentUser) {
            authService.handleUnauthorized()
            return
        }

        setUser(currentUser)
    }, []);

    if(!selectedParty){
        console.log("error, no selected party")
        navigate("/overview/discover")
        return <div>error, selected party was null</div>
    }

    if(!user){
        console.log("user was null")
        return <div>Loading...</div>
    }



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
        if (!startTime?.toDate()) {
            newFeedbacks.startTime = 'invalid time format';
            valid = false;
        }
        if (!isAccessCodeEnabled && accessCode) {
            newFeedbacks.isAccessCodeEnabled = 'to use access code, you should enable it';
            valid = false;
        }
        if (isAccessCodeEnabled && !accessCode) {
            newFeedbacks.accessCode = 'access code is required if you enable it';
            valid = false;
        }
        if (accessCode && accessCode.length < 6) {
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
        if (!validate()) return

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
                navigate("/visitParty/manageParty")
            })
            .catch(err => {
                if (err.response) {
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
            <div style={styles.outerContainer}>
                {/*<VisitPartyNavBar onProfileClick={() => setProfileOpen(true)}/>*/}
                {/*<VisitPartyProfile isOpen={profileOpen} onClose={() => setProfileOpen(false)} currentParty={selectedParty} user={user} onLeaveParty={() => console.log("leaveparty")} />*/}

                <div style={styles.container}>
                        <h2 style={styles.h2}>Create Party</h2>

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
                            </div>

                            {/* Access Code Enable Slider */}
                            <div style={styles.checkbox}>
                                <label style={styles.label}>Access Code Enabled</label>
                                <Checkbox
                                    checked={isAccessCodeEnabled}
                                    onChange={(e) => setIsAccessCodeEnabled(e.target.checked)}
                                    style={styles.slider}
                                />
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

                        {/* Buttons */}
                        <div style={styles.buttonsContainer}>
                            <Button type="primary" style={styles.button} onClick={handleCreate}>
                                Save
                            </Button>
                            <Button type="primary" style={styles.resetButton} onClick={handleCancel}>
                                Reset
                            </Button>
                        </div>
                        {feedbacks.button && <p style={styles.error}>{feedbacks.button}</p>}
                    </div>
            </div>
    );
};

const styles: { [key: string]: CSSProperties } = {
    outerContainer: {
        backgroundImage: `url(${backgroundImage})`,
        position: 'fixed',
        backgroundSize: 'cover',
        backgroundPosition: 'center',
        overflowY: 'auto',
        height: '100vh',
        width: '100vw',
        display: 'flex',
        flexDirection: 'column',
        color: '#ffffff',
    },
    container: {
        width: "min(80%, 1000px)",
        margin: "20px auto",
        padding: "20px",
        display: "flex",
        flexDirection: "column",
        // backgroundColor: "#2c2c2c", // Darker gray background for content box
        backgroundColor: 'rgba(33, 33, 33, 0.95)',
        borderRadius: "8px",
        boxShadow: "0 4px 8px rgba(0, 0, 0, 0.4)", // Slightly stronger shadow for depth
        color: "#007bff", // Ensure text is white for readability
    },
    h2: {
        color: "#d3d3d3", // Light gray for headings
        fontSize: "2.5rem",
        fontWeight: "bold",
        textAlign: "left",
        marginBottom: "20px"
    },
    inputContainer: {
        display: "flex",
        flexDirection: "column",
        alignItems: "flex-start",
        marginBottom: "20px",
        gap: "10px",
    },
    input: {
        padding: "8px 12px",
        fontSize: "1rem",
        borderRadius: "5px",
        border: "1px solid #444", // Darker border to blend with dark mode
        backgroundColor: "#3a3a3a", // Dark input background
        color: "#ffffff", // Light input text
        width: "60%",
        marginBottom: "20px",

    },
    buttonsContainer: {
        display: "flex",
        flexDirection: "row",
        gap: "20px"
    },
    button: {
        width: 'auto',
        minWidth: '120px',
        padding: '10px 20px',
        borderRadius: '5px',
        marginBottom: '20px',
        fontWeight: 'bold',
        color: '#ffffff',
        backgroundColor: '#007bff', // Accent color to match link color in header
        boxShadow: '0 4px 8px rgba(0, 0, 0, 0.2)',
    },
    resetButton: {
        width: 'auto',
        minWidth: '120px',
        padding: '10px 20px',
        borderRadius: '5px',
        marginBottom: '20px',
        fontWeight: 'bold',
        color: '#ffffff',
        backgroundColor: '#007bff',
        boxShadow: '0 4px 8px rgba(0, 0, 0, 0.2)',
    },
    success: {
        color: "#66ff66", // Light green for success messages
        fontSize: "1rem",
        marginTop: "5px",
    },
    error: {
        color: "#ff6666", // Light red for error messages
        fontSize: "1rem",
        marginTop: "5px",
    },
    loading: {
        textAlign: "center",
        fontSize: "1rem",
        color: "#d3d3d3",
    },
    errorMessage: {
        textAlign: "center",
        fontSize: "1rem",
        color: "#ff6666",
    },
    checkbox: {
        marginBottom: "20px",
    }
};


export default PartyOptions