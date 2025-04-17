import { useDispatch, useSelector } from 'react-redux';
import { RootState } from '../../store/store.ts';
import {useEffect} from "react";
import {closePartyProfileDrawer} from "../../store/slices/partyProfileDrawerSlice.ts";
import classes from "./PartyProfileDrawer.module.scss"

export const PartyProfileDrawer = () => {
  const dispatch = useDispatch();
  const isOpen = useSelector((state: RootState) => state.profileDrawer.isOpen);

  useEffect(() => {
    const handleEscape = (e: KeyboardEvent) => {
      if (e.key === 'Escape') {
        dispatch(closePartyProfileDrawer());
      }
    };

    if (isOpen) {
      document.addEventListener('keydown', handleEscape);
      document.body.style.overflow = 'hidden';
    } else {
      document.removeEventListener('keydown', handleEscape);
      document.body.style.overflow = '';
    }

    return () => { //important for ssr, but we keep this here just in case
      document.removeEventListener('keydown', handleEscape);
      document.body.style.overflow = '';
    };
  }, [isOpen, dispatch]);

  return (
      <div>
        {isOpen && (
            <div
                className={classes.overlay}
                onClick={() => dispatch(closePartyProfileDrawer())}
                aria-hidden="true"
            />
        )}

        <div
            className={`${classes.drawer} ${isOpen ? classes.open : ''}`}
            aria-hidden={!isOpen}
        >
          <div className={classes.drawerHeader}>
            <h2>Profile</h2>
            <button
                className={classes.closeButton}
                onClick={() => dispatch(closePartyProfileDrawer())}
                aria-label="Close profile drawer"
            >
              &times;
            </button>
          </div>

          <div className={classes.drawerContent}>
            <div className={classes.profileSection}>
              <div className={classes.avatar}></div>
              <h3 className={classes.userName}>John Doe</h3>
              <p className={classes.userEmail}>john.doe@example.com</p>
            </div>

            <nav className={classes.profileMenu}>
              <button className={classes.menuItem}>Edit Profile</button>
              <button className={classes.menuItem}>Settings</button>
              <button className={classes.menuItem}>Logout</button>
            </nav>
          </div>
        </div>
      </div>
  );
};