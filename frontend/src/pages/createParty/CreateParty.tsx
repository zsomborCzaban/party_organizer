import React, { useState } from 'react';
import { Input, Button, DatePicker, Checkbox } from 'antd';
import dayjs from 'dayjs';
import 'antd/dist/reset.css';
import { Party } from '../../data/types/Party';
import { useNavigate } from 'react-router-dom';
import { ApiError } from '../../data/types/ApiResponseTypes';
import { createParty } from '../../api/apis/PartyApi';
import {toast} from "sonner";

// Inline CSS styles
const styles: { [key: string]: React.CSSProperties } = {
  outerOuterContainer: {
    backgroundImage: `url(${'background image url'})`,
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
    width: '100vw',
    height: 'auto',
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
    textAlign: 'center',
  },
};

interface Feedbacks {
  Name?: string;
  Place?: string;
  GoogleMapsLink?: string;
  StartTime?: string;
  FacebookLink?: string;
  WhatsappLink?: string;
  IsPrivate?: string;
  AccessCodeEnabled?: string;
  AccessCode?: string;
  buttonError?: string;

  [key: string]: string | undefined;
}

const CreateParty: React.FC = () => {
  const navigate = useNavigate();

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

  const handleErrors = (errs: ApiError[]) => {
    const newFeedbacks: Feedbacks = {
      Name: '',
      Place: '',
      GoogleMapsLink: '',
      StartTime: '',
      FacebookLink: '',
      WhatsappLink: '',
      IsPrivate: '',
      AccessCodeEnabled: '',
      AccessCode: '',
    };

    Array.from(errs).forEach((err) => {
      if (newFeedbacks[err.field] !== undefined) {
        newFeedbacks[err.field] = err.err;
      }
    });

    console.log(newFeedbacks)
    setFeedbacks(newFeedbacks);
  };

  const handleCreate = () => {
    const fixStartTime = startTime || dayjs()
    const party: Party = {
      name: partyName,
      place: displayedPlace,
      google_maps_link: location,
      facebook_link: facebookLink,
      whatsapp_link: whatsAppLink,
      start_time: fixStartTime.toDate(),
      is_private: isPrivate,
      access_code_enabled: isAccessCodeEnabled,
      access_code: accessCode,
    };

    createParty(party)
      .then((returnedParty) => {
        localStorage.setItem('partyName', returnedParty.name)
        localStorage.setItem('partyId', returnedParty.ID.toString())
        localStorage.setItem('partyOrganizerName', returnedParty.organizer.username)
        navigate('/partyHome');
        toast.success('Party created')
      })
      .catch((err) => {
        if (err.response?.data?.errors?.Errors) {
          handleErrors(err.response.data.errors.Errors);
        } else {
          toast.error('Unexpected error')
        }
      });
  };

  const handleCancel = () => {
    navigate('/');
  };

  return (
    <div style={styles.outerOuterContainer}>
      <div style={styles.outerContainer}>
        {/* <LocationPicker/> todo */}

        <div style={styles.formContainer}>
          <h2 style={styles.formTitle}>Create Party</h2>
          {/* Party Name */}
          <label style={styles.label}>Party Name</label>
          <Input
            placeholder='Enter Party Name'
            value={partyName}
            onChange={(e) => setPartyName(e.target.value)}
            style={styles.input}
          />
          {feedbacks.Name && <p style={styles.error}>{feedbacks.Name}</p>}

          {/* Displayed Place */}
          <label style={styles.label}>Displayed Place</label>
          <Input
            placeholder='Enter Displayed Place'
            value={displayedPlace}
            onChange={(e) => setDisplayedPlace(e.target.value)}
            style={styles.input}
          />
          {feedbacks.Place && <p style={styles.error}>{feedbacks.Place}</p>}

          {/* Actual Location */}
          <label style={styles.label}>Actual Location</label>
          <Input
            placeholder='Enter googlemaps plus code'
            value={location}
            onChange={(e) => setLocation(e.target.value)}
            style={styles.input}
          />
          {feedbacks.GoogleMapsLink && <p style={styles.error}>{feedbacks.GoogleMapsLink}</p>}

          {/* Time Picker */}
          <label style={styles.label}>Time</label>
          <DatePicker
            showTime
            style={styles.input}
            onChange={(date) => setStartTime(date)}
          />
          {feedbacks.StartTime && <p style={styles.error}>{feedbacks.StartTime}</p>}

          {/* Facebook Link */}
          <label style={styles.label}>Facebook Link</label>
          <Input
            placeholder='Enter Facebook Link'
            value={facebookLink}
            onChange={(e) => setFacebookLink(e.target.value)}
            style={styles.input}
          />
          {feedbacks.FacebookLink && <p style={styles.error}>{feedbacks.FacebookLink}</p>}

          {/* WhatsApp Link */}
          <label style={styles.label}>WhatsApp Link</label>
          <Input
            placeholder='Enter WhatsApp Link'
            value={whatsAppLink}
            onChange={(e) => setWhatsAppLink(e.target.value)}
            style={styles.input}
          />
          {feedbacks.WhatsappLink && <p style={styles.error}>{feedbacks.WhatsappLink}</p>}

          {/* Private Slider */}
          <div style={styles.checkboxContainer}>
            <div style={styles.checkbox}>
              <label style={styles.label}>Private</label>
              <Checkbox
                checked={isPrivate}
                onChange={(e) => setIsPrivate(e.target.checked)}
                style={styles.slider}
              />
              {feedbacks.IsPrivate && <p style={styles.error}>{feedbacks.IsPrivate}</p>}
            </div>

            <div style={styles.checkbox}>
              {/* Access Code Enable Slider */}
              <label style={styles.label}>Access Code Enabled</label>
              <Checkbox
                checked={isAccessCodeEnabled}
                onChange={(e) => setIsAccessCodeEnabled(e.target.checked)}
                style={styles.slider}
              />
              {feedbacks.AccessCodeEnabled && <p style={styles.error}>{feedbacks.AccessCodeEnabled}</p>}
            </div>
          </div>

          {/* Access Code */}
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

          {/* Create Button */}
          <div style={styles.buttonsContainer}>
            <Button
              type='primary'
              style={styles.createButton}
              onClick={handleCreate}
            >
              Create
            </Button>
            <Button
              type='primary'
              style={styles.cancelButton}
              onClick={handleCancel}
            >
              Cancel
            </Button>
          </div>
          {feedbacks.buttonError && <p style={styles.error}>{feedbacks.buttonError}</p>}
        </div>
      </div>
    </div>
  );
};

export default CreateParty;
