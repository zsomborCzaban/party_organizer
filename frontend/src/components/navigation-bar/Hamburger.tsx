import {NavigationButton} from "./navigation-button/NavigationButton.tsx";
import classes from "./NavigationBar.module.scss";
import {useAppSelector} from "../../store/store-helper.ts";
import {isUserLoggedIn} from "../../store/slices/UserSlice.ts";

export const Hamburger = () => {
    const userLoggedIn = useAppSelector(isUserLoggedIn);

    return (
        <>
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
        </>
    )
}