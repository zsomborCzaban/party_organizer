import {useNavigate} from "react-router-dom"
import {authService} from "./AuthService";
import {useEffect} from "react";

const useSessionCheck = () => {
    const navigate = useNavigate()

    useEffect(() => {
        if(!authService.isAuthenticated()) {
            navigate("/login")
        }
    }, [navigate]);
}

export default useSessionCheck