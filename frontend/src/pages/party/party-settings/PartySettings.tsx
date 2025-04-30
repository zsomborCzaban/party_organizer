import { Checkbox, ConfigProvider, DatePicker, Input, theme } from 'antd';
import dayjs from 'dayjs';
import { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { ApiError } from '../../../data/types/ApiResponseTypes.ts';
import {EMPTY_PARTY_POPULATED, Party, PartyPopulated} from '../../../data/types/Party.ts';
import { updateParty } from '../../../api/apis/PartyApi.ts';
import {useApi} from "../../../context/ApiContext.ts";
import {toast} from "sonner";
import classes from './PartySettings.module.scss';
import {setForTime} from "../../../data/utils/timeoutSetterUtils.ts";

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
  buttonSuccess?: string;

  [key: string]: string | undefined;
}

const PartySettings = () => {
  const api = useApi()
  const navigate = useNavigate();
  const partyId = Number(localStorage.getItem('partyId') || '-1')


  const [party, setParty] = useState<PartyPopulated>(EMPTY_PARTY_POPULATED)
  const [partyName, setPartyName] = useState('');
  const [place, setPlace] = useState('');
  const [googlemapsLink, setGoogleMapsLink] = useState('');
  const [startTime, setStartTime] = useState<dayjs.Dayjs>(dayjs(''));
  const [facebookLink, setFacebookLink] = useState('');
  const [whatsAppLink, setWhatsappLink] = useState('');
  const [isPrivate, setIsPrivate] = useState( false);
  const [isAccessCodeEnabled, setAccessCodeEnabled] = useState(false);
  const [accessCode, setAccessCode] = useState('');
  const [feedbacks, setFeedbacks] = useState<Feedbacks>({});

  useEffect(() => {
    api.partyApi.getParty(partyId)
        .then(result => {
          if(result === 'error'){
            toast.error('Unable to load party')
            return
          }
          if(result === 'private party'){ //could also be because of the user is not logged in
            toast.error('Navigation error')
            navigate('/partyHome')
            return
          }
          setParty(result.data)
        })
        .catch(() => {
          toast.error('Unexpected error')
        })
  }, [api.partyApi, navigate, partyId]);

  const handleReset = () => {
    setPartyName(party.name);
    setPlace(party.place);
    setGoogleMapsLink(party.google_maps_link);
    setStartTime(dayjs(party.start_time));
    setFacebookLink(party.facebook_link);
    setWhatsappLink(party.whatsapp_link);
    setIsPrivate(party.is_private);
    setAccessCodeEnabled(party.access_code_enabled);
    setAccessCode(party.access_code);

    setFeedbacks({});
  };

  useEffect(() => {
    handleReset()
  }, [party]);

  // const validate = (): boolean => {
  //   let valid = true;
  //   const newFeedbacks: Feedbacks = {};
  //
  //   if (!partyName) {
  //     newFeedbacks.PartyName = 'party name is required.';
  //     valid = false;
  //   }
  //   if (!place) {
  //     newFeedbacks.Place = 'display name is required.';
  //     valid = false;
  //   }
  //   if (!googlemapsLink) {
  //     newFeedbacks.GoogleMapsLink = 'googlemapsLink is required.';
  //     valid = false;
  //   }
  //   if (!startTime) {
  //     newFeedbacks.StartTime = 'party time is required.';
  //     valid = false;
  //   }
  //   if (!startTime?.toDate()) {
  //     newFeedbacks.StartTime = 'invalid time format';
  //     valid = false;
  //   }
  //   if (isAccessCodeEnabled && !accessCode) {
  //     newFeedbacks.AccessCode = 'access code is required if you enable it';
  //     valid = false;
  //   }
  //   if (isAccessCodeEnabled && accessCode && accessCode.length < 6) {
  //     newFeedbacks.AccessCode = 'access code must be at least 6 characters long';
  //     valid = false;
  //   }
  //
  //   setFeedbacks(newFeedbacks);
  //   return valid;
  // };

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
      buttonError: '',
    };

    Array.from(errs).forEach((err) => {
      if (newFeedbacks[err.field] !== undefined) {
        newFeedbacks[err.field] = err.err;
      } else {
        newFeedbacks.buttonError = err.err;
      }
    });
    setFeedbacks(newFeedbacks);
  };

  const handleCreate = () => {
    // if (!validate()) return;

    const partyToUpdate: Party = {
      ID: party.ID,
      name: partyName,
      place: place,
      google_maps_link: googlemapsLink,
      facebook_link: facebookLink,
      whatsapp_link: whatsAppLink,
      start_time: startTime.toDate()!,
      is_private: isPrivate,
      access_code_enabled: isAccessCodeEnabled,
      access_code: isAccessCodeEnabled ?  accessCode: '',
    };

    updateParty(partyToUpdate)
      .then((returnedParty) => {
        setParty(returnedParty)
        localStorage.setItem("partyName", returnedParty.name)
        toast.success('Party saved')
      })
      .catch((err) => {
        if (err.response?.data?.errors?.Errors) {
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
      <div className={classes.background}>
        <div className={classes.outerContainer}>
          <ConfigProvider theme={{algorithm: theme.darkAlgorithm}}>
            <div className={classes.container}>
              <h2 className={classes.h2}>Party Settings</h2>

              <div className={classes.inputDiv}>
                <label className={classes.label}>Party Name *</label>
                <Input
                    placeholder='Enter Party Name'
                    value={partyName}
                    onChange={(e) => setPartyName(e.target.value)}
                    className={classes.input}
                />
                {feedbacks.Name && <p className={classes.error}>{feedbacks.Name}</p>}
              </div>

              <div className={classes.inputDiv}>
                <label className={classes.label}>Displayed Place *</label>
                <Input
                    placeholder='Enter Displayed Place'
                    value={place}
                    onChange={(e) => setPlace(e.target.value)}
                    className={classes.input}
                />
                {feedbacks.Place && <p className={classes.error}>{feedbacks.Place}</p>}
              </div>

              <div className={classes.inputDiv}>
                <label className={classes.label}>Actual Location</label>
                <Input
                    placeholder='Enter googlemaps plus code'
                    value={googlemapsLink}
                    onChange={(e) => setGoogleMapsLink(e.target.value)}
                    className={classes.input}
                />
                {feedbacks.GoogleMapsLink && <p className={classes.error}>{feedbacks.GoogleMapsLink}</p>}
              </div>

              <div className={classes.inputDiv}>
                <label className={classes.label}>Time *</label>
                <DatePicker
                    showTime
                    value={startTime}
                    className={classes.input}
                    onChange={(date) => setStartTime(date)}
                />
                {feedbacks.StartTime && <p className={classes.error}>{feedbacks.StartTime}</p>}
              </div>

              <div className={classes.inputDiv}>
                <label className={classes.label}>Facebook Link</label>
                <Input
                    placeholder='Enter Facebook Link'
                    value={facebookLink}
                    onChange={(e) => setFacebookLink(e.target.value)}
                    className={classes.input}
                />
                {feedbacks.FacebookLink && <p className={classes.error}>{feedbacks.FacebookLink}</p>}
              </div>

              <div className={classes.inputDiv}>
                <label className={classes.label}>WhatsApp Link</label>
                <Input
                    placeholder='Enter WhatsApp Link'
                    value={whatsAppLink}
                    onChange={(e) => setWhatsappLink(e.target.value)}
                    className={classes.input}
                />
                {feedbacks.WhatsappLink && <p className={classes.error}>{feedbacks.WhatsappLink}</p>}
              </div>

              <div className={classes.inputDiv}>
                <div className={classes.checkboxContainer}>
                  <div className={classes.checkbox}>
                    <Checkbox
                        id='isPrivate'
                        checked={isPrivate}
                        onChange={(e) => setIsPrivate(e.target.checked)}
                    > Private
                    </Checkbox>
                  </div>

                  <div className={classes.checkbox}>
                    <Checkbox
                        id='isAccessCodeEnabled'
                        checked={isAccessCodeEnabled}
                        onChange={(e) => setAccessCodeEnabled(e.target.checked)}
                        disabled={!isPrivate}
                    > Access Code Enabled
                    </Checkbox>
                  </div>
                </div>
                {feedbacks.AccessCodeEnabled && <p className={classes.error}>{feedbacks.AccessCodeEnabled}</p>}
              </div>

              <div className={classes.inputDiv}>
                {isAccessCodeEnabled && (
                    <>
                      <label className={classes.label}>Access Code</label>
                      <Input
                          placeholder='Enter Access Code'
                          value={accessCode}
                          onChange={(e) => setAccessCode(e.target.value)}
                          className={classes.input}
                          prefix={party.ID.toString() + "_"}
                      />
                      {feedbacks.AccessCode && <p className={classes.error}>{feedbacks.AccessCode}</p>}
                    </>
                )}
              </div>

              {/* Buttons */}
              <div className={classes.buttonsContainer}>
                <button
                    className={classes.button}
                    onClick={handleCreate}
                >
                  Save
                </button>
                <button
                    className={classes.resetButton}
                    onClick={handleReset}
                >
                  Reset
                </button>
              </div>
              {feedbacks.buttonError && <p className={classes.error}>{feedbacks.buttonError}</p>}
              {feedbacks.buttonSuccess && <p className={classes.success}>{feedbacks.buttonSuccess}</p>}
            </div>
          </ConfigProvider>
        </div>
      </div>
        );
        };

        export default PartySettings;
