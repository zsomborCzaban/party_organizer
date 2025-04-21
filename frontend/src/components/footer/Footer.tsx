import { useNavigate } from 'react-router-dom';
import classes from './Footer.module.scss';

export const Footer = () => {
  const navigate = useNavigate();
  const currentYear = new Date().getFullYear();

  return (
    <footer className={classes.footer}>
      <div className={classes.footerContent}>
        <div className={classes.footerSection}>
          <h3 className={classes.sectionTitle}>Party Organizer</h3>
          <p className={classes.description}>
            Making event planning and party organization easier for everyone.
            Connect, celebrate, and create memories together.
          </p>
        </div>

        <div className={classes.footerSection}>
          <h4>Quick Links</h4>
          <ul>
            <li>
              <button onClick={() => navigate('/')}>Discover</button>
            </li>
            <li>
              <button onClick={() => navigate('/parties')}>Parties</button>
            </li>
            <li>
              <button onClick={() => navigate('/friends')}>Friends</button>
            </li>
          </ul>
        </div>

        {/*<div className={classes.footerSection}>*/}
        {/*  <h4>Support</h4>*/}
        {/*  <ul>*/}
        {/*    <li>*/}
        {/*      <button onClick={() => navigate('/help')}>Help Center</button>*/}
        {/*    </li>*/}
        {/*    <li>*/}
        {/*      <button onClick={() => navigate('/contact')}>Contact Us</button>*/}
        {/*    </li>*/}
        {/*    <li>*/}
        {/*      <button onClick={() => navigate('/privacy')}>Privacy Policy</button>*/}
        {/*    </li>*/}
        {/*    <li>*/}
        {/*      <button onClick={() => navigate('/terms')}>Terms of Service</button>*/}
        {/*    </li>*/}
        {/*  </ul>*/}
        {/*</div>*/}

        <div className={classes.footerSection}>
          <h4>Connect With Us</h4>
          <div className={classes.socialLinks}>
            <a href="https://www.instagram.com/party_organizer_the_thesis/" target="_blank" rel="noopener noreferrer">
              Instagram
            </a>
            <a href="mailto:partyorganizercustomerservice@gmail.com" className={classes.emailLink}>
              Email Us
            </a>
          </div>
        </div>
      </div>

      <div className={classes.footerBottom}>
        <div className={classes.copyright}>
          © {currentYear} Party Organizer. All rights reserved.
        </div>
      </div>
    </footer>
  );
};
