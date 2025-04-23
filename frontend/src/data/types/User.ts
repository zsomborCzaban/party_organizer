export interface User {
    ID: number,
    username: string;
    email: string;
    profile_picture_url: string;
}

export const EMPTY_USER: User = {
    ID: 0,
    username: '',
    email: '',
    profile_picture_url: '',
}