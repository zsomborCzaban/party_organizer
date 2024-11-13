import {authService} from './AuthService';
import {Navigate, Outlet} from 'react-router-dom';

const PrivateRoute = () => authService.isAuthenticated() ? <Outlet /> : <Navigate to='/login' />;

export default PrivateRoute;