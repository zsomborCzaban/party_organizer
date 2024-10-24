import {BrowserRouter, Route, Routes} from "react-router-dom";
import Login from "../features/authtentication/Login";
import PrivateRoute from "../auth/PrivateRoute";
import Login2 from "../features/authtentication/Login2";
import Register from "../features/authtentication/Register";
import Discover from "../features/overView/discover/Discover";
import PartiesPage from "../features/overView/partiesPage/PartiesPage";
import Friends from "../features/overView/friends/Friends";
import CreateParty from "../features/createParty/CreateParty";
import SetupParty from "../features/createParty/SetupParty";
import PartyHome from "../features/visitParty/partyHome/PartyHome";

const AppRoutes = () => {
    return (
        <BrowserRouter>
            <Routes>
                <Route path="/overview/discover" element={<PrivateRoute />}>
                    <Route path="/overview/discover" element={<Discover/>}/>
                </Route>
                <Route path="/overview/parties" element={<PrivateRoute />}>
                    <Route path="/overview/parties" element={<PartiesPage />}/>
                </Route>
                <Route path="/overview/friends" element={<PrivateRoute />}>
                    <Route path="/overview/friends" element={<Friends />}/>
                </Route>
                <Route path="/createParty" element={<PrivateRoute />}>
                    <Route path="/createParty" element={<CreateParty />}/>
                </Route>
                <Route path="/setupParty" element={<PrivateRoute />}>
                    <Route path="/setupParty" element={<SetupParty />}/>
                </Route>
                <Route path="/visitParty/partyHome" element={<PrivateRoute />}>
                    <Route path="/visitParty/partyHome" element={<PartyHome />}/>
                </Route>
                <Route path="/login" element={<Login/>} />
                <Route path="/login2" element={<Login2/>} />
                <Route path="/register" element={<Register/>} />
            </Routes>
        </BrowserRouter>
    )
}

export default AppRoutes;