import { isUserLoggedIn } from '../store/sclices/UserSlice';
import { useAppSelector } from '../store/store-helper';
import { Navigate, Outlet } from 'react-router-dom';
import {toast} from "sonner";

export const RequireAuthForRoute = () => {
  const userLoggedIn = useAppSelector(isUserLoggedIn);
  if(!userLoggedIn) toast.warning('Login to access that page')
  return userLoggedIn ? <Outlet /> : <Navigate to='/login' />;
};
