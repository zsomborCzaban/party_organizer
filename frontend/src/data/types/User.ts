export interface User {
    id: number,
    username: string;
    email: string;
    profile_picture_url: string;
}

export const EMPTY_USER: User = {
    id: 0,
    username: '',
    email: '',
    profile_picture_url: '',
}