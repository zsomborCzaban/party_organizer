import { useLocation, useNavigate } from 'react-router-dom';
import classes from './NavigationBar.module.scss';
import { NavigationButton } from './navigation-button/NavigationButton';

export const NavigationBar = () => {
  const navigate = useNavigate();
  const location = useLocation();
  // check if we are logged in

  return (
    <div className={classes.navBar}>
      <span
        className={classes.pageTitle}
        onClick={() => navigate('/')}
      >
        Party Organizer
      </span>
      <div className={classes.buttonContainer}>
        <NavigationButton
          buttonText='Discover'
          navigateToLink='/'
        />
        <NavigationButton
          buttonText='My Parties' // visible if logged in
          navigateToLink='/'
        />
        <NavigationButton
          buttonText='Friends' // visivle if logged in
          navigateToLink='/'
        />
      </div>
      <div className={classes.authInformationContainer}>
        <NavigationButton
          buttonText='Profile' // if logged in, open profile modal
          navigateToLink='/'
        />

        {/* TODO: add user logged in to also hide button */}
        {location.pathname !== '/login' && (
          <NavigationButton
            buttonText='Login' // if not logged in, login page
            navigateToLink='/login'
          />
        )}
      </div>
    </div>
  );
};
