import { isUserLoggedIn } from '../store/sclices/UserSlice';
import { useAppSelector } from '../store/store-helper';
import { Navigate, Outlet } from 'react-router-dom';

export const RequireAuthForRoute = () => {
  const userLoggedIn = useAppSelector(isUserLoggedIn);
  return userLoggedIn ? <Outlet /> : <Navigate to='/login' />;
};
