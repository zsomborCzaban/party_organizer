import { useLocation, useNavigate } from 'react-router-dom';
import classes from './NavigationBar.module.scss';
import { NavigationButton } from './navigation-button/NavigationButton';
import { useAppSelector } from '../../store/store-helper';
import { isUserLoggedIn } from '../../store/slices/UserSlice';

export const NavigationBar = () => {
  const navigate = useNavigate();
  const location = useLocation();
  const userLoggedIn = useAppSelector(isUserLoggedIn);

  return (
    <nav className={classes.navBar}>
      <div className={classes.leftSection}>
        <span
          className={classes.pageTitle}
          onClick={() => navigate('/')}
        >
          Party Organizer
        </span>
      </div>

      <div className={classes.centerSection}>
        <NavigationButton
          buttonText='Discover'
          navigateToLink='/'
          isActive={location.pathname === '/'}
        />
        {userLoggedIn && (
          <>
            <NavigationButton
              buttonText='Parties'
              navigateToLink='/parties'
              isActive={location.pathname === '/parties'}
            />
            <NavigationButton
              buttonText='Friends'
              navigateToLink='/friends'
              isActive={location.pathname === '/friends'}
            />
          </>
        )}
      </div>

      <div className={classes.rightSection}>
        {userLoggedIn && (
          <NavigationButton
            buttonText='Profile'
            navigateToLink='/profile'
            isActive={location.pathname === '/profile'}
          />
        )}
        {location.pathname !== '/login' && !userLoggedIn && (
          <button
            className={classes.authButton}
            onClick={() => navigate('/login')}
          >
            Sign In
          </button>
        )}
      </div>
    </nav>
  );
};
