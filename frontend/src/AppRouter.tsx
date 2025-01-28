import { BrowserRouter, Route, Routes } from 'react-router-dom';
import PrivateRoute from './auth/PrivateRoute';
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
import Register from './pages/authtentication/Register';
import { Homepage } from './pages/HomePage';
import { NavigationBar } from './components/navigation-bar/NavigationBar';
import classes from './AppRouter.module.scss';

export const AppRouter = () => (
  <BrowserRouter>
    <NavigationBar />
    <div className={classes.appContent}>
      <Routes>
        <Route
          path='/'
          element={<Homepage />}
        />
        <Route
          path='/overview/discover'
          element={<PrivateRoute />}
        >
          <Route
            path='/overview/discover'
            element={<Discover />}
          />
        </Route>
        <Route
          path='/overview/parties'
          element={<PrivateRoute />}
        >
          <Route
            path='/overview/parties'
            element={<PartiesPage />}
          />
        </Route>
        <Route
          path='/overview/friends'
          element={<PrivateRoute />}
        >
          <Route
            path='/overview/friends'
            element={<Friends />}
          />
        </Route>
        <Route
          path='/createParty'
          element={<PrivateRoute />}
        >
          <Route
            path='/createParty'
            element={<CreateParty />}
          />
        </Route>
        <Route
          path='/visitParty/partyHome'
          element={<PrivateRoute />}
        >
          <Route
            path='/visitParty/partyHome'
            element={<PartyHome />}
          />
        </Route>
        <Route
          path='/visitParty/contributions'
          element={<PrivateRoute />}
        >
          <Route
            path='/visitParty/contributions'
            element={<Contributions />}
          />
        </Route>
        <Route
          path='/visitParty/hallOfFame'
          element={<PrivateRoute />}
        >
          <Route
            path='/visitParty/hallOfFame'
            element={<HallOfFame />}
          />
        </Route>
        <Route
          path='/visitParty/manageParty'
          element={<PrivateRoute />}
        >
          <Route
            path='/visitParty/manageParty'
            element={<ManageParty />}
          />
        </Route>
        <Route
          path='/visitParty/partySettings'
          element={<PrivateRoute />}
        >
          <Route
            path='/visitParty/partySettings'
            element={<PartySettings />}
          />
        </Route>
        <Route
          path='/login'
          element={<Login />}
        />
        <Route
          path='/register'
          element={<Register />}
        />
      </Routes>
    </div>
  </BrowserRouter>
);
