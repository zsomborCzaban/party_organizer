import { useLocation, useNavigate } from 'react-router-dom';
import classes from './NavigationBar.module.scss';
import { NavigationButton } from './navigation-button/NavigationButton';
import { useAppSelector } from '../../store/store-helper';
import { isUserLoggedIn } from '../../store/slices/UserSlice';
import {toggleDefaultProfileDrawer} from "../../store/slices/profileDrawersSlice.ts";
import {NavigationDrawerButton} from "./navigation-button/NavigationDrawerButton.tsx";
import {useDispatch} from "react-redux";
import {AppDispatch} from "../../store/store.ts";
import {useState} from 'react';
import {Menu} from '@mui/icons-material';

export const NavigationBar = () => {
    const dispatch = useDispatch<AppDispatch>();
    const navigate = useNavigate();
    const location = useLocation();
    const userLoggedIn = useAppSelector(isUserLoggedIn);
    const [isMobileMenuOpen, setIsMobileMenuOpen] = useState(false);

    const handleMobileMenuToggle = () => {
        setIsMobileMenuOpen(!isMobileMenuOpen);
    };

    const handleNavigation = (path: string) => {
        navigate(path);
        setIsMobileMenuOpen(false);
    };

    return (
        <>
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
                    <button
                        className={classes.mobileMenuButton}
                        onClick={handleMobileMenuToggle}
                    >
                        <Menu className={classes.menuIcon} />
                    </button>
                    {userLoggedIn && (
                        <NavigationDrawerButton
                            buttonText='Profile'
                            onClick={() => dispatch(toggleDefaultProfileDrawer())}
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

            {/* Mobile Menu */}
            <div className={`${classes.mobileMenu} ${isMobileMenuOpen ? classes.open : ''}`}>
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
                {location.pathname !== '/login' && !userLoggedIn && (
                    <button
                        className={classes.authButton}
                        onClick={() => handleNavigation('/login')}
                    >
                        Sign In
                    </button>
                )}
            </div>
        </>
    );
};
