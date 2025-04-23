import { useLocation, useNavigate } from 'react-router-dom';
import classes from './NavigationBarDarkTheme.module.scss';
import { NavigationButtonDarkTheme } from './navigation-button/NavigationButtonDarkTheme.tsx';
import { useAppSelector } from '../../store/store-helper';
import { isUserLoggedIn } from '../../store/slices/UserSlice';
import { getUserName } from "../../auth/AuthUserUtil.ts";
import {NavigationDrawerButtonDarkTheme} from "./navigation-button/NavigationDrawerButtonDarkTheme.tsx";
import {useDispatch} from "react-redux";
import {AppDispatch} from "../../store/store.ts";
import {togglePartyProfileDrawer} from "../../store/slices/profileDrawersSlice.ts";

export const PartyNavigationBar = () => {
    const dispatch = useDispatch<AppDispatch>()
    const navigate = useNavigate();
    const location = useLocation();
    const userLoggedIn = useAppSelector(isUserLoggedIn);
    const partyName = localStorage.getItem('partyName') || 'Unexpected error';
    const organizerName = localStorage.getItem('partyOrganizerName'); //cannot do || 'Unexpected error' because if a user would choose the name unexpected error, he could see the buttons ^^.
    const userName = getUserName()
    const isUserOrganizer = userName && organizerName && userName === organizerName

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
                {userLoggedIn && isUserOrganizer && (
                        <>
                            <NavigationButtonDarkTheme
                                buttonText='Manage Party'
                                navigateToLink='/manageParty'
                                isActive={location.pathname === '/manageParty'}
                            />
                            <NavigationButtonDarkTheme
                                buttonText='Party Settings'
                                navigateToLink='/partySettings'
                                isActive={location.pathname === '/partySettings'}
                            />
                        </>
                    )}
                <NavigationButtonDarkTheme
                    buttonText='Home'
                    navigateToLink='/partyHome'
                    isActive={location.pathname === '/partyHome'}
                />
                {userLoggedIn && (
                    <>
                        <NavigationButtonDarkTheme
                            buttonText='Contributions'
                            navigateToLink='/contributions'
                            isActive={location.pathname === '/contributions'}
                        />
                        <NavigationButtonDarkTheme
                            buttonText='Hall Of Fame'
                            navigateToLink='/hallOfFame'
                            isActive={location.pathname === '/hallOfFame'}
                        />
                        <NavigationButtonDarkTheme
                            buttonText='Cocktails'
                            navigateToLink='/cocktails'
                            isActive={location.pathname === '/cocktails'}
                        />
                    </>
                )}
            </div>

            <div className={classes.rightSection}>
                {userLoggedIn ? (
                    <NavigationDrawerButtonDarkTheme
                        buttonText='Profile'
                        onClick={() => dispatch(togglePartyProfileDrawer())}
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