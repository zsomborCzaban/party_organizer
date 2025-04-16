import { useLocation, useNavigate } from 'react-router-dom';
import classes from './NavigationBarDarkTheme.module.scss';
import { NavigationButton } from './navigation-button/NavigationButton';
import { useAppSelector } from '../../store/store-helper';
import { isUserLoggedIn } from '../../store/sclices/UserSlice';

export const PartyNavigationBar = () => {
    const navigate = useNavigate();
    const location = useLocation();
    const userLoggedIn = useAppSelector(isUserLoggedIn);
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
                {userLoggedIn && (
                    <>
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
                    </>
                )}
            </div>

            <div className={classes.rightSection}>
                {userLoggedIn ? (
                    <NavigationButton
                        buttonText='Profile'
                        navigateToLink='/profile'
                        isActive={location.pathname === '/profile'}
                    />
                ) : (
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