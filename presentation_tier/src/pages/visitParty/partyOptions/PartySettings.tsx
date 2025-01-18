import {CSSProperties, useEffect, useState} from 'react';
import {useDispatch, useSelector} from 'react-redux';
import {AppDispatch, RootState} from 'store/store';
import {useNavigate} from 'react-router-dom';
import {Button, Checkbox, ConfigProvider, DatePicker, Input, theme} from 'antd';
import backgroundImage from 'data/resources/images/gears.png';
import dayjs from 'dayjs';
import {ApiError} from 'api/ApiResponse';
import {Party} from 'data/types/Party';
import {updateParty} from 'data/apis/PartyApi';
import {setSelectedParty} from 'data/sclices/PartySlice';
import VisitPartyNavBar from 'components/navbar/VisitPartyNavBar';
import VisitPartyProfile from 'components/drawer/VisitPartyProfile';
import {User} from 'data/types/User';
import {getUser} from 'auth/AuthUserUtil';
import {authService} from 'auth/AuthService';
import {setForTime} from 'data/utils/timeoutSetterUtils';


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
        width: 'min(80%, 1000px)',
        margin: '20px auto',
        padding: '20px',
        display: 'flex',
        flexDirection: 'column',
        // backgroundColor: "#2c2c2c", // Darker gray background for content box
        backgroundColor: 'rgba(33, 33, 33, 0.95)',
        borderRadius: '8px',
        boxShadow: '0 4px 8px rgba(0, 0, 0, 0.4)', // Slightly stronger shadow for depth
        color: '#007bff', // Ensure text is white for readability
    },
    h2: {
        color: '#d3d3d3', // Light gray for headings
        fontSize: '2.5rem',
        fontWeight: 'bold',
        textAlign: 'left',
        marginBottom: '20px',
    },
    inputDiv: {
        display: 'flex',
        flexDirection: 'column',
        alignItems: 'flex-start',
        marginBottom: '20px',
    },
    label: {
        marginBottom: '5px',
    },
    input: {
        padding: '8px 12px',
        fontSize: '1rem',
        borderRadius: '5px',
        border: '1px solid #444', // Darker border to blend with dark mode
        backgroundColor: '#3a3a3a', // Dark input background
        color: '#ffffff', // Light input text
        width: '60%',
    },
    checkboxContainer: {
        display: 'flex',
        flexDirection: 'row',
        gap: '20px',
        width: '60%',
    },
    checkbox: {
        display: 'flex',
        flexDirection: 'row',
        gap: '5px',
    },
    buttonsContainer: {
        display: 'flex',
        flexDirection: 'row',
        gap: '20px',
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
        color: '#66ff66', // Light green for success messages
        fontSize: '1rem',
        marginTop: '5px',
    },
    error: {
        color: '#ff6666', // Light red for error messages
        fontSize: '1rem',
        marginTop: '5px',
        marginBottom: '0px',
    },
    loading: {
        textAlign: 'center',
        fontSize: '1rem',
        color: '#d3d3d3',
    },
    errorMessage: {
        textAlign: 'center',
        fontSize: '1rem',
        color: '#ff6666',
    },
};

interface Feedbacks{
    PartyName?: string;
    Place?: string;
    GooglemapsLink?: string;
    StartTime?: string;
    FacebookLink?: string;
    WhatsAppLink?: string;
    IsPrivate?: string;
    IsAccessCodeEnabled?: string;
    AccessCode?: string;
    buttonError?: string;
    buttonSuccess?: string;

    [key: string]: string | undefined;
}

