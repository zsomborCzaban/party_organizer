import {EMPTY_USER, User} from './User';

export interface Party {
    ID?: number;
    place: string;
    name: string;
    google_maps_link: string;
    facebook_link: string;
    whatsapp_link: string;
    start_time: Date;
    is_private: boolean;
    access_code_enabled: boolean;
    access_code: string;
    organizer?: User;
}

export interface PartyPopulated {
    ID: number;
    place: string;
    name: string;
    google_maps_link: string;
    facebook_link: string;
    whatsapp_link: string;
    start_time: Date;
    is_private: boolean;
    access_code_enabled: boolean;
    access_code: string;
    organizer: User;
}

export const EMPTY_PARTY_POPULATED: PartyPopulated = {
    ID: 0,
    place: '',
    name: '',
    google_maps_link: '',
    facebook_link: '',
    whatsapp_link: '',
    start_time: new Date(),
    is_private: false,
    access_code_enabled: false,
    access_code: '',
    organizer: EMPTY_USER,
}