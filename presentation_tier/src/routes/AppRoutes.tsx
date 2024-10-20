import {BrowserRouter, Route, Routes} from "react-router-dom";
import Login from "../features/authtentication/Login";
import PrivateRoute from "../auth/PrivateRoute";
import Home from "../features/home/Home_temporary";

const AppRoutes = () => {
    return (
        <BrowserRouter>
            <Routes>
                <Route path="/" element={<PrivateRoute />}>
                    <Route path="/" element={<Home />}/>
                </Route>
                <Route path="/login" element={<Login/>} />
            </Routes>
        </BrowserRouter>
    )
}

export default AppRoutes;