import { BrowserRouter, Route, Routes } from 'react-router-dom';
import Discover from './pages/overView/discover/Discover';
import { PartyHome } from './pages/party/party-home/PartyHome';
import Contributions from './pages/visitParty/contribution/Contributions';
import HallOfFame from './pages/visitParty/hallOfFame/HallOfFame';
import ManageParty from './pages/party/manage-party/ManageParty.tsx';
import PartySettings from './pages/party/party-settings/PartySettings.tsx';
import { Login } from './pages/authtentication/login/Login';
import Register from './pages/authtentication/register/Register';
import { Homepage } from './pages/HomePage';
import { RequireAuthForRoute } from './auth/RequireAuthForRoute';
import { Parties } from './pages/party/parties/Parties';
import CreateParty from './pages/createParty/CreateParty';
import { Friends } from './pages/friends/Friends';
import {MainLayout, PartyLayout} from "./layouts/Layouts.tsx";

export const AppRouter = () => (
  <BrowserRouter>
    <Routes>
      <Route element={<MainLayout />}>
        <Route path="/login" element={<Login />} />
        <Route path="/register" element={<Register />} />
      </Route>

      <Route element={<PartyLayout />}>
        <Route path="/partyHome" element={<PartyHome />} />
        <Route element={<RequireAuthForRoute />}>
          <Route path="/contributions" element={<Contributions />} />
          <Route path="/manageParty" element={<ManageParty />} />
          <Route path="/partySettings" element={<PartySettings />} />
          <Route path="/hallOfFame" element={<HallOfFame />} />
        </Route>
      </Route>

      <Route element={<RequireAuthForRoute />}>
        <Route path="/createParty" element={<CreateParty />} />
      </Route>

      <Route element={<MainLayout />}>
        <Route path="/" element={<Homepage />} />
        <Route element={<RequireAuthForRoute />}>
          <Route path="/overview/discover" element={<Discover />} />
          <Route path="/parties" element={<Parties />} />
          <Route path="/friends" element={<Friends />} />
        </Route>
      </Route>
    </Routes>
  </BrowserRouter>
);
