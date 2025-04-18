import { useDispatch, useSelector } from 'react-redux';
import {AppDispatch, RootState} from '../../../store/store.ts';
import {useEffect} from "react";
import {closePartyProfileDrawer} from "../../../store/slices/profileDrawersSlice.ts";
import classes from "./PartyProfileDrawer.module.scss"
import {getUser} from "../../../auth/AuthUserUtil.ts";
import {EMPTY_USER} from "../../../data/types/User.ts";
import {toast} from "sonner";
import {
  handleLogoutUtil,
  handleUploadProfilePictureUtil
} from "../../../data/utils/ProfileDrawerUtils.ts";
import {useNavigate} from "react-router-dom";

export const DefaultProfileDrawer = () => {
  const dispatch = useDispatch<AppDispatch>();
  const navigate = useNavigate()
  const isOpen = useSelector((state: RootState) => state.profileDrawers.isOpen);
  const user = getUser() || EMPTY_USER

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

  useEffect(() => {
    if(!user){
      toast.error('Failed to load user')
    }
  }, [user]);

  return (<div>
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
              <img
                  src={user.profile_picture_url}
                  alt={user.username}
                  className={classes.avatar}
              />
              <h3 className={classes.userName}>{user.username}</h3>
              <p className={classes.userEmail}>{user.email}</p>
            </div>

            <nav className={classes.profileMenu}>
              <input className={classes.profileInput} id='file-input' type='file' accept='image/*'
                     onChange={handleUploadProfilePictureUtil}/>
              <label htmlFor='file-input' className={classes.menuItem}>
                Upload profile picture
              </label>
              <button className={classes.menuItem} onClick={() => handleLogoutUtil(navigate)}>Logout</button>
            </nav>
          </div>
        </div>
      </div>
  );
};