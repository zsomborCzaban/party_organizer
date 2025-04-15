import { BrowserRouter, Route, Routes } from 'react-router-dom';
import Discover from './pages/overView/discover/Discover';
import PartiesPage from './pages/overView/partiesPage/PartiesPage';
import { FriendsOld } from './pages/overView/friends/Friends';
import PartyHome from './pages/visitParty/partyHome/PartyHome';
import Contributions from './pages/visitParty/contribution/Contributions';
import HallOfFame from './pages/visitParty/hallOfFame/HallOfFame';
import ManageParty from './pages/visitParty/manageParty/ManageParty';
import PartySettings from './pages/visitParty/partyOptions/PartySettings';
import { Login } from './pages/authtentication/login/Login';
import Register from './pages/authtentication/register/Register';
import { Homepage } from './pages/HomePage';
import { NavigationBar } from './components/navigation-bar/NavigationBar';
import classes from './AppRouter.module.scss';
import { Footer } from './components/footer/Footer';
import { RequireAuthForRoute } from './auth/RequireAuthForRoute';
import { Parties } from './pages/party/parties/Parties';
import { CreateParty } from './pages/party/create-party/CreateParty';
import { Friends } from './pages/friends/Friends';

export const AppRouter = () => (
  <BrowserRouter>
    <div className={classes.app}>
      <NavigationBar />
      <div className={classes.appContent}>
        <Routes>
          {/* Public routes */}
          <Route
            path='/'
            element={<Homepage />}
          />
          <Route
            path='/login'
            element={<Login />}
          />
          <Route
            path='/register'
            element={<Register />}
          />
          {/* Private routes */}
          <Route element={<RequireAuthForRoute />}>
            <Route
              path='/overview/discover'
              element={<Discover />}
            />
            <Route
              path='/createParty'
              element={<CreateParty />}
            />
            <Route
              path='/partyHome'
              element={<PartyHome />}
            />
            <Route
              path='/visitParty/contributions'
              element={<Contributions />}
            />

            <Route
              path='/visitParty/manageParty'
              element={<ManageParty />}
            />
            <Route
              path='/visitParty/partySettings'
              element={<PartySettings />}
            />
            <Route
              path='/visitParty/hallOfFame'
              element={<HallOfFame />}
            />
            <Route
              path='/parties'
              element={<Parties />}
            />
            <Route
              path='/friends'
              element={
                <>
                  <Friends />
                </>
              }
            />
            <Route
              path='/party/create'
              element={<CreateParty />}
            />
          </Route>
        </Routes>
      </div>
      <Footer />
    </div>
  </BrowserRouter>
);
