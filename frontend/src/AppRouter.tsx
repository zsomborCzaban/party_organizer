import { BrowserRouter, Route, Routes } from 'react-router-dom';
import Discover from './pages/overView/discover/Discover';
import PartiesPage from './pages/overView/partiesPage/PartiesPage';
import Friends from './pages/overView/friends/Friends';
import CreateParty from './pages/createParty/CreateParty';
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
              path='/overview/parties'
              element={<PartiesPage />}
            />
            <Route
              path='/overview/friends'
              element={<Friends />}
            />
            <Route
              path='/createParty'
              element={<CreateParty />}
            />
            <Route
              path='/visitParty/partyHome'
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
          </Route>
        </Routes>
      </div>
      <Footer />
    </div>
  </BrowserRouter>
);
