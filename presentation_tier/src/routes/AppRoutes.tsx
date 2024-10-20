import {BrowserRouter, Route, Routes} from "react-router-dom";
import Login from "../features/authtentication/Login";
import PrivateRoute from "../auth/PrivateRoute";
import Home from "../features/home/Home_temporary";
import Login2 from "../features/authtentication/Login2";
import Register from "../features/authtentication/Register";
import Discover from "../features/overView/discover/Discover";

const AppRoutes = () => {
    return (
        <BrowserRouter>
            <Routes>
                <Route path="/" element={<PrivateRoute />}>
                    <Route path="/" element={<Home />}/>
                </Route>
                <Route path="/overview/discover" element={<PrivateRoute />}>
                    <Route path="/overview/discover" element={<Discover/>}/>
                </Route>
                <Route path="/login" element={<Login/>} />
                <Route path="/login2" element={<Login2/>} />
                <Route path="/register" element={<Register/>} />
            </Routes>
        </BrowserRouter>
    )
}

export default AppRoutes;