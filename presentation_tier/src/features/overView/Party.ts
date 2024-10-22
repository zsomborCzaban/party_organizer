import {User} from "./User";

export interface Party {
    ID: number;
    place: string;
    name: string;
    start_time: Date;
    private: Boolean;
    organizer_id: number;
    organizer: User;
}