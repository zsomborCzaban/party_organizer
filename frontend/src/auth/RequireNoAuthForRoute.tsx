import { isUserLoggedIn } from '../store/slices/UserSlice';
import { useAppSelector } from '../store/store-helper';
import { Navigate, Outlet } from 'react-router-dom';
import {toast} from "sonner";

export const RequireNoAuthForRoute = () => {
    const userLoggedIn = useAppSelector(isUserLoggedIn);
    if(userLoggedIn) toast.info('Already logged in')
    return !userLoggedIn ? <Outlet /> : <Navigate to='/' />;
};
