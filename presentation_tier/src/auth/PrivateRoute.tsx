import {authService} from "./AuthService";
import {Navigate, Outlet} from "react-router-dom";

const PrivateRoute = () => {
    return authService.isAuthenticated() ? <Outlet /> : <Navigate to="/login" />
}

export default PrivateRoute