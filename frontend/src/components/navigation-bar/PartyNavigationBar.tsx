import { useLocation, useNavigate } from 'react-router-dom';
import classes from './PartyNavigationBar.module.scss';
import { NavigationButton } from './navigation-button/NavigationButton';

export const PartyNavigationBar = () => {
    const navigate = useNavigate();
    const location = useLocation();
    const partyName = localStorage.getItem('partyName') || 'Unexpected error';

    return (
        <nav className={classes.navBar}>
            <div className={classes.leftSection}>
        <span
            className={classes.pageTitle}
            onClick={() => navigate('/')}
        >
          Party organizer
        </span>
                <span
                    className={classes.pageTitle}
                    onClick={() => navigate('/partyHome')}
                >
          {partyName}
        </span>
            </div>

            <div className={classes.centerSection}>
                <NavigationButton
                    buttonText='Home'
                    navigateToLink='/partyHome'
                    isActive={location.pathname === '/partyHome'}
                />
                <NavigationButton
                    buttonText='Contributions'
                    navigateToLink='/contributions'
                    isActive={location.pathname === '/contributions'}
                />
                <NavigationButton
                    buttonText='Hall Of Fame'
                    navigateToLink='/hallOfFame'
                    isActive={location.pathname === '/hallOfFame'}
                />
                <NavigationButton
                    buttonText='Cocktails'
                    navigateToLink='/cocktails'
                    isActive={location.pathname === '/cocktails'}
                />
            </div>

            <div className={classes.rightSection}>

                <NavigationButton
                    buttonText='Profile'
                    navigateToLink='/profile'
                    isActive={location.pathname === '/profile'}
                />
            </div>
        </nav>
    );
};