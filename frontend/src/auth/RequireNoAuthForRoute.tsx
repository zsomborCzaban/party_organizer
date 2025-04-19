import { isUserLoggedIn } from '../store/slices/UserSlice';
import { useAppSelector } from '../store/store-helper';
import { Navigate, Outlet } from 'react-router-dom';

export const RequireNoAuthForRoute = () => {
    const userLoggedIn = useAppSelector(isUserLoggedIn);
    return !userLoggedIn ? <Outlet /> : <Navigate to='/' />;
};
