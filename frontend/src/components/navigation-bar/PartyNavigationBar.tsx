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
import { useState } from 'react';
import { Menu } from '@mui/icons-material';

export const PartyNavigationBar = () => {
    const dispatch = useDispatch<AppDispatch>();
    const navigate = useNavigate();
    const location = useLocation();
    const userLoggedIn = useAppSelector(isUserLoggedIn);
    const [isMobileMenuOpen, setIsMobileMenuOpen] = useState(false);
    const partyName = localStorage.getItem('partyName') || 'Unexpected error';
    const organizerName = localStorage.getItem('partyOrganizerName');
    const userName = getUserName();
    const isUserOrganizer = userName && organizerName && userName === organizerName;

    const handleMobileMenuToggle = () => {
        setIsMobileMenuOpen(!isMobileMenuOpen);
    };


    return (
        <>
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
                    <button
                        className={classes.mobileMenuButton}
                        onClick={handleMobileMenuToggle}
                        aria-label="Toggle menu"
                    >
                        <Menu className={classes.menuIcon} />
                    </button>
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

            <div
                className={`${classes.mobileMenuBackdrop} ${isMobileMenuOpen ? classes.open : ''}`}
                onClick={handleMobileMenuToggle}
            />

            <div className={`${classes.mobileMenu} ${isMobileMenuOpen ? classes.open : ''}`}>
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
                {!userLoggedIn && (
                    <button
                        className={classes.authButton}
                        onClick={() => {
                            navigate('/login')
                            setIsMobileMenuOpen(false)
                        }}
                    >
                        Sign In
                    </button>
                )}
            </div>
        </>
    );
};