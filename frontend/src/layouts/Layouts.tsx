import { NavigationBar } from '../components/navigation-bar/NavigationBar';
import { Footer } from '../components/footer/Footer';
import { Outlet } from 'react-router-dom';
import classes from './Layouts.module.scss';
import {PartyNavigationBar} from "../components/navigation-bar/PartyNavigationBar.tsx";
import {PartyProfileDrawer} from "../components/drawer/PartyProfileDrawer.tsx";

export const MainLayout = () => {
    return (
        <div className={classes.app}>
            <NavigationBar />
            <main className={classes.appContent}>
                <Outlet />
            </main>
            <Footer />
        </div>
    );
};

export const PartyLayout = () => {
    return (
        <div>
            <PartyNavigationBar />
            <PartyProfileDrawer />
            <main className={classes.appContent}>
                <Outlet />
            </main>
        </div>
    )
}