const PartySettings = () => {
    const navigate = useNavigate();
    const dispatch = useDispatch<AppDispatch>();

    const {selectedParty} = useSelector((state: RootState)=> state.selectedPartyStore);

    const [profileOpen, setProfileOpen] = useState(false);
    const [user, setUser] = useState<User>();
    const [partyName, setPartyName] = useState(selectedParty ? selectedParty.name : '');
    const [place, setPlace] = useState(selectedParty ? selectedParty.place : '');
    const [googlemapsLink, setGoogleMapsLink] = useState(selectedParty ? selectedParty.google_maps_link : '');
    const [startTime, setStartTime] = useState<dayjs.Dayjs>(dayjs(selectedParty? selectedParty.start_time : ''));
    const [facebookLink, setFacebookLink] = useState(selectedParty ? selectedParty.facebook_link : '');
    const [whatsAppLink, setWhatsAppLink] = useState(selectedParty ? selectedParty.whatsapp_link : '');
    const [isPrivate, setIsPrivate] = useState(selectedParty ? selectedParty.is_private : false);
    const [isAccessCodeEnabled, setIsAccessCodeEnabled] = useState(selectedParty ? selectedParty.access_code_enabled : false);
    const [accessCode, setAccessCode] = useState(selectedParty ? selectedParty.access_code : '');
    const [feedbacks, setFeedbacks] = useState<Feedbacks>({});


    useEffect(() => {
        const currentUser = getUser();

        if(!currentUser) {
            authService.handleUnauthorized();
            return;
        }

        setUser(currentUser);
    }, []);

    if(!selectedParty){
        console.log('error, no selected party');
        navigate('/overview/discover');
        return <div>error, selected party was null</div>;
    }

    if(!user){
        console.log('user was null');
        return <div>Loading...</div>;
    }

    const handleReset = () => {
        setPartyName(selectedParty.name);
        setPlace(selectedParty.place);
        setGoogleMapsLink(selectedParty.google_maps_link);
        setStartTime(dayjs(selectedParty.start_time));
        setFacebookLink(selectedParty.facebook_link);
        setWhatsAppLink(selectedParty.whatsapp_link);
        setIsPrivate(selectedParty.is_private);
        setIsAccessCodeEnabled(selectedParty.access_code_enabled);
        setAccessCode(selectedParty.access_code);

        setFeedbacks({});
    };

    const validate = (): boolean => {
        let valid = true;
        const newFeedbacks: Feedbacks = {};

        if (!partyName) {
            newFeedbacks.PartyName = 'party name is required.';
            valid = false;
        }
        if (!place) {
            newFeedbacks.Place = 'display name is required.';
            valid = false;
        }
        if (!googlemapsLink) {
            newFeedbacks.GooglemapsLink = 'googlemapsLink is required.';
            valid = false;
        }
        if (!startTime) {
            newFeedbacks.StartTime = 'party time is required.';
            valid = false;
        }
        if (!startTime?.toDate()) {
            newFeedbacks.StartTime = 'invalid time format';
            valid = false;
        }
        if (!isAccessCodeEnabled && accessCode) {
            newFeedbacks.IsAccessCodeEnabled = 'to use access code, you should enable it';
            valid = false;
        }
        if (isAccessCodeEnabled && !accessCode) {
            newFeedbacks.AccessCode = 'access code is required if you enable it';
            valid = false;
        }
        if (accessCode && accessCode.length < 6) {
            newFeedbacks.AccessCode = 'access code must be at least 6 characters long';
            valid = false;
        }

        setFeedbacks(newFeedbacks);
        return valid;
    };

    const handleErrors = (errs: ApiError[]) => {
        const newFeedbacks: Feedbacks ={
            PartyName: '',
            Place: '',
            GooglemapsLink: '',
            StartTime: '',
            FacebookLink: '',
            WhatsAppLink: '',
            IsPrivate: '',
            IsAccessCodeEnabled: '',
            AccessCode: '',
            buttonError: '',
            buttonSuccess: '',
        };

        Array.from(errs).forEach(err => {
            if (newFeedbacks[err.field] !== undefined){
                newFeedbacks[err.field] = err.err;
            } else {
                newFeedbacks.buttonError = err.err;
            }
        });
        setFeedbacks(newFeedbacks);
    };

    const handleCreate = () => {
        if (!validate()) return;

        const party: Party = {
            ID: selectedParty.ID,
            name: partyName,
            place,
            google_maps_link: googlemapsLink,
            facebook_link: facebookLink,
            whatsapp_link: whatsAppLink,
            start_time: startTime?.toDate()!,
            is_private: isPrivate,
            access_code_enabled: isAccessCodeEnabled,
            access_code: accessCode,
        };

        updateParty(party)
            .then((returnedParty) => {
                dispatch(setSelectedParty(returnedParty));

                const newFeedbacks: Feedbacks = {};
                newFeedbacks.buttonSuccess = 'saved successfully';
                setForTime(setFeedbacks, newFeedbacks, {}, 3000);
                
            })
            .catch(err => {
                if (err.response) {
                    const errors: ApiError[] = err.response.data.errors.Errors;
                    handleErrors(errors);
                    
                } else {
                    const newFeedbacks: Feedbacks = {};
                    newFeedbacks.buttonError = 'Something unexpected happened. Try again later!';
                    setForTime(setFeedbacks, newFeedbacks, {}, 3000);
                    
                }
            });
    };

    return (
      <div style={styles.outerContainer}>
        <ConfigProvider
          theme={{algorithm: theme.darkAlgorithm}}
        >
          <VisitPartyNavBar onProfileClick={() => setProfileOpen(true)} />
          <VisitPartyProfile isOpen={profileOpen} onClose={() => setProfileOpen(false)} currentParty={selectedParty} user={user} onLeaveParty={() => console.log('leaveparty')} />

          <div style={styles.container}>
            <h2 style={styles.h2}>Party Settings</h2>

            <div style={styles.inputDiv}>
              <label style={styles.label}>Party Name</label>
              <Input
                placeholder='Enter Party Name'
                value={partyName}
                onChange={(e) => setPartyName(e.target.value)}
                style={styles.input}
              />
              {feedbacks.PartyName && <p style={styles.error}>{feedbacks.PartyName}</p>}
            </div>

            <div style={styles.inputDiv}>

              <label style={styles.label}>Displayed Place</label>
              <Input
                placeholder='Enter Displayed Place'
                value={place}
                onChange={(e) => setPlace(e.target.value)}
                style={styles.input}
              />
              {feedbacks.Place && <p style={styles.error}>{feedbacks.Place}</p>}
            </div>


            <div style={styles.inputDiv}>
              <label style={styles.label}>Actual Location</label>
              <Input
                placeholder='Enter googlemaps plus code'
                value={googlemapsLink}
                onChange={(e) => setGoogleMapsLink(e.target.value)}
                style={styles.input}
              />
              {feedbacks.GooglemapsLink && <p style={styles.error}>{feedbacks.GooglemapsLink}</p>}
            </div>

            <div style={styles.inputDiv}>
              <label style={styles.label}>Time</label>
              <DatePicker
                showTime
                value={startTime}
                style={styles.input}
                onChange={(date) => setStartTime(date)}
              />
              {feedbacks.StartTime && <p style={styles.error}>{feedbacks.StartTime}</p>}
            </div>

            <div style={styles.inputDiv}>
              <label style={styles.label}>Facebook Link</label>
              <Input
                placeholder='Enter Facebook Link'
                value={facebookLink}
                onChange={(e) => setFacebookLink(e.target.value)}
                style={styles.input}
              />
              {feedbacks.FacebookLink && <p style={styles.error}>{feedbacks.FacebookLink}</p>}
            </div>

            <div style={styles.inputDiv}>
              <label style={styles.label}>WhatsApp Link</label>
              <Input
                placeholder='Enter WhatsApp Link'
                value={whatsAppLink}
                onChange={(e) => setWhatsAppLink(e.target.value)}
                style={styles.input}
              />
              {feedbacks.WhatsAppLink && <p style={styles.error}>{feedbacks.WhatsAppLink}</p>}
            </div>

            <div style={styles.inputDiv}>
              <div style={styles.checkboxContainer}>
                <div style={styles.checkbox}>
                  <label htmlFor='isPrivate' style={styles.label}>Private</label>
                  <Checkbox
                    id='isPrivate'
                    checked={isPrivate}
                    onChange={(e) => setIsPrivate(e.target.checked)}
                  />
                </div>

                <div style={styles.checkbox}>
                  <label htmlFor='isAccessCodeEnabled' style={styles.label}>Access Code Enabled</label>
                  <Checkbox
                    id='isAccessCodeEnabled'
                    checked={isAccessCodeEnabled}
                    onChange={(e) => setIsAccessCodeEnabled(e.target.checked)}
                  />
                </div>
              </div>
            </div>

            <div style={styles.inputDiv}>
              {isAccessCodeEnabled && (
              <>
                <label style={styles.label}>Access Code</label>
                <Input
                  placeholder='Enter Access Code'
                  value={accessCode}
                  onChange={(e) => setAccessCode(e.target.value)}
                  style={styles.input}
                />
                {feedbacks.AccessCode && <p style={styles.error}>{feedbacks.AccessCode}</p>}
              </>
                            )}
            </div>

            {/* Buttons */}
            <div style={styles.buttonsContainer}>
              <Button type='primary' style={styles.button} onClick={handleCreate}>
                Save
              </Button>
              <Button type='primary' style={styles.resetButton} onClick={handleReset}>
                Reset
              </Button>
            </div>
            {feedbacks.buttonError && <p style={styles.error}>{feedbacks.buttonError}</p>}
            {feedbacks.buttonSuccess && <p style={styles.success}>{feedbacks.buttonSuccess}</p>}

          </div>
        </ConfigProvider>
      </div>
    );
};

export default PartySettings;