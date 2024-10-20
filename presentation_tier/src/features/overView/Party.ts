import {User} from "./User";

export interface Party {
    ID: number;
    Place: string;
    Name: string;
    StartTime: Date;
    Private: Boolean;
    OrganizerID: number;
    Organizer: User;
}