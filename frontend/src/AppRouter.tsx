import { BrowserRouter, Route, Routes } from 'react-router-dom';
import { PartyHome } from './pages/party/party-home/PartyHome';
import { Contributions } from './pages/contributions/Contributions.tsx';
import { HallOfFame } from './pages/hallOfFame/HallOfFame.tsx';
import ManageParty from './pages/party/manage-party/ManageParty.tsx';
import PartySettings from './pages/party/party-settings/PartySettings.tsx';
import { Login } from './pages/authtentication/login/Login';
import Register from './pages/authtentication/register/Register';
import { Homepage } from './pages/home/HomePage.tsx';
import { RequireAuthForRoute } from './auth/RequireAuthForRoute';
import { Parties } from './pages/party/parties/Parties';
import CreateParty from './pages/party/create-party/CreateParty.tsx';
import { Friends } from './pages/friends/Friends';
import { MainLayout, PartyLayout } from "./layouts/Layouts.tsx";
import { Cocktails } from "./pages/cocktails/Cocktails.tsx";
import {RequireNoAuthForRoute} from "./auth/RequireNoAuthForRoute.tsx";
import {ForgotPassword} from "./pages/authtentication/forgotPassword/ForgotPassword.tsx";
import {ConfirmEmail} from "./pages/authtentication/confirmEmail/ConfirmEmail.tsx";
import {ChangePassword} from "./pages/authtentication/changePassword/ChangePassword.tsx";

export const AppRouter = () => (
  <BrowserRouter>
    <Routes>
      <Route element={<MainLayout />}>
        <Route element={<RequireNoAuthForRoute />} >
          <Route path="/login" element={<Login />} />
          <Route path="/register" element={<Register />} />
          <Route path="/forgotPassword" element={<ForgotPassword />} />
          <Route path="/confirmEmail" element={<ConfirmEmail />} />
          <Route path="/changePassword" element={<ChangePassword />} />
        </Route>
      </Route>

      <Route element={<PartyLayout />}>
        <Route path="/partyHome" element={<PartyHome />} />
        <Route element={<RequireAuthForRoute />}>
          <Route path="/contributions" element={<Contributions />} />
          <Route path="/manageParty" element={<ManageParty />} />
          <Route path="/partySettings" element={<PartySettings />} />
          <Route path="/hallOfFame" element={<HallOfFame />} />
          <Route path="/cocktails" element={<Cocktails />} />
        </Route>
      </Route>

      <Route element={<RequireAuthForRoute />}>
        <Route path="/createParty" element={<CreateParty />} />
      </Route>

      <Route element={<MainLayout />}>
        <Route path="/" element={<Homepage />} />
        <Route element={<RequireAuthForRoute />}>
          <Route path="/parties" element={<Parties />} />
          <Route path="/friends" element={<Friends />} />
        </Route>
      </Route>
    </Routes>
  </BrowserRouter>
);
