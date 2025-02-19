import { useLocation, useNavigate } from 'react-router-dom';
import classes from './NavigationBar.module.scss';
import { NavigationButton } from './navigation-button/NavigationButton';
import { useAppSelector } from '../../store/store-helper';
import { getUserJwt } from '../../store/sclices/UserSlice';

export const NavigationBar = () => {
  const navigate = useNavigate();
  const location = useLocation();
  const userJwt = useAppSelector(getUserJwt);

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
        <NavigationButton
          buttonText='My Parties'
          navigateToLink='/my-parties'
          isActive={location.pathname === '/my-parties'}
        />
        <NavigationButton
          buttonText='Friends'
          navigateToLink='/friends'
          isActive={location.pathname === '/friends'}
        />
      </div>

      <div className={classes.rightSection}>
        <NavigationButton
          buttonText='Profile'
          navigateToLink='/profile'
          isActive={location.pathname === '/profile'}
        />
        {location.pathname !== '/login' && !userJwt && (
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
