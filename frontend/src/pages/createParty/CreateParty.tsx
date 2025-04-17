import React, { useState, useRef } from 'react';
import { Input, DatePicker, Checkbox, ConfigProvider, theme } from 'antd';
import dayjs from 'dayjs';
import 'antd/dist/reset.css';
import { Party } from '../../data/types/Party';
import { useNavigate } from 'react-router-dom';
import { ApiError } from '../../data/types/ApiResponseTypes';
import { createParty } from '../../api/apis/PartyApi';
import { toast } from "sonner";
import classes from './CreateParty.module.scss';
import artlistVideo from '../../data/resources/videos/artlist_video.mp4';

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
  const videoRef = useRef<HTMLVideoElement>(null);
  const [showVideo, setShowVideo] = useState(false);

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
        toast.success('Party created')
        
        setShowVideo(true);
        if (videoRef.current) {
          videoRef.current.play();
        }
      })
      .catch((err) => {
        if (err.response?.data?.errors?.Errors) {
          handleErrors(err.response.data.errors.Errors);
        } else {
          toast.error('Unexpected error')
        }
      });
  };

  const handleVideoEnded = () => {
    navigate('/partyHome');
  };

  const handleCancel = () => {
    navigate('/');
  };

  return (
    <ConfigProvider
      theme={{
        algorithm: theme.darkAlgorithm,
        token: {
          colorBgContainer: '#3a3a3a',
          colorBorder: '#444',
          borderRadius: 5,
          fontSize: 15,
          controlHeight: 40,
          paddingContentHorizontal: 14,
        },
        components: {
          Input: {
            colorBgContainer: '#3a3a3a',
            hoverBorderColor: '#007bff',
            hoverBg: '#000000',
            activeBorderColor: '#007bff',
            paddingBlock: 10,
            paddingInline: 14,
          },
          DatePicker: {
            colorBgContainer: '#3a3a3a',
            hoverBorderColor: '#007bff',
            activeBorderColor: '#007bff',
            paddingBlock: 10,
            paddingInline: 14,
          },
        },
      }}
    >
      <div className={classes.pageContainer}>
        {showVideo && (
          <video
            ref={videoRef}
            className={classes.videoBackground}
            onEnded={handleVideoEnded}
            autoPlay
            muted
            playsInline
            preload="auto"
            style={{
              width: '100%',
              height: '100%',
              objectFit: 'cover',
              transform: 'scale(1.01)',
              backfaceVisibility: 'hidden',
              WebkitBackfaceVisibility: 'hidden',
              imageRendering: 'crisp-edges'
            }}
          >
            <source src={artlistVideo} type="video/mp4" />
          </video>
        )}
        <div className={classes.formWrapper}>
          <div className={classes.formContainer}>
            <h2 className={classes.formTitle}>Create Party</h2>
            
            <div className={classes.inputGroup}>
              <label className={classes.label}>Party Name</label>
              <Input
                placeholder='Enter Party Name'
                value={partyName}
                onChange={(e) => setPartyName(e.target.value)}
              />
              {feedbacks.Name && <p className={classes.error}>{feedbacks.Name}</p>}
            </div>

            <div className={classes.inputGroup}>
              <label className={classes.label}>Displayed Place</label>
              <Input
                placeholder='Enter Displayed Place'
                value={displayedPlace}
                onChange={(e) => setDisplayedPlace(e.target.value)}
              />
              {feedbacks.Place && <p className={classes.error}>{feedbacks.Place}</p>}
            </div>

            <div className={classes.inputGroup}>
              <label className={classes.label}>Actual Location</label>
              <Input
                placeholder='Enter googlemaps plus code'
                value={location}
                onChange={(e) => setLocation(e.target.value)}
              />
              {feedbacks.GoogleMapsLink && (
                <p className={classes.error}>{feedbacks.GoogleMapsLink}</p>
              )}
            </div>

            <div className={classes.inputGroup}>
              <label className={classes.label}>Time</label>
              <DatePicker
                showTime
                onChange={(date) => setStartTime(date)}
              />
              {feedbacks.StartTime && (
                <p className={classes.error}>{feedbacks.StartTime}</p>
              )}
            </div>

            <div className={classes.inputGroup}>
              <label className={classes.label}>Facebook Link</label>
              <Input
                placeholder='Enter Facebook Link'
                value={facebookLink}
                onChange={(e) => setFacebookLink(e.target.value)}
              />
              {feedbacks.FacebookLink && (
                <p className={classes.error}>{feedbacks.FacebookLink}</p>
              )}
            </div>

            <div className={classes.inputGroup}>
              <label className={classes.label}>WhatsApp Link</label>
              <Input
                placeholder='Enter WhatsApp Link'
                value={whatsAppLink}
                onChange={(e) => setWhatsAppLink(e.target.value)}
              />
              {feedbacks.WhatsappLink && (
                <p className={classes.error}>{feedbacks.WhatsappLink}</p>
              )}
            </div>

            <div className={classes.checkboxContainer}>
              <div className={classes.checkboxGroup}>
                <Checkbox
                  checked={isPrivate}
                  onChange={(e) => setIsPrivate(e.target.checked)}
                >
                  Private Party
                </Checkbox>
              </div>
              <div className={classes.checkboxGroup}>
                <Checkbox
                  checked={isAccessCodeEnabled}
                  onChange={(e) => setIsAccessCodeEnabled(e.target.checked)}
                >
                  Enable Access Code
                </Checkbox>
              </div>
            </div>

            {isAccessCodeEnabled && (
              <div className={classes.inputGroup}>
                <label className={classes.label}>Access Code</label>
                <Input
                  placeholder='Enter Access Code'
                  value={accessCode}
                  onChange={(e) => setAccessCode(e.target.value)}
                />
                {feedbacks.AccessCode && (
                  <p className={classes.error}>{feedbacks.AccessCode}</p>
                )}
              </div>
            )}

            <div className={classes.buttonsContainer}>
              <button onClick={handleCreate} className={classes.createButton}>
                Create Party
              </button>
              <button onClick={handleCancel} className={classes.cancelButton}>
                Cancel
              </button>
            </div>
          </div>
        </div>
      </div>
    </ConfigProvider>
  );
};

export default CreateParty;
