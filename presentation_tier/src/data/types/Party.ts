import {User} from "./User";

export interface Party {
    ID?: number;
    place: string;
    name: string;
    google_maps_link: string;
    facebook_link: string;
    whatsapp_link: string;
    start_time: Date;
    is_private: Boolean;
    access_code_enabled: Boolean;
    access_code: string;
    organizer?: User;
}