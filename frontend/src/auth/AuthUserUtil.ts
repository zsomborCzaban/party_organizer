import { getJwtAuthToken } from './AuthStorageUtils';
import { jwtDecode, JwtPayload } from 'jwt-decode';
import { User } from '../data/types/User';

interface UserJwtPayload extends JwtPayload {
  id: string;
  email: string;
  username: string;
  profilePictureUrl: string;
}

export const getUserId = () => {
  const authToken = getJwtAuthToken();
  if (!authToken) {
    return null;
  }

  try {
    const decoded: UserJwtPayload = jwtDecode(authToken);
    return decoded.sub;
  } catch (e) {
    return null;
  }
};

export const getUserEmail = () => {
  const authToken = getJwtAuthToken();
  if (!authToken) {
    return null;
  }

  try {
    const decoded: UserJwtPayload = jwtDecode(authToken);
    return decoded.email;
  } catch (e) {
    return null;
  }
};

export const getUserName = () => {
  const authToken = getJwtAuthToken();
  if (!authToken) {
    return null;
  }

  try {
    const decoded: UserJwtPayload = jwtDecode(authToken);
    return decoded.username;
  } catch (e) {
    return null;
  }
};


export const getUserProfilePicture = () => {
  const authToken = getJwtAuthToken();
  if (!authToken) {
    return null;
  }

  try {
    const decoded: UserJwtPayload = jwtDecode(authToken);
    return decoded.profilePictureUrl;
  } catch (e) {
    return null;
  }
};

export const getUser = () => {
  const authToken = getJwtAuthToken();
  if (!authToken) {
    return null;
  }

  try {
    const decoded: UserJwtPayload = jwtDecode(authToken);
    const user: User = {
      ID: Number(decoded.id),
      username: decoded.username,
      email: decoded.email,
      profile_picture_url: decoded.profilePictureUrl,
    };
    return user;
  } catch (e) {
    return null;
  }
};
