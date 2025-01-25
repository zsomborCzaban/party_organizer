import { useNavigate } from 'react-router-dom';
import classes from './NavigationBar.module.scss';
import { NavigationButton } from './navigation-button/NavigationButton';
import { AuthApi } from '../../api/apis/AuthenticationApi';
import apiClient from '../../api/Api';

export const NavigationBar = () => {
  const autapi = new AuthApi(apiClient);
  const navigate = useNavigate();
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
        <NavigationButton
          buttonText='Profile' // if logged in, open profile modal
          navigateToLink='/'
        />

        <NavigationButton
          buttonText='Login' // if not logged in, login page
          navigateToLink='/'
        />
        <button onClick={()=> autapi.postLogin('daniel', 'sandor')}>sandor</button>
      </div>
      <div className={classes.authInformationContainer}>
      <NavigationButton
          buttonText='Profile' // if logged in, open profile modal
          navigateToLink='/'
        />

        <NavigationButton
          buttonText='Login' // if not logged in, login page
          navigateToLink='/'
        />
      </div>
    </div>
  );
};
