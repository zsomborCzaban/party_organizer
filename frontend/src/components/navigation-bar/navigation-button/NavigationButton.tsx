import { useNavigate } from 'react-router-dom';
import classes from './NavigationButton.module.scss';

export const NavigationButton = ({ buttonText, navigateToLink }: { buttonText: string; navigateToLink: string }) => {
  const navigate = useNavigate();
  return (
    <div
      onClick={() => navigate(navigateToLink)}
      className={classes.navigationButton}
    >
      {buttonText}
    </div>
  );
};